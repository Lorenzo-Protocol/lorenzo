package keeper

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Lorenzo-Protocol/lorenzo/v3/x/token/types"
)

// RemoveTokenPair removes a token pair and its id mappings.
func (k Keeper) RemoveTokenPair(ctx sdk.Context, tokenPair types.TokenPair) {
	id := tokenPair.GetID()
	k.DeleteTokenPair(ctx, id)
	k.DeleteTokenPairIdByERC20(ctx, tokenPair.GetERC20ContractAddress())
	k.DeleteTokenPairIdByDenom(ctx, tokenPair.Denom)
}

// IsRegisteredByDenom checks if a token pair is registered by coin denom.
func (k Keeper) IsRegisteredByDenom(ctx sdk.Context, denom string) bool {
	store := ctx.KVStore(k.storeKey)
	key := types.PrefixTokenPairIdByDenomStoreKey(denom)
	return store.Has(key)
}

// IsRegisteredByERC20 checks if a token pair is registered by erc20 address
func (k Keeper) IsRegisteredByERC20(ctx sdk.Context, erc20Addr common.Address) bool {
	store := ctx.KVStore(k.storeKey)
	key := types.PrefixTokenPairIdByERC20StoreKey(erc20Addr)
	return store.Has(key)
}

// GetTokenPairs returns all token pairs.
func (k Keeper) GetTokenPairs(ctx sdk.Context) []types.TokenPair {
	var tokenPairs []types.TokenPair

	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefixTokenPair)
	defer iterator.Close() //nolint:errcheck

	for ; iterator.Valid(); iterator.Next() {
		var tokenPair types.TokenPair
		k.cdc.MustUnmarshal(iterator.Value(), &tokenPair)
		tokenPairs = append(tokenPairs, tokenPair)
	}

	return tokenPairs
}

// GetTokenPairId returns the token pair id by either denom or erc20 address.
func (k Keeper) GetTokenPairId(ctx sdk.Context, token string) []byte {
	if common.IsHexAddress(token) {
		addr := common.HexToAddress(token)
		return k.GetTokenPairIdByERC20(ctx, addr)
	}
	return k.GetTokenPairIdByDenom(ctx, token)
}

// SetTokenPair sets a token pair in the store.
func (k Keeper) SetTokenPair(ctx sdk.Context, tokenPair types.TokenPair) {
	store := ctx.KVStore(k.storeKey)
	key := types.PrefixTokenPairStoreKey(tokenPair.GetID())
	bz := k.cdc.MustMarshal(&tokenPair)
	store.Set(key, bz)
}

// GetTokenPair gets a token pair by its id.
func (k Keeper) GetTokenPair(ctx sdk.Context, id []byte) (types.TokenPair, bool) {
	var pair types.TokenPair
	if id == nil {
		return pair, false
	}

	store := ctx.KVStore(k.storeKey)
	key := types.PrefixTokenPairStoreKey(id)
	bz := store.Get(key)
	if len(bz) == 0 {
		return pair, false
	}

	k.cdc.MustUnmarshal(bz, &pair)
	return pair, true
}

// DeleteTokenPair deletes a token pair by its id
func (k Keeper) DeleteTokenPair(ctx sdk.Context, id []byte) {
	store := ctx.KVStore(k.storeKey)
	keys := types.PrefixTokenPairStoreKey(id)
	store.Delete(keys)
}

// SetTokenPairIdByERC20 sets the token pair id by the ERC20 address.
func (k Keeper) SetTokenPairIdByERC20(ctx sdk.Context, erc20Addr common.Address, id []byte) {
	store := ctx.KVStore(k.storeKey)
	key := types.PrefixTokenPairIdByERC20StoreKey(erc20Addr)
	store.Set(key, id)
}

// GetTokenPairIdByERC20 gets the token pair id by the ERC20 address.
func (k Keeper) GetTokenPairIdByERC20(ctx sdk.Context, erc20Addr common.Address) []byte {
	store := ctx.KVStore(k.storeKey)
	key := types.PrefixTokenPairIdByERC20StoreKey(erc20Addr)
	return store.Get(key)
}

// DeleteTokenPairIdByERC20 deletes the token pair id by the ERC20 address.
func (k Keeper) DeleteTokenPairIdByERC20(ctx sdk.Context, erc20Addr common.Address) {
	store := ctx.KVStore(k.storeKey)
	key := types.PrefixTokenPairIdByERC20StoreKey(erc20Addr)
	store.Delete(key)
}

// SetTokenPairIdByDenom sets the token pair id by coin denom.
func (k Keeper) SetTokenPairIdByDenom(ctx sdk.Context, denom string, id []byte) {
	store := ctx.KVStore(k.storeKey)
	key := types.PrefixTokenPairIdByDenomStoreKey(denom)
	store.Set(key, id)
}

// GetTokenPairIdByDenom gets the token pair id by coin denom.
func (k Keeper) GetTokenPairIdByDenom(ctx sdk.Context, denom string) []byte {
	store := ctx.KVStore(k.storeKey)
	key := types.PrefixTokenPairIdByDenomStoreKey(denom)
	return store.Get(key)
}

// DeleteTokenPairIdByDenom deletes the token pair id by coin denom.
func (k Keeper) DeleteTokenPairIdByDenom(ctx sdk.Context, denom string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixTokenPairIdByDenom)
	store.Delete([]byte(denom))
}
