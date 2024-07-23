package token_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"

	"github.com/Lorenzo-Protocol/lorenzo/app"
	ibctesting "github.com/Lorenzo-Protocol/lorenzo/testutil/ibc"
)

type MiddlewareTestSuite struct {
	suite.Suite

	chainA *app.LorenzoApp
	chainB *app.LorenzoApp

	Coordinator *ibctesting.Coordinator

	LorenzoChainA *ibctesting.TestChain
	LorenzoChainB *ibctesting.TestChain

	Path *ibctesting.Path
}

func (suite *MiddlewareTestSuite) SetupTest() {
	suite.Coordinator = ibctesting.NewCoordinator(suite.T(), 2)
	suite.LorenzoChainA = suite.Coordinator.GetChain(ibctesting.GetChainID(1))
	suite.LorenzoChainB = suite.Coordinator.GetChain(ibctesting.GetChainID(2))

	suite.Path = ibctesting.NewPath(suite.LorenzoChainA, suite.LorenzoChainB)
	suite.Path.EndpointA.ChannelConfig.PortID = ibctransfertypes.ModuleName
	suite.Path.EndpointB.ChannelConfig.PortID = ibctransfertypes.ModuleName
	suite.Path.EndpointA.ChannelConfig.Version = ibctransfertypes.Version
	suite.Path.EndpointB.ChannelConfig.Version = ibctransfertypes.Version

	suite.Coordinator.Setup(suite.Path)

	suite.chainA = suite.LorenzoChainA.App.(*app.LorenzoApp) //nolint:errcheck
	suite.chainB = suite.LorenzoChainB.App.(*app.LorenzoApp) //nolint:errcheck
}

func TestMiddlewareTestSuite(t *testing.T) {
	suite.Run(t, new(MiddlewareTestSuite))
}

// NewMockPacket returns a new mock packet sending from chain-a to chain-b
func (suite *MiddlewareTestSuite) NewMockTransferPacket(data []byte) channeltypes.Packet {
	return channeltypes.NewPacket(
		data,
		suite.LorenzoChainA.SenderAccount.GetSequence(),
		suite.Path.EndpointA.ChannelConfig.PortID,
		suite.Path.EndpointA.ChannelID,
		suite.Path.EndpointB.ChannelConfig.PortID,
		suite.Path.EndpointB.ChannelID,
		clienttypes.NewHeight(0, 100),
		0,
	)
}

// utilsCreateIBCDenom creates an IBC denom from the given channel and port identifiers and raw denom.
// NOTE: it's only used for test purpose and not safe for multiple hop denom.
func (suite *MiddlewareTestSuite) utilsCreateIBCDenom(destChan, destPort, baseDenom string) string {
	trace := ibctransfertypes.DenomTrace{
		Path:      fmt.Sprintf("%s/%s", destPort, destChan),
		BaseDenom: baseDenom,
	}
	return trace.IBCDenom()
}

func (suite *MiddlewareTestSuite) utilsMockAcknowledgement(success bool) []byte {
	if !success {
		ack := channeltypes.NewErrorAcknowledgement(errors.New("ics-20 error acknowledgement"))
		return ack.Acknowledgement()
	}
	return channeltypes.NewResultAcknowledgement([]byte{byte(1)}).Acknowledgement()
}
