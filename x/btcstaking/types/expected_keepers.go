package types

import (
	big "math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	evmtypes "github.com/ethereum/go-ethereum/core/types"

	lrz "github.com/Lorenzo-Protocol/lorenzo/v3/types"
	agenttypes "github.com/Lorenzo-Protocol/lorenzo/v3/x/agent/types"
	"github.com/Lorenzo-Protocol/lorenzo/v3/x/bnblightclient/types"
	btclctypes "github.com/Lorenzo-Protocol/lorenzo/v3/x/btclightclient/types"
	ccevtypes "github.com/Lorenzo-Protocol/lorenzo/v3/x/ccev/types"
	plantypes "github.com/Lorenzo-Protocol/lorenzo/v3/x/plan/types"
	"github.com/btcsuite/btcd/chaincfg"
)

type BTCLightClientKeeper interface {
	GetBaseBTCHeader(ctx sdk.Context) *btclctypes.BTCHeaderInfo
	GetTipInfo(ctx sdk.Context) *btclctypes.BTCHeaderInfo
	GetHeaderByHash(ctx sdk.Context, hash *lrz.BTCHeaderHashBytes) *btclctypes.BTCHeaderInfo
	GetBTCNet() *chaincfg.Params
	// GetFeeRate(ctx sdk.Context) uint64
}

type BNBLightClientKeeper interface {
	VerifyReceiptProof(ctx sdk.Context, number uint64, receipt *evmtypes.Receipt, proof *types.Proof) ([]types.CrossChainEvent, error)
}

type CCEVKeeper interface {
	VerifyAndCallback(ctx sdk.Context, chainID uint32, number uint64, receiptBz []byte, proofBz []byte, handler ccevtypes.EventHandler) error
}

type PlanKeeper interface {
	Mint(ctx sdk.Context, planId uint64, to common.Address, amount *big.Int) error
	GetPlan(ctx sdk.Context, planId uint64) (plantypes.Plan, bool)
}
type AgentKeeper interface {
	GetAgent(ctx sdk.Context, id uint64) (agenttypes.Agent, bool)
}

type EvmKeeper interface {
	ChainID() *big.Int
}
