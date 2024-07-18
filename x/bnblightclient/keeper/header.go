package keeper

import (
	"slices"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Lorenzo-Protocol/lorenzo/x/bnblightclient/types"
)

// UploadHeaders adds a batch of headers to the bnb light client chain
func(k Keeper) UploadHeaders(ctx sdk.Context, headers []*types.Header) error {
	slices.SortFunc(headers, func(a, b *types.Header) int {
		return int(a.Number - b.Number)
	})

	vHeader := headers
	latestedHeader, exist := k.GetLatestedHeader(ctx)
	if exist {
		vHeader = append([]*types.Header{latestedHeader}, headers...)
	}
	
	// verify headers
	if err := types.VeryHeaders(vHeader); err != nil { 
		return err 
	}

	for _, header := range headers {
		k.setHeader(ctx, header)
	}
	k.setLatestedHeaderNumber(ctx, headers[len(headers) - 1].Number)
	return nil
}

// UpdateHeader updates the header in the Keeper.
//
// Parameters:
// - ctx: the context object.
// - header: the header to be updated.
//
// Returns:
// - error: an error if the header update fails.
func(k Keeper) UpdateHeader(ctx sdk.Context, header *types.Header) error {
	if err := types.VeryHeaders([]*types.Header{header}); err != nil {
		return err
	}

	if !k.HasHeader(ctx, header.Number) {
		return errorsmod.Wrapf(types.ErrHeaderNotFound, "header %d not found, cannot update", header.Number)
	}

	k.setHeader(ctx, header)
	return nil
}


// GetLatestedHeader retrieves the latested header from the store.
//
// Parameters:
// - ctx: the context object
//
// Returns:
// - types.Header: the latested header
// - bool: true if the header was found, false otherwise
func(k Keeper) GetLatestedHeader(ctx sdk.Context) (*types.Header, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyLatestedHeaderNumber())
	if bz == nil {
		return nil, false
	}

	number := sdk.BigEndianToUint64(bz)
	return k.GetHeader(ctx, number)
}

// GetHeader retrieves the header for a specific number from the store.
//
// Parameters:
// - ctx: the context object
// - number: the number of the header to retrieve
//
// Returns:
// - types.Header: the header object
// - bool: true if the header was found, false otherwise
func(k Keeper) GetHeader(ctx sdk.Context, number uint64) (*types.Header, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyHeader(number))
	if bz == nil {
		return nil, false
	}
	var header types.Header
	k.cdc.MustUnmarshal(bz, &header)
	return &header, true
}

// GetHeaderByHash retrieves a header from the store based on its hash.
//
// Parameters:
// - ctx: the context object
// - hash: the hash of the header to retrieve
//
// Returns:
// - *types.Header: the header object, or nil if not found
// - bool: true if the header was found, false otherwise
func(k Keeper) GetHeaderByHash(ctx sdk.Context, hash []byte) (*types.Header, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyHeaderHash(hash))
	if bz == nil {
		return nil, false
	}

	number := sdk.BigEndianToUint64(bz)
	return k.GetHeader(ctx, number)
}

// HasHeader checks if a header with the given number exists in the store.
//
// Parameters:
// - ctx: the context object
// - number: the number of the header to check
// Return type: bool
func(k Keeper) HasHeader(ctx sdk.Context, number uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.KeyHeader(number))
}

func(k Keeper) setHeader(ctx sdk.Context, header *types.Header) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(header)
	store.Set(types.KeyHeader(header.Number), bz)

	numberBz := sdk.Uint64ToBigEndian(header.Number)
	store.Set(types.KeyHeaderHash(header.Hash), numberBz)
}

func(k Keeper) setLatestedHeaderNumber(ctx sdk.Context, number uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := sdk.Uint64ToBigEndian(number)
	store.Set(types.KeyLatestedHeaderNumber(), bz)
}