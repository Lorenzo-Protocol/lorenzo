package keeper

import (
	errorsmod "cosmossdk.io/errors"

	"github.com/ethereum/go-ethereum/common"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"

	transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/v7/modules/core/exported"

	"github.com/Lorenzo-Protocol/lorenzo/x/token/types"
)

// OnRecvPacket is post-posting logic for the IBC transfer module
func (k Keeper) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	ack exported.Acknowledgement,
) exported.Acknowledgement {
	var data transfertypes.FungibleTokenPacketData
	if err := transfertypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		err = errorsmod.Wrapf(errortypes.ErrInvalidType, "cannot unmarshal ICS-20 transfer packet data")
		return channeltypes.NewErrorAcknowledgement(err)
	}

	receiver, err := sdk.AccAddressFromBech32(data.Receiver)
	if err != nil {
		return channeltypes.NewErrorAcknowledgement(err)
	}

	// avoid extra costs for relayers.
	ctx = ctx.WithKVGasConfig(storetypes.GasConfig{}).WithTransientKVGasConfig(storetypes.GasConfig{})

	coin := k.getReceivedCoin(
		packet.SourcePort, packet.SourceChannel,
		packet.DestinationPort, packet.DestinationChannel,
		data.Denom, data.Amount)

	if _, err := k.ConvertCoin(sdk.WrapSDKContext(ctx), &types.MsgConvertCoin{
		Coin: sdk.Coin{
			Denom:  coin.Denom,
			Amount: coin.Amount,
		},
		Receiver: common.BytesToAddress(receiver.Bytes()).String(),
		Sender:   receiver.String(),
	}); err != nil {
		return channeltypes.NewErrorAcknowledgement(err)
	}

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

// getReceivedCoin returns the transferred coin from an ICS20 FungibleTokenPacketData
func (k Keeper) getReceivedCoin(srcPort, srcChannel, dstPort, dstChannel, rawDenom, rawAmt string) sdk.Coin {
	// NOTE: Denom and amount are already validated
	amount, _ := sdk.NewIntFromString(rawAmt)

	// if coin arrives at source chain, remove its prefix
	if transfertypes.ReceiverChainIsSource(srcPort, srcChannel, rawDenom) {
		voucherPrefix := transfertypes.GetDenomPrefix(srcPort, srcChannel)
		unprefixedDenom := rawDenom[len(voucherPrefix):]

		denom := unprefixedDenom

		denomTrace := transfertypes.ParseDenomTrace(unprefixedDenom)
		if denomTrace.Path != "" {
			denom = denomTrace.IBCDenom()
		}

		return sdk.Coin{
			Denom:  denom,
			Amount: amount,
		}
	}

	// otherwise add the prefix
	sourcePrefix := transfertypes.GetDenomPrefix(dstPort, dstChannel)
	prefixedDenom := sourcePrefix + rawDenom
	denomTrace := transfertypes.ParseDenomTrace(prefixedDenom)
	voucherDenom := denomTrace.IBCDenom()

	return sdk.Coin{
		Denom:  voucherDenom,
		Amount: amount,
	}
}
