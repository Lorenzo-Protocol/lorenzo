package v400

import (
	"errors"

	"github.com/Lorenzo-Protocol/lorenzo/v3/app/upgrades"
	bnblightclienttypes "github.com/Lorenzo-Protocol/lorenzo/v3/x/bnblightclient/types"
	"github.com/Lorenzo-Protocol/lorenzo/v3/x/ccev"
	ccevtypes "github.com/Lorenzo-Protocol/lorenzo/v3/x/ccev/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Upgrade defines a struct containing necessary fields that a SoftwareUpgradeProposal
var Upgrade = upgrades.Upgrade{
	UpgradeName:               "v4.0",
	UpgradeHandlerConstructor: upgradeHandlerConstructor,
	StoreUpgrades: &storetypes.StoreUpgrades{
		Added: []string{
			ccevtypes.StoreKey,
		},
	},
}

func upgradeHandlerConstructor(
	m *module.Manager,
	c module.Configurator,
	app upgrades.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		fromVM[ccevtypes.ModuleName] = ccev.AppModule{}.ConsensusVersion()

		// merge bnb light client module
		if err := mergeBnbClient(ctx, app); err != nil {
			return nil, err
		}
		return app.ModuleManager.RunMigrations(ctx, c, fromVM)
	}
}

func mergeBnbClient(ctx sdk.Context, app upgrades.AppKeepers) error {
	bnbclient := app.BNBLightClientKeeper.GetParams(ctx)
	// ccev module init
	// 1. set params
	ccevParams := &ccevtypes.Params{
		AllowList: bnbclient.AllowList,
	}
	if err := app.CCEVkeeper.SetParams(ctx, ccevParams); err != nil {
		return errors.New("failed to set bnb light client params")
	}

	latestHeader, has := app.BNBLightClientKeeper.GetLatestHeader(ctx)
	if has {
		// 2. merge bnb light client
		client := &ccevtypes.Client{
			ChainId:   bnbclient.ChainId,
			ChainName: "Binance Smart Chain",
			InitialBlock: ccevtypes.TinyHeader{
				Hash:        hexutil.Encode(latestHeader.Hash),
				Number:      latestHeader.Number,
				ReceiptRoot: hexutil.Encode(latestHeader.ReceiptRoot),
			},
		}
		if err := app.CCEVkeeper.CreateClient(ctx, client); err != nil {
			return errors.New("failed to merge bnb light client")
		}
	}

	// 3. merge bnb staking plan contract
	return app.CCEVkeeper.UploadContract(
		ctx,
		bnbclient.ChainId,
		bnbclient.StakePlanHubAddress,
		bnbclient.EventName,
		bnblightclienttypes.StakePlanHubContractABIJSON,
	)
}
