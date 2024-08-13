package keeper

import (
	"fmt"

	contractsplan "github.com/Lorenzo-Protocol/lorenzo/v3/contracts/plan"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	errorsmod "cosmossdk.io/errors"
	"github.com/Lorenzo-Protocol/lorenzo/v3/x/plan/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// DeployBeaconForPlan deploys a new Yield Accruing Token (YAT) contract.
//
// Parameters:
// - ctx: the SDK context.
// - deployer: the address of the account deploying the contract.
// Returns:
// - common.Address: the address of the deployed contract.
// - error: an error if the deployment fails.
func (k Keeper) DeployBeaconForPlan(
	ctx sdk.Context,
	implementation common.Address,
) (common.Address, error) {
	deployer := k.getModuleEthAddress(ctx)
	contractArgs, err := contractsplan.BeaconContract.ABI.Pack(
		"",
		implementation,
		deployer,
	)
	if err != nil {
		return common.Address{}, errorsmod.Wrap(types.ErrABIPack, fmt.Sprintf("failed to pack contract arguments: %s", err))
	}

	data := make([]byte, len(contractsplan.BeaconContract.Bin)+len(contractArgs))
	copy(data[:len(contractsplan.BeaconContract.Bin)], contractsplan.BeaconContract.Bin)
	copy(data[len(contractsplan.BeaconContract.Bin):], contractArgs)

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

// UpgradeBeaconForPlan upgrades the Plan contract to a new implementation.
//
// Parameters:
// - ctx: the SDK context.
// - implementation: the address of the new implementation contract.
//
// Returns:
// - error: an error if the upgrade fails.
func (k Keeper) UpgradeBeaconForPlan(
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
		contractsplan.BeaconContract.ABI,
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

// GetPlanImplementationFromBeacon returns the implementation address of the Plan Logic contract.
//
// Parameters:
// - ctx: the SDK context.
// Returns:
// - common.Address: the address of the implementation contract.
// - error: an error if the implementation address is not found.
func (k Keeper) GetPlanImplementationFromBeacon(
	ctx sdk.Context,
) (common.Address, error) {
	params := k.GetParams(ctx)

	if len(params.Beacon) == 0 {
		return common.Address{}, errorsmod.Wrapf(types.ErrBeaconNotSet, "beacon not set")
	}
	contractABI := contractsplan.BeaconContract.ABI

	beacon := common.HexToAddress(params.Beacon)

	caller := k.getModuleEthAddress(ctx)

	res, err := k.CallEVM(
		ctx,
		contractABI,
		caller,
		beacon,
		false,
		types.BeaconMethodImplementation,
	)
	if err != nil {
		return common.Address{}, err
	}
	if res.Failed() {
		return common.Address{}, errorsmod.Wrapf(
			types.ErrVMExecution, "failed to get implementation address reason: %s",
			res.Revert(),
		)
	}

	unpacked, err := contractABI.Unpack(types.BeaconMethodImplementation, res.Ret)
	if err != nil || len(unpacked) == 0 {
		return common.Address{}, errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to unpack plan description from contract %s", beacon.Hex(),
		)
	}
	implementation, ok := unpacked[0].(common.Address)
	if !ok {
		return common.Address{}, errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to unpack plan description from contract %s", beacon.Hex(),
		)
	}

	return implementation, nil
}
