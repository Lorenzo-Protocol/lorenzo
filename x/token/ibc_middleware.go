package token

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	"github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v7/modules/core/05-port/types"
	ibcexported "github.com/cosmos/ibc-go/v7/modules/core/exported"

	ibctransfer "github.com/Lorenzo-Protocol/lorenzo/x/ibctransfer"
	"github.com/Lorenzo-Protocol/lorenzo/x/token/keeper"
)

var _ porttypes.IBCModule = &IBCMiddleware{}

// NewIBCMiddleware creates a new IBCMiddleware given the keeper and underlying application
func NewIBCMiddleware(module *ibctransfer.IBCModule, k *keeper.Keeper) IBCMiddleware {
	return IBCMiddleware{
		IBCModule: module,
		keeper:    k,
	}
}

// IBCMiddleware implements the ICS26 callbacks for the ibc-transfer module.
type IBCMiddleware struct {
	*ibctransfer.IBCModule
	keeper *keeper.Keeper
}

// OnRecvPacket implements the ICS-26 interface. If it successfully handles OnRecvPacket, a
// post-processing handler will try converting the coin to an ERC20 token.
func (im IBCMiddleware) OnRecvPacket(
	ctx sdk.Context,
	packet types.Packet,
	relayer sdk.AccAddress,
) ibcexported.Acknowledgement {
	ack := im.IBCModule.OnRecvPacket(ctx, packet, relayer)
	if !ack.Success() {
		return ack
	}

	// post-processing
	return im.keeper.OnRecvPacket(ctx, packet, ack)
}

// OnAcknowledgementPacket implements the ICS-26 interface. If it successfully handles OnAcknowledgementPacket,
// a post-processing handler will try refunding the token transferred and convert the coin to an ERC20 token.
func (im IBCMiddleware) OnAcknowledgementPacket(
	ctx sdk.Context,
	packet types.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	if err := im.IBCModule.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer); err != nil {
		return err
	}

	var ack channeltypes.Acknowledgement
	var data transfertypes.FungibleTokenPacketData
	transfertypes.ModuleCdc.MustUnmarshalJSON(packet.GetData(), &data)
	transfertypes.ModuleCdc.MustUnmarshalJSON(acknowledgement, &ack)

	// post-processing
	return im.keeper.OnAcknowledgementPacket(ctx, packet, data, ack)
}

// OnTimeoutPacket implements the ICS-26 interface. If it successfully handles OnTimeoutPacket,
// a post-processing handler will try refunding the token converted on previous sending.
func (im IBCMiddleware) OnTimeoutPacket(
	ctx sdk.Context,
	packet types.Packet,
	relayer sdk.AccAddress,
) error {
	if err := im.IBCModule.OnTimeoutPacket(ctx, packet, relayer); err != nil {
		return err
	}

	var data transfertypes.FungibleTokenPacketData
	transfertypes.ModuleCdc.MustUnmarshalJSON(packet.GetData(), &data)

	// post-processing
	return im.keeper.OnTimeoutPacket(ctx, packet, data)
}
