package keeper_test

import (
	"fmt"

	sdkmath "cosmossdk.io/math"
	"github.com/Lorenzo-Protocol/lorenzo/x/plan/types"
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

func (suite *KeeperTestSuite) TestPlansQuery() {
	testCases := []struct {
		name       string
		request    *types.QueryPlansRequest
		malleate   func()
		validation func()
		expectErr  bool
	}{
		{
			name:      "success",
			request:   &types.QueryPlansRequest{},
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

				yatAddr, err := suite.lorenzoApp.PlanKeeper.DeployYATContract(
					suite.ctx, "lorenzo", "ALRZ")
				suite.Require().NoError(err)
				// create plan
				planReq := types.Plan{
					Name:               "lorenzo-stake-plan",
					PlanDescUri:        "https://lorenzo-protocol.io/lorenzo-stake-plan",
					AgentId:            uint64(1),
					PlanStartBlock:     sdkmath.NewInt(1000),
					PeriodBlocks:       sdkmath.NewInt(1000),
					YatContractAddress: yatAddr.Hex(),
				}

				_, err = suite.lorenzoApp.PlanKeeper.AddPlan(suite.ctx, planReq)
				suite.Require().NoError(err)
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("PlansQuery - %s", tc.name), func() {
			suite.SetupTest()
			if tc.malleate != nil {
				tc.malleate()
			}
			_, err := suite.queryClient.Plans(suite.ctx, tc.request)
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

func (suite *KeeperTestSuite) TestPlanQuery() {
	testCases := []struct {
		name       string
		request    *types.QueryPlanRequest
		malleate   func()
		validation func()
		expectErr  bool
	}{
		{
			name: "success",
			request: &types.QueryPlanRequest{
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

				yatAddr, err := suite.lorenzoApp.PlanKeeper.DeployYATContract(
					suite.ctx, "lorenzo", "ALRZ")
				suite.Require().NoError(err)
				// create plan
				planReq := types.Plan{
					Name:               "lorenzo-stake-plan",
					PlanDescUri:        "https://lorenzo-protocol.io/lorenzo-stake-plan",
					AgentId:            uint64(1),
					PlanStartBlock:     sdkmath.NewInt(1000),
					PeriodBlocks:       sdkmath.NewInt(1000),
					YatContractAddress: yatAddr.Hex(),
				}

				_, err = suite.lorenzoApp.PlanKeeper.AddPlan(suite.ctx, planReq)
				suite.Require().NoError(err)
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("PlanQuery - %s", tc.name), func() {
			suite.SetupTest()
			if tc.malleate != nil {
				tc.malleate()
			}
			_, err := suite.queryClient.Plan(suite.ctx, tc.request)
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

func (suite *KeeperTestSuite) TestClaimLeafNodeQuery() {
	//testCases := []struct {
	//	name       string
	//	request    *types.QueryClaimLeafNodeRequest
	//	malleate   func()
	//	validation func()
	//	expectErr  bool
	//}{
	//	{
	//		name:      "success",
	//		request:   &types.QueryClaimLeafNodeRequest{},
	//		expectErr: false,
	//		malleate: func() {
	//			suite.Commit()
	//		},
	//	},
	//}
	//
	//for _, tc := range testCases {
	//	suite.Run(fmt.Sprintf("ClaimLeafNodeQuery - %s", tc.name), func() {
	//		suite.SetupTest()
	//		if tc.malleate != nil {
	//			tc.malleate()
	//		}
	//		_, err := suite.queryClient.ClaimLeafNode(suite.ctx, tc.request)
	//		if tc.expectErr {
	//			suite.Require().Error(err)
	//		} else {
	//			suite.Require().NoError(err)
	//		}
	//		if tc.validation != nil {
	//			tc.validation()
	//		}
	//	})
	//}
}
