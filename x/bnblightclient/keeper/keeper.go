package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)

// Keeper is the keeper struct for the x/bnblightclient module
type Keeper struct {
	storeKey storetypes.StoreKey
	cdc      codec.BinaryCodec

	// the address capable of executing a MsgUpdateParams message. Typically, this should be the x/gov module account.
	authority string
}

// NewKeeper creates a new Keeper object
func NewKeeper(cdc codec.BinaryCodec, storeKey storetypes.StoreKey, authority string) Keeper {
	return Keeper{
		storeKey:  storeKey,
		cdc:       cdc,
		authority: authority,
	}
}
