package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"

	"github.com/Lorenzo-Protocol/lorenzo/v3/x/fee/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper of the fee store
type Keeper struct {
	cdc       codec.BinaryCodec
	storeKey  storetypes.StoreKey
	authority string
}

// NewKeeper initializes a new Keeper.
//
// cdc - binary codec for the Keeper
// storeKey - store key for the Keeper
// authority - authority for the Keeper
// Returns a Keeper
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	authority string,
) *Keeper {
	return &Keeper{
		cdc:       cdc,
		storeKey:  storeKey,
		authority: authority,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
