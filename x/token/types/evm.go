package types

import "fmt"

// ERC20Data represents the ERC20 token details used to map
// the token to a Cosmos Coin
type ERC20Data struct {
	Address  string
	Name     string
	Symbol   string
	Decimals uint8
}

// ERC20StringResponse defines the string value from the call response
type ERC20StringResponse struct {
	Value string
}

// ERC20Uint8Response defines the uint8 value from the call response
type ERC20Uint8Response struct {
	Value uint8
}

// ERC20BoolResponse defines the bool value from the call response
type ERC20BoolResponse struct {
	Value bool
}

// NewERC20Data creates a new ERC20Data instance
func NewERC20Data(address, name, symbol string, decimals uint8) ERC20Data {
	return ERC20Data{
		Address:  address,
		Name:     name,
		Symbol:   symbol,
		Decimals: decimals,
	}
}

// BaseDenom constructs coin base denom for an erc20 contract.
func (d ERC20Data) BaseDenom() string {
	return fmt.Sprintf("%s/%s", DenomPrefix, d.Address)
}

// Description returns coin description for an erc20 contract.
func (d ERC20Data) Description() string {
	return fmt.Sprintf("erc20/%s mapping to cosmos sdk coin", d.Address)
}

// SanitizedName sanitizes the erc20 name to be an acceptable denom.
func (d ERC20Data) SanitizedName() string {
	return SanitizeERC20Name(d.Name)
}
