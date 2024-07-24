package keeper

import (
	"github.com/Lorenzo-Protocol/lorenzo/x/bnblightclient/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) prune(ctx sdk.Context) {
	// get the latested header
	latestedNumber := k.GetLatestNumber(ctx)
	params := k.GetParams(ctx)
	pruneEndNumber := int64(latestedNumber - params.RetainedBlocks)
	if pruneEndNumber <= 0 {
		return
	}
	k.pruneHeaders(ctx, uint64(pruneEndNumber))
}

func (k Keeper) pruneHeaders(ctx sdk.Context, pruneEndNumber uint64) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStoreReversePrefixIterator(store, types.KeyPrefixHeader)
	defer func() {
		_ = iterator.Close()
	}()

	for ; iterator.Valid(); iterator.Next() {
		iterKey := iterator.Key()
		number := sdk.BigEndianToUint64(iterKey[1:])
		if number <= pruneEndNumber {
			header, exist := k.GetHeader(ctx, number)
			if !exist {
				continue
			}

			// delete header
			store.Delete(iterKey)
			// delete header hash
			store.Delete(types.KeyHeaderHash(header.Hash))

			// delete event record
			prefix := append(types.KeyPrefixEvmEvent, sdk.Uint64ToBigEndian(number)...)
			iterator2 := sdk.KVStoreReversePrefixIterator(store, prefix)
			defer func() {
				_ = iterator2.Close()
			}()

			for ; iterator2.Valid(); iterator2.Next() {
				iterKey2 := iterator2.Key()
				store.Delete(iterKey2)
			}
		}
	}
}
