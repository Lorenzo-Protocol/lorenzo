package keeper

import (
	"github.com/Lorenzo-Protocol/lorenzo/v2/x/plan/types"

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

	// set plan
	var maxNumber uint64
	for _, plan := range genState.Plans {
		k.setPlan(ctx, plan)
		if plan.Id > maxNumber {
			maxNumber = plan.Id
		}
	}

	k.setNextNumber(ctx, maxNumber+1)
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	params := k.GetParams(ctx)
	nextNumber := k.GetNextNumber(ctx)
	plans := k.GetPlans(ctx)
	return types.NewGenesisState(params, nextNumber, plans)
}
