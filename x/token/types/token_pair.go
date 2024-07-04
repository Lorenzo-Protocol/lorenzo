package types

import (
	"github.com/ethereum/go-ethereum/common"
)

func (tp *TokenPair) IsNativeCoin() bool {
	return tp.Ownership == OWNER_MODULE
}

func (tp *TokenPair) IsNativeERC20() bool {
	return tp.Ownership == OWNER_EXTERNAL
}

func (tp *TokenPair) GetERC20ContractAddress() common.Address {
	panic("implement me")
}
