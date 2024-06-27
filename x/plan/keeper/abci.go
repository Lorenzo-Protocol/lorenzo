package keeper

import (
	"github.com/Lorenzo-Protocol/lorenzo/x/plan/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// EndBlocker handles block beginning logic for plan module
func (k Keeper) EndBlocker(ctx sdk.Context) {
	logger := k.Logger(ctx)
	params := k.GetParams(ctx)

	// deploy a new beacon contract if the current beacon contract is empty
	if len(params.Beacon) == 0 {
		// deploy a new beacon proxy contract & deploy a new plan logic contract
		// 1. deploy a new plan logic contract
		logicAddr, err := k.DeployStakePlanLogicContract(ctx)
		if err != nil {
			panic(err)
		}
		// 2. deploy a new plan beacon contract
		beaconAddr, err := k.DeployBeacon(ctx, logicAddr)
		if err != nil {
			panic(err)
		}
		params.Beacon = beaconAddr.Hex()
		if err := k.SetParams(ctx, params); err != nil {
			panic(err)
		}
		logger.Info(
			"deploy plan contract",
			"beacon",
			beaconAddr.Hex(),
			"logic",
			logicAddr.Hex(),
		)
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventSetParams,
				sdk.NewAttribute(types.AttributeKeyBeaconAddr, beaconAddr.String()),
				sdk.NewAttribute(types.AttributeKeyLogicAddr, logicAddr.String()),
			),
		)
	}
}
