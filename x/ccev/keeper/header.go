package keeper

import (
	"golang.org/x/exp/slices"

	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Lorenzo-Protocol/lorenzo/v3/x/ccev/types"
)

// CreateClient creates a new client in store
//
// Args:
// - client: the new client to be created
//
// Returns:
// - error: if the client already exists, or other error occurs
func (k Keeper) CreateClient(ctx sdk.Context,client *types.Client) error {
	if k.hasClient(ctx, client.ChainId) {
		return errorsmod.Wrapf(types.ErrDuplicateClient, "client %d already exists", client.ChainId)
	}
	k.setClient(ctx, client)
	return nil
}

// UploadHeaders adds a batch of headers to the bnb light client chain
func (k Keeper) UploadHeaders(ctx sdk.Context, chainID uint32, headers []*types.TinyHeader) error {
	if len(headers) == 0 {
		return errorsmod.Wrap(types.ErrInvalidHeader, "header is empty")
	}

	slices.SortFunc(headers, func(a, b *types.TinyHeader) bool {
		return a.Number < b.Number
	})

	for _, header := range headers {
		k.setHeader(ctx, chainID, header)
	}
	k.setLatestNumber(ctx, chainID, headers[len(headers)-1].Number)
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
func (k Keeper) UpdateHeader(ctx sdk.Context, header *types.TinyHeader, deleteSubsequentHeaders bool) error {
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
func (k Keeper) GetLatestHeader(ctx sdk.Context, chainID uint32) (*types.TinyHeader, bool) {
	number := k.GetLatestNumber(ctx, chainID)
	return k.GetHeader(ctx, chainID, number)
}

// GetLatestNumber retrieves the latest number from the store.
//
// Parameters:
// - ctx: the context object
//
// Returns:
// - uint64: the latest number
func (k Keeper) GetLatestNumber(ctx sdk.Context, chainID uint32) uint64 {
	store := k.clientStore(ctx, chainID)
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
func (k Keeper) GetHeader(ctx sdk.Context, chainID uint32, number uint64) (*types.TinyHeader, bool) {
	store := k.clientStore(ctx, chainID)
	bz := store.Get(types.KeyHeader(number))
	if bz == nil {
		return nil, false
	}
	var header types.TinyHeader
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
// func (k Keeper) GetAllHeaders(ctx sdk.Context) (headers []*types.TinyHeader) {
// 	store := ctx.KVStore(k.storeKey)

// 	it := sdk.KVStorePrefixIterator(store, types.KeyPrefixHeader)
// 	defer it.Close() //nolint:errcheck

// 	for ; it.Valid(); it.Next() {
// 		var header types.TinyHeader
// 		k.cdc.MustUnmarshal(it.Value(), &header)
// 		headers = append(headers, &header)
// 	}
// 	return
// }

// GetHeaderByHash retrieves a header from the store based on its hash.
//
// Parameters:
// - ctx: the context object
// - hash: the hash of the header to retrieve
//
// Returns:
// - *types.Header: the header object, or nil if not found
// - bool: true if the header was found, false otherwise
func (k Keeper) GetHeaderByHash(ctx sdk.Context, chainID uint32, hash []byte) (*types.TinyHeader, bool) {
	store := k.clientStore(ctx, chainID)
	bz := store.Get(types.KeyHeaderHash(hash))
	if bz == nil {
		return nil, false
	}

	number := sdk.BigEndianToUint64(bz)
	return k.GetHeader(ctx, chainID, number)
}

func (k Keeper) setHeader(ctx sdk.Context, chainID uint32, header *types.TinyHeader) {
	store := k.clientStore(ctx, chainID)
	bz := k.cdc.MustMarshal(header)
	store.Set(types.KeyHeader(header.Number), bz)

	numberBz := sdk.Uint64ToBigEndian(header.Number)
	store.Set(types.KeyHeaderHash(header.Hash), numberBz)
}

func (k Keeper) setLatestNumber(ctx sdk.Context, chainID uint32, number uint64) {
	store := k.clientStore(ctx, chainID)
	bz := sdk.Uint64ToBigEndian(number)
	store.Set(types.KeyLatestHeaderNumber(), bz)
}

func (k Keeper) clientStore(ctx sdk.Context, chainID uint32) prefix.Store {
	store := ctx.KVStore(k.storeKey)
	return prefix.NewStore(store, sdk.Uint64ToBigEndian(uint64(chainID)))
}

func (k Keeper) setClient(ctx sdk.Context, client *types.Client) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(client)
	store.Set(types.KeyClient(client.ChainId), bz)

	k.setHeader(ctx, client.ChainId, &client.InitialBlock)
}

func (k Keeper) hasClient(ctx sdk.Context,chainID uint32) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.KeyClient(chainID))
}

func (k Keeper) getClient(ctx sdk.Context, chainID uint32) *types.Client {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyClient(chainID))
	if bz == nil {
		return nil
	}
	var client types.Client
	k.cdc.MustUnmarshal(bz, &client)
	return &client
}
