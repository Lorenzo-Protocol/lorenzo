package keeper

import (
	"github.com/Lorenzo-Protocol/lorenzo/v2/x/btcstaking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetParams sets the x/btcstaking module parameters.
//
// ctx - Context object.
// p - Params object to be set.
// error - Returns an error if validation fails.
func (k Keeper) SetParams(ctx sdk.Context, p *types.Params) error {
	store := ctx.KVStore(k.storeKey)
	paramsBytes := k.cdc.MustMarshal(p)
	store.Set(types.ParamsKey, paramsBytes)
	return nil
}

// GetParams retrieves the x/btcstaking module parameters.
//
// ctx - Context object.
// returns - Params object.
func (k Keeper) GetParams(ctx sdk.Context) *types.Params {
	store := ctx.KVStore(k.storeKey)
	var params types.Params
	bz := store.Get(types.ParamsKey)
	if bz == nil {
		return &params
	}

	k.cdc.MustUnmarshal(bz, &params)
	return &params
}
