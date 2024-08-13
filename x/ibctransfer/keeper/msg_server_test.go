package keeper_test

import (
	"fmt"

	"github.com/stretchr/testify/mock"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v7/modules/core/24-host"

	"github.com/Lorenzo-Protocol/lorenzo/v3/x/ibctransfer/keeper"
	tokentypes "github.com/Lorenzo-Protocol/lorenzo/v3/x/token/types"
)

func (suite *KeeperTestSuite) TestTransfer() {
	mockChannelKeeper := &MockChannelKeeper{}
	mockICS4Wrapper := &MockICS4Wrapper{}

	mockChannelKeeper.On("GetNextSequenceSend",
		mock.Anything, mock.Anything, mock.Anything).
		Return(1, true)
	mockChannelKeeper.On("GetChannel",
		mock.Anything, mock.Anything, mock.Anything).
		Return(channeltypes.Channel{Counterparty: channeltypes.NewCounterparty("transfer", "channel-1")}, true)
	mockICS4Wrapper.On("SendPacket",
		mock.Anything, mock.Anything, mock.Anything).Return(nil)

	testCases := []struct {
		name     string
		malleate func() *types.MsgTransfer
		expPass  bool
	}{
		{
			"success: native transfer",
			func() *types.MsgTransfer {
				senderAcc := sdk.AccAddress(suite.address.Bytes())

				// fund coins
				coins := sdk.NewCoins(sdk.NewCoin("coin", sdk.NewInt(10)))
				err := suite.app.BankKeeper.MintCoins(suite.ctx, tokentypes.ModuleName, coins)
				suite.Require().NoError(err)
				err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, tokentypes.ModuleName, senderAcc, coins)
				suite.Require().NoError(err)
				suite.Commit()

				transferMsg := types.NewMsgTransfer(
					"transfer",
					"channel-0",
					sdk.NewCoin("coin", sdk.NewInt(10)), senderAcc.String(),
					"", timeoutHeight, 0, "")

				return transferMsg
			},
			true,
		},
		{
			"fail: invalid sender",
			func() *types.MsgTransfer {
				contractAddr := suite.utilsERC20Deploy("coin", "token", uint8(6))
				suite.Commit()

				transferMsg := types.NewMsgTransfer("transfer", "channel-0",
					sdk.NewCoin("erc20/"+contractAddr.String(), sdk.NewInt(10)), "",
					"", timeoutHeight, 0, "")
				return transferMsg
			},
			false,
		},
		{
			"success: with token module disabled, sufficient balance",
			func() *types.MsgTransfer {
				contractAddr := suite.utilsERC20Deploy("coin", "token", uint8(6))
				suite.Commit()

				pair, err := suite.app.TokenKeeper.RegisterERC20(suite.ctx, contractAddr)
				suite.Require().NoError(err)
				suite.Commit()

				senderAcc := sdk.AccAddress(suite.address.Bytes())
				suite.utilsERC20Mint(contractAddr, tokentypes.ModuleAddress, suite.address, 10)
				suite.Commit()

				coin := sdk.NewCoin(pair.Denom, sdk.NewInt(10))
				coins := sdk.NewCoins(coin)

				err = suite.app.BankKeeper.MintCoins(suite.ctx, tokentypes.ModuleName, coins)
				suite.Require().NoError(err)
				suite.Commit()

				err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, tokentypes.ModuleName, senderAcc, coins)
				suite.Require().NoError(err)
				suite.Commit()

				params := suite.app.TokenKeeper.GetParams(suite.ctx)
				params.EnableConversion = false
				suite.app.TokenKeeper.SetParams(suite.ctx, params)
				suite.Commit()

				transferMsg := types.NewMsgTransfer("transfer", "channel-0",
					sdk.NewCoin(pair.Denom, sdk.NewInt(10)), senderAcc.String(), "",
					timeoutHeight, 0, "")

				return transferMsg
			},
			true,
		},
		{
			"fail: token module disabled, insufficient balance",
			func() *types.MsgTransfer {
				contractAddr := suite.utilsERC20Deploy("coin", "token", uint8(6))
				suite.Commit()

				pair, err := suite.app.TokenKeeper.RegisterERC20(suite.ctx, contractAddr)
				suite.Require().NoError(err)
				suite.Commit()

				senderAcc := sdk.AccAddress(suite.address.Bytes())
				suite.utilsERC20Mint(contractAddr, tokentypes.ModuleAddress, suite.address, 10)
				suite.Commit()

				params := suite.app.TokenKeeper.GetParams(suite.ctx)
				params.EnableConversion = false
				suite.app.TokenKeeper.SetParams(suite.ctx, params)
				suite.Commit()

				transferMsg := types.NewMsgTransfer("transfer", "channel-0",
					sdk.NewCoin(pair.Denom, sdk.NewInt(10)), senderAcc.String(), "",
					timeoutHeight, 0, "")

				return transferMsg
			},
			false,
		},
		{
			"success: native coin transfer",
			func() *types.MsgTransfer {
				senderAcc := sdk.AccAddress(suite.address.Bytes())

				coin := sdk.NewCoin("test", sdk.NewInt(10))
				coins := sdk.NewCoins(coin)

				err := suite.app.BankKeeper.MintCoins(suite.ctx, tokentypes.ModuleName, coins)
				suite.Require().NoError(err)

				err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, tokentypes.ModuleName, senderAcc, coins)
				suite.Require().NoError(err)
				suite.Commit()

				transferMsg := types.NewMsgTransfer("transfer", "channel-0",
					coin, senderAcc.String(), "", timeoutHeight, 0, "")

				return transferMsg
			},
			true,
		},
		{
			"success: token pair disabled",
			func() *types.MsgTransfer {
				contractAddr := suite.utilsERC20Deploy("coin", "token", uint8(6))
				suite.Commit()

				pair, err := suite.app.TokenKeeper.RegisterERC20(suite.ctx, contractAddr)
				suite.Require().NoError(err)
				pair.Enabled = false
				suite.app.TokenKeeper.SetTokenPair(suite.ctx, *pair)

				coin := sdk.NewCoin(pair.Denom, sdk.NewInt(10))
				senderAcc := sdk.AccAddress(suite.address.Bytes())
				transferMsg := types.NewMsgTransfer("transfer", "channel-0",
					coin, senderAcc.String(), "", timeoutHeight, 0, "")

				// mint coins to perform the regular transfer without conversions
				err = suite.app.BankKeeper.MintCoins(suite.ctx, tokentypes.ModuleName, sdk.NewCoins(coin))
				suite.Require().NoError(err)

				err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, tokentypes.ModuleName, senderAcc, sdk.NewCoins(coin))
				suite.Require().NoError(err)
				suite.Commit()

				return transferMsg
			},
			true,
		},
		{
			"success: convert happens",
			func() *types.MsgTransfer {
				contractAddr := suite.utilsERC20Deploy("coin", "token", uint8(6))
				suite.Commit()

				pair, err := suite.app.TokenKeeper.RegisterERC20(suite.ctx, contractAddr)
				suite.Require().NoError(err)
				suite.Commit()

				suite.utilsERC20Mint(contractAddr, tokentypes.ModuleAddress, suite.address, 10)
				suite.Commit()

				senderAcc := sdk.AccAddress(suite.address.Bytes())
				coins := sdk.NewCoins(sdk.NewCoin(pair.Denom, sdk.NewInt(1)))
				err = suite.app.BankKeeper.MintCoins(suite.ctx, tokentypes.ModuleName, coins)
				suite.Require().NoError(err)
				err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, tokentypes.ModuleName, senderAcc, coins)
				suite.Require().NoError(err)
				suite.Commit()

				transferMsg := types.NewMsgTransfer("transfer", "channel-0",
					sdk.NewCoin(pair.Denom, sdk.NewInt(11)), senderAcc.String(),
					"", timeoutHeight, 0, "")

				return transferMsg
			},
			true,
		},
		{
			"success: has enough balance in coins",
			func() *types.MsgTransfer {
				contractAddr := suite.utilsERC20Deploy("coin", "token", uint8(6))
				suite.Commit()

				pair, err := suite.app.TokenKeeper.RegisterERC20(suite.ctx, contractAddr)
				suite.Require().NoError(err)
				suite.Commit()

				senderAcc := sdk.AccAddress(suite.address.Bytes())

				coins := sdk.NewCoins(sdk.NewCoin(pair.Denom, sdk.NewInt(10)))
				err = suite.app.BankKeeper.MintCoins(suite.ctx, tokentypes.ModuleName, coins)
				suite.Require().NoError(err)
				err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, tokentypes.ModuleName, senderAcc, coins)
				suite.Require().NoError(err)
				suite.Commit()

				transferMsg := types.NewMsgTransfer("transfer", "channel-0",
					sdk.NewCoin(pair.Denom, sdk.NewInt(10)), senderAcc.String(),
					"", timeoutHeight, 0, "")

				return transferMsg
			},
			true,
		},
		{
			"fail: balance not enough",
			func() *types.MsgTransfer {
				contractAddr := suite.utilsERC20Deploy("coin", "token", uint8(6))
				suite.Commit()

				pair, err := suite.app.TokenKeeper.RegisterERC20(suite.ctx, contractAddr)
				suite.Require().NoError(err)
				suite.Commit()

				senderAcc := sdk.AccAddress(suite.address.Bytes())
				transferMsg := types.NewMsgTransfer("transfer", "channel-0",
					sdk.NewCoin(pair.Denom, sdk.NewInt(10)), senderAcc.String(), "",
					timeoutHeight, 0, "")

				return transferMsg
			},
			false,
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.SetupTest()

			_, err := suite.app.ScopedTransferKeeper.NewCapability(suite.ctx,
				host.ChannelCapabilityPath("transfer", "channel-0"))
			suite.Require().NoError(err)

			suite.app.ICS20WrapperKeeper = keeper.NewKeeper(
				suite.app.AppCodec(), suite.app.GetKey(types.StoreKey), suite.app.GetSubspace(types.ModuleName),
				&MockICS4Wrapper{}, mockChannelKeeper, &suite.app.IBCKeeper.PortKeeper,
				suite.app.AccountKeeper, suite.app.BankKeeper, suite.app.ScopedTransferKeeper,
				suite.app.TokenKeeper)

			msg := tc.malleate()

			_, err = suite.app.ICS20WrapperKeeper.Transfer(sdk.WrapSDKContext(suite.ctx), msg)
			if tc.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
				suite.T().Log(err)
			}
		})
	}
}
