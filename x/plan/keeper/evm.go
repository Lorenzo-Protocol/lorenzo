package keeper

import (
	"encoding/json"
	"math/big"

	"github.com/Lorenzo-Protocol/lorenzo/v2/x/plan/types"

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

		// k.evmKeeper.
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
