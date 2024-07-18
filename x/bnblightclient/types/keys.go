package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
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
	// KeyPrefixLatestedNumber defines the prefix to retrieve the latest header number
	KeyPrefixLatestedNumber = []byte{0x03}

	// KeyPrefixHeadHash defines the prefix to retrieve the head hash
	KeyPrefixHeadHash = []byte{0x04}
	// KeyPrefixTxHash defines the prefix to retrieve the tx hash
	KeyPrefixTxHash = []byte{0x05}
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

// KeyLatestedHeaderNumber returns the key for the latest header number
func KeyLatestedHeaderNumber() []byte {
	return KeyPrefixLatestedNumber
}

// KeyTxHash returns the key for the tx hash
func KeyTxHash(blockNumber uint64, txHash common.Hash) []byte {
	key := make([]byte, 0)
	copy(key, KeyPrefixTxHash)

	bumberBz := sdk.Uint64ToBigEndian(blockNumber)
	key = append(key, bumberBz...)

	txHashBz := txHash.Bytes()
	key = append(key, txHashBz...)
	return key
}
