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
		// TODO: StakePlanHubAddress, EventName, RetainedBlocks, AllowList, ChainId should be updated according to the mainnet.
		// Note: StakePlanHubAddress is Binance Smart Chain testnet address. It should be updated to mainnet address for mainnet.
		// Note:  ChainId Is Binance Smart Chain testnet chain id. It should be updated to mainnet chain id for mainnet.
		bnbLightClientParams := &bnblightclienttypes.Params{
			StakePlanHubAddress: "0x9ADb675bc89d9EC5d829709e85562b7c99658D59",
			EventName:           "StakeBTC2JoinStakePlan",
			RetainedBlocks:      10000,
			AllowList:           []string{"lrz1xa40j022h2rcmnte47gyjg8688grln94pp84lc"},
			ChainId:             97,
		}

		err := app.BNBLightClientKeeper.SetParams(ctx, bnbLightClientParams)
		if err != nil {
			panic("failed to set bnb light client params")
		}

		return app.ModuleManager.RunMigrations(ctx, c, fromVM)
	}
}
