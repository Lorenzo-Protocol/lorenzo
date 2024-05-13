package keeper

import (
	"github.com/Lorenzo-Protocol/lorenzo/x/plan/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
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

	authzKeeper authzkeeper.Keeper
}

func NewKeeper(
	storeKey storetypes.StoreKey,
	cdc codec.BinaryCodec,
	authority string,
	ak types.AccountKeeper,
	bk bankkeeper.Keeper,
	evmKeeper types.EVMKeeper,
	authzKeeper authzkeeper.Keeper,
) Keeper {

	return Keeper{
		authority:     authority,
		storeKey:      storeKey,
		cdc:           cdc,
		accountKeeper: ak,
		bankKeeper:    bk,
		evmKeeper:     evmKeeper,
		authzKeeper:   authzKeeper,
	}
}
