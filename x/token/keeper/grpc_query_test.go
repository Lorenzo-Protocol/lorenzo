package keeper_test

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/ethereum/go-ethereum/common"

	"github.com/Lorenzo-Protocol/lorenzo/v2/x/token/types"
)

func (suite *KeeperTestSuite) TestQueryParams() {
	resp, err := suite.queryClient.Params(suite.ctx, &types.QueryParamsRequest{})
	suite.Require().NoError(err)
	suite.Require().NotNil(resp)

	actual := suite.app.TokenKeeper.GetParams(suite.ctx)
	suite.Require().Equal(resp.Params.EnableEVMHook, actual.EnableEVMHook)
	suite.Require().Equal(resp.Params.EnableConversion, actual.EnableConversion)
}

func (suite *KeeperTestSuite) TestQueryTokenPair() {
	pair := suite.utilsFundAndRegisterCoin(coinMetadata, tester, 10000)

	testCases := []struct {
		name       string
		expectPass bool
		token      string
	}{
		{
			name:       "token pair found by denom",
			expectPass: true,
			token:      pair.Denom,
		},
		{
			name:       "token pair found by erc20 address",
			expectPass: true,
			token:      pair.ContractAddress,
		},
		{
			name:       "token pair not found",
			expectPass: false,
			token:      "any-token",
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			// register a toke pair
			resp, err := suite.queryClient.TokenPair(suite.ctx, &types.QueryTokenPairRequest{
				Token: tc.token,
			})
			if tc.expectPass {
				suite.Require().NoError(err)
				suite.Require().NotNil(resp)
				suite.Require().Equal(pair.GetID(), resp.TokenPair.GetID())
				suite.Require().Equal(pair.ContractAddress, resp.TokenPair.ContractAddress)
				suite.Require().Equal(pair.Denom, resp.TokenPair.Denom)
			} else {
				suite.Require().Error(err)
				suite.Require().Nil(resp)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryTokenPairs() {
	pairs := make(map[string]types.TokenPair)
	for i := 1; i <= 10; i++ {
		metadata := banktypes.Metadata{
			Description: "",
			DenomUnits: []*banktypes.DenomUnit{
				{
					Denom:    coinBaseDenom + strconv.Itoa(i),
					Exponent: coinBaseExponent,
				},
				{
					Denom:    coinDisplay + strconv.Itoa(i),
					Exponent: coinMaxExponent,
				},
			},
			Base:    coinBaseDenom + strconv.Itoa(i),
			Display: coinDisplay,
			Name:    erc20Name,
			Symbol:  erc20Symbol,
		}
		pair := suite.utilsFundAndRegisterCoin(metadata, tester, 10000)
		pairs[string(pair.GetID())] = pair
	}

	resp, err := suite.queryClient.TokenPairs(suite.ctx, &types.QueryTokenPairsRequest{})
	suite.Require().NoError(err)
	suite.Require().NotNil(resp)

	for _, pair := range resp.TokenPairs {
		suite.Require().Equal(pairs[string(pair.GetID())].ContractAddress, pair.ContractAddress)
		suite.Require().Equal(pairs[string(pair.GetID())].Denom, pair.Denom)
	}

	suite.Require().Equal(len(pairs), len(resp.TokenPairs))
}

func (suite *KeeperTestSuite) TestQueryBalance() {
	// register a toke pair and convert some coin to erc20
	pair := suite.utilsFundAndRegisterCoin(coinMetadata, tester, 10000)
	_, err := suite.msgServer.ConvertCoin(suite.ctx, &types.MsgConvertCoin{
		Coin:     sdk.NewCoin(pair.Denom, sdk.NewInt(1000)),
		Receiver: common.BytesToAddress(tester.Bytes()).String(),
		Sender:   tester.String(),
	})
	suite.Require().NoError(err)
	suite.Commit()

	testCases := []struct {
		name       string
		token      string
		address    string
		expectPass bool
	}{
		{
			name:       "success: by coin & bech32 address",
			expectPass: true,
			token:      pair.Denom,
			address:    tester.String(),
		},
		{
			name:       "success: by coin & hex address",
			expectPass: true,
			token:      pair.Denom,
			address:    common.BytesToAddress(tester.Bytes()).String(),
		},
		{
			name:       "success: by erc20 address & bech32 address",
			expectPass: true,
			token:      pair.ContractAddress,
			address:    tester.String(),
		},
		{
			name:       "success: by erc20 address & hex account address",
			expectPass: true,
			token:      pair.ContractAddress,
			address:    common.BytesToAddress(tester.Bytes()).String(),
		},
		{
			name:       "fail: invalid token",
			expectPass: false,
			token:      "0123456",
			address:    tester.String(),
		},
		{
			name:       "fail: invalid address",
			expectPass: false,
			token:      pair.Denom,
			address:    "invalid",
		},
		{
			name:       "fail: denom not found",
			expectPass: false,
			token:      "unknown",
			address:    tester.String(),
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			resp, err := suite.queryClient.Balance(suite.ctx, &types.QueryBalanceRequest{
				Token:          tc.token,
				AccountAddress: tc.address,
			})
			if tc.expectPass {
				suite.Require().NoError(err)
				suite.Require().NotNil(resp)
				suite.Require().Equal(pair.ContractAddress, resp.Erc20Address)
				suite.Require().Equal("1000", resp.Erc20TokenAmount)
				suite.Require().Equal(pair.Denom, resp.Coin.Denom)
				suite.Require().Equal(int64(9000), resp.Coin.Amount.Int64())
			} else {
				suite.Require().Error(err)
				suite.Require().Nil(resp)
			}
		})
	}
}
