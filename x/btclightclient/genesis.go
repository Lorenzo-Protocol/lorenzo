package btclightclient

import (
	"github.com/Lorenzo-Protocol/lorenzo/x/btclightclient/keeper"
	"github.com/Lorenzo-Protocol/lorenzo/x/btclightclient/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	if err := genState.Validate(); err != nil {
		panic(err)
	}

	k.SetBaseBTCHeader(ctx, genState.BaseBtcHeader)
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	baseBTCHeader := k.GetBaseBTCHeader(ctx)
	if baseBTCHeader == nil {
		panic("A base BTC Header has not been set")
	}

	genesis.BaseBtcHeader = *baseBTCHeader
	genesis.Params = k.GetParams(ctx)

	return genesis
}
