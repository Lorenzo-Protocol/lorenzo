package keeper

import (
	"github.com/Lorenzo-Protocol/lorenzo/x/btcstaking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetBTCReceivingAddr sets the x/btcstaking module parameters.
func (k Keeper) SetBTCReceivingAddr(ctx sdk.Context, p string) error {
	store := ctx.KVStore(k.storeKey)
	bz := []byte(p)
	store.Set(types.BTCReceivingAddrKey, bz)
	return nil
}

// GetBTCReceivingAddr returns the current x/btcstaking module parameters.
func (k Keeper) GetBTCReceivingAddr(ctx sdk.Context) (p string) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.BTCReceivingAddrKey)
	if bz == nil {
		return p
	}
	p = string(bz)
	return p
}
