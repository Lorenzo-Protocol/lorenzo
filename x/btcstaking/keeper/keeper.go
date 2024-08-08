package keeper

import (
	"fmt"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/Lorenzo-Protocol/lorenzo/v2/x/btcstaking/types"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey storetypes.StoreKey

		bankKeeper bankkeeper.Keeper

		btclcKeeper types.BTCLightClientKeeper
		bnblcKeeper types.BNBLightClientKeeper
		agentKeeper types.AgentKeeper

		planKeeper types.PlanKeeper

		evmKeeper types.EvmKeeper

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		authority string
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,

	btclcKeeper types.BTCLightClientKeeper,
	bnblcKeeper types.BNBLightClientKeeper,
	agentKeeper types.AgentKeeper,
	bankKeeper bankkeeper.Keeper,
	planKeeper types.PlanKeeper,
	evmKeeper types.EvmKeeper,

	authority string,
) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,

		btclcKeeper: btclcKeeper,
		bnblcKeeper: bnblcKeeper,
		agentKeeper: agentKeeper,
		bankKeeper:  bankKeeper,
		planKeeper:  planKeeper,
		evmKeeper:   evmKeeper,

		authority: authority,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
