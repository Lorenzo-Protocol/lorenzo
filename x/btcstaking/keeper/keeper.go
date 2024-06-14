package keeper

import (
	"fmt"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/Lorenzo-Protocol/lorenzo/x/btcstaking/types"
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

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		authority string
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,

	btclcKeeper types.BTCLightClientKeeper,
	bankKeeper bankkeeper.Keeper,

	authority string,
) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,

		btclcKeeper: btclcKeeper,
		bankKeeper:  bankKeeper,

		authority: authority,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
	//return ctx.Logger().With("module", fmt.Sprintf("x/btcstaking"))
}

// BeginBlocker is invoked upon `BeginBlock` of the system. The function
// iterates over all BTC delegations under non-slashed finality providers
// to 1) record the voting power table for the current height, and 2) record
// the voting power distribution cache used for computing voting power table
// and distributing rewards once the block is finalised by finality providers.
/*func (k Keeper) BeginBlocker(ctx context.Context) error {
	// index BTC height at the current height
	k.IndexBTCHeight(ctx)
	// update voting power distribution
	k.UpdatePowerDist(ctx)

	return nil
}*/
