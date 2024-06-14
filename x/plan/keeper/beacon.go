package keeper

import (
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	errorsmod "cosmossdk.io/errors"
	"github.com/Lorenzo-Protocol/lorenzo/contracts"
	"github.com/Lorenzo-Protocol/lorenzo/x/plan/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// DeployBeacon deploys a new Yield Accruing Token (YAT) contract.
//
// Parameters:
// - ctx: the SDK context.
// - deployer: the address of the account deploying the contract.
// Returns:
// - common.Address: the address of the deployed contract.
// - error: an error if the deployment fails.
func (k Keeper) DeployBeacon(
	ctx sdk.Context,
	implementation common.Address,
) (common.Address, error) {
	deployer := k.getModuleEthAddress(ctx)
	contractArgs, err := contracts.BeaconContract.ABI.Pack(
		"",
		implementation,
		deployer,
	)
	if err != nil {
		return common.Address{}, errorsmod.Wrap(types.ErrABIPack, fmt.Sprintf("failed to pack contract arguments: %s", err))
	}

	data := make([]byte, len(contracts.BeaconContract.Bin)+len(contractArgs))
	copy(data[:len(contracts.BeaconContract.Bin)], contracts.BeaconContract.Bin)
	copy(data[len(contracts.BeaconContract.Bin):], contractArgs)

	nonce, err := k.accountKeeper.GetSequence(ctx, deployer.Bytes())
	if err != nil {
		return common.Address{}, err
	}
	// generate contract address
	contractAddr := crypto.CreateAddress(deployer, nonce)
	result, err := k.CallEVMWithData(ctx, deployer, nil, data, true)
	if err != nil {
		return common.Address{}, errorsmod.Wrapf(err, "failed to deploy contract for beacon")
	}

	if result.Failed() {
		return common.Address{}, errorsmod.Wrapf(
			types.ErrVMExecution,
			"failed to deploy contract for beacon, reason: %s", result.Revert())
	}
	return contractAddr, nil
}

// UpgradeYAT upgrades the YAT contract to a new implementation.
//
// Parameters:
// - ctx: the SDK context.
// - implementation: the address of the new implementation contract.
//
// Returns:
// - error: an error if the upgrade fails.
func (k Keeper) UpgradeYAT(
	ctx sdk.Context,
	implementation common.Address,
) error {
	if !common.IsHexAddress(implementation.Hex()) {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid implementation address: %s", implementation.Hex())
	}

	params := k.GetParams(ctx)

	if len(params.Beacon) == 0 {
		return errorsmod.Wrapf(types.ErrBeaconNotSet, "beacon not set")
	}

	beacon := common.HexToAddress(params.Beacon)

	caller := k.getModuleEthAddress(ctx)

	res, err := k.CallEVM(
		ctx,
		contracts.BeaconContract.ABI,
		caller,
		beacon,
		true,
		types.BeaconMethodUpgradeTo,
		implementation,
	)
	if err != nil {
		return err
	}
	if res.Failed() {
		return errorsmod.Wrapf(
			types.ErrVMExecution, "failed to upgrade contract reason: %s",
			res.Revert(),
		)
	}
	return nil
}
