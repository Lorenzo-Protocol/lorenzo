package keeper

import (
	"github.com/Lorenzo-Protocol/lorenzo/x/btcstaking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetBTCReceivingAddr sets the x/btcstaking module parameters.
func (k Keeper) SetParams(ctx sdk.Context, p *types.Params) error {
	store := ctx.KVStore(k.storeKey)
	paramsBytes := k.cdc.MustMarshal(p)
	store.Set(types.ParamsKey, paramsBytes)
	return nil
}

// GetBTCReceivingAddr returns the current x/btcstaking module parameters.
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
