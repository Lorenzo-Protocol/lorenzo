package keeper

import (
	"github.com/Lorenzo-Protocol/lorenzo/v3/x/plan/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetParams sets the parameters for the given context.
//
// ctx - Context object.
// p - Params object to be set.
// error - Returns an error if validation fails.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) error {
	if err := params.Validate(); err != nil {
		return err
	}
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&params)
	store.Set(types.ParamsKey, bz)
	return nil
}

// GetParams retrieves the parameters of the Keeper.
//
// ctx sdk.Context - Context
// types.Params - Parameters
// returns p types.Params - Parameters
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ParamsKey)
	if bz == nil {
		return params
	}
	k.cdc.MustUnmarshal(bz, &params)
	return params
}
