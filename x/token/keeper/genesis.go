package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Lorenzo-Protocol/lorenzo/x/token/types"
)

// InitGenesis initializes the token module's state from a given genesis state
func (k Keeper) InitGenesis(ctx sdk.Context, gs *types.GenesisState) {
	k.SetParams(ctx, gs.Params)

	for _, pair := range gs.TokenPairs {
		id := pair.GetID()
		k.SetTokenPair(ctx, pair)
		k.SetTokenPairIdByDenom(ctx, pair.Denom, id)
		k.SetTokenPairIdByERC20(ctx, pair.GetERC20ContractAddress(), id)
	}

	k.accountKeeper.GetModuleAccount(ctx, types.ModuleName)
}

// ExportGenesis exports the genesis state
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return &types.GenesisState{
		Params:     k.GetParams(ctx),
		TokenPairs: k.GetTokenPairs(ctx),
	}
}
