package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Lorenzo-Protocol/lorenzo/x/agent/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	var maxNumber uint64
	for _, agent := range genState.Agents {
		k.setAgent(ctx, agent)

		if agent.Id > maxNumber {
			maxNumber = agent.Id
		}
	}
	if len(genState.Admin) > 0 {
		admin := sdk.MustAccAddressFromBech32(genState.Admin)
		k.setAdmin(ctx, admin)
	}
	k.setNextNumber(ctx, maxNumber+1)
}

// ExportGenesis returns the capability module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return &types.GenesisState{
		Admin:  k.GetAdmin(ctx).String(),
		Agents: k.getAgents(ctx),
	}
}
