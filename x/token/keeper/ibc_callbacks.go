package keeper

import (
	"github.com/Lorenzo-Protocol/lorenzo/x/token/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/ibc-go/v7/modules/core/exported"
	"github.com/ethereum/go-ethereum/common"

	transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
)

// OnRecvPacket is post-posting logic for the IBC transfer module
// TODO: we haven't decide what to do with post-processing
func (k Keeper) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	ack exported.Acknowledgement,
) exported.Acknowledgement {
	return ack
}

// OnAcknowledgementPacket is post-posting logic for the IBC transfer module
func (k Keeper) OnAcknowledgementPacket(
	ctx sdk.Context,
	_ channeltypes.Packet,
	data transfertypes.FungibleTokenPacketData,
	ack channeltypes.Acknowledgement,
) error {
	switch ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:
		return k.ConvertCoinFromPacket(ctx, data)
	default:
		return nil
	}
}

// OnTimeoutPacket is post-posting logic for the IBC transfer module
func (k Keeper) OnTimeoutPacket(
	ctx sdk.Context,
	_ channeltypes.Packet,
	data transfertypes.FungibleTokenPacketData,
) error {
	return k.ConvertCoinFromPacket(ctx, data)
}

// ConvertCoinFromPacket converts the coin from the packet
func (k Keeper) ConvertCoinFromPacket(
	ctx sdk.Context,
	data transfertypes.FungibleTokenPacketData,
) error {
	sender, err := sdk.AccAddressFromBech32(data.Sender)
	if err != nil {
		return err
	}

	// avoid extra costs for relayers.
	ctx = ctx.WithKVGasConfig(storetypes.GasConfig{}).
		WithTransientKVGasConfig(storetypes.GasConfig{})

	// get ibc (or just original) denom
	denom := transfertypes.ParseDenomTrace(data.Denom).IBCDenom()
	amount, _ := sdk.NewIntFromString(data.Amount)

	// ConvertCoin will help to check if the denom is registered
	if _, err := k.ConvertCoin(sdk.WrapSDKContext(ctx), &types.MsgConvertCoin{
		Coin:     sdk.Coin{Denom: denom, Amount: amount},
		Receiver: common.BytesToAddress(sender).String(),
		Sender:   sender.String(),
	}); err != nil {
		return err
	}

	return nil
}
