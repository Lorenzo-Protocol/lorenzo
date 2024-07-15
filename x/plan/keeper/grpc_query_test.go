package keeper_test

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/common"

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
					PlanStartTime:      1000,
					PeriodTime:         1000,
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
					PlanStartTime:      1000,
					PeriodTime:         1000,
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
	testCases := []struct {
		name         string
		request      *types.QueryClaimLeafNodeRequest
		malleate     func()
		getArgs      func() string
		expectErr    bool
		expectResult bool
	}{
		{
			name: "success - leaf node exists",
			request: &types.QueryClaimLeafNodeRequest{
				Id:       1,
				RoundId:  sdkmath.NewInt(0),
				LeafNode: "",
			},
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
					PlanStartTime:      1000,
					PeriodTime:         1000,
					YatContractAddress: yatAddr.Hex(),
				}

				planResult, err := suite.lorenzoApp.PlanKeeper.AddPlan(suite.ctx, planReq)
				suite.Require().NoError(err)

				// set merkle root
				merkelRoot := "0x39c19150c14c397b133682e95742b651babde3418edaaa4375a3197604159346"
				err = suite.lorenzoApp.PlanKeeper.SetMerkleRoot(
					suite.ctx,
					common.HexToAddress(planResult.ContractAddress),
					merkelRoot)
				suite.Require().NoError(err)

				// set minter
				err = suite.lorenzoApp.PlanKeeper.SetMinter(suite.ctx, yatAddr, common.HexToAddress(planResult.ContractAddress))
				suite.Require().NoError(err)
				account := common.HexToAddress("0xc07ed08685d3F2D3c351755854EFE7ab8fEa398F")
				err = suite.lorenzoApp.PlanKeeper.ClaimYATToken(
					suite.ctx,
					common.HexToAddress(planResult.ContractAddress),
					account,
					sdkmath.NewInt(0).BigInt(),
					sdkmath.NewInt(100).BigInt(),
					"0x365cc96c249dc95f3f2e4934371b55ee1c5ef9e6f6da6407b1ec26aa6cd12109",
				)
				suite.Require().NoError(err)
			},
			getArgs: func() string {
				account := common.HexToAddress("0xc07ed08685d3F2D3c351755854EFE7ab8fEa398F")
				amount := big.NewInt(100)
				leafNode := crypto.Keccak256Hash(
					account.Bytes(),
					common.LeftPadBytes(amount.Bytes(), 32),
				)
				return leafNode.Hex()
			},
			expectErr:    false,
			expectResult: true,
		},

		{
			name: "success - leaf node not exists",
			request: &types.QueryClaimLeafNodeRequest{
				Id:       1,
				RoundId:  sdkmath.NewInt(0),
				LeafNode: "",
			},
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
					PlanStartTime:      1000,
					PeriodTime:         1000,
					YatContractAddress: yatAddr.Hex(),
				}

				planResult, err := suite.lorenzoApp.PlanKeeper.AddPlan(suite.ctx, planReq)
				suite.Require().NoError(err)

				// set merkle root
				merkelRoot := "0x39c19150c14c397b133682e95742b651babde3418edaaa4375a3197604159346"
				err = suite.lorenzoApp.PlanKeeper.SetMerkleRoot(
					suite.ctx,
					common.HexToAddress(planResult.ContractAddress),
					merkelRoot)
				suite.Require().NoError(err)

				// set minter
				err = suite.lorenzoApp.PlanKeeper.SetMinter(suite.ctx, yatAddr, common.HexToAddress(planResult.ContractAddress))
				suite.Require().NoError(err)
			},
			getArgs: func() string {
				account := common.HexToAddress("0xdC29190A78C1533a0780D9fF8A801f095FA5B615")
				amount := sdkmath.NewInt(200).BigInt()
				leafNode := crypto.Keccak256Hash(
					account.Bytes(),
					common.LeftPadBytes(amount.Bytes(), 32),
				)
				return leafNode.Hex()
			},
			expectErr:    false,
			expectResult: false,
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("ClaimLeafNodeQuery - %s", tc.name), func() {
			suite.SetupTest()
			if tc.getArgs != nil {
				tc.request.LeafNode = tc.getArgs()
			}
			if tc.malleate != nil {
				tc.malleate()
			}
			result, err := suite.queryClient.ClaimLeafNode(suite.ctx, tc.request)
			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
			suite.Require().Equal(tc.expectResult, result.Success)
		})
	}
}
