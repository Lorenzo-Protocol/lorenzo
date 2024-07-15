package app

import (
	"fmt"

	"github.com/Lorenzo-Protocol/lorenzo/app/upgrades"
	v200 "github.com/Lorenzo-Protocol/lorenzo/app/upgrades/v200"

	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

var router = upgrades.NewUpgradeRouter()

func init() {
	// register v2.0 upgrade plan
	router.Register(v200.Upgrade)
}

// RegisterUpgradePlans register a handler of upgrade plan
func (app *LorenzoApp) RegisterUpgradePlans() {
	app.setupUpgradeStoreLoaders()
	app.setupUpgradeHandlers()
}

func (app *LorenzoApp) appKeepers() upgrades.AppKeepers {
	return upgrades.AppKeepers{
		AppCodec:        app.AppCodec(),
		BankKeeper:      app.BankKeeper,
		AccountKeeper:   app.AccountKeeper,
		GetKey:          app.GetKey,
		ModuleManager:   app.mm,
		EvmKeeper:       app.EvmKeeper,
		FeeMarketKeeper: app.FeeMarketKeeper,
		ReaderWriter:    app,
	}
}

// configure store loader that checks if version == upgradeHeight and applies store upgrades
func (app *LorenzoApp) setupUpgradeStoreLoaders() {
	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Sprintf("failed to read upgrade info from disk %s", err))
	}

	if app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		return
	}

	app.SetStoreLoader(
		upgradetypes.UpgradeStoreLoader(
			upgradeInfo.Height,
			router.UpgradeInfo(upgradeInfo.Name).StoreUpgrades,
		),
	)
}

func (app *LorenzoApp) setupUpgradeHandlers() {
	// SAFE: upgrade handlers are registered in the init function
	for upgradeName, upgrade := range router.Routers() {
		app.UpgradeKeeper.SetUpgradeHandler(
			upgradeName,
			upgrade.UpgradeHandlerConstructor(
				app.mm,
				app.configurator,
				app.appKeepers(),
			),
		)
	}
}
