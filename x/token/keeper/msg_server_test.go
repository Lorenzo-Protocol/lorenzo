package keeper_test

import (
	"github.com/Lorenzo-Protocol/lorenzo/v2/app"
	"github.com/ethereum/go-ethereum/common"
	evmtypes "github.com/evmos/ethermint/x/evm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/Lorenzo-Protocol/lorenzo/v2/x/token/types"
)

const (
	// erc20 metadata
	erc20Name     = "Coin Token"
	erc20Symbol   = "COIN"
	erc20Decimals = uint8(18)

	// coin metadata
	coinBaseDenom    = "acoin"
	coinDisplay      = "coin"
	coinBaseExponent = uint32(0)
	coinMaxExponent  = uint32(18)
)

var (
	authority = authtypes.NewModuleAddress(govtypes.ModuleName)
	tester    = app.CreateTestAddrs(1)[0]
	tester2   = app.CreateTestAddrs(1)[0]

	coinMetadata = banktypes.Metadata{
		Description: "",
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    coinBaseDenom,
				Exponent: coinBaseExponent,
			},
			{
				Denom:    coinDisplay,
				Exponent: coinMaxExponent,
			},
		},
		Base:    coinBaseDenom,
		Display: coinDisplay,
		Name:    erc20Name,
		Symbol:  erc20Symbol,
	}
)

func (suite *KeeperTestSuite) TestRegisterCoin() {
	testcases := []struct {
		name       string
		expectPass bool
		sender     string
		malleate   func()
	}{
		{
			name:       "fail: sender is not authority",
			sender:     tester.String(),
			expectPass: false,
		},
		{
			name:       "fail: token module is disabled",
			sender:     authority.String(),
			expectPass: false,
			malleate: func() {
				suite.app.TokenKeeper.SetParams(suite.ctx, types.Params{
					EnableConversion: false,
					EnableEVMHook:    true,
				})
				suite.Commit()
			},
		},
		{
			name:       "fail: coin has no supply",
			sender:     authority.String(),
			expectPass: false,
			malleate: func() {
				err := suite.app.BankKeeper.BurnCoins(
					suite.ctx,
					types.ModuleName,
					sdk.NewCoins(sdk.NewInt64Coin(coinBaseDenom, 10000)),
				)
				suite.Require().NoError(err)
				suite.Commit()
			},
		},
		{
			name:       "fail: inconsistent metadata",
			sender:     authority.String(),
			expectPass: false,
			malleate: func() {
				coinMetadataCopy := coinMetadata
				coinMetadataCopy.Display = "coin2"
				suite.app.BankKeeper.SetDenomMetaData(suite.ctx, coinMetadataCopy)
				suite.Commit()
			},
		},
		{
			name:       "fail: coin already registered",
			sender:     authority.String(),
			expectPass: false,
			malleate: func() {
				_, err := suite.msgServer.RegisterCoin(suite.ctx, &types.MsgRegisterCoin{
					Authority: authority.String(),
					Metadata:  []banktypes.Metadata{coinMetadata},
				})
				suite.Require().NoError(err)
				suite.Commit()
			},
		},
		{
			name:       "success: register coin",
			sender:     authority.String(),
			expectPass: true,
		},
		// TODO: force evm to fail
	}

	for _, tc := range testcases {
		suite.Run(tc.name, func() {
			suite.SetupTest() // reset before each case

			// supply for coin
			err := suite.app.BankKeeper.MintCoins(
				suite.ctx,
				types.ModuleName,
				sdk.NewCoins(sdk.NewInt64Coin(coinBaseDenom, 10000)),
			)
			suite.Require().NoError(err)
			suite.Commit()

			if tc.malleate != nil {
				tc.malleate()
			}

			_, err = suite.msgServer.RegisterCoin(suite.ctx, &types.MsgRegisterCoin{
				Authority: tc.sender,
				Metadata:  []banktypes.Metadata{coinMetadata},
			})

			if tc.expectPass {
				suite.Require().NoError(err, tc.name)
				// metadata
				metadata, found := suite.app.BankKeeper.GetDenomMetaData(suite.ctx, coinMetadata.Base)
				suite.Require().True(found)
				suite.Require().Equal(coinMetadata, metadata)
				// token pair
				resp, err := suite.queryClient.TokenPair(suite.ctx, &types.QueryTokenPairRequest{Token: coinMetadata.Base})
				suite.Require().NoError(err)
				suite.Require().Equal(resp.TokenPair.Denom, coinMetadata.Base)
				// contract addr
				resp2, err := suite.queryClientEvm.Account(suite.ctx, &evmtypes.QueryAccountRequest{Address: resp.TokenPair.ContractAddress})
				suite.Require().NoError(err)
				suite.Require().NotNil(resp2.CodeHash)
			} else {
				suite.Require().Error(err, tc.name)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestRegisterERC20() {
	testcases := []struct {
		name           string
		sender         string
		malleate       func(string)
		malleateDeploy func() string
		expectPass     bool
	}{
		{
			name:       "fail: sender is not authority",
			expectPass: false,
			sender:     tester.String(),
		},
		{
			name:       "fail: token module is disabled",
			sender:     authority.String(),
			expectPass: false,
			malleate: func(_ string) {
				suite.app.TokenKeeper.SetParams(suite.ctx, types.Params{
					EnableConversion: false,
					EnableEVMHook:    true,
				})
				suite.Commit()
			},
		},
		{
			name:       "fail: contract not exist",
			sender:     authority.String(),
			expectPass: false,
			malleateDeploy: func() string {
				// don't deploy contract
				return "0x1D1530e3A3719BE0BEe1abba5016Cf2e236f3277"
			},
		},
		{
			name:       "fail: denom metadata already exist",
			sender:     authority.String(),
			expectPass: false,
			malleate: func(addr string) {
				// create erc20/addr
				coinMetadataCopy := coinMetadata
				coinMetadataCopy.Base = types.DenomPrefix + "/" + addr
				suite.app.BankKeeper.SetDenomMetaData(suite.ctx, coinMetadataCopy)
				suite.Commit()
			},
		},
		{
			name:       "fail: denom already registered",
			sender:     authority.String(),
			expectPass: false,
			malleate: func(addr string) {
				tokenPair := types.NewTokenPair(common.HexToAddress(addr), coinMetadata.Name, types.OWNER_MODULE)
				tokenPairId := tokenPair.GetID()

				denom := types.DenomPrefix + "/" + addr
				suite.app.TokenKeeper.SetTokenPairIdByDenom(suite.ctx, denom, tokenPairId)
			},
		},
		{
			name:       "fail: erc20 token metadata invalid",
			sender:     authority.String(),
			expectPass: false,
			malleateDeploy: func() string {
				// empty symbol!
				contractAddr := suite.utilsERC20Deploy(erc20Name, "", erc20Decimals)
				suite.T().Log(contractAddr.String())
				return contractAddr.String()
			},
		},
		{
			name:       "success: register erc20",
			expectPass: true,
			sender:     authority.String(),
		},
	}

	for _, tc := range testcases {
		suite.Run(tc.name, func() {
			suite.SetupTest() // reset before each case

			var contractAddr string
			if tc.malleateDeploy != nil {
				contractAddr = tc.malleateDeploy()
			} else {
				addr := suite.utilsERC20Deploy(erc20Name, erc20Symbol, erc20Decimals)
				contractAddr = addr.String()
			}

			if tc.malleate != nil {
				tc.malleate(contractAddr)
			}

			// register erc20
			_, err := suite.msgServer.RegisterERC20(suite.ctx, &types.MsgRegisterERC20{
				Authority:         tc.sender,
				ContractAddresses: []string{contractAddr},
			})

			// deploy contract
			if tc.expectPass {
				suite.Require().NoError(err, tc.name)
				// metadata
				denom := types.DenomPrefix + "/" + contractAddr
				metadata, found := suite.app.BankKeeper.GetDenomMetaData(suite.ctx, denom)
				suite.Require().True(found)
				suite.Require().Equal(denom, metadata.Base)
				suite.Require().Equal(denom, metadata.Name)
				suite.Require().Equal(erc20Symbol, metadata.Symbol)
				suite.Require().Equal(types.SanitizeERC20Name(erc20Name), metadata.Display)
				suite.Require().Equal(len(metadata.DenomUnits), 2)
				suite.Require().Equal(metadata.DenomUnits[0].Denom, denom)
				suite.Require().Equal(metadata.DenomUnits[0].Exponent, coinBaseExponent)
				suite.Require().Equal(metadata.DenomUnits[1].Denom, types.SanitizeERC20Name(erc20Name))
				suite.Require().Equal(metadata.DenomUnits[1].Exponent, uint32(erc20Decimals))

			} else {
				suite.Require().Error(err, tc.name)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestToggleConversion() {
	testcases := []struct {
		name       string
		sender     string
		expectPass bool
		malleate   func()
	}{
		{
			name:       "fail: sender is not authority",
			expectPass: false,
			sender:     tester.String(),
		},
		{
			name:       "fail: token module is disabled",
			expectPass: false,
			sender:     authority.String(),
			malleate: func() {
				suite.app.TokenKeeper.SetParams(suite.ctx, types.Params{
					EnableConversion: false,
				})
				suite.Commit()
			},
		},
		{
			name:       "fail: token pair not found",
			expectPass: false,
			sender:     authority.String(),
		},
		{
			name:       "success: toggle token pair conversion",
			expectPass: true,
			sender:     authority.String(),
			malleate: func() {
				err := suite.app.BankKeeper.MintCoins(
					suite.ctx,
					types.ModuleName,
					sdk.NewCoins(sdk.NewInt64Coin(coinBaseDenom, 10000)),
				)
				suite.Require().NoError(err)
				suite.Commit()

				_, err = suite.msgServer.RegisterCoin(suite.ctx, &types.MsgRegisterCoin{
					Authority: authority.String(),
					Metadata:  []banktypes.Metadata{coinMetadata},
				})
				suite.Require().NoError(err)
				suite.Commit()
			},
		},
	}

	for _, tc := range testcases {
		suite.Run(tc.name, func() {
			suite.SetupTest()

			if tc.malleate != nil {
				tc.malleate()
			}

			_, err := suite.msgServer.ToggleConversion(suite.ctx, &types.MsgToggleConversion{
				Authority: tc.sender,
				Token:     coinBaseDenom,
			})

			if tc.expectPass {
				suite.Require().NoError(err, tc.name)
				tokenPair, err := suite.queryClient.TokenPair(suite.ctx, &types.QueryTokenPairRequest{Token: coinBaseDenom})
				suite.Require().NoError(err)
				suite.Require().Equal(tokenPair.TokenPair.Enabled, false)
				suite.Require().Equal(tokenPair.TokenPair.Denom, coinBaseDenom)

			} else {
				suite.Require().Error(err, tc.name)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestUpdateParams() {
	testcases := []struct {
		name       string
		sender     string
		expectPass bool
	}{
		{
			name:       "fail: sender is not authority",
			expectPass: false,
			sender:     tester.String(),
		},
		{
			name:       "success: update params",
			expectPass: true,
			sender:     authority.String(),
		},
	}

	for _, tc := range testcases {
		suite.Run(tc.name, func() {
			suite.SetupTest()

			_, err := suite.msgServer.UpdateParams(suite.ctx, &types.MsgUpdateParams{
				Authority: tc.sender,
				Params: types.Params{
					EnableConversion: false,
					EnableEVMHook:    true,
				},
			})

			if tc.expectPass {
				suite.Require().NoError(err, tc.name)
			} else {
				suite.Require().Error(err, tc.name)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestConvertNativeCoinToVoucherERC20() {
	testcases := []struct {
		name               string
		sender             string
		expectPass         bool
		malleateMintEnable func()
		malleateMsg        func(coin *types.MsgConvertCoin)
	}{
		{
			name:       "fail: mint disabled due to token module disabled",
			expectPass: false,
			malleateMintEnable: func() {
				suite.app.TokenKeeper.SetParams(suite.ctx, types.Params{
					EnableConversion: false,
				})
				suite.Commit()
			},
		},
		{
			name:       "fail: mint disabled due to token pair is disabled",
			expectPass: false,
			malleateMintEnable: func() {
				_, err := suite.msgServer.ToggleConversion(suite.ctx, &types.MsgToggleConversion{
					Authority: authority.String(),
					Token:     coinBaseDenom,
				})
				suite.Require().NoError(err)
				suite.Commit()
			},
		},
		//{
		//	name:       "fail: mint disabled due to receiver is blocked",
		//	expectPass: false,
		//	malleate: func() {
		//		// TODO: set block addr when init bank keeper.
		//	},
		//},
		{
			name:       "fail: mint disabled due to receiver is not sender and coin is not allowed to send",
			expectPass: false,
			malleateMsg: func(coin *types.MsgConvertCoin) {
				coin.Receiver = common.BytesToAddress(authority).String()
			},
			malleateMintEnable: func() {
				suite.app.BankKeeper.SetSendEnabled(suite.ctx, coinBaseDenom, false)
				suite.Commit()
			},
		},
		{
			name:       "success: convert native coin to voucher erc20",
			expectPass: true,
		},
		{
			name:       "fail: native coin is not enough",
			expectPass: false,
			malleateMsg: func(coin *types.MsgConvertCoin) {
				coin.Coin = sdk.NewInt64Coin(coinMetadata.Base, 1_000_001)
			},
		},
	}

	for _, tc := range testcases {
		suite.Run(tc.name, func() {
			suite.SetupTest()
			sender := tester
			receiver := tester2

			pair := suite.utilsFundAndRegisterCoin(coinMetadata, sender, 1_000_000)

			msg := &types.MsgConvertCoin{
				Sender:   sender.String(),
				Receiver: common.BytesToAddress(receiver).String(),
				Coin:     sdk.NewInt64Coin(coinMetadata.Base, 1_000),
			}

			if tc.malleateMintEnable != nil {
				tc.malleateMintEnable()
			}

			if tc.malleateMsg != nil {
				tc.malleateMsg(msg)
			}

			_, err := suite.msgServer.ConvertCoin(suite.ctx, msg)

			if tc.expectPass {
				suite.Require().NoError(err)

				// check balance
				coin := suite.app.BankKeeper.GetBalance(suite.ctx, sender, coinMetadata.Base)
				suite.Require().Equal(int64(999_000), coin.Amount.Int64())

				erc20balance := suite.utilsERC20BalanceOf(common.HexToAddress(pair.ContractAddress), common.BytesToAddress(receiver))
				suite.Require().Equal(int64(1_000), erc20balance)

			} else {
				suite.Require().Error(err)
				suite.T().Log(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestConvertVoucherERC20ToNativeCoin() {
	testcases := []struct {
		name        string
		sender      string
		expectPass  bool
		malleateMsg func(msg *types.MsgConvertERC20)
	}{
		{
			name:       "success: convert voucher erc20 to native coin",
			expectPass: true,
		},
		{
			name:       "fail: voucher token is not enough",
			expectPass: false,
			malleateMsg: func(msg *types.MsgConvertERC20) {
				msg.Amount = sdk.NewInt(1_000_001)
			},
		},
	}

	for _, tc := range testcases {
		suite.Run(tc.name, func() {
			suite.SetupTest()
			// convert all native coin to voucher erc20
			pair := suite.utilsFundAndRegisterCoin(coinMetadata, tester, 1_000_000)
			_, err := suite.msgServer.ConvertCoin(suite.ctx, &types.MsgConvertCoin{
				Sender:   tester.String(),
				Receiver: common.BytesToAddress(tester).String(),
				Coin:     sdk.NewInt64Coin(coinMetadata.Base, 1_000_000),
			})
			suite.Require().NoError(err)
			coin := suite.app.BankKeeper.GetBalance(suite.ctx, tester, coinMetadata.Base)
			suite.Require().Equal(int64(0), coin.Amount.Int64())
			erc20balance := suite.utilsERC20BalanceOf(common.HexToAddress(pair.ContractAddress), common.BytesToAddress(tester))
			suite.Require().Equal(int64(1_000_000), erc20balance)

			msg := &types.MsgConvertERC20{
				Sender:          common.BytesToAddress(tester).String(),
				Receiver:        tester.String(),
				ContractAddress: pair.ContractAddress,
				Amount:          sdk.NewInt(1_000),
			}

			if tc.malleateMsg != nil {
				tc.malleateMsg(msg)
			}

			_, err = suite.msgServer.ConvertERC20(suite.ctx, msg)

			if tc.expectPass {
				suite.Require().NoError(err)

				// check balance
				erc20balance := suite.utilsERC20BalanceOf(common.HexToAddress(pair.ContractAddress), common.BytesToAddress(tester))
				suite.Require().Equal(int64(999_000), erc20balance)

				coin2 := suite.app.BankKeeper.GetBalance(suite.ctx, tester, coinMetadata.Base)
				suite.Require().Equal(int64(1_000), coin2.Amount.Int64())

				supply := suite.app.BankKeeper.GetSupply(suite.ctx, coinMetadata.Base)
				suite.Require().Equal(int64(1_000_000), supply.Amount.Int64())

			} else {
				suite.Require().Error(err)
				suite.T().Log(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestConvertNativeERC20ToVoucherCoin() {
	testcases := []struct {
		name               string
		sender             string
		expectPass         bool
		malleateMintEnable func(pair types.TokenPair)
		malleateMsg        func(coin *types.MsgConvertERC20)
	}{
		{
			name:       "fail: mint disabled due to token module disabled",
			expectPass: false,
			malleateMintEnable: func(pair types.TokenPair) {
				suite.app.TokenKeeper.SetParams(suite.ctx, types.Params{
					EnableConversion: false,
				})
				suite.Commit()
			},
		},
		{
			name:       "fail: mint disabled due to token pair is disabled",
			expectPass: false,
			malleateMintEnable: func(pair types.TokenPair) {
				_, err := suite.msgServer.ToggleConversion(suite.ctx, &types.MsgToggleConversion{
					Authority: authority.String(),
					Token:     pair.Denom,
				})
				suite.Require().NoError(err)
				suite.Commit()
			},
		},
		//{
		//	name:       "fail: mint disabled due to receiver is blocked",
		//	expectPass: false,
		//	malleate: func() {
		//		// TODO: set block addr when init bank keeper.
		//	},
		//},
		{
			name:       "fail: mint disabled due to receiver is not sender and coin is not allowed to send",
			expectPass: false,
			malleateMsg: func(msg *types.MsgConvertERC20) {
				msg.Receiver = authority.String()
			},
			malleateMintEnable: func(pair types.TokenPair) {
				suite.app.BankKeeper.SetSendEnabled(suite.ctx, pair.Denom, false)
				suite.Commit()
			},
		},
		{
			name:       "success: convert native erc20 to voucher coin",
			expectPass: true,
		},
		{
			name:       "fail: native erc20 is not enough",
			expectPass: false,
			malleateMsg: func(msg *types.MsgConvertERC20) {
				msg.Amount = sdk.NewInt(1_000_001)
			},
		},
	}

	for _, tc := range testcases {
		suite.Run(tc.name, func() {
			suite.SetupTest()
			// deploy erc20, fund tester, and register pair for erc20
			pair, contractAddr := suite.utilsFundAndRegisterERC20(
				erc20Name, erc20Symbol, erc20Decimals,
				common.BytesToAddress(tester), 1_000_000)

			msg := &types.MsgConvertERC20{
				ContractAddress: contractAddr.String(),
				Amount:          sdk.NewInt(1_000),
				Receiver:        tester.String(),
				Sender:          common.BytesToAddress(tester).String(),
			}

			if tc.malleateMintEnable != nil {
				tc.malleateMintEnable(pair)
			}

			if tc.malleateMsg != nil {
				tc.malleateMsg(msg)
			}

			// set account
			suite.app.AccountKeeper.SetAccount(suite.ctx,
				suite.app.AccountKeeper.NewAccountWithAddress(suite.ctx, tester))

			_, err := suite.msgServer.ConvertERC20(suite.ctx, msg)

			if tc.expectPass {
				suite.T().Log(tester.String())
				suite.Require().NoError(err)
				// query balance
				token := suite.utilsERC20BalanceOf(contractAddr, common.BytesToAddress(tester))
				suite.Require().Equal(int64(999_000), token)
				coin := suite.app.BankKeeper.GetBalance(suite.ctx, tester, pair.Denom)
				suite.Require().Equal(int64(1_000), coin.Amount.Int64())
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestConvertVoucherCoinToNativeERC20() {
	testcases := []struct {
		name        string
		sender      string
		expectPass  bool
		malleateMsg func(msg *types.MsgConvertCoin)
	}{
		{
			name:       "success: convert native erc20 to voucher coin",
			expectPass: true,
		},
		{
			name:       "fail: native erc20 is not enough",
			expectPass: false,
			malleateMsg: func(msg *types.MsgConvertCoin) {
				msg.Coin.Amount = sdk.NewInt(1_000_001)
			},
		},
	}

	for _, tc := range testcases {
		suite.Run(tc.name, func() {
			suite.SetupTest()

			// deploy erc20, fund tester, and register pair for erc20
			pair, contractAddr := suite.utilsFundAndRegisterERC20(
				erc20Name, erc20Symbol, erc20Decimals,
				common.BytesToAddress(tester), 1_000_000)

			// set account
			suite.app.AccountKeeper.SetAccount(suite.ctx,
				suite.app.AccountKeeper.NewAccountWithAddress(suite.ctx, tester))

			// convert to voucher coin first.
			_, err := suite.msgServer.ConvertERC20(suite.ctx, &types.MsgConvertERC20{
				ContractAddress: contractAddr.String(),
				Amount:          sdk.NewInt(1_000_000),
				Receiver:        tester.String(),
				Sender:          common.BytesToAddress(tester).String(),
			})
			suite.Require().NoError(err)

			msg := &types.MsgConvertCoin{
				Sender:   tester.String(),
				Receiver: common.BytesToAddress(tester).String(),
				Coin:     sdk.NewInt64Coin(pair.Denom, 1_000),
			}

			if tc.malleateMsg != nil {
				tc.malleateMsg(msg)
			}

			_, err = suite.msgServer.ConvertCoin(suite.ctx, msg)

			if tc.expectPass {
				suite.Require().NoError(err)
				// query balance
				coin := suite.app.BankKeeper.GetBalance(suite.ctx, tester, pair.Denom)
				suite.Require().Equal(int64(999_000), coin.Amount.Int64())
				token := suite.utilsERC20BalanceOf(contractAddr, common.BytesToAddress(tester))
				suite.Require().Equal(int64(1_000), token)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}
