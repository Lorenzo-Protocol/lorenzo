package keeper_test

import (
	"strconv"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

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
