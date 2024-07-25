package keeper

import (
	"slices"

	"github.com/Lorenzo-Protocol/lorenzo/v2/x/bnblightclient/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	if err := genState.Validate(); err != nil {
		panic(err)
	}

	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}

	slices.SortFunc(genState.Headers, func(a, b *types.Header) int {
		return int(a.Number - b.Number)
	})
	for _, header := range genState.Headers {
		k.setHeader(ctx, header)
	}
	k.setLatestNumber(ctx, genState.Headers[len(genState.Headers)-1].Number)

	for _, record := range genState.Events {
		k.setEvmEvent(ctx, record)
	}
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return &types.GenesisState{
		Params:  k.GetParams(ctx),
		Headers: k.GetAllHeaders(ctx),
		Events:  k.GetAllEvmEvents(ctx),
	}
}
