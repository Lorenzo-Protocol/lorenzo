package v300

import (
	"github.com/Lorenzo-Protocol/lorenzo/v3/app/upgrades"
	"github.com/Lorenzo-Protocol/lorenzo/v3/x/bnblightclient"
	bnblightclienttypes "github.com/Lorenzo-Protocol/lorenzo/v3/x/bnblightclient/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

// Upgrade defines a struct containing necessary fields that a SoftwareUpgradeProposal
var Upgrade = upgrades.Upgrade{
	UpgradeName:               "v3.0",
	UpgradeHandlerConstructor: upgradeHandlerConstructor,
	StoreUpgrades: &storetypes.StoreUpgrades{
		Added: []string{
			bnblightclienttypes.StoreKey,
		},
	},
}

func upgradeHandlerConstructor(
	m *module.Manager,
	c module.Configurator,
	app upgrades.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		fromVM[bnblightclienttypes.ModuleName] = bnblightclient.AppModule{}.ConsensusVersion()

		// bnb light client module init
		// 1. set params
		bnbLightClientParams := &bnblightclienttypes.Params{
			StakePlanHubAddress: "0x8BCd1CCDA853677Ac865C882B60FBaF5030EeF50",
			EventName:           "StakeBTC2JoinStakePlan",
			RetainedBlocks:      1000000,
			AllowList:           []string{"lrz1lh79r9v0gtzljm6wa4ya6dfrz2jappy406wflw"},
			ChainId:             56,
		}

		err := app.BNBLightClientKeeper.SetParams(ctx, bnbLightClientParams)
		if err != nil {
			panic("failed to set bnb light client params")
		}

		feeParams := app.FeeKeeper.GetParams(ctx)
		newNonFeeMsgs := []string{
			"/lorenzo.btcstaking.v1.MsgCreateBTCStaking",
			"/lorenzo.btclightclient.v1.MsgInsertHeaders",
			"/lorenzo.plan.v1.MsgSetMerkleRoot",
			"/lorenzo.bnblightclient.v1.MsgUploadHeaders",
			"/lorenzo.bnblightclient.v1.MsgUpdateHeader",
			"/lorenzo.btcstaking.v1.MsgCreateBTCBStaking",
		}
		feeParams.NonFeeMsgs = newNonFeeMsgs
		if err := app.FeeKeeper.SetParams(ctx, feeParams); err != nil {
			panic("failed to set fee params")
		}

		return app.ModuleManager.RunMigrations(ctx, c, fromVM)
	}
}
