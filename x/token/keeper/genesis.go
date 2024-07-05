package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Lorenzo-Protocol/lorenzo/x/token/types"
)

// ImportGenesis initializes the token module's state from a given genesis state
func (k Keeper) ImportGenesis(ctx sdk.Context, gs *types.GenesisState) {
	if addr := k.accountKeeper.GetModuleAddress(types.ModuleName); addr == nil {
		panic("the token module account has not been set")
	}

	k.SetParams(ctx, gs.Params)

	for _, pair := range gs.TokenPairs {
		id := pair.GetID()
		k.SetTokenPair(ctx, pair)
		k.SetTokenPairIdByDenom(ctx, pair.Denom, id)
		k.SetTokenPairIdByERC20(ctx, pair.GetERC20ContractAddress(), id)
	}
}

// ExportGenesis exports the genesis state
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return &types.GenesisState{
		Params:     k.GetParams(ctx),
		TokenPairs: k.GetTokenPairs(ctx),
	}
}
