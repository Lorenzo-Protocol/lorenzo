package types

import (
	"fmt"

	"github.com/cometbft/cometbft/crypto/tmhash"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

// NewTokenPair creates a new TokenPair
func NewTokenPair(erc20Addr common.Address, denom string, source Source) TokenPair {
	return TokenPair{
		ContractAddress: erc20Addr.String(),
		Denom:           denom,
		Enabled:         true,
		Source:          source,
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

	if tp.Source != OWNER_MODULE && tp.Source != OWNER_CONTRACT {
		return fmt.Errorf("invalid token source: %s", tp.Source)
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
	return tp.Source == OWNER_MODULE
}

// IsNativeERC20 checks if the token is erc20 contract originated
func (tp *TokenPair) IsNativeERC20() bool {
	return tp.Source == OWNER_CONTRACT
}

// GetERC20ContractAddress return the common address
func (tp *TokenPair) GetERC20ContractAddress() common.Address {
	return common.HexToAddress(tp.ContractAddress)
}
