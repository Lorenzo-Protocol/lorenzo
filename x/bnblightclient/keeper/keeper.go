package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Lorenzo-Protocol/lorenzo/v2/x/bnblightclient/types"
)

// Keeper is the keeper struct for the x/bnblightclient module
type Keeper struct {
	storeKey storetypes.StoreKey
	cdc      codec.BinaryCodec
	
	// the address capable of executing a MsgUpdateParams message. Typically, this should be the x/gov module account.
	authority string
}

// NewKeeper creates a new Keeper object
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	authority string,
) Keeper {
	return Keeper{
		storeKey:  storeKey,
		cdc:       cdc,
		authority: authority,
	}
}

// ChainID returns the chain id
func (k Keeper) ChainID(ctx sdk.Context) uint32 {
	return k.GetParams(ctx).ChainId
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
