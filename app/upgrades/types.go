package upgrades

import (
	agentkeeper "github.com/Lorenzo-Protocol/lorenzo/v2/x/agent/keeper"
	btcstakingkeeper "github.com/Lorenzo-Protocol/lorenzo/v2/x/btcstaking/keeper"
	plankeeper "github.com/Lorenzo-Protocol/lorenzo/v2/x/plan/keeper"
	tokenkeeper "github.com/Lorenzo-Protocol/lorenzo/v2/x/token/keeper"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	store "github.com/cosmos/cosmos-sdk/store/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ibckeeper "github.com/cosmos/ibc-go/v7/modules/core/keeper"
	evmkeeper "github.com/evmos/ethermint/x/evm/keeper"
	feemarketkeeper "github.com/evmos/ethermint/x/feemarket/keeper"
)

// Upgrade defines a struct containing necessary fields that a SoftwareUpgradeProposal
// must have written, in order for the state migration to go smoothly.
// An upgrade must implement this struct, and then set it in the app.go.
// The app.go will then define the handler.
type Upgrade struct {
	// Upgrade version name, for the upgrade handler, e.g. `v7`
	UpgradeName string

	// UpgradeHandlerConstructor defines the function that creates an upgrade handler
	UpgradeHandlerConstructor func(*module.Manager, module.Configurator, AppKeepers) upgradetypes.UpgradeHandler

	// Store upgrades, should be used for any new modules introduced, new modules deleted, or store names renamed.
	StoreUpgrades *store.StoreUpgrades
}

type ConsensusParamsReaderWriter interface {
	StoreConsensusParams(ctx sdk.Context, cp *tmproto.ConsensusParams)
	GetConsensusParams(ctx sdk.Context) *tmproto.ConsensusParams
}

type AppKeepers struct {
	AppCodec         codec.Codec
	AccountKeeper    authkeeper.AccountKeeper
	BankKeeper       bankkeeper.Keeper
	GetKey           func(moduleName string) *storetypes.KVStoreKey
	ModuleManager    *module.Manager
	IBCKeeper        *ibckeeper.Keeper
	EvmKeeper        *evmkeeper.Keeper
	FeeMarketKeeper  feemarketkeeper.Keeper
	AgentKeeper      agentkeeper.Keeper
	PlanKeeper       *plankeeper.Keeper
	BTCStakingKeeper btcstakingkeeper.Keeper
	TokenKeeper      *tokenkeeper.Keeper

	ReaderWriter ConsensusParamsReaderWriter
}

type UpgradeRouter struct {
	mu map[string]Upgrade
}

func NewUpgradeRouter() *UpgradeRouter {
	return &UpgradeRouter{make(map[string]Upgrade)}
}

func (r *UpgradeRouter) Register(u Upgrade) *UpgradeRouter {
	if _, has := r.mu[u.UpgradeName]; has {
		panic(u.UpgradeName + " already registered")
	}
	r.mu[u.UpgradeName] = u
	return r
}

func (r *UpgradeRouter) Routers() map[string]Upgrade {
	return r.mu
}

func (r *UpgradeRouter) UpgradeInfo(planName string) Upgrade {
	return r.mu[planName]
}
