package types

import (
	"strings"

	tokentypes "github.com/Lorenzo-Protocol/lorenzo/v3/x/token/types"
)

// GetTokenFromDenom returns the token arg from the denom:
// - erc20/hex: return erc20 contract address
// - others: return the denom
func GetTokenFromDenom(denom string) string {
	err := tokentypes.ValidateERC20Denom(denom)
	if err != nil {
		return denom
	}

	denomSplit := strings.SplitN(denom, "/", 2)
	return denomSplit[1]
}
