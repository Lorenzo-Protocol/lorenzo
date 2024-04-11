package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	lrz "github.com/Lorenzo-Protocol/lorenzo/types"
	btclctypes "github.com/Lorenzo-Protocol/lorenzo/x/btclightclient/types"
	"github.com/btcsuite/btcd/chaincfg"
)

type BTCLightClientKeeper interface {
	GetBaseBTCHeader(ctx sdk.Context) *btclctypes.BTCHeaderInfo
	GetTipInfo(ctx sdk.Context) *btclctypes.BTCHeaderInfo
	GetHeaderByHash(ctx sdk.Context, hash *lrz.BTCHeaderHashBytes) *btclctypes.BTCHeaderInfo
	GetBTCNet() *chaincfg.Params
	GetFeeRate(ctx sdk.Context) uint64
}
