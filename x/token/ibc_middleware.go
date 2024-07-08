package token

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"

	"github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v7/modules/core/05-port/types"
	ibcexported "github.com/cosmos/ibc-go/v7/modules/core/exported"

	"github.com/Lorenzo-Protocol/lorenzo/x/token/keeper"
)

var _ porttypes.IBCModule = &IBCMiddleware{}

// IBCMiddleware implements the ICS26 callbacks for the transfer middleware given
// the erc20 keeper and the underlying application.
type IBCMiddleware struct {
	keeper.Keeper
}

func (I IBCMiddleware) OnChanOpenInit(ctx sdk.Context, order types.Order, connectionHops []string, portID string, channelID string, channelCap *capabilitytypes.Capability, counterparty types.Counterparty, version string) (string, error) {
	// TODO implement me
	panic("implement me")
}

func (I IBCMiddleware) OnChanOpenTry(ctx sdk.Context, order types.Order, connectionHops []string, portID, channelID string, channelCap *capabilitytypes.Capability, counterparty types.Counterparty, counterpartyVersion string) (version string, err error) {
	// TODO implement me
	panic("implement me")
}

func (I IBCMiddleware) OnChanOpenAck(ctx sdk.Context, portID, channelID string, counterpartyChannelID string, counterpartyVersion string) error {
	// TODO implement me
	panic("implement me")
}

func (I IBCMiddleware) OnChanOpenConfirm(ctx sdk.Context, portID, channelID string) error {
	// TODO implement me
	panic("implement me")
}

func (I IBCMiddleware) OnChanCloseInit(ctx sdk.Context, portID, channelID string) error {
	// TODO implement me
	panic("implement me")
}

func (I IBCMiddleware) OnChanCloseConfirm(ctx sdk.Context, portID, channelID string) error {
	// TODO implement me
	panic("implement me")
}

func (I IBCMiddleware) OnRecvPacket(ctx sdk.Context, packet types.Packet, relayer sdk.AccAddress) ibcexported.Acknowledgement {
	// TODO implement me
	panic("implement me")
}

func (I IBCMiddleware) OnAcknowledgementPacket(ctx sdk.Context, packet types.Packet, acknowledgement []byte, relayer sdk.AccAddress) error {
	// TODO implement me
	panic("implement me")
}

func (I IBCMiddleware) OnTimeoutPacket(ctx sdk.Context, packet types.Packet, relayer sdk.AccAddress) error {
	// TODO implement me
	panic("implement me")
}
