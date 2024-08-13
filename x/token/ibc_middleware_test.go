package token_test

import (
	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"

	"github.com/Lorenzo-Protocol/lorenzo/v3/contracts/erc20"
	tokentypes "github.com/Lorenzo-Protocol/lorenzo/v3/x/token/types"
)

// TestOnRecvPacket tests the OnRecvPacket function of the middleware.
// Note: mock packet received instead sending an actual packet.
func (suite *MiddlewareTestSuite) TestOnRecvPacket() {
	testCases := []struct {
		name            string
		expectConverted bool
		baseDenom       string
		pair            tokentypes.TokenPair
	}{
		{
			name:            "token pair not registered",
			baseDenom:       "coin",
			expectConverted: false,
		},
		{
			name:            "token pair registered",
			baseDenom:       "coin",
			expectConverted: true,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest()

			module, _, err := suite.LorenzoChainB.App.GetIBCKeeper().PortKeeper.LookupModuleByPort(
				suite.LorenzoChainB.GetContext(), ibctransfertypes.ModuleName)
			suite.Require().NoError(err)

			cbs, ok := suite.LorenzoChainB.App.GetIBCKeeper().Router.GetRoute(module)
			suite.Require().True(ok)

			sender := suite.LorenzoChainA.SenderAccount.GetAddress()
			receiver := suite.LorenzoChainB.SenderAccount.GetAddress()

			// construct package sent from chain-a
			data := ibctransfertypes.FungibleTokenPacketData{
				Denom:    tc.baseDenom,
				Amount:   "1000",
				Sender:   sender.String(),
				Receiver: receiver.String(),
				Memo:     "none",
			}
			packet := suite.NewMockTransferPacket(data.GetBytes())

			// register in advance
			if tc.expectConverted {
				ibcDenom := suite.utilsCreateIBCDenom(
					suite.Path.EndpointB.ChannelID,
					suite.Path.EndpointB.ChannelConfig.PortID,
					"coin")

				// NOTE: mint coin before register, avoid empty supply.
				err := suite.chainB.BankKeeper.MintCoins(suite.LorenzoChainB.GetContext(),
					tokentypes.ModuleName, sdk.NewCoins(sdk.NewCoin(ibcDenom, sdk.NewInt(1))))
				suite.Require().NoError(err)

				// register pair for this token.
				pair, err := suite.chainB.TokenKeeper.RegisterCoin(suite.LorenzoChainB.GetContext(), banktypes.Metadata{
					Description: "",
					DenomUnits: []*banktypes.DenomUnit{
						{
							Denom:    ibcDenom,
							Exponent: 0,
						},
					},
					Base:    ibcDenom,
					Display: ibcDenom,
					Name:    "coin",
					Symbol:  "coin",
				})
				suite.Require().NoError(err)
				resPair, found := suite.chainB.TokenKeeper.GetTokenPair(suite.LorenzoChainB.GetContext(), pair.GetID())
				suite.Require().True(found)
				suite.Require().Equal(pair, &resPair)
				tc.pair = *pair
			}

			ack := cbs.OnRecvPacket(suite.LorenzoChainB.GetContext(), packet, suite.LorenzoChainB.SenderAccount.GetAddress())
			suite.Require().True(ack.Success())

			if tc.expectConverted {
				balance := suite.chainB.TokenKeeper.ERC20BalanceOf(
					suite.LorenzoChainB.GetContext(),
					erc20.ERC20MinterBurnerDecimalsContract.ABI,
					common.HexToAddress(tc.pair.ContractAddress),
					common.BytesToAddress(receiver.Bytes()),
				)
				suite.Require().Equal(data.Amount, balance.String())
			}
		})
	}
}

// TestOnAcknowledgementPacket tests the OnAcknowledgementPacket function of the middleware.
// Note: mock packet timeout instead sending an actual packet.
func (suite *MiddlewareTestSuite) TestOnAcknowledgementPacket() {
	testCases := []struct {
		name          string
		malleate      func()
		ack           []byte
		invalidAck    bool
		expectConvert bool
	}{
		{
			name:       "invalid ack",
			ack:        []byte("any-value"),
			invalidAck: true,
		},
		{
			name: "ack confirmed",
			ack:  suite.utilsMockAcknowledgement(true),
		},
		{
			name: "ack error, token registered",
			ack:  suite.utilsMockAcknowledgement(false),
			malleate: func() {
				_, err := suite.chainA.TokenKeeper.RegisterCoin(suite.LorenzoChainA.GetContext(), banktypes.Metadata{
					Description: "",
					DenomUnits: []*banktypes.DenomUnit{
						{
							Denom:    "coin",
							Exponent: 0,
						},
					},
					Base:    "coin",
					Display: "coin",
					Name:    "coin",
					Symbol:  "coin",
				})
				suite.Require().NoError(err)
			},
			expectConvert: true,
		},
		{
			name: "ack error, token not registered",
			ack:  suite.utilsMockAcknowledgement(false),
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest()
			// get route callback
			module, _, err := suite.LorenzoChainA.App.GetIBCKeeper().PortKeeper.LookupModuleByPort(
				suite.LorenzoChainA.GetContext(), ibctransfertypes.ModuleName)
			suite.Require().NoError(err)

			cbs, ok := suite.LorenzoChainA.App.GetIBCKeeper().Router.GetRoute(module)
			suite.Require().True(ok)

			sender := suite.LorenzoChainA.SenderAccount.GetAddress()
			receiver := suite.LorenzoChainB.SenderAccount.GetAddress()

			// fund sender
			err = suite.chainA.BankKeeper.MintCoins(suite.LorenzoChainA.GetContext(),
				tokentypes.ModuleName, sdk.NewCoins(sdk.NewCoin("coin", sdk.NewInt(1000))))
			suite.Require().NoError(err)
			err = suite.chainA.BankKeeper.SendCoinsFromModuleToAccount(suite.LorenzoChainA.GetContext(),
				tokentypes.ModuleName, sender, sdk.NewCoins(sdk.NewCoin("coin", sdk.NewInt(1000))))
			suite.Require().NoError(err)

			if tc.malleate != nil {
				tc.malleate()
			}

			// lock coin
			msg := &ibctransfertypes.MsgTransfer{
				SourcePort:       suite.Path.EndpointA.ChannelConfig.PortID,
				SourceChannel:    suite.Path.EndpointA.ChannelID,
				Token:            sdk.NewCoin("coin", sdk.NewInt(1000)),
				Sender:           sender.String(),
				Receiver:         receiver.String(),
				TimeoutHeight:    clienttypes.NewHeight(0, 100),
				TimeoutTimestamp: 0,
				Memo:             "",
			}

			_, _ = suite.chainA.ICS20WrapperKeeper.Transfer(suite.LorenzoChainA.GetContext(), msg) // nolint:errcheck

			// get packet
			data := ibctransfertypes.FungibleTokenPacketData{
				Denom:    "coin",
				Amount:   "1000",
				Sender:   sender.String(),
				Receiver: receiver.String(),
				Memo:     "none",
			}
			packet := suite.NewMockTransferPacket(data.GetBytes())

			err = cbs.OnAcknowledgementPacket(suite.LorenzoChainA.GetContext(), packet, tc.ack, sender)
			if tc.invalidAck {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				if tc.expectConvert {
					id := suite.chainA.TokenKeeper.GetTokenPairIdByDenom(suite.LorenzoChainA.GetContext(), "coin")
					pair, found := suite.chainA.TokenKeeper.GetTokenPair(suite.LorenzoChainA.GetContext(), id)
					suite.Require().True(found)

					balance := suite.chainA.TokenKeeper.ERC20BalanceOf(
						suite.LorenzoChainA.GetContext(),
						erc20.ERC20MinterBurnerDecimalsContract.ABI,
						common.HexToAddress(pair.ContractAddress),
						common.BytesToAddress(sender.Bytes()),
					)
					suite.Require().Equal("1000", balance.String())
				}
			}
		})
	}
}
