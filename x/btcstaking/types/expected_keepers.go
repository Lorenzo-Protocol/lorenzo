package types

import (
	big "math/big"

	"github.com/btcsuite/btcd/chaincfg"
	sdk "github.com/cosmos/cosmos-sdk/types"
	evmtypes "github.com/ethereum/go-ethereum/core/types"

	lrz "github.com/Lorenzo-Protocol/lorenzo/v2/types"
	bnblightclienttypes "github.com/Lorenzo-Protocol/lorenzo/v2/x/bnblightclient/types"
	btclctypes "github.com/Lorenzo-Protocol/lorenzo/v2/x/btclightclient/types"
)

type BTCLightClientKeeper interface {
	GetBaseBTCHeader(ctx sdk.Context) *btclctypes.BTCHeaderInfo
	GetTipInfo(ctx sdk.Context) *btclctypes.BTCHeaderInfo
	GetHeaderByHash(ctx sdk.Context, hash *lrz.BTCHeaderHashBytes) *btclctypes.BTCHeaderInfo
	GetBTCNet() *chaincfg.Params
	GetFeeRate(ctx sdk.Context) uint64
}

// BNBLightClientKeeper is an expected keeper for the bnblightclient module
type BNBLightClientKeeper interface {
	VerifyReceiptProof(ctx sdk.Context, receipt *evmtypes.Receipt, proof *bnblightclienttypes.Proof) ([]bnblightclienttypes.CrossChainEvent, error)
}

type EvmKeeper interface {
	ChainID() *big.Int
}
