package keeper_test

import (
	"fmt"

	"github.com/Lorenzo-Protocol/lorenzo/v3/x/agent/types"
)

func (suite *KeeperTestSuite) TestParamsQuery() {
	testCases := []struct {
		name       string
		request    *types.QueryParamsRequest
		malleate   func()
		validation func()
		expectErr  bool
	}{
		{
			name:      "success",
			request:   &types.QueryParamsRequest{},
			expectErr: false,
			malleate: func() {
				suite.Commit()
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("ParamsQuery - %s", tc.name), func() {
			suite.SetupTest()
			if tc.malleate != nil {
				tc.malleate()
			}
			_, err := suite.queryClient.Params(suite.ctx, tc.request)
			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
			if tc.validation != nil {
				tc.validation()
			}
		})
	}
}

func (suite *KeeperTestSuite) TestAgentsQuery() {
	testCases := []struct {
		name       string
		request    *types.QueryAgentsRequest
		malleate   func()
		validation func()
		expectErr  bool
	}{
		{
			name:      "success",
			request:   &types.QueryAgentsRequest{},
			expectErr: false,
			malleate: func() {
				suite.Commit()
				// create agent
				name := "sinohope4"
				btcReceivingAddress := "3C7VPws9fMW3kcwRJvMkSVdqMs4SAhQCqq"
				ethAddr := "0x6508d68f4e5931f93fadc3b7afac5092e195b80f"
				description := "lorenzo"
				url := "https://lorenzo-protocol.io"
				agentId := suite.lorenzoApp.AgentKeeper.AddAgent(
					suite.ctx,
					name, btcReceivingAddress, ethAddr, description, url)
				suite.Require().NotEqual(agentId, 0)
				suite.Require().Equal(agentId, uint64(1))
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("PlansQuery - %s", tc.name), func() {
			suite.SetupTest()
			if tc.malleate != nil {
				tc.malleate()
			}
			_, err := suite.queryClient.Agents(suite.ctx, tc.request)
			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
			if tc.validation != nil {
				tc.validation()
			}
		})
	}
}

func (suite *KeeperTestSuite) TestAgentQuery() {
	testCases := []struct {
		name       string
		request    *types.QueryAgentRequest
		malleate   func()
		validation func()
		expectErr  bool
	}{
		{
			name: "success",
			request: &types.QueryAgentRequest{
				Id: 1,
			},
			expectErr: false,
			malleate: func() {
				suite.Commit()
				// create agent
				name := "sinohope4"
				btcReceivingAddress := "3C7VPws9fMW3kcwRJvMkSVdqMs4SAhQCqq"
				ethAddr := "0x6508d68f4e5931f93fadc3b7afac5092e195b80f"
				description := "lorenzo"
				url := "https://lorenzo-protocol.io"
				agentId := suite.lorenzoApp.AgentKeeper.AddAgent(
					suite.ctx,
					name, btcReceivingAddress, ethAddr, description, url)
				suite.Require().NotEqual(agentId, 0)
				suite.Require().Equal(agentId, uint64(1))
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("PlanQuery - %s", tc.name), func() {
			suite.SetupTest()
			if tc.malleate != nil {
				tc.malleate()
			}
			_, err := suite.queryClient.Agent(suite.ctx, tc.request)
			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
			if tc.validation != nil {
				tc.validation()
			}
		})
	}
}
