package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName is the name of the module
	ModuleName = "bnblightclient"

	// StoreKey is the string store representation
	StoreKey = ModuleName
)

var (
	// Delimiter is the delimiter for the keys
	Delimiter = []byte{0x00}
	// ParamsKey defines the key to store the Params object
	ParamsKey = []byte{0x01}
	// KeyPrefixHeader defines the prefix to retrieve all headers
	KeyPrefixHeader = []byte{0x02}
	// KeyPrefixLatestNumber defines the prefix to retrieve the latest header number
	KeyPrefixLatestNumber = []byte{0x03}

	// KeyPrefixHeadHash defines the prefix to retrieve the head hash
	KeyPrefixHeadHash = []byte{0x04}
	// KeyPrefixEvmEvent defines the prefix to retrieve the cross chain event
	KeyPrefixEvmEvent = []byte{0x05}
)

// KeyHeader returns the key for a header
func KeyHeader(blockNumber uint64) []byte {
	bz := sdk.Uint64ToBigEndian(blockNumber)
	return append(KeyPrefixHeader, bz...)
}

// KeyHeaderHash returns the key for the header hash
func KeyHeaderHash(hash []byte) []byte {
	return append(KeyPrefixHeadHash, hash...)
}

// KeyLatestHeaderNumber returns the key for the latest header number
func KeyLatestHeaderNumber() []byte {
	return KeyPrefixLatestNumber
}

// KeyEvmEvent returns the key for the cross chain event index
func KeyEvmEvent(chainID uint32, blockNumber uint64, contract []byte, idx uint64) []byte {
	key := append([]byte{}, KeyPrefixEvmEvent...)

	chainIDBz := sdk.Uint64ToBigEndian(uint64(chainID))
	key = append(key, chainIDBz...)
	key = append(key, Delimiter...)

	bumberBz := sdk.Uint64ToBigEndian(blockNumber)
	key = append(key, bumberBz...)
	key = append(key, Delimiter...)

	key = append(key, contract...)
	key = append(key, Delimiter...)

	idxBz := sdk.Uint64ToBigEndian(idx)
	key = append(key, idxBz...)
	return key
}

// PrefixKeyEvmEvent returns the prefix for the cross chain event index
func PrefixKeyEvmEvent(chainID uint32, blockNumber uint64) []byte {
	key := append([]byte{}, KeyPrefixEvmEvent...)

	chainIDBz := sdk.Uint64ToBigEndian(uint64(chainID))
	key = append(key, chainIDBz...)
	key = append(key, Delimiter...)

	bumberBz := sdk.Uint64ToBigEndian(blockNumber)
	key = append(key, bumberBz...)
	key = append(key, Delimiter...)
	return key
}
