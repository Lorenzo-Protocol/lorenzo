package types

import (
	big "math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"

	lrz "github.com/Lorenzo-Protocol/lorenzo/v2/types"
	agenttypes "github.com/Lorenzo-Protocol/lorenzo/v2/x/agent/types"
	btclctypes "github.com/Lorenzo-Protocol/lorenzo/v2/x/btclightclient/types"
	plantypes "github.com/Lorenzo-Protocol/lorenzo/v2/x/plan/types"
	"github.com/btcsuite/btcd/chaincfg"
)

type BTCLightClientKeeper interface {
	GetBaseBTCHeader(ctx sdk.Context) *btclctypes.BTCHeaderInfo
	GetTipInfo(ctx sdk.Context) *btclctypes.BTCHeaderInfo
	GetHeaderByHash(ctx sdk.Context, hash *lrz.BTCHeaderHashBytes) *btclctypes.BTCHeaderInfo
	GetBTCNet() *chaincfg.Params
	// GetFeeRate(ctx sdk.Context) uint64
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
