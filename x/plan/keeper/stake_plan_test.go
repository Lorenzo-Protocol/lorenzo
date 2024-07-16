package keeper_test

import (
	"fmt"
	"math/big"

	sdkmath "cosmossdk.io/math"
	"github.com/Lorenzo-Protocol/lorenzo/x/plan/types"

	"github.com/ethereum/go-ethereum/common"
)

func (suite *KeeperTestSuite) TestStakePlan() {
	testCases := []struct {
		name       string
		plan       *types.Plan
		malleate   func(plan *types.Plan) common.Address
		validation func(*types.Plan, common.Address)
	}{
		{
			name: "success - query plan",
			plan: &types.Plan{
				Id:                 1,
				Name:               "lorenzo-stake-plan",
				PlanDescUri:        "https://lorenzo-protocol.io/lorenzo-stake-plan",
				AgentId:            1,
				PlanStartBlock:     1000,
				PeriodBlocks:       1000,
				YatContractAddress: "",
				Enabled:            types.PlanStatus_Unpause,
			},
			malleate: func(plan *types.Plan) common.Address {
				// create plan
				suite.Commit()

				// deploy yat
				yatAddr, err := suite.lorenzoApp.PlanKeeper.DeployYATContract(
					suite.ctx, "lorenzo", "ALRZ")
				suite.Require().NoError(err)
				plan.YatContractAddress = yatAddr.Hex()
				//
				sdkmath.NewIntFromUint64(plan.AgentId)
				planAddr, err := suite.keeper.DeployStakePlanProxyContract(
					suite.ctx,
					plan.Name,
					plan.PlanDescUri,
					sdkmath.NewIntFromUint64(plan.Id).BigInt(),
					sdkmath.NewIntFromUint64(plan.AgentId).BigInt(),
					big.NewInt(int64(plan.PlanStartBlock)),
					big.NewInt(int64(plan.PeriodBlocks)),
					yatAddr,
				)
				suite.Require().NoError(err)
				return planAddr
			},
			validation: func(plan *types.Plan, address common.Address) {
				// plan id
				planId, err := suite.keeper.PlanId(suite.ctx, address)
				suite.Require().NoError(err)
				suite.Require().Equal(plan.Id, planId)

				// plan name
				planName, err := suite.keeper.StakePlanName(suite.ctx, address)
				suite.Require().NoError(err)
				suite.Require().Equal(plan.Name, planName)

				// plan agent id
				agentId, err := suite.keeper.AgentId(suite.ctx, address)
				suite.Require().NoError(err)
				suite.Require().Equal(plan.AgentId, agentId)

				// Plan desc
				planDesc, err := suite.keeper.PlanDesc(suite.ctx, address)
				suite.Require().NoError(err)
				suite.Require().Equal(plan.PlanDescUri, planDesc)

				// Plan start block
				planStartBlock, err := suite.keeper.PlanStartBlock(suite.ctx, address)
				suite.Require().NoError(err)
				suite.Require().Equal(plan.PlanStartBlock, planStartBlock)

				// Period blocks
				periodBlocks, err := suite.keeper.PeriodBlocks(suite.ctx, address)
				suite.Require().NoError(err)
				suite.Require().Equal(plan.PeriodBlocks, periodBlocks)
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("QueryStakePlan - %s", tc.name), func() {
			suite.SetupTest() // reset
			contractAddr := tc.malleate(tc.plan)
			tc.validation(tc.plan, contractAddr)
		})
	}
}
