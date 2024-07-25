package types

import (
	big "math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"

	lrz "github.com/Lorenzo-Protocol/lorenzo/v2/types"
	btclctypes "github.com/Lorenzo-Protocol/lorenzo/v2/x/btclightclient/types"
	"github.com/btcsuite/btcd/chaincfg"
)

type BTCLightClientKeeper interface {
	GetBaseBTCHeader(ctx sdk.Context) *btclctypes.BTCHeaderInfo
	GetTipInfo(ctx sdk.Context) *btclctypes.BTCHeaderInfo
	GetHeaderByHash(ctx sdk.Context, hash *lrz.BTCHeaderHashBytes) *btclctypes.BTCHeaderInfo
	GetBTCNet() *chaincfg.Params
	GetFeeRate(ctx sdk.Context) uint64
}

type EvmKeeper interface {
	ChainID() *big.Int
}
