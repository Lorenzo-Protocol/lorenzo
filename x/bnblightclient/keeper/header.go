package keeper

import (
	"golang.org/x/exp/slices"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Lorenzo-Protocol/lorenzo/v2/x/bnblightclient/types"
)

// UploadHeaders adds a batch of headers to the bnb light client chain
func (k Keeper) UploadHeaders(ctx sdk.Context, headers []*types.Header) error {
	if len(headers) == 0 {
		return errorsmod.Wrap(types.ErrInvalidHeader, "header is empty")
	}

	slices.SortFunc(headers, func(a, b *types.Header) bool {
		return a.Number < b.Number
	})

	vHeader := headers
	latestedHeader, exist := k.GetLatestHeader(ctx)
	if exist {
		vHeader = append([]*types.Header{latestedHeader}, headers...)
	}

	// verify headers
	if err := types.VerifyHeaders(vHeader); err != nil {
		return err
	}

	for _, header := range headers {
		k.setHeader(ctx, header)
	}
	k.setLatestNumber(ctx, headers[len(headers)-1].Number)
	return nil
}

// UpdateHeader updates the header in the Keeper.
//
// Parameters:
// - ctx: the context object.
// - header: the header to be updated.
// - deleteSubsequentHeaders: whether to delete subsequent headers.
//
// Returns:
// - error: an error if the header update fails.
func (k Keeper) UpdateHeader(ctx sdk.Context, header *types.Header, deleteSubsequentHeaders bool) error {
	if !k.HasHeader(ctx, header.Number) {
		return errorsmod.Wrapf(types.ErrHeaderNotFound, "header %d not found, cannot update", header.Number)
	}

	vHeader := []*types.Header{header}
	preHeader, exist := k.GetHeader(ctx, header.Number-1)
	if exist {
		vHeader = []*types.Header{preHeader, header}
	}

	if err := types.VerifyHeaders(vHeader); err != nil {
		return err
	}

	// delete subsequent headers if deleteSubsequentHeaders is true
	if deleteSubsequentHeaders {
		latestNumber := k.GetLatestNumber(ctx)
		for i := header.Number + 1; i < latestNumber; i++ {
			k.deleteHeader(ctx, i)
		}
		k.setLatestNumber(ctx, header.Number)
	}

	k.setHeader(ctx, header)
	return nil
}

// GetLatestHeader retrieves the latest header from the store.
//
// Parameters:
// - ctx: the context object
//
// Returns:
// - types.Header: the latest header
// - bool: true if the header was found, false otherwise
func (k Keeper) GetLatestHeader(ctx sdk.Context) (*types.Header, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyLatestHeaderNumber())
	if bz == nil {
		return nil, false
	}

	number := sdk.BigEndianToUint64(bz)
	return k.GetHeader(ctx, number)
}

// GetLatestNumber retrieves the latest number from the store.
//
// Parameters:
// - ctx: the context object
//
// Returns:
// - uint64: the latest number
func (k Keeper) GetLatestNumber(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyLatestHeaderNumber())
	if bz == nil {
		return 0
	}
	return sdk.BigEndianToUint64(bz)
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
func (k Keeper) GetHeader(ctx sdk.Context, number uint64) (*types.Header, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyHeader(number))
	if bz == nil {
		return nil, false
	}
	var header types.Header
	k.cdc.MustUnmarshal(bz, &header)
	return &header, true
}

// GetAllHeaders retrieves all headers from the store.
//
// Parameters:
// - ctx: the context object
//
// Returns:
// - headers: a slice of Header objects
func (k Keeper) GetAllHeaders(ctx sdk.Context) (headers []*types.Header) {
	store := ctx.KVStore(k.storeKey)

	it := sdk.KVStorePrefixIterator(store, types.KeyPrefixHeader)
	defer it.Close() //nolint:errcheck

	for ; it.Valid(); it.Next() {
		var header types.Header
		k.cdc.MustUnmarshal(it.Value(), &header)
		headers = append(headers, &header)
	}
	return
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
func (k Keeper) GetHeaderByHash(ctx sdk.Context, hash []byte) (*types.Header, bool) {
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
func (k Keeper) HasHeader(ctx sdk.Context, number uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.KeyHeader(number))
}

func (k Keeper) setHeader(ctx sdk.Context, header *types.Header) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(header)
	store.Set(types.KeyHeader(header.Number), bz)

	numberBz := sdk.Uint64ToBigEndian(header.Number)
	store.Set(types.KeyHeaderHash(header.Hash), numberBz)
}

func (k Keeper) deleteHeader(ctx sdk.Context, number uint64) {
	header, exist := k.GetHeader(ctx, number)
	if !exist {
		return
	}

	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyHeader(number))

	numberBz := sdk.Uint64ToBigEndian(header.Number)
	store.Set(types.KeyHeaderHash(header.Hash), numberBz)
}

func (k Keeper) setLatestNumber(ctx sdk.Context, number uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := sdk.Uint64ToBigEndian(number)
	store.Set(types.KeyLatestHeaderNumber(), bz)
}
