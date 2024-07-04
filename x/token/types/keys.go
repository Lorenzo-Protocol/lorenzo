package types

import (
	"github.com/ethereum/go-ethereum/common"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

const (
	ModuleName = "token"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName
)

// ModuleAddress is the native module address for EVM
var ModuleAddress common.Address

func init() {
	ModuleAddress = common.BytesToAddress(authtypes.NewModuleAddress(ModuleName).Bytes())
}

const (
	prefixParams = iota + 1
	prefixTokenPair
	prefixTokenPairIdByERC20
	prefixTokenPairIdByDenom
)

var (
	KeyPrefixParams             = []byte{prefixParams}
	KeyPrefixTokenPair          = []byte{prefixTokenPair}
	KeyPrefixTokenPairIdByERC20 = []byte{prefixTokenPairIdByERC20}
	KeyPrefixTokenPairIdByDenom = []byte{prefixTokenPairIdByDenom}
)

// PrefixTokenPairStoreKey returns the key for the token pair.
// Items are stored with the following key: values
// <prefix><token_pair_id_bz> -> <token_pair>
func PrefixTokenPairStoreKey(id []byte) []byte {
	key := make([]byte, len(KeyPrefixTokenPair)+len(id))
	copy(key, KeyPrefixTokenPair)
	copy(key[len(KeyPrefixTokenPair):], id)
	return key
}

// PrefixTokenPairIdByERC20StoreKey returns the key for the token pair by ERC20 address.
// Items are stored with the following key: values
// <prefix><erc20_addr_bz> -> <token_pair_id_bz>
func PrefixTokenPairIdByERC20StoreKey(erc20Addr common.Address) []byte {
	erc20Bz := erc20Addr.Bytes()
	key := make([]byte, len(KeyPrefixTokenPairIdByERC20)+len(erc20Bz))
	copy(key, KeyPrefixTokenPairIdByERC20)
	copy(key[len(KeyPrefixTokenPairIdByERC20):], erc20Bz)
	return key
}

// PrefixTokenPairIdByDenomStoreKey returns the key for the token pair by denomination.
// Items are stored with the following key: values
// <prefix><denom> -> <token_pair_id_bz>
func PrefixTokenPairIdByDenomStoreKey(denom string) []byte {
	key := make([]byte, len(KeyPrefixTokenPairIdByDenom)+len(denom))
	copy(key, KeyPrefixTokenPairIdByDenom)
	copy(key[len(KeyPrefixTokenPairIdByDenom):], denom)
	return key
}
