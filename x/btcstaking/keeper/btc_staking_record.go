package keeper

import (
	"github.com/Lorenzo-Protocol/lorenzo/x/btcstaking/types"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) addBTCStakingRecord(ctx sdk.Context, btcStk *types.BTCStakingRecord) error {
	store := k.btcStakingRecordStore(ctx)
	btcStkKey := btcStk.TxHash
	store.Set(btcStkKey, k.cdc.MustMarshal(btcStk))
	return nil
}

func (k Keeper) getBTCStakingRecord(ctx sdk.Context, txHash chainhash.Hash) *types.BTCStakingRecord {
	store := k.btcStakingRecordStore(ctx)
	btcStakingRecordBytes := store.Get(txHash[:])
	if len(btcStakingRecordBytes) == 0 {
		return nil
	}
	var btcStkRecord types.BTCStakingRecord
	k.cdc.MustUnmarshal(btcStakingRecordBytes, &btcStkRecord)
	return &btcStkRecord
}

func (k Keeper) btcStakingRecordStore(ctx sdk.Context) prefix.Store {
	kvStore := ctx.KVStore(k.storeKey)
	return prefix.NewStore(kvStore, types.BTCStakingRecordKey)
}
