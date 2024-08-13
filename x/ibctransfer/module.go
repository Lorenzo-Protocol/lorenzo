package transfer

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/module"
	ibctransferkeeper "github.com/cosmos/ibc-go/v7/modules/apps/transfer/keeper"

	"github.com/Lorenzo-Protocol/lorenzo/v3/x/ibctransfer/keeper"
	ibctransfer "github.com/cosmos/ibc-go/v7/modules/apps/transfer"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
)

var (
	_ module.AppModule           = AppModule{}
	_ module.HasGenesis          = AppModule{}
	_ module.HasConsensusVersion = AppModule{}
	_ module.HasServices         = AppModule{}
)

// AppModuleBasic embeds the IBC transfer AppModuleBasic
type AppModuleBasic struct {
	*ibctransfer.AppModuleBasic
}

// AppModule is a wrapper around the ibc transfer module
type AppModule struct {
	*ibctransfer.AppModule
	keeper *keeper.Keeper
}

// NewAppModule creates a new AppModule
func NewAppModule(k *keeper.Keeper) AppModule {
	ics20AppModule := ibctransfer.NewAppModule(*k.Keeper)

	return AppModule{
		AppModule: &ics20AppModule,
		keeper:    k,
	}
}

func (am AppModule) RegisterServices(cfg module.Configurator) {
	ibctransfertypes.RegisterMsgServer(cfg.MsgServer(), am.keeper)
	ibctransfertypes.RegisterQueryServer(cfg.QueryServer(), am.keeper)

	m := ibctransferkeeper.NewMigrator(*am.keeper.Keeper)
	if err := cfg.RegisterMigration(ibctransfertypes.ModuleName, 1, m.MigrateTraces); err != nil {
		panic(fmt.Sprintf("failed to migrate transfer app from version 1 to 2: %v", err))
	}

	if err := cfg.RegisterMigration(ibctransfertypes.ModuleName, 2, m.MigrateTotalEscrowForDenom); err != nil {
		panic(fmt.Sprintf("failed to migrate transfer app from version 2 to 3: %v", err))
	}
}
