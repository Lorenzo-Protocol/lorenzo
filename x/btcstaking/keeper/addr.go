package keeper

import (
	"bytes"
	"encoding/hex"

	"github.com/Lorenzo-Protocol/lorenzo/x/btcstaking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"golang.org/x/crypto/sha3"
)

func validateMintToAddr40(addrStr []byte) (sdk.AccAddress, error) {
	addr, err := hex.DecodeString(string(addrStr))
	if err != nil {
		return nil, types.ErrMintToAddr.Wrapf("%s", hex.EncodeToString(addrStr))
	}
	hasUpper, hasLower := false, false
	for i := 0; i < len(addrStr); i++ {
		// 'a' > 'A' > '0'
		if addrStr[i] >= 'a' {
			hasLower = true
		} else if addrStr[i] >= 'A' {
			hasUpper = true
		}
	}
	if hasLower && hasUpper {
		// checksummed address
		hasher := sha3.NewLegacyKeccak256()
		hasher.Write(bytes.ToLower(addrStr))
		addrHash := hex.EncodeToString(hasher.Sum(nil))
		for i := 0; i < len(addrStr); i++ {
			if addrStr[i] <= '9' {
			} else if addrHash[i] >= '8' && addrStr[i] >= 'a' && addrStr[i] <= 'z' {
				// error should be uppercase
				return nil, types.ErrMintToAddr.Wrapf("checksum error: should be uppercase %d %s", i, string(addrStr))
			} else if addrHash[i] < '8' && addrStr[i] >= 'A' && addrStr[i] <= 'Z' {
				// error should be lowercase
				return nil, types.ErrMintToAddr.Wrapf("checksum error: %s", string(addrStr))
			}
		}
	}
	return addr, nil
}
func validateMintToAddress(addr []byte) (sdk.AccAddress, error) {
	if len(addr) == 20 {
		// raw bytes, not hex encoded, no check required
		return addr, nil
	} else if len(addr) == 42 {
		if addr[0] == '0' && (addr[1] == 'x' || addr[1] == 'X') {
			return validateMintToAddr40(addr[2:])
		} else {
			return nil, types.ErrMintToAddr.Wrapf("unknown prefix: %s", string(addr))
		}
	} else if len(addr) == 40 {
		return validateMintToAddr40(addr)
	}
	return nil, types.ErrMintToAddr.Wrapf("invalid mint to address length: %d", len(addr))
}
