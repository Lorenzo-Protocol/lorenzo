package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

const (
	// ModuleName is the name of the module
	ModuleName = "ccev"

	// StoreKey is the string store representation
	StoreKey = ModuleName
)

var (
	// ParamsKey defines the key to store the Params object
	ParamsKey = []byte{0x01}
	// KeyPrefixClient  defines the prefix to retrieve the client
	KeyPrefixClient = []byte{0x02}
	// KeyPrefixHeader defines the prefix to retrieve all headers
	KeyPrefixHeader = []byte{0x03}
	// KeyPrefixLatestNumber defines the prefix to retrieve the latest header number
	KeyPrefixLatestNumber = []byte{0x04}
	// KeyPrefixHeadHash defines the prefix to retrieve the head hash
	KeyPrefixHeadHash = []byte{0x05}
	// KeyPrefixEvent defines the prefix to retrieve the event
	KeyPrefixEvent = []byte{0x06}
	// KeyPrefixCrossChainContract defines the prefix to retrieve the cross chain contract
	KeyPrefixCrossChainContract = []byte{0x07}
)

// KeyClient returns the key for a client
func KeyClient(chainID uint32) []byte {
	bz := sdk.Uint64ToBigEndian(uint64(chainID))
	return append(KeyPrefixClient, bz...)
}

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

// KeyEvent returns the key for the event
func KeyEvent(contract []byte, identify string) []byte {
	temp := append(contract, []byte(identify)...)
	return append(KeyPrefixHeadHash, temp...)
}

// KeyCrossChainContract returns the key for the CrossChainContract
func KeyCrossChainContract(address common.Address) []byte {
	return append(KeyPrefixCrossChainContract, address.Bytes()...)
}
