package keeper_test

import (
	"fmt"

	"github.com/Lorenzo-Protocol/lorenzo/v3/x/agent/types"
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
					AllowList: []string{
						testAdmin.String(),
					},
				},
				[]types.Agent{
					{
						Id:                  1,
						Name:                "lorenzo-stake-plan",
						BtcReceivingAddress: "tb1ptt9gnjnvje343y47z2wd7r8w6mnuylp8z0w74qftv7p5",
						EthAddr:             "0xBAb28FF7659481F1c8516f616A576339936AFB06",
						Description:         "test agent",
						Url:                 "https://xxx.com",
					},
				},
			),
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("InitGenesis - %s", tc.name), func() {
			suite.Require().NotPanics(func() {
				suite.keeper.InitGenesis(
					suite.ctx, *tc.genesisState,
				)
			})
			params := suite.lorenzoApp.AgentKeeper.GetParams(suite.ctx)
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
					AllowList: []string{
						testAdmin.String(),
					},
				},
				[]types.Agent{
					{
						Id:                  1,
						Name:                "lorenzo-stake-plan",
						BtcReceivingAddress: "tb1ptt9gnjnvje343y47z2wd7r8w6mnuylp8z0w74qftv7p5",
						EthAddr:             "0xBAb28FF7659481F1c8516f616A576339936AFB06",
						Description:         "test agent",
						Url:                 "https://xxx.com",
					},
				},
			),
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("ExportGenesis - %s", tc.name), func() {
			suite.lorenzoApp.AgentKeeper.InitGenesis(
				suite.ctx, *tc.genesisState,
			)
			genesisExported := suite.lorenzoApp.AgentKeeper.ExportGenesis(suite.ctx)
			params := suite.lorenzoApp.AgentKeeper.GetParams(suite.ctx)
			suite.Require().Equal(genesisExported.Params, params)

			agentList := suite.lorenzoApp.AgentKeeper.GetAgents(suite.ctx)
			if len(agentList) > 0 {
				suite.Require().Equal(tc.genesisState.Agents, agentList)
			} else {
				suite.Require().Len(tc.genesisState.Agents, 0)
			}
		})
	}
}
