package keeper_test

import (
	"github.com/Lorenzo-Protocol/lorenzo/x/agent/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestAgentQuery() {
	suite.SetupTest()
	wctx := sdk.WrapSDKContext(suite.ctx)

	response, err := suite.queryClient.Agent(wctx, &types.QueryAgentRequest{
		Id: 1,
	})
	suite.Require().NoError(err)
	suite.Require().NotNil(response)
}

func (suite *KeeperTestSuite) TestAgentsQuery() {
	suite.SetupTest()
	wctx := sdk.WrapSDKContext(suite.ctx)

	response, err := suite.queryClient.Agents(wctx, &types.QueryAgentsRequest{})
	suite.Require().NoError(err)
	suite.Require().NotNil(response)
}

func (suite *KeeperTestSuite) TestAdminQuery() {
	suite.SetupTest()
	wctx := sdk.WrapSDKContext(suite.ctx)

	response, err := suite.queryClient.Admin(wctx, &types.QueryAdminRequest{})
	suite.Require().NoError(err)
	suite.Require().NotNil(response)
}
