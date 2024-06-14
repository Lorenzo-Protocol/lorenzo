package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"

	"github.com/Lorenzo-Protocol/lorenzo/x/plan/types"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

// Keeper of this module maintains collections of erc20.
type Keeper struct {
	storeKey storetypes.StoreKey
	cdc      codec.BinaryCodec

	// the address capable of executing a MsgUpdateParams message. Typically, this should be the x/gov module account.
	authority string

	accountKeeper types.AccountKeeper
	bankKeeper    bankkeeper.Keeper
	evmKeeper     types.EVMKeeper
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	authority string,
	ak types.AccountKeeper,
	bk bankkeeper.Keeper,
	evmKeeper types.EVMKeeper,
) *Keeper {

	return &Keeper{
		authority:     authority,
		storeKey:      storeKey,
		cdc:           cdc,
		accountKeeper: ak,
		bankKeeper:    bk,
		evmKeeper:     evmKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
