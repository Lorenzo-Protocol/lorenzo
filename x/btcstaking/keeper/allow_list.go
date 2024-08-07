package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

// Authorized checks if the given address is authorized.
//
// ctx: the SDK context.
// address: the address to check.
// bool: true if the address is authorized, false otherwise.
func (k Keeper) Authorized(ctx sdk.Context, address sdk.AccAddress) bool {
	params := k.GetParams(ctx)
	for _, addr := range params.MinterAllowList {
		if sdk.MustAccAddressFromBech32(addr).Equals(address) {
			return true
		}
	}
	return false
}
