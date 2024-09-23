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

	// MintYatSuccess defines mint yat success
	MintYatSuccess = "success"
	// MintYatFailed defines mint yat failed
	MintYatFailed = "failed"
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

// KeyxBTCStakingRecord returns the key for the BTCB staking record
func KeyxBTCStakingRecord(chainID uint32, contract []byte, stakingIdx uint64) []byte {
	key := append([]byte{}, BTCBStakingRecordKey...)

	chainIDBz := sdk.Uint64ToBigEndian(uint64(chainID))
	key = append(key, chainIDBz...)
	key = append(key, Delimiter...)

	stakingIdxBz := sdk.Uint64ToBigEndian(stakingIdx)
	key = append(key, stakingIdxBz...)
	key = append(key, Delimiter...)

	key = append(key, contract...)
	return key
}
