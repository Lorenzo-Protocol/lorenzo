package keeper

import (
	"math/big"

	errorsmod "cosmossdk.io/errors"
	"github.com/Lorenzo-Protocol/lorenzo/contracts"
	"github.com/Lorenzo-Protocol/lorenzo/x/plan/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// DeployYATContract deploys a new Yield Accruing Token (YAT) contract.
//
// Parameters:
// - ctx: the SDK context.
// - name: the name of the YAT contract.
// - symbol: the symbol of the YAT contract.
// Returns:
// - common.Address: the address of the deployed contract.
// - error: an error if the deployment fails.
func (k Keeper) DeployYATContract(
	ctx sdk.Context,
	name,
	symbol string,
) (common.Address, error) {
	contractArgs, err := contracts.YieldAccruingTokenContract.ABI.Pack(
		"",
		name,
		symbol,
		types.ModuleAddress,
	)
	if err != nil {
		return common.Address{}, errorsmod.Wrap(
			types.ErrABIPack,
			errorsmod.Wrap(err, "failed to create transaction data").Error(),
		)
	}
	data := make([]byte, len(contracts.YieldAccruingTokenContract.Bin)+len(contractArgs))
	copy(data[:len(contracts.YieldAccruingTokenContract.Bin)], contracts.YieldAccruingTokenContract.Bin)
	copy(data[len(contracts.YieldAccruingTokenContract.Bin):], contractArgs)

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
			"failed to deploy contract for yat logic contract")
	}
	if result.Failed() {
		return common.Address{}, errorsmod.Wrapf(
			types.ErrVMExecution,
			"failed to deploy contract for yat logic contract, reason: %s", result.Revert())
	}

	return contractAddr, nil
}

// ClaimReward mints YAT tokens to an account.
//
// Parameters:
// - ctx: the SDK context.
// - contractAddress: the address of the YAT contract.
// - erc20Addr: the address of the ERC20 token to claim rewards for.
// - amount: the amount of tokens to claim.
//
// Returns:
// - error: an error if the Claim Reward fails.
func (k Keeper) ClaimReward(
	ctx sdk.Context,
	contractAddress common.Address,
	erc20Addr common.Address,
	amount *big.Int,
) error {
	contractABI := contracts.YieldAccruingTokenContract.ABI
	res, err := k.CallEVM(
		ctx,
		contractABI,
		k.getModuleEthAddress(ctx),
		contractAddress,
		true,
		types.YATMethodClaimReward,
		// args
		erc20Addr,
		amount,
	)
	if err != nil {
		return err
	}
	if res.Failed() {
		return errorsmod.Wrapf(
			types.ErrVMExecution, "failed to yat contract: %s, reason: %s",
			contractAddress.String(),
			res.Revert(),
		)
	}
	return nil
}

// Mint mints YAT tokens to an account.
//
// Parameters:
// - ctx: the SDK context.
// - contractAddress: the address of the YAT contract.
// - to: the address of the account to mint tokens for.
// - amount: the amount of tokens to mint.
//
// Returns:
// - error: an error if the mint fails.
func (k Keeper) Mint(
	ctx sdk.Context,
	contractAddress, to common.Address,
	amount *big.Int,
) error {
	contractABI := contracts.YieldAccruingTokenContract.ABI
	res, err := k.CallEVM(
		ctx,
		contractABI,
		k.getModuleEthAddress(ctx),
		contractAddress,
		true,
		types.YATMethodMint,
		// args
		to,
		amount,
	)
	if err != nil {
		return err
	}
	if res.Failed() {
		return errorsmod.Wrapf(
			types.ErrVMExecution, "failed to yat contract: %s, reason: %s",
			contractAddress.String(),
			res.Revert(),
		)
	}
	return nil
}

func (k Keeper) SetMinter(
	ctx sdk.Context,
	contractAddress common.Address,
	minter common.Address,
) error {
	contractABI := contracts.YieldAccruingTokenContract.ABI
	res, err := k.CallEVM(
		ctx,
		contractABI,
		k.getModuleEthAddress(ctx),
		contractAddress,
		true,
		types.YATMethodSetMinter,
		// args
		minter,
	)
	if err != nil {
		return err
	}
	if res.Failed() {
		return errorsmod.Wrapf(
			types.ErrVMExecution, "failed to yat contract: %s, reason: %s",
			contractAddress.String(),
			res.Revert(),
		)
	}
	return nil
}

func (k Keeper) RemoveMinter(
	ctx sdk.Context,
	contractAddress common.Address,
	minter common.Address,
) error {
	contractABI := contracts.YieldAccruingTokenContract.ABI
	res, err := k.CallEVM(
		ctx,
		contractABI,
		k.getModuleEthAddress(ctx),
		contractAddress,
		true,
		types.YATMethodRemoveMinter,
		// args
		minter,
	)
	if err != nil {
		return err
	}
	if res.Failed() {
		return errorsmod.Wrapf(
			types.ErrVMExecution, "failed to yat contract: %s, reason: %s",
			contractAddress.String(),
			res.Revert(),
		)
	}
	return nil
}
