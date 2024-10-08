package keeper

import (
	"bytes"
	"encoding/json"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/Lorenzo-Protocol/lorenzo/v3/x/plan/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/evmos/ethermint/server/config"
	evmtypes "github.com/evmos/ethermint/x/evm/types"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

var (
	RevertSelectorInvalidParams      = crypto.Keccak256([]byte("InvalidParams()"))[:4]
	RevertSelectorEmptyMerkleRoot    = crypto.Keccak256([]byte("EmptyMerkleRoot()"))[:4]
	RevertSelectorAlreadyClaimed     = crypto.Keccak256([]byte("AlreadyClaimed()"))[:4]
	RevertSelectorInvalidMerkleProof = crypto.Keccak256([]byte("InvalidMerkleProof()"))[:4]
)

// CallEVM performs a smart contract method call using given args
func (k Keeper) CallEVM(
	ctx sdk.Context,
	contractABI abi.ABI,
	from, to common.Address,
	commit bool,
	method string,
	args ...interface{},
) (*evmtypes.MsgEthereumTxResponse, error) {
	data, err := contractABI.Pack(method, args...)
	if err != nil {
		return nil, errorsmod.Wrap(
			types.ErrABIPack,
			errorsmod.Wrap(err, "failed to create transaction data").Error(),
		)
	}

	resp, err := k.CallEVMWithData(ctx, from, &to, data, commit)
	if err != nil {
		return nil, errorsmod.Wrapf(err, "contract call failed: method '%s', contract '%s'", method, to)
	}
	return resp, nil
}

// CallEVMWithData performs a smart contract method call using contract data
func (k Keeper) CallEVMWithData(
	ctx sdk.Context,
	from common.Address,
	contractAddress *common.Address,
	data []byte,
	commit bool,
) (*evmtypes.MsgEthereumTxResponse, error) {
	nonce, err := k.accountKeeper.GetSequence(ctx, from.Bytes())
	if err != nil {
		return nil, err
	}

	gasCap := config.DefaultGasCap
	if commit {
		args, err := json.Marshal(evmtypes.TransactionArgs{
			From: &from,
			To:   contractAddress,
			Data: (*hexutil.Bytes)(&data),
		})
		if err != nil {
			return nil, errorsmod.Wrapf(errortypes.ErrJSONMarshal, "failed to marshal tx args: %s", err.Error())
		}

		// k.evmKeeper.EthCall
		callRes, err := k.evmKeeper.EthCall(sdk.WrapSDKContext(ctx), &evmtypes.EthCallRequest{
			Args:    args,
			GasCap:  config.DefaultGasCap,
			ChainId: k.evmKeeper.ChainID().Int64(),
		})
		if err != nil {
			return nil, err
		}
		if len(callRes.VmError) != 0 {
			revertRetErr := UnpackRevert(callRes.Ret)
			return nil, revertRetErr
		}

		// k.evmKeeper.EstimateGas
		gasRes, err := k.evmKeeper.EstimateGas(sdk.WrapSDKContext(ctx), &evmtypes.EthCallRequest{
			Args:    args,
			GasCap:  config.DefaultGasCap,
			ChainId: k.evmKeeper.ChainID().Int64(),
		})
		if err != nil {
			return nil, err
		}
		gasCap = gasRes.Gas
	}

	msg := ethtypes.NewMessage(
		from,
		contractAddress,
		nonce,
		big.NewInt(0), // amount
		gasCap,        // gasLimit
		big.NewInt(0), // gasFeeCap
		big.NewInt(0), // gasTipCap
		big.NewInt(0), // gasPrice
		data,
		ethtypes.AccessList{}, // AccessList
		!commit,               // isFake
	)

	res, err := k.evmKeeper.ApplyMessage(ctx, msg, evmtypes.NewNoOpTracer(), commit)
	if err != nil {
		return nil, err
	}

	if res.Failed() {
		return nil, errorsmod.Wrap(evmtypes.ErrVMExecution, res.VmError)
	}

	return res, nil
}

// DeployContract deploys a new Stake Plan Logic contract.
//
// Parameters:
// - ctx: the SDK context.
// - compiledContract: the compiled contract to deploy.
// Returns:
// - common.Address: the address of the deployed contract.
// - error: an error if the deployment fails.
func (k Keeper) DeployContract(
	ctx sdk.Context,
	compiledContract evmtypes.CompiledContract,
) (common.Address, error) {
	// pack proxy contract arguments
	contractArgs, err := compiledContract.ABI.Pack(
		"",
	)
	if err != nil {
		return common.Address{}, errorsmod.Wrapf(
			types.ErrABIPack, "contract arguments are invalid: %s", err.Error())
	}
	data := make([]byte, len(compiledContract.Bin)+len(contractArgs))
	copy(data[:len(compiledContract.Bin)], compiledContract.Bin)
	copy(data[len(compiledContract.Bin):], contractArgs)

	deployer := k.getModuleEthAddress(ctx)
	nonce, err := k.accountKeeper.GetSequence(ctx, deployer.Bytes())
	if err != nil {
		return common.Address{}, err
	}

	// generate contract address
	contractAddr := crypto.CreateAddress(deployer, nonce)
	result, err := k.CallEVMWithData(ctx, deployer, nil, data, true)
	if err != nil {
		return common.Address{}, errorsmod.Wrapf(
			err,
			"failed to deploy contract for stake plan logic contract")
	}
	if result.Failed() {
		return common.Address{}, errorsmod.Wrapf(
			types.ErrVMExecution,
			"failed to deploy contract for stake plan logic contract, reason: %s", result.Revert())
	}

	return contractAddr, nil
}

func UnpackRevert(data []byte) error {
	if len(data) < 4 {
		return errors.New("invalid data for unpacking")
	}
	switch {
	case bytes.Equal(data[:4], RevertSelectorInvalidParams):
		return errors.New("InvalidParams")
	case bytes.Equal(data[:4], RevertSelectorEmptyMerkleRoot):
		return errors.New("EmptyMerkleRoot")
	case bytes.Equal(data[:4], RevertSelectorAlreadyClaimed):
		return errors.New("AlreadyClaimed")
	case bytes.Equal(data[:4], RevertSelectorInvalidMerkleProof):
		return errors.New("InvalidMerkleProof")
	}

	return evmtypes.NewExecErrorWithReason(data)
}
