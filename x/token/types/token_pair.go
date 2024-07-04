package types

import (
	"github.com/cometbft/cometbft/crypto/tmhash"
	"github.com/ethereum/go-ethereum/common"
)

// NewTokenPair creates a new TokenPair
func NewTokenPair(erc20Addr common.Address, denom string, ownership Ownership) TokenPair {
	return TokenPair{
		ContractAddress: erc20Addr.String(),
		Denom:           denom,
		Enabled:         true,
		Ownership:       ownership,
	}
}

func (tp *TokenPair) GetID() []byte {
	id := tp.ContractAddress + "|" + tp.Denom
	return tmhash.Sum([]byte(id))
}

func (tp *TokenPair) IsNativeCoin() bool {
	return tp.Ownership == OWNER_MODULE
}

func (tp *TokenPair) IsNativeERC20() bool {
	return tp.Ownership == OWNER_EXTERNAL
}

func (tp *TokenPair) GetERC20ContractAddress() common.Address {
	panic("implement me")
}
