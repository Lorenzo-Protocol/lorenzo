package keeper_test

import (
	"fmt"

	"github.com/Lorenzo-Protocol/lorenzo/v3/x/agent/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func (suite *KeeperTestSuite) TestMsgServer_UpdateParams() {
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
			name: "fail - AllowList address is not valid",
			request: &types.MsgUpdateParams{
				Authority: authtypes.NewModuleAddress(govtypes.ModuleName).String(),
				Params: types.Params{
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

func (suite *KeeperTestSuite) TestMsgServer_AddAgent() {
	testCases := []struct {
		name       string
		request    *types.MsgAddAgent
		malleate   func(request *types.MsgAddAgent)
		validation func(request *types.MsgAddAgent)
		expectErr  bool
	}{
		{
			name: "fail - invalid sender",
			request: &types.MsgAddAgent{
				Name:                "sinohope4",
				BtcReceivingAddress: "tb1ptt9gnjnvje343y47z2wd7r8w6mnuylp8z0w74qftv7p5x323vxeq9jrn6f",
				EthAddr:             "0xbCC0CdF7683120a1965A343245FA602314C13b9A",
				Description:         "test",
				Url:                 "https://sinohope.com",
			},
			expectErr: true,
		},
		{
			name:      "fail - sender not authorized",
			request:   &types.MsgAddAgent{Sender: "lrz1tffj9qp3wpdnuds443c86wffrac4jkapkjmmcy"},
			expectErr: true,
		},
		{
			name: "success - valid create agent",
			request: &types.MsgAddAgent{
				Name:                "sinohope4",
				BtcReceivingAddress: "tb1ptt9gnjnvje343y47z2wd7r8w6mnuylp8z0w74qftv7p5x323vxeq9jrn6f",
				EthAddr:             "0xbCC0CdF7683120a1965A343245FA602314C13b9A",
				Description:         "test",
				Url:                 "https://sinohope.com",
				Sender:              testAdmin.String(),
			},
			malleate: func(request *types.MsgAddAgent) {
				suite.Commit()
			},
			validation: func(request *types.MsgAddAgent) {
				agent, found := suite.lorenzoApp.AgentKeeper.GetAgent(suite.ctx, uint64(1))
				suite.Require().True(found)
				suite.Require().Equal(request.Name, agent.Name)
				suite.Require().Equal(request.BtcReceivingAddress, agent.BtcReceivingAddress)
				suite.Require().Equal(request.EthAddr, agent.EthAddr)
				suite.Require().Equal(request.Description, agent.Description)
				suite.Require().Equal(request.Url, agent.Url)
				suite.Require().Equal(agent.Id, uint64(1))
			},
			expectErr: false,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("MsgAddAgent - %s", tc.name), func() {
			suite.SetupTest()
			if tc.malleate != nil {
				tc.malleate(tc.request)
			}
			_, err := suite.msgServer.AddAgent(suite.ctx, tc.request)
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

func (suite *KeeperTestSuite) TestMsgServer_EditAgent() {
	testCases := []struct {
		name       string
		request    *types.MsgEditAgent
		malleate   func(request *types.MsgEditAgent)
		validation func(request *types.MsgEditAgent)
		expectErr  bool
	}{
		{
			name:      "fail - invalid sender",
			request:   &types.MsgEditAgent{Sender: "foobar"},
			expectErr: true,
		},
		{
			name:      "fail - sender not authorized",
			request:   &types.MsgEditAgent{Sender: "lrz1tffj9qp3wpdnuds443c86wffrac4jkapkjmmcy"},
			expectErr: true,
		},
		{
			name: "success - valid edit agent",
			request: &types.MsgEditAgent{
				Id:          1,
				Name:        "lorenzo",
				Description: "lorenzo is a protocol",
				Url:         "https://lorenzo.com",
				Sender:      testAdmin.String(),
			},
			malleate: func(request *types.MsgEditAgent) {
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
			validation: func(request *types.MsgEditAgent) {
				agent, found := suite.lorenzoApp.AgentKeeper.GetAgent(suite.ctx, uint64(1))
				suite.Require().True(found)
				suite.Require().Equal(request.Name, agent.Name)
				suite.Require().Equal(request.Description, agent.Description)
				suite.Require().Equal(request.Url, agent.Url)
				suite.Require().Equal(agent.Id, uint64(1))
			},
			expectErr: false,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("MsgEditAgent - %s", tc.name), func() {
			suite.SetupTest()
			if tc.malleate != nil {
				tc.malleate(tc.request)
			}
			_, err := suite.msgServer.EditAgent(suite.ctx, tc.request)
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

func (suite *KeeperTestSuite) TestMsgServer_RemoveAgent() {
	testCases := []struct {
		name       string
		request    *types.MsgRemoveAgent
		malleate   func(request *types.MsgRemoveAgent)
		validation func(request *types.MsgRemoveAgent)
		expectErr  bool
	}{
		{
			name:      "fail - invalid sender",
			request:   &types.MsgRemoveAgent{Sender: "foobar"},
			expectErr: true,
		},
		{
			name:      "fail - sender not authorized",
			request:   &types.MsgRemoveAgent{Sender: "lrz1tffj9qp3wpdnuds443c86wffrac4jkapkjmmcy"},
			expectErr: true,
		},
		{
			name: "fail - agent not found",
			request: &types.MsgRemoveAgent{
				Id:     10,
				Sender: testAdmin.String(),
			},
			expectErr: true,
		},
		{
			name: "success - valid remove agent",
			request: &types.MsgRemoveAgent{
				Id:     1,
				Sender: testAdmin.String(),
			},
			malleate: func(request *types.MsgRemoveAgent) {
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
			validation: func(request *types.MsgRemoveAgent) {
				_, found := suite.lorenzoApp.AgentKeeper.GetAgent(suite.ctx, uint64(1))
				suite.Require().False(found)
			},
			expectErr: false,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("MsgEditAgent - %s", tc.name), func() {
			suite.SetupTest()
			if tc.malleate != nil {
				tc.malleate(tc.request)
			}
			_, err := suite.msgServer.RemoveAgent(suite.ctx, tc.request)
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
