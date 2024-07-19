package types

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/cometbft/cometbft/crypto/tmhash"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

// Validate validates the token pair
func (tp *TokenPair) Validate() error {
	if err := sdk.ValidateDenom(tp.Denom); err != nil {
		return err
	}

	if !common.IsHexAddress(tp.ContractAddress) {
		return fmt.Errorf("invalid contract address: %s", tp.ContractAddress)
	}

	if tp.Ownership != OWNER_MODULE && tp.Ownership != OWNER_EXTERNAL {
		return fmt.Errorf("invalid ownership: %s", tp.Ownership)
	}

	return nil
}

// GetID returns token pair id
func (tp *TokenPair) GetID() []byte {
	id := tp.ContractAddress + "|" + tp.Denom
	return tmhash.Sum([]byte(id))
}

// IsNativeCoin checks if the token is sdk coin originated
func (tp *TokenPair) IsNativeCoin() bool {
	return tp.Ownership == OWNER_MODULE
}

// IsNativeERC20 checks if the token is erc20 contract originated
func (tp *TokenPair) IsNativeERC20() bool {
	return tp.Ownership == OWNER_EXTERNAL
}

// GetERC20ContractAddress return the common address
func (tp *TokenPair) GetERC20ContractAddress() common.Address {
	return common.HexToAddress(tp.ContractAddress)
}
