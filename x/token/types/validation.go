package types

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ethereum/go-ethereum/common"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// ValidateERC20Denom validates if an erc20 denom follows erc/hex spec:
// - erc20 denom can only be created as the result of executing MsgRegisterCoin
// - erc20 denom must be prefixed with DenomPrefix and formed in "DenomPrefix/ContractAddress"
// - denom that doesn't follow rules as considered as non erc20 denom
func ValidateERC20Denom(denom string) error {
	denomSplit := strings.SplitN(denom, "/", 2)
	if denomSplit[0] != DenomPrefix || len(denomSplit) != 2 {
		return fmt.Errorf("denom deosn't follow rules of erc20 denom")
	}

	if !common.IsHexAddress(denomSplit[1]) {
		return fmt.Errorf("denom dosen't follow rules of erc20 denom")
	}

	return nil
}

// CompareMetadata checks if all the fields of the provided coin metadata are equal.
func CompareMetadata(a, b banktypes.Metadata) error {
	if a.Base == b.Base && a.Description == b.Description &&
		a.Display == b.Display && a.Name == b.Name && a.Symbol == b.Symbol {
		if len(a.DenomUnits) != len(b.DenomUnits) {
			return fmt.Errorf(
				"metadata provided has different denom units from stored, %d ≠ %d",
				len(a.DenomUnits), len(b.DenomUnits))
		}

		for i, v := range a.DenomUnits {
			if (v.Exponent != b.DenomUnits[i].Exponent) ||
				(v.Denom != b.DenomUnits[i].Denom) ||
				!CompareStringSlice(v.Aliases, b.DenomUnits[i].Aliases) {
				return fmt.Errorf(
					"metadata provided has different denom unit from stored, %s ≠ %s",
					a.DenomUnits[i], b.DenomUnits[i])
			}
		}

		return nil
	}
	return fmt.Errorf("metadata provided is different from stored")
}

// CompareStringSlice checks if two string slices are equal.
func CompareStringSlice(aliasesA, aliasesB []string) bool {
	if len(aliasesA) != len(aliasesB) {
		return false
	}

	for i := 0; i < len(aliasesA); i++ {
		if aliasesA[i] != aliasesB[i] {
			return false
		}
	}

	return true
}

// SanitizeERC20Name sanitize erc20 name to be a acceptable denom.
// NOTE: when erc20 contract has decimals > 0, an additional denom unit for this erc20 denom is created.
// The denom unit must compromise rules as follows:
// - denom must not be prefixed with ibc
// - denom must not be prefixed with DenomPrefix
// - denom must not be prefixed with digital numbers
// - denom must not contain any unexpected char
func SanitizeERC20Name(name string) string {
	name = removeLeadingNumbers(name)
	name = removeSpecialChars(name)
	if len(name) > 128 {
		name = name[:128]
	}
	name = removeInvalidPrefixes(name)
	return name
}

const (
	// (?m)^(\d+) remove leading numbers
	reLeadingNumbers = `(?m)^(\d+)`
	// ^[^A-Za-z] forces first chars to be letters
	// [^a-zA-Z0-9/-] deletes special characters
	reDnmString = `^[^A-Za-z]|[^a-zA-Z0-9/-]`
)

func removeLeadingNumbers(str string) string {
	re := regexp.MustCompile(reLeadingNumbers)
	return re.ReplaceAllString(str, "")
}

func removeSpecialChars(str string) string {
	re := regexp.MustCompile(reDnmString)
	return re.ReplaceAllString(str, "")
}

// recursively remove every invalid prefix
func removeInvalidPrefixes(str string) string {
	if strings.HasPrefix(str, "ibc/") {
		return removeInvalidPrefixes(str[4:])
	}
	if strings.HasPrefix(str, DenomPrefix+"/") {
		return removeInvalidPrefixes(str[6:])
	}
	return str
}
