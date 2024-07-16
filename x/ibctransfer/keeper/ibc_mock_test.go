package keeper_test

import (
	"github.com/stretchr/testify/mock"

	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"

	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/v7/modules/core/exported"
)

type MockChannelKeeper struct {
	mock.Mock
}

func (b *MockChannelKeeper) GetChannel(
	ctx sdk.Context,
	srcPort,
	srcChan string,
) (channel channeltypes.Channel, found bool) {
	args := b.Called(mock.Anything, mock.Anything, mock.Anything)
	return args.Get(0).(channeltypes.Channel), true
}

func (b *MockChannelKeeper) GetNextSequenceSend(
	ctx sdk.Context,
	portID,
	channelID string,
) (uint64, bool) {
	_ = b.Called(mock.Anything, mock.Anything, mock.Anything)
	return 1, true
}

func (b *MockChannelKeeper) GetAllChannelsWithPortPrefix(
	ctx sdk.Context,
	portPrefix string,
) []channeltypes.IdentifiedChannel {
	return []channeltypes.IdentifiedChannel{}
}

type MockICS4Wrapper struct {
	mock.Mock
}

func (b *MockICS4Wrapper) WriteAcknowledgement(
	_ sdk.Context,
	_ *capabilitytypes.Capability,
	_ exported.PacketI,
	_ exported.Acknowledgement,
) error {
	return nil
}

func (b *MockICS4Wrapper) GetAppVersion(
	ctx sdk.Context,
	portID string,
	channelID string,
) (string, bool) {
	return "", false
}

func (b *MockICS4Wrapper) SendPacket(
	ctx sdk.Context,
	channelCap *capabilitytypes.Capability,
	sourcePort string,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
	data []byte,
) (sequence uint64, err error) {
	return 0, nil
}
