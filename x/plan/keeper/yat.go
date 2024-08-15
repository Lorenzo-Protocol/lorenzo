package keeper

import (
	"math/big"

	contractsplan "github.com/Lorenzo-Protocol/lorenzo/v3/contracts/plan"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/Lorenzo-Protocol/lorenzo/v3/x/plan/types"
)

const (
	UpdateMinterTypeAdd = iota + 1
	UpdateMinterTypeRemove
)

// UpdateMinter updates the minter of a YAT contract.
//
// Parameters:
// - ctx: the SDK context.
// - contractAddress: the address of the YAT contract.
// - minter: the address of the minter to update.
// - added: a boolean indicating if the minter is being added or removed.
//
// Returns:
// - error: an error if the Update Minter fails.
func (k Keeper) UpdateMinter(
	ctx sdk.Context,
	contractAddress string,
	minter string,
	updateType int,
) error {
	// Check if the contract address is a valid address
	if !common.IsHexAddress(contractAddress) {
		return errorsmod.Wrap(types.ErrContractAddress, "invalid contract address")
	}
	contractAddr := common.HexToAddress(contractAddress)

	// check if the contract address does exist
	yatAcct := k.evmKeeper.GetAccountWithoutBalance(ctx, contractAddr)
	if yatAcct == nil {
		return types.ErrYatContractNotFound
	}
	if !yatAcct.IsContract() {
		return types.ErrYatContractNotContract
	}

	planID := k.GetPlanIdByContractAddr(ctx, minter)
	if planID == 0 {
		return types.ErrStakePlanContractNotFound
	}
	// Check if the minter address is a valid address
	if !common.IsHexAddress(minter) {
		return errorsmod.Wrap(types.ErrEthAddress, "invalid Ethereum address")
	}
	minterAddr := common.HexToAddress(minter)

	// check if the minter address does exist
	minterAcct := k.evmKeeper.GetAccountWithoutBalance(ctx, minterAddr)
	if minterAcct == nil {
		return types.ErrStakePlanContractNotFound
	}
	if !minterAcct.IsContract() {
		return types.ErrStakePlanContractNotContract
	}

	switch updateType {
	case UpdateMinterTypeAdd:
		return k.SetMinter(ctx, contractAddr, minterAddr)
	case UpdateMinterTypeRemove:
		return k.RemoveMinter(ctx, contractAddr, minterAddr)
	default:
		return errorsmod.Wrap(types.ErrInvalidUpdateMinterType, "invalid update type")
	}
}

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
	contractArgs, err := contractsplan.YieldAccruingTokenContract.ABI.Pack(
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
	data := make([]byte, len(contractsplan.YieldAccruingTokenContract.Bin)+len(contractArgs))
	copy(data[:len(contractsplan.YieldAccruingTokenContract.Bin)], contractsplan.YieldAccruingTokenContract.Bin)
	copy(data[len(contractsplan.YieldAccruingTokenContract.Bin):], contractArgs)

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
	contractABI := contractsplan.YieldAccruingTokenContract.ABI
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

// MintFromYat mints YAT tokens to an account.
//
// Parameters:
// - ctx: the SDK context.
// - contractAddress: the address of the YAT contract.
// - to: the address of the account to mint tokens for.
// - amount: the amount of tokens to mint.
//
// Returns:
// - error: an error if the mint fails.
func (k Keeper) MintFromYat(
	ctx sdk.Context,
	contractAddress, to common.Address,
	amount *big.Int,
) error {
	// check if contractAddress is zero address
	zeroAddr := common.Address{}
	if contractAddress == zeroAddr {
		return errorsmod.Wrap(types.ErrContractAddress, "invalid contract address")
	}
	contractABI := contractsplan.YieldAccruingTokenContract.ABI
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

// SetMinter sets a minter for a YAT contract.
//
// Parameters:
// - ctx: the SDK context.
// - contractAddress: the address of the YAT contract.
// - minter: the address of the minter to set.
//
// Returns:
// - error: an error if the Set Minter fails.
func (k Keeper) SetMinter(
	ctx sdk.Context,
	contractAddress common.Address,
	minter common.Address,
) error {
	contractABI := contractsplan.YieldAccruingTokenContract.ABI
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

// RemoveMinter removes a minter from a YAT contract.
//
// Parameters:
// - ctx: the SDK context.
// - contractAddress: the address of the YAT contract.
// - minter: the address of the minter to remove.
//
// Returns:
// - error: an error if the Remove Minter fails.
func (k Keeper) RemoveMinter(
	ctx sdk.Context,
	contractAddress common.Address,
	minter common.Address,
) error {
	contractABI := contractsplan.YieldAccruingTokenContract.ABI
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

func (k Keeper) GetOwner(
	ctx sdk.Context,
	contractAddress common.Address,
) (common.Address, error) {
	contractABI := contractsplan.YieldAccruingTokenContract.ABI
	res, err := k.CallEVM(
		ctx,
		contractABI,
		k.getModuleEthAddress(ctx),
		contractAddress,
		true,
		types.YATMethodOwner,
	)
	if err != nil {
		return common.Address{}, err
	}
	unpacked, err := contractABI.Unpack("owner", res.Ret)
	if err != nil {
		return common.Address{}, errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to unpack owner: %s", err.Error(),
		)
	}
	yatContractAddress, ok := unpacked[0].(common.Address)
	if !ok {
		return common.Address{}, errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to convert YAT contract address to address from contract %s", contractAddress.Hex(),
		)
	}
	if res.Failed() {
		return common.Address{}, errorsmod.Wrapf(
			types.ErrVMExecution, "failed to yat contract: %s, reason: %s",
			contractAddress.String(),
			res.Revert(),
		)
	}
	return yatContractAddress, nil
}

func (k Keeper) HasRoleFromYAT(
	ctx sdk.Context,
	contractAddress common.Address,
	roleName string,
	account common.Address,
) (bool, error) {
	contractABI := contractsplan.YieldAccruingTokenContract.ABI
	role := crypto.Keccak256Hash([]byte(roleName))
	res, err := k.CallEVM(
		ctx,
		contractABI,
		k.getModuleEthAddress(ctx),
		contractAddress,
		true,
		types.YATMethodHasRole,
		role,
		account,
	)
	if err != nil {
		return false, err
	}
	unpacked, err := contractABI.Unpack(types.YATMethodHasRole, res.Ret)
	if err != nil {
		return false, errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to unpack hasRole: %s", err.Error(),
		)
	}
	hasRole, ok := unpacked[0].(bool)
	if !ok {
		return false, errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to convert hasRole to bool from contract %s", contractAddress.Hex(),
		)
	}
	if res.Failed() {
		return false, errorsmod.Wrapf(
			types.ErrVMExecution, "failed to yat contract: %s, reason: %s",
			contractAddress.String(),
			res.Revert(),
		)
	}
	return hasRole, nil
}

func (k Keeper) BalanceOfFromYAT(
	ctx sdk.Context,
	contractAddress common.Address,
	account common.Address,
) (*big.Int, error) {
	contractABI := contractsplan.YieldAccruingTokenContract.ABI
	res, err := k.CallEVM(
		ctx,
		contractABI,
		k.getModuleEthAddress(ctx),
		contractAddress,
		true,
		types.YATMethodBalanceOf,
		account,
	)
	if err != nil {
		return nil, err
	}
	unpacked, err := contractABI.Unpack(types.YATMethodBalanceOf, res.Ret)
	if err != nil {
		return nil, errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to unpack balanceOf: %s", err.Error(),
		)
	}
	balance, ok := unpacked[0].(*big.Int)
	if !ok {
		return nil, errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to convert balance to big.Int from contract %s", contractAddress.Hex(),
		)
	}
	if res.Failed() {
		return nil, errorsmod.Wrapf(
			types.ErrVMExecution, "failed to yat contract: %s, reason: %s",
			contractAddress.String(),
			res.Revert(),
		)
	}
	return balance, nil
}
