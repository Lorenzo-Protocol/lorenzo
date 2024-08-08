package keeper

import (
	"github.com/Lorenzo-Protocol/lorenzo/v2/x/bnblightclient/types"
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
	logger := k.Logger(ctx)
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefixHeader)
	defer func() {
		if err := iterator.Close(); err != nil {
			logger.Error("close iterator error", "err", err)
		}
	}()

	logger.Info("prune headers", "pruneEndNumber", pruneEndNumber)
	for ; iterator.Valid(); iterator.Next() {
		iterKey := iterator.Key()
		number := sdk.BigEndianToUint64(iterKey[1:])
		if number > pruneEndNumber {
			header, exist := k.GetHeader(ctx, number)
			if !exist {
				continue
			}

			// delete header
			store.Delete(iterKey)
			// delete header hash
			store.Delete(types.KeyHeaderHash(header.Hash))
		}
	}
}
