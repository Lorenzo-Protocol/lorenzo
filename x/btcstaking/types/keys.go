package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "btcstaking"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_btcstaking"

	// denom used by this module
	NativeTokenDenom = "stBTC"
)

var (
	Delimiter            = []byte{0x00}
	ParamsKey            = []byte{0x01} // key prefix for the BTC receiving address
	BTCStakingRecordKey  = []byte{0x02} // key prefix for the BTC staking record
	BTCBStakingRecordKey = []byte{0x03} // key prefix for the BTCB staking record
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// KeyBTCBStakingRecord returns the key for the BTCB staking record
func KeyBTCBStakingRecord(txHash []byte, eventIdx uint64) []byte {
	key := append([]byte{}, BTCBStakingRecordKey...)
	key = append(key, txHash...)
	key = append(key, Delimiter...)

	chainIDBz := sdk.Uint64ToBigEndian(eventIdx)
	key = append(key, chainIDBz...)
	return key
}
