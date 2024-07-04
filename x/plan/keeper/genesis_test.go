package keeper_test

import (
	"fmt"

	sdkmath "cosmossdk.io/math"
	"github.com/Lorenzo-Protocol/lorenzo/x/plan/types"
)

func (suite *KeeperTestSuite) TestInitGenesis() {
	testCases := []struct {
		name         string
		genesisState *types.GenesisState
	}{
		{
			"empty genesis",
			&types.GenesisState{},
		},
		{
			"default genesis",
			types.DefaultGenesisState(),
		},
		{
			"custom genesis",
			types.NewGenesisState(
				types.Params{
					Beacon: "0x1cd55261EC1e167bf7c7201EE79517B7334F575c",
					AllowList: []string{
						testAdmin.String(),
					},
				},
				1,
				[]types.Plan{
					{
						Id:                 1,
						Name:               "lorenzo-stake-plan",
						PlanDescUri:        "https://lorenzo-protocol.io/lorenzo-stake-plan",
						AgentId:            uint64(1),
						PlanStartBlock:     sdkmath.NewInt(1000),
						PeriodBlocks:       sdkmath.NewInt(1000),
						YatContractAddress: "0x5dCA2483280D9727c80b5518faC4556617fb19ZZ",
						ContractAddress:    "0xa0DD3937820aE374cEC68b131E19b54c574b6710",
						Enabled:            types.PlanStatus_Unpause,
					},
				},
			),
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("InitGenesis - %s", tc.name), func() {
			suite.Require().NotPanics(func() {
				suite.lorenzoApp.PlanKeeper.InitGenesis(
					suite.ctx, *tc.genesisState,
				)
			})
			params := suite.lorenzoApp.PlanKeeper.GetParams(suite.ctx)
			suite.Require().Equal(tc.genesisState.Params, params)
		})
	}
}

func (suite *KeeperTestSuite) TestExportGenesis() {
	testCases := []struct {
		name         string
		genesisState *types.GenesisState
	}{
		{
			"empty genesis",
			&types.GenesisState{},
		},
		{
			"default genesis",
			types.DefaultGenesisState(),
		},
		{
			"custom genesis",
			types.NewGenesisState(
				types.Params{
					Beacon: "0x1cd55261EC1e167bf7c7201EE79517B7334F575c",
					AllowList: []string{
						testAdmin.String(),
					},
				},
				1,
				[]types.Plan{
					{
						Id:                 1,
						Name:               "lorenzo-stake-plan",
						PlanDescUri:        "https://lorenzo-protocol.io/lorenzo-stake-plan",
						AgentId:            uint64(1),
						PlanStartBlock:     sdkmath.NewInt(1000),
						PeriodBlocks:       sdkmath.NewInt(1000),
						YatContractAddress: "0x5dCA2483280D9727c80b5518faC4556617fb19ZZ",
						ContractAddress:    "0xa0DD3937820aE374cEC68b131E19b54c574b6710",
						Enabled:            types.PlanStatus_Unpause,
					},
				},
			),
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("ExportGenesis - %s", tc.name), func() {
			suite.lorenzoApp.PlanKeeper.InitGenesis(
				suite.ctx, *tc.genesisState,
			)
			genesisExported := suite.lorenzoApp.PlanKeeper.ExportGenesis(suite.ctx)
			params := suite.lorenzoApp.PlanKeeper.GetParams(suite.ctx)
			suite.Require().Equal(genesisExported.Params, params)

			plans := suite.lorenzoApp.PlanKeeper.GetPlans(suite.ctx)
			if len(plans) > 0 {
				suite.Require().Equal(tc.genesisState.Plans, plans)
			} else {
				suite.Require().Len(tc.genesisState.Plans, 0)
			}
		})
	}
}
