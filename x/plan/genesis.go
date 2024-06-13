package plan

import (
	"github.com/Lorenzo-Protocol/lorenzo/x/plan/keeper"
	"github.com/Lorenzo-Protocol/lorenzo/x/plan/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis new fee genesis
func InitGenesis(ctx sdk.Context, k *keeper.Keeper, genState types.GenesisState) {
	if err := genState.Validate(); err != nil {
		panic(err)
	}

	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper *keeper.Keeper) *types.GenesisState {
	params := keeper.GetParams(ctx)
	nextNumber := keeper.GetNextNumber(ctx)
	plans := keeper.GetPlans(ctx)
	return types.NewGenesisState(params, nextNumber, plans)
}
