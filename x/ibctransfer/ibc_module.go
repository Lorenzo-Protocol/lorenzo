package transfer

import (
	ibctransfer "github.com/cosmos/ibc-go/v7/modules/apps/transfer"
	porttypes "github.com/cosmos/ibc-go/v7/modules/core/05-port/types"

	"github.com/Lorenzo-Protocol/lorenzo/x/ibctransfer/keeper"
)

var _ porttypes.IBCModule = IBCModule{}

// IBCModule defines the IBC module that wraps the ibc transfer module
type IBCModule struct {
	*ibctransfer.IBCModule
}

// NewIBCModule creates a new IBCModule given the keeper
func NewIBCModule(k *keeper.Keeper) IBCModule {
	transferModule := ibctransfer.NewIBCModule(*k.Keeper)
	return IBCModule{
		IBCModule: &transferModule,
	}
}
