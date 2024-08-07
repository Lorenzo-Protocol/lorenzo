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