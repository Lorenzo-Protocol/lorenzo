package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Lorenzo-Protocol/lorenzo/x/token/types"
)

type Keeper struct {
	storeKey  storetypes.StoreKey
	cdc       codec.BinaryCodec
	authority sdk.AccAddress

	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
	evmKeeper     types.EVMKeeper
}

// NewKeeper creates a new token keeper
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	authority sdk.AccAddress,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	evmKeeper types.EVMKeeper,
) *Keeper {
	return &Keeper{
		storeKey:      storeKey,
		cdc:           cdc,
		authority:     authority,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
		evmKeeper:     evmKeeper,
	}
}

// MintEnabled checks whether the token is allowed to mint and convert:
// It returns the token pair without error if:
//  1. global conversion is enabled
//  2. token pair conversion is enabled
//  3. receiver address is not blocked by bank module.
//  4. coins are enabled for bank module transfers
func (k Keeper) MintEnabled(
	ctx sdk.Context,
	sender, receiver sdk.AccAddress,
	token string,
) (types.TokenPair, error) {
	panic("implement me")
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
