package types

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// ValidateERC20Denom validates if an erc20 denom follows erc/hex spec:
// - erc20 denom can only be created as the result of executing MsgRegisterCoin
// - erc20 denom must be prefixed with DenomPrefix and formed in "DenomPrefix/ContractAddress"
// - denom that doesn't follow rules as considered as non erc20 denom
func ValidateERC20Denom(denom string) error {
	denomSplit := strings.SplitN(denom, "/", 2)

	if denomSplit[0] != DenomPrefix || len(denomSplit) != 2 {
		return fmt.Errorf("denom %s deosn't follow rules of erc20 denom")
	}

	if !common.IsHexAddress(denomSplit[1]) {
		return fmt.Errorf("denom dosen't follow rules of erc20 denom")
	}

	return nil
}
