package keeper_test

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	sdkmath "cosmossdk.io/math"

	"github.com/Lorenzo-Protocol/lorenzo/x/plan/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func (suite *KeeperTestSuite) TestUpdateParams() {
	suite.SetupTest()
	testCases := []struct {
		name      string
		request   *types.MsgUpdateParams
		expectErr bool
	}{
		{
			name:      "fail - invalid authority",
			request:   &types.MsgUpdateParams{Authority: "foobar"},
			expectErr: true,
		},
		{
			name: "fail - beacon is not a valid beacon address",
			request: &types.MsgUpdateParams{
				Authority: authtypes.NewModuleAddress(govtypes.ModuleName).String(),
				Params:    types.Params{Beacon: "0x123456"},
			},
			expectErr: true,
		},
		{
			name: "fail - beacon is not a valid allow_list address",
			request: &types.MsgUpdateParams{
				Authority: authtypes.NewModuleAddress(govtypes.ModuleName).String(),
				Params: types.Params{
					Beacon: "0xbCC0CdF7683120a1965A343245FA602314C13b9A",
					AllowList: []string{
						"cosmos1qperwt9wrnkg5k9e5gzfgjppzpqqu8t3q4yjx9",
						"lrz1tffj9qp3wpdnuds443c86wffrac4jkapkjmmcy",
					},
				},
			},
			expectErr: true,
		},
		{
			name: "pass - valid Update msg",
			request: &types.MsgUpdateParams{
				Authority: authtypes.NewModuleAddress(govtypes.ModuleName).String(),
				Params: types.Params{
					Beacon: "0xbCC0CdF7683120a1965A343245FA602314C13b9A",
					AllowList: []string{
						"lrz1tffj9qp3wpdnuds443c86wffrac4jkapkjmmcy",
					},
				},
			},
			expectErr: false,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("MsgUpdateParams - %s", tc.name), func() {
			_, err := suite.msgServer.UpdateParams(suite.ctx, tc.request)
			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestCreatPlan() {
	testCases := []struct {
		name       string
		request    *types.MsgCreatePlan
		malleate   func(request *types.MsgCreatePlan)
		validation func()
		expectErr  bool
	}{
		{
			name:      "fail - invalid sender",
			request:   &types.MsgCreatePlan{Sender: "foobar"},
			expectErr: true,
		},
		{
			name:      "fail - sender not authorized",
			request:   &types.MsgCreatePlan{Sender: "lrz1tffj9qp3wpdnuds443c86wffrac4jkapkjmmcy"},
			expectErr: true,
		},
		{
			name: "fail - invalid yat contract address",
			request: &types.MsgCreatePlan{
				Name:               "lorenzo-stake-plan",
				PlanDescUri:        "https://lorenzo-protocol.io/lorenzo-stake-plan",
				AgentId:            uint64(1),
				PlanStartBlock:     sdkmath.NewInt(1000),
				PeriodBlocks:       sdkmath.NewInt(1000),
				YatContractAddress: "0xbCC0CdF7683120a1965A343245FA602314C13b9A",
				Sender:             testAdmin.String(),
			},
			malleate: func(request *types.MsgCreatePlan) {
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
			expectErr: true,
		},
		{
			name: "success - valid create plan",
			request: &types.MsgCreatePlan{
				Name:           "lorenzo-stake-plan",
				PlanDescUri:    "https://lorenzo-protocol.io/lorenzo-stake-plan",
				AgentId:        uint64(1),
				PlanStartBlock: sdkmath.NewInt(1000),
				PeriodBlocks:   sdkmath.NewInt(1000),
				Sender:         testAdmin.String(),
			},
			malleate: func(request *types.MsgCreatePlan) {
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

				request.YatContractAddress = yatAddr.Hex()
			},
			validation: func() {
				plan, found := suite.lorenzoApp.PlanKeeper.GetPlan(suite.ctx, uint64(1))
				suite.Require().True(found)
				suite.Require().Equal(plan.Name, "lorenzo-stake-plan")
				suite.Require().Equal(plan.PlanDescUri, "https://lorenzo-protocol.io/lorenzo-stake-plan")
				suite.Require().Equal(plan.AgentId, uint64(1))
				suite.Require().Equal(plan.PlanStartBlock, sdkmath.NewInt(1000))
				suite.Require().Equal(plan.PeriodBlocks, sdkmath.NewInt(1000))

				planContractAddress := common.HexToAddress(plan.ContractAddress)

				// YatContractAddress
				yatContractAddress, err := suite.lorenzoApp.PlanKeeper.YatContractAddress(suite.ctx, planContractAddress)
				suite.Require().NoError(err)
				suite.Require().Equal(yatContractAddress, common.HexToAddress(plan.YatContractAddress))
				// StakePlanName
				stakePlanName, err := suite.lorenzoApp.PlanKeeper.StakePlanName(suite.ctx, planContractAddress)
				suite.Require().NoError(err)
				suite.Require().Equal(stakePlanName, plan.Name)

				agentId, err := suite.lorenzoApp.PlanKeeper.AgentId(suite.ctx, planContractAddress)
				suite.Require().NoError(err)
				suite.Require().Equal(agentId, plan.AgentId)

				planId, err := suite.lorenzoApp.PlanKeeper.PlanId(suite.ctx, planContractAddress)
				suite.Require().NoError(err)
				suite.Require().Equal(planId, plan.Id)

				// PlanDesc
				planDesc, err := suite.lorenzoApp.PlanKeeper.PlanDesc(suite.ctx, planContractAddress)
				suite.Require().NoError(err)
				suite.Require().Equal(planDesc, plan.PlanDescUri)

				// PlanStartBlock
				planStartBlock, err := suite.lorenzoApp.PlanKeeper.PlanStartBlock(suite.ctx, planContractAddress)
				suite.Require().NoError(err)
				suite.Require().Equal(planStartBlock, plan.PlanStartBlock.Uint64())

				// PeriodBlocks
				periodBlocks, err := suite.lorenzoApp.PlanKeeper.PeriodBlocks(suite.ctx, planContractAddress)
				suite.Require().NoError(err)
				suite.Require().Equal(periodBlocks, plan.PeriodBlocks.Uint64())

				// NextRewardReceiveBlock
				nextRewardReceiveBlock, err := suite.lorenzoApp.PlanKeeper.NextRewardReceiveBlock(suite.ctx, planContractAddress)
				suite.Require().NoError(err)
				suite.Require().Equal(nextRewardReceiveBlock, plan.PlanStartBlock.Uint64()+plan.PeriodBlocks.Uint64())

				// ClaimRoundId
				claimRoundId, err := suite.lorenzoApp.PlanKeeper.ClaimRoundId(suite.ctx, planContractAddress)
				suite.Require().NoError(err)
				suite.Require().Equal(claimRoundId, uint64(0))
			},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("MsgCreatPlan - %s", tc.name), func() {
			suite.SetupTest()
			if tc.malleate != nil {
				tc.malleate(tc.request)
			}
			_, err := suite.msgServer.CreatePlan(suite.ctx, tc.request)
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

func (suite *KeeperTestSuite) TestUpgradePlan() {
	testCases := []struct {
		name       string
		request    *types.MsgUpgradePlan
		malleate   func(request *types.MsgUpgradePlan)
		validation func(request *types.MsgUpgradePlan)
		expectErr  bool
	}{
		{
			name:      "fail - invalid sender",
			request:   &types.MsgUpgradePlan{Authority: "foobar"},
			expectErr: true,
		},
		{
			name:      "fail - sender not authorized",
			request:   &types.MsgUpgradePlan{Authority: "lrz1tffj9qp3wpdnuds443c86wffrac4jkapkjmmcy"},
			expectErr: true,
		},
		{
			name: "fail - invalid implementation address",
			request: &types.MsgUpgradePlan{
				Authority:      testAdmin.String(),
				Implementation: "0x123456",
			},
			expectErr: true,
		},
		{
			name: "fail - old implementation address is empty",
			request: &types.MsgUpgradePlan{
				Authority: testAdmin.String(),
			},
			malleate: func(request *types.MsgUpgradePlan) {
				newImplementation, err := suite.lorenzoApp.PlanKeeper.DeployStakePlanLogicContract(suite.ctx)
				suite.Require().NoError(err)
				request.Implementation = newImplementation.Hex()
			},
			expectErr: true,
		},
		{
			name: "success - valid upgrade plan",
			request: &types.MsgUpgradePlan{
				Authority: testAdmin.String(),
			},
			malleate: func(request *types.MsgUpgradePlan) {
				suite.Commit()
				newImplementation, err := suite.lorenzoApp.PlanKeeper.DeployStakePlanLogicContract(suite.ctx)
				suite.Require().NoError(err)
				request.Implementation = newImplementation.Hex()
			},
			expectErr: false,
			validation: func(request *types.MsgUpgradePlan) {
				planImplementation, err := suite.lorenzoApp.PlanKeeper.GetPlanImplementationFromBeacon(suite.ctx)
				suite.Require().NoError(err)
				suite.Require().Equal(planImplementation.Hex(), request.Implementation)
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("MsgUpgradePlan - %s", tc.name), func() {
			suite.SetupTest()
			if tc.malleate != nil {
				tc.malleate(tc.request)
			}
			_, err := suite.msgServer.UpgradePlan(suite.ctx, tc.request)
			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
			if tc.validation != nil {
				tc.validation(tc.request)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestUpdatePlanStatus() {
	testCases := []struct {
		name       string
		request    *types.MsgUpdatePlanStatus
		malleate   func()
		validation func()
		expectErr  bool
	}{
		{
			name:      "fail - invalid sender",
			request:   &types.MsgUpdatePlanStatus{Sender: "foobar"},
			expectErr: true,
		},
		{
			name:      "fail - sender not authorized",
			request:   &types.MsgUpdatePlanStatus{Sender: "lrz1tffj9qp3wpdnuds443c86wffrac4jkapkjmmcy"},
			expectErr: true,
		},
		{
			name: "fail - plan not found",
			request: &types.MsgUpdatePlanStatus{
				Sender: testAdmin.String(),
				PlanId: 1,
				Status: types.PlanStatus_Pause,
			},
			expectErr: true,
		},
		{
			name: "fail - plan status equals to current status",
			request: &types.MsgUpdatePlanStatus{
				Sender: testAdmin.String(),
				PlanId: 1,
				Status: types.PlanStatus_Unpause,
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
					PlanStartBlock:     sdkmath.NewInt(1000),
					PeriodBlocks:       sdkmath.NewInt(1000),
					YatContractAddress: yatAddr.Hex(),
				}

				_, err = suite.lorenzoApp.PlanKeeper.AddPlan(suite.ctx, planReq)
				suite.Require().NoError(err)
			},
			expectErr: true,
		},
		{
			name: "success - valid update plan status, plan pause",
			request: &types.MsgUpdatePlanStatus{
				Sender: testAdmin.String(),
				PlanId: 1,
				Status: types.PlanStatus_Pause,
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
					PlanStartBlock:     sdkmath.NewInt(1000),
					PeriodBlocks:       sdkmath.NewInt(1000),
					YatContractAddress: yatAddr.Hex(),
				}

				_, err = suite.lorenzoApp.PlanKeeper.AddPlan(suite.ctx, planReq)
				suite.Require().NoError(err)
			},
			expectErr: false,
			validation: func() {
				plan, found := suite.lorenzoApp.PlanKeeper.GetPlan(suite.ctx, uint64(1))
				suite.Require().True(found)
				suite.Require().Equal(plan.Enabled, types.PlanStatus_Pause)
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("MsgUpdatePlanStatus - %s", tc.name), func() {
			suite.SetupTest()
			if tc.malleate != nil {
				tc.malleate()
			}
			_, err := suite.msgServer.UpdatePlanStatus(suite.ctx, tc.request)
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

func (suite *KeeperTestSuite) TestSetMerkleRoot() {
	testCases := []struct {
		name       string
		request    *types.MsgSetMerkleRoot
		malleate   func() string
		validation func(string, *types.MsgSetMerkleRoot)
		expectErr  bool
	}{
		{
			name:      "fail - invalid sender",
			request:   &types.MsgSetMerkleRoot{Sender: "foobar"},
			expectErr: true,
		},
		{
			name:      "fail - sender not authorized",
			request:   &types.MsgSetMerkleRoot{Sender: "lrz1tffj9qp3wpdnuds443c86wffrac4jkapkjmmcy"},
			expectErr: true,
		},
		{
			name: "fail - plan not found",
			request: &types.MsgSetMerkleRoot{
				Sender:     testAdmin.String(),
				PlanId:     1,
				MerkleRoot: "0x34337eb06160f22cfc735517076cb8d69f60afae27700d20e918cfb41f9faca7",
			},
			expectErr: true,
		},
		{
			name: "success - valid set merkle root",
			request: &types.MsgSetMerkleRoot{
				Sender:     testAdmin.String(),
				PlanId:     1,
				MerkleRoot: "0x34337eb06160f22cfc735517076cb8d69f60afae27700d20e918cfb41f9faca7",
			},
			malleate: func() string {
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

				planResult, err := suite.lorenzoApp.PlanKeeper.AddPlan(suite.ctx, planReq)
				suite.Require().NoError(err)
				return planResult.ContractAddress
			},
			validation: func(contractAddrHex string, request *types.MsgSetMerkleRoot) {
				contractAddr := common.HexToAddress(contractAddrHex)
				roundId, err := suite.lorenzoApp.PlanKeeper.ClaimRoundId(suite.ctx, contractAddr)
				suite.Require().NoError(err)
				suite.Require().Equal(roundId, uint64(1))
				merkleRoot, err := suite.lorenzoApp.PlanKeeper.MerkleRoot(
					suite.ctx, contractAddr, sdkmath.NewIntFromUint64(roundId-1).BigInt())
				suite.Require().NoError(err)
				suite.Require().Equal(merkleRoot, request.MerkleRoot)
			},
			expectErr: false,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("MsgSetMerkleRoot - %s", tc.name), func() {
			suite.SetupTest()
			var contractAddrHex string
			if tc.malleate != nil {
				contractAddrHex = tc.malleate()
			}
			_, err := suite.msgServer.SetMerkleRoot(suite.ctx, tc.request)
			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
			if tc.validation != nil {
				tc.validation(contractAddrHex, tc.request)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestClaims() {
	testCases := []struct {
		name       string
		request    *types.MsgClaims
		malleate   func()
		validation func()
		expectErr  bool
	}{
		{
			name:      "fail - invalid sender",
			request:   &types.MsgClaims{Sender: "foobar"},
			expectErr: true,
		},
		{
			name: "fail - plan not found",
			request: &types.MsgClaims{
				Sender:      "lrz1cpldpp5960ed8s63w4v9fml84w875wv0emcda5",
				PlanId:      1,
				Receiver:    "lrz1cpldpp5960ed8s63w4v9fml84w875wv0emcda5",
				RoundId:     sdkmath.NewInt(0),
				Amount:      sdkmath.NewInt(100),
				MerkleProof: "0x1764cb495e1c2565f6d033e298a2d46a527c93a5a48c8b318fa05e9b07489b33",
			},
			expectErr: true,
		},
		{
			name: "success - valid claims",
			request: &types.MsgClaims{
				Sender:      "lrz1cpldpp5960ed8s63w4v9fml84w875wv0emcda5",
				PlanId:      1,
				Receiver:    "0xc07ed08685d3F2D3c351755854EFE7ab8fEa398F",
				RoundId:     sdkmath.NewInt(0),
				Amount:      sdkmath.NewInt(100),
				MerkleProof: "0x365cc96c249dc95f3f2e4934371b55ee1c5ef9e6f6da6407b1ec26aa6cd12109",
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
					PlanStartBlock:     sdkmath.NewInt(1000),
					PeriodBlocks:       sdkmath.NewInt(1000),
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
			validation: func() {
				plan, found := suite.lorenzoApp.PlanKeeper.GetPlan(suite.ctx, uint64(1))
				suite.Require().True(found)
				amount, err := suite.lorenzoApp.PlanKeeper.BalanceOfFromYAT(
					suite.ctx, common.HexToAddress(plan.YatContractAddress),
					common.HexToAddress("0xc07ed08685d3F2D3c351755854EFE7ab8fEa398F"))
				suite.Require().NoError(err)
				suite.Require().Equal(amount, sdkmath.NewInt(100).BigInt())
			},
			expectErr: false,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("MsgClaims - %s", tc.name), func() {
			suite.SetupTest()
			if tc.malleate != nil {
				tc.malleate()
			}
			_, err := suite.msgServer.Claims(suite.ctx, tc.request)
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

func (suite *KeeperTestSuite) TestCreateYAT() {
	testCases := []struct {
		name       string
		request    *types.MsgCreateYAT
		malleate   func()
		validation func(string)
		expectErr  bool
	}{
		{
			name:      "fail - invalid sender",
			request:   &types.MsgCreateYAT{Sender: "foobar"},
			expectErr: true,
		},
		{
			name:      "fail - sender not authorized",
			request:   &types.MsgCreateYAT{Sender: "lrz1tffj9qp3wpdnuds443c86wffrac4jkapkjmmcy"},
			expectErr: true,
		},
		{
			name: "success - valid create yat",
			request: &types.MsgCreateYAT{
				Sender: testAdmin.String(),
				Name:   "lorenzo",
				Symbol: "ALRZ",
			},
			expectErr: false,
			validation: func(yatAddressHex string) {
				yatAddress := common.HexToAddress(yatAddressHex)
				owner, err := suite.lorenzoApp.PlanKeeper.GetOwner(suite.ctx, yatAddress)
				suite.Require().NoError(err)
				suite.Require().Equal(owner, types.ModuleAddress)
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("MsgCreateYAT - %s", tc.name), func() {
			suite.SetupTest()
			if tc.malleate != nil {
				tc.malleate()
			}
			result, err := suite.msgServer.CreateYAT(suite.ctx, tc.request)
			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
			if tc.validation != nil {
				tc.validation(result.ContractAddress)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestSetMinter() {
	testCases := []struct {
		name       string
		request    *types.MsgSetMinter
		malleate   func(*types.MsgSetMinter)
		validation func(*types.MsgSetMinter)
		expectErr  bool
	}{
		{
			name:      "fail - invalid sender",
			request:   &types.MsgSetMinter{Sender: "foobar"},
			expectErr: true,
		},
		{
			name:      "fail - sender not authorized",
			request:   &types.MsgSetMinter{Sender: "lrz1tffj9qp3wpdnuds443c86wffrac4jkapkjmmcy"},
			expectErr: true,
		},
		{
			name: "fail - invalid minter address",
			request: &types.MsgSetMinter{
				Sender:          testAdmin.String(),
				Minter:          "0x123456",
				ContractAddress: "0xbCC0CdF7683120a1965A343245FA602314C13b9A",
			},
			expectErr: true,
		},
		{
			name: "fail - invalid yat contract address",
			request: &types.MsgSetMinter{
				Sender:          testAdmin.String(),
				Minter:          types.ModuleAddress.Hex(),
				ContractAddress: "0x123456",
			},
			expectErr: true,
		},
		{
			name: "fail - yat not exists",
			request: &types.MsgSetMinter{
				Sender:          testAdmin.String(),
				ContractAddress: "0xbCC0CdF7683120a1965A343245FA602314C13b9A",
			},
			expectErr: true,
		},
		{
			name: "fail - minter not exists",
			request: &types.MsgSetMinter{
				Sender: testAdmin.String(),
				Minter: "0xbCC0CdF7683120a1965A343245FA602314C13b9A",
			},
			expectErr: true,
			malleate: func(request *types.MsgSetMinter) {
				yatAddr, err := suite.lorenzoApp.PlanKeeper.DeployYATContract(
					suite.ctx, "lorenzo", "ALRZ")
				suite.Require().NoError(err)
				request.ContractAddress = yatAddr.Hex()
			},
		},
		{
			name: "success - valid set minter",
			request: &types.MsgSetMinter{
				Sender: testAdmin.String(),
			},
			malleate: func(request *types.MsgSetMinter) {
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

				// deploy yat contract
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

				planResult, err := suite.lorenzoApp.PlanKeeper.AddPlan(suite.ctx, planReq)
				suite.Require().NoError(err)
				request.ContractAddress = yatAddr.Hex()
				request.Minter = planResult.ContractAddress
			},
			expectErr: false,
			validation: func(request *types.MsgSetMinter) {
				yatAddress := common.HexToAddress(request.ContractAddress)
				minterAddress := common.HexToAddress(request.Minter)
				found, err := suite.lorenzoApp.PlanKeeper.HasRoleFromYAT(
					suite.ctx,
					yatAddress,
					"YAT_MINTER_ROLE",
					minterAddress,
				)
				suite.Require().NoError(err)
				suite.Require().True(found)
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("MsgSetMinter - %s", tc.name), func() {
			suite.SetupTest()
			if tc.malleate != nil {
				tc.malleate(tc.request)
			}
			_, err := suite.msgServer.SetMinter(suite.ctx, tc.request)
			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
			if tc.validation != nil {
				tc.validation(tc.request)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestRemoveMinter() {
	testCases := []struct {
		name       string
		request    *types.MsgRemoveMinter
		malleate   func(*types.MsgRemoveMinter)
		validation func(*types.MsgRemoveMinter)
		expectErr  bool
	}{
		{
			name:      "fail - invalid sender",
			request:   &types.MsgRemoveMinter{Sender: "foobar"},
			expectErr: true,
		},
		{
			name:      "fail - sender not authorized",
			request:   &types.MsgRemoveMinter{Sender: "lrz1tffj9qp3wpdnuds443c86wffrac4jkapkjmmcy"},
			expectErr: true,
		},
		{
			name: "fail - invalid minter address",
			request: &types.MsgRemoveMinter{
				Sender:          testAdmin.String(),
				Minter:          "0x123456",
				ContractAddress: "0xbCC0CdF7683120a1965A343245FA602314C13b9A",
			},
			expectErr: true,
		},
		{
			name: "fail - invalid yat contract address",
			request: &types.MsgRemoveMinter{
				Sender:          testAdmin.String(),
				Minter:          types.ModuleAddress.Hex(),
				ContractAddress: "0x123456",
			},
			expectErr: true,
		},
		{
			name: "fail - yat not exists",
			request: &types.MsgRemoveMinter{
				Sender:          testAdmin.String(),
				ContractAddress: "0xbCC0CdF7683120a1965A343245FA602314C13b9A",
			},
			expectErr: true,
		},
		{
			name: "fail - minter not exists",
			request: &types.MsgRemoveMinter{
				Sender: testAdmin.String(),
				Minter: "0xbCC0CdF7683120a1965A343245FA602314C13b9A",
			},
			expectErr: true,
			malleate: func(request *types.MsgRemoveMinter) {
				yatAddr, err := suite.lorenzoApp.PlanKeeper.DeployYATContract(
					suite.ctx, "lorenzo", "ALRZ")
				suite.Require().NoError(err)
				request.ContractAddress = yatAddr.Hex()
			},
		},
		{
			name: "success - valid remove minter",
			request: &types.MsgRemoveMinter{
				Sender: testAdmin.String(),
			},
			malleate: func(request *types.MsgRemoveMinter) {
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

				// deploy yat contract
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

				planResult, err := suite.lorenzoApp.PlanKeeper.AddPlan(suite.ctx, planReq)
				suite.Require().NoError(err)
				// set minter
				err = suite.lorenzoApp.PlanKeeper.SetMinter(
					suite.ctx,
					yatAddr,
					common.HexToAddress(planResult.ContractAddress),
				)
				suite.Require().NoError(err)
				request.ContractAddress = yatAddr.Hex()
				request.Minter = planResult.ContractAddress
			},
			expectErr: false,
			validation: func(request *types.MsgRemoveMinter) {
				yatAddress := common.HexToAddress(request.ContractAddress)
				minterAddress := common.HexToAddress(request.Minter)
				found, err := suite.lorenzoApp.PlanKeeper.HasRoleFromYAT(
					suite.ctx,
					yatAddress,
					"YAT_MINTER_ROLE",
					minterAddress,
				)
				suite.Require().NoError(err)
				suite.Require().False(found)
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("MsgRemoveMinter - %s", tc.name), func() {
			suite.SetupTest()
			if tc.malleate != nil {
				tc.malleate(tc.request)
			}
			_, err := suite.msgServer.RemoveMinter(suite.ctx, tc.request)
			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
			if tc.validation != nil {
				tc.validation(tc.request)
			}
		})
	}
}
