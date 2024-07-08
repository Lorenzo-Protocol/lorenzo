package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/cosmos/ibc-go/v7/modules/apps/transfer/keeper"
	transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	porttypes "github.com/cosmos/ibc-go/v7/modules/core/05-port/types"
	"github.com/cosmos/ibc-go/v7/modules/core/exported"

	"github.com/Lorenzo-Protocol/lorenzo/x/ibctransfer/types"
)

// Keeper defines the modified IBC transfer keeper that embeds the original one.
type Keeper struct {
	*keeper.Keeper
	bankKeeper    types.BankKeeper
	tokenKeeper   types.TokenKeeper
	accountKeeper types.AccountKeeper
}

// NewKeeper creates a new IBC transfer Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	paramSpace paramtypes.Subspace,

	ics4Wrapper porttypes.ICS4Wrapper,
	channelKeeper transfertypes.ChannelKeeper,
	portKeeper transfertypes.PortKeeper,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	scopedKeeper exported.ScopedKeeper,
	tokenKeeper types.TokenKeeper,
) Keeper {
	transferKeeper := keeper.NewKeeper(
		cdc, storeKey, paramSpace,
		ics4Wrapper, channelKeeper, portKeeper,
		accountKeeper, bankKeeper, scopedKeeper,
	)

	return Keeper{
		Keeper:        &transferKeeper,
		bankKeeper:    bankKeeper,
		tokenKeeper:   tokenKeeper,
		accountKeeper: accountKeeper,
	}
}
