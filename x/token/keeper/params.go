package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Lorenzo-Protocol/lorenzo/v2/x/token/types"
)

// SetParams sets the paras for the token module.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&params)
	store.Set(types.KeyPrefixParams, bz)
}

// GetParams returns the params of the token module.
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	var params types.Params
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyPrefixParams)
	if bz == nil {
		return params
	}
	k.cdc.MustUnmarshal(bz, &params)
	return params
}

// IsConvertEnabled returns true if the ERC20 module is enabled.
func (k Keeper) IsConvertEnabled(ctx sdk.Context) bool {
	return k.GetParams(ctx).EnableConversion
}

// IsEVMHookEnabled returns true if the EVM hook is enabled.
func (k Keeper) IsEVMHookEnabled(ctx sdk.Context) bool {
	return k.GetParams(ctx).EnableEVMHook
}
