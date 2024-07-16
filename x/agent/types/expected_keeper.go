package types

import (
	"github.com/btcsuite/btcd/chaincfg"
)

// BTCLightClientKeeper is an expected keeper for btc light client
type BTCLightClientKeeper interface {
	GetBTCNet() *chaincfg.Params
}
