package keeper

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	"github.com/cometbft/cometbft/libs/log"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"

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
	if addr := accountKeeper.GetModuleAddress(types.ModuleName); addr == nil {
		panic(fmt.Sprintf("the x/%s module account has not been set", types.ModuleName))
	}

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
//  4. if receiver is not equal to sender, then the coin must be allowed to send.
func (k Keeper) MintEnabled(
	ctx sdk.Context,
	sender, receiver sdk.AccAddress,
	token string,
) (types.TokenPair, error) {
	if !k.IsConvertEnabled(ctx) {
		return types.TokenPair{}, errorsmod.Wrapf(types.ErrConvertDisabled,
			"token module is disabled")
	}

	id := k.GetTokenPairId(ctx, token)
	pair, found := k.GetTokenPair(ctx, id)
	if !found {
		return types.TokenPair{}, errorsmod.Wrapf(types.ErrTokenPairNotFound,
			"token pair not found for token %s", token)
	}

	if !pair.Enabled {
		return types.TokenPair{}, errorsmod.Wrapf(types.ErrTokenPairDisabled,
			"token pair is disabled for token %s", token)
	}

	if k.bankKeeper.BlockedAddr(receiver.Bytes()) {
		return types.TokenPair{}, errorsmod.Wrapf(errortypes.ErrUnauthorized,
			"%s is not allowed to receive transactions", receiver,
		)
	}

	if !sender.Equals(receiver) && !k.bankKeeper.IsSendEnabledCoin(ctx, sdk.Coin{Denom: pair.Denom}) {
		return types.TokenPair{}, errorsmod.Wrapf(errortypes.ErrUnauthorized,
			"coin is not allowed to be sent")
	}

	return pair, nil
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
