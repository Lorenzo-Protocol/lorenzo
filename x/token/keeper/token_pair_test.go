package keeper_test

import (
	utiltx "github.com/Lorenzo-Protocol/lorenzo/testutil/tx"
	"github.com/Lorenzo-Protocol/lorenzo/x/token/types"
	"github.com/ethereum/go-ethereum/common"
)

func (suite *KeeperTestSuite) TestRemoveTokenPair() {
	contract := utiltx.GenerateAddress()
	pair := types.NewTokenPair(contract, "coin", types.OWNER_MODULE)
	suite.app.TokenKeeper.SetTokenPair(suite.ctx, pair)
	suite.app.TokenKeeper.SetTokenPairIdByDenom(suite.ctx, pair.Denom, pair.GetID())
	suite.app.TokenKeeper.SetTokenPairIdByERC20(suite.ctx, contract, pair.GetID())

	// before remove
	res, found := suite.app.TokenKeeper.GetTokenPair(suite.ctx, pair.GetID())
	suite.Require().True(found)
	suite.Require().Equal(pair, res)
	id := suite.app.TokenKeeper.GetTokenPairIdByDenom(suite.ctx, pair.Denom)
	suite.Require().Equal(pair.GetID(), id)
	id = suite.app.TokenKeeper.GetTokenPairIdByERC20(suite.ctx, contract)
	suite.Require().Equal(pair.GetID(), id)

	// after remove
	suite.app.TokenKeeper.RemoveTokenPair(suite.ctx, pair)
	_, found = suite.app.TokenKeeper.GetTokenPair(suite.ctx, pair.GetID())
	suite.Require().False(found)
	id = suite.app.TokenKeeper.GetTokenPairIdByDenom(suite.ctx, pair.Denom)
	suite.Require().Nil(id)
	id = suite.app.TokenKeeper.GetTokenPairIdByERC20(suite.ctx, contract)
	suite.Require().Nil(id)
}

func (suite *KeeperTestSuite) TestIsRegisteredByDenom() {
	contract := utiltx.GenerateAddress()
	pair := types.NewTokenPair(contract, coinBaseDenom, types.OWNER_MODULE)
	suite.app.TokenKeeper.SetTokenPair(suite.ctx, pair)
	suite.app.TokenKeeper.SetTokenPairIdByDenom(suite.ctx, pair.Denom, pair.GetID())
	suite.app.TokenKeeper.SetTokenPairIdByERC20(suite.ctx, contract, pair.GetID())

	testCases := []struct {
		name             string
		denom            string
		expectRegistered bool
	}{
		{"not registered", "ccoin", false},
		{"registered", coinBaseDenom, true},
	}

	for _, tc := range testCases {
		registered := suite.app.TokenKeeper.IsRegisteredByDenom(suite.ctx, tc.denom)
		if tc.expectRegistered {
			suite.Require().True(registered)
		} else {
			suite.Require().False(registered)
		}
	}
}

func (suite *KeeperTestSuite) TestIsRegisteredByERC20() {
	contract := utiltx.GenerateAddress()
	pair := types.NewTokenPair(contract, coinBaseDenom, types.OWNER_MODULE)
	suite.app.TokenKeeper.SetTokenPair(suite.ctx, pair)
	suite.app.TokenKeeper.SetTokenPairIdByDenom(suite.ctx, pair.Denom, pair.GetID())
	suite.app.TokenKeeper.SetTokenPairIdByERC20(suite.ctx, contract, pair.GetID())

	testCases := []struct {
		name             string
		contract         common.Address
		expectRegistered bool
	}{
		{"not registered", utiltx.GenerateAddress(), false},
		{"registered", contract, true},
	}

	for _, tc := range testCases {
		registered := suite.app.TokenKeeper.IsRegisteredByERC20(suite.ctx, tc.contract)
		if tc.expectRegistered {
			suite.Require().True(registered)
		} else {
			suite.Require().False(registered)
		}
	}
}

func (suite *KeeperTestSuite) TestGetTokenPairs() {
	var pairs []types.TokenPair

	testCases := []struct {
		name     string
		malleate func()
	}{
		{
			"no pair registered", func() {
				pairs = []types.TokenPair{}
			},
		},
		{
			"1 pair registered",
			func() {
				pair := types.NewTokenPair(utiltx.GenerateAddress(), "coin", types.OWNER_MODULE)
				suite.app.TokenKeeper.SetTokenPair(suite.ctx, pair)
				pairs = []types.TokenPair{pair}
			},
		},
		{
			"2 pairs registered",
			func() {
				pair := types.NewTokenPair(utiltx.GenerateAddress(), "coin", types.OWNER_MODULE)
				pair2 := types.NewTokenPair(utiltx.GenerateAddress(), "coin2", types.OWNER_MODULE)
				suite.app.TokenKeeper.SetTokenPair(suite.ctx, pair)
				suite.app.TokenKeeper.SetTokenPair(suite.ctx, pair2)
				pairs = []types.TokenPair{pair, pair2}
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest() // reset
			tc.malleate()
			res := suite.app.TokenKeeper.GetTokenPairs(suite.ctx)
			suite.Require().ElementsMatch(pairs, res, tc.name)
		})
	}
}

func (suite *KeeperTestSuite) TestGetTokenPairId() {
	contract := utiltx.GenerateAddress()
	pair := types.NewTokenPair(contract, coinBaseDenom, types.OWNER_MODULE)
	suite.app.TokenKeeper.SetTokenPair(suite.ctx, pair)
	suite.app.TokenKeeper.SetTokenPairIdByDenom(suite.ctx, pair.Denom, pair.GetID())
	suite.app.TokenKeeper.SetTokenPairIdByERC20(suite.ctx, contract, pair.GetID())

	testCases := []struct {
		name  string
		token string
		expID []byte
	}{
		{"nil token", "", nil},
		{"by denom", coinBaseDenom, pair.GetID()},
		{"by contract", contract.String(), pair.GetID()},
	}

	for _, tc := range testCases {
		id := suite.app.TokenKeeper.GetTokenPairId(suite.ctx, tc.token)
		if id != nil {
			suite.Require().Equal(tc.expID, id)
		} else {
			suite.Require().Nil(id)
		}
	}
}

func (suite *KeeperTestSuite) TestGetTokenPair() {
	contract := utiltx.GenerateAddress()
	pair := types.NewTokenPair(contract, coinBaseDenom, types.OWNER_MODULE)
	suite.app.TokenKeeper.SetTokenPair(suite.ctx, pair)
	suite.app.TokenKeeper.SetTokenPairIdByDenom(suite.ctx, pair.Denom, pair.GetID())
	suite.app.TokenKeeper.SetTokenPairIdByERC20(suite.ctx, contract, pair.GetID())

	testCases := []struct {
		name        string
		id          string
		expectFound bool
	}{
		{"unknown pair id", "whatever", false},
		{"registered token", string(pair.GetID()), true},
	}

	for _, tc := range testCases {
		res, found := suite.app.TokenKeeper.GetTokenPair(suite.ctx, []byte(tc.id))
		if tc.expectFound {
			suite.Require().True(found)
			suite.Require().Equal(pair, res)
		} else {
			suite.Require().False(found)
		}
	}
}

func (suite *KeeperTestSuite) TestGetTokenPairIdByERC20() {
	contract := utiltx.GenerateAddress()
	pair := types.NewTokenPair(contract, coinBaseDenom, types.OWNER_MODULE)
	suite.app.TokenKeeper.SetTokenPair(suite.ctx, pair)
	suite.app.TokenKeeper.SetTokenPairIdByDenom(suite.ctx, pair.Denom, pair.GetID())
	suite.app.TokenKeeper.SetTokenPairIdByERC20(suite.ctx, contract, pair.GetID())

	testCases := []struct {
		name     string
		contract common.Address
	}{
		{"by contract", contract},
	}

	for _, tc := range testCases {
		id := suite.app.TokenKeeper.GetTokenPairIdByERC20(suite.ctx, tc.contract)
		suite.Require().Equal(pair.GetID(), id)
	}
}

func (suite *KeeperTestSuite) TestGetTokenPairIdByDenom() {
	contract := utiltx.GenerateAddress()
	pair := types.NewTokenPair(contract, coinBaseDenom, types.OWNER_MODULE)
	suite.app.TokenKeeper.SetTokenPair(suite.ctx, pair)
	suite.app.TokenKeeper.SetTokenPairIdByDenom(suite.ctx, pair.Denom, pair.GetID())
	suite.app.TokenKeeper.SetTokenPairIdByERC20(suite.ctx, contract, pair.GetID())

	testCases := []struct {
		name  string
		denom string
	}{
		{"by denom", coinBaseDenom},
	}

	for _, tc := range testCases {
		id := suite.app.TokenKeeper.GetTokenPairIdByDenom(suite.ctx, tc.denom)
		suite.Require().Equal(pair.GetID(), id)
	}
}
