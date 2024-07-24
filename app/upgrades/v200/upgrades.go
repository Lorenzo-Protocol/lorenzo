package v200

import (
	"github.com/Lorenzo-Protocol/lorenzo/app/upgrades"
	"github.com/Lorenzo-Protocol/lorenzo/x/agent"
	agenttypes "github.com/Lorenzo-Protocol/lorenzo/x/agent/types"
	"github.com/Lorenzo-Protocol/lorenzo/x/plan"
	plantypes "github.com/Lorenzo-Protocol/lorenzo/x/plan/types"
	"github.com/Lorenzo-Protocol/lorenzo/x/token"
	tokentypes "github.com/Lorenzo-Protocol/lorenzo/x/token/types"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

// Upgrade defines a struct containing necessary fields that a SoftwareUpgradeProposal
var Upgrade = upgrades.Upgrade{
	UpgradeName:               "v2.0",
	UpgradeHandlerConstructor: upgradeHandlerConstructor,
	StoreUpgrades: &storetypes.StoreUpgrades{
		Added: []string{
			agenttypes.StoreKey,
			plantypes.StoreKey,
			tokentypes.StoreKey,
		},
	},
}

func upgradeHandlerConstructor(
	m *module.Manager,
	c module.Configurator,
	app upgrades.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		fromVM[agenttypes.ModuleName] = agent.AppModule{}.ConsensusVersion()
		fromVM[plantypes.ModuleName] = plan.AppModule{}.ConsensusVersion()
		fromVM[tokentypes.ModuleName] = token.AppModule{}.ConsensusVersion()

		// agent module init
		// 1. set admin
		admin, err := sdk.AccAddressFromBech32("lrz1xa40j022h2rcmnte47gyjg8688grln94pp84lc")
		if err != nil {
			return nil, err
		}

		if err := app.AgentKeeper.SetAdmin(ctx, admin); err != nil {
			return nil, err
		}

		// 2. set agents
		if err := migrateAgentFromBTCStakingToAgent(ctx, app); err != nil {
			return nil, err
		}

		// plan module init
		planParams := plantypes.Params{
			AllowList: []string{"lrz1xa40j022h2rcmnte47gyjg8688grln94pp84lc"},
		}

		if err := app.PlanKeeper.SetParams(ctx, planParams); err != nil {
			return nil, err
		}

		// 3. set token params
		tokenParams := tokentypes.DefaultParams()
		app.TokenKeeper.SetParams(ctx, tokenParams)

		return app.ModuleManager.RunMigrations(ctx, c, fromVM)
	}
}
