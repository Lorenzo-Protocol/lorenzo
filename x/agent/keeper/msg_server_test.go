package keeper_test

import (
	"github.com/Lorenzo-Protocol/lorenzo/v2/app"
	"github.com/Lorenzo-Protocol/lorenzo/v2/x/agent/keeper"
	"github.com/Lorenzo-Protocol/lorenzo/v2/x/agent/types"
)

func (suite *KeeperTestSuite) TestMsgServer_AddAgent() {
	msgServer := keeper.NewMsgServerImpl(suite.keeper)

	suite.Run("invalid btc address", func() {
		_, err := msgServer.AddAgent(suite.ctx, &types.MsgAddAgent{
			Sender:              testAdmin.String(),
			Name:                "agent1",
			BtcReceivingAddress: "tb1ptt9gnjnvje343y47z2wd7r8w6mnuylp8z0w74qftv7p5",
			EthAddr:             "0xBAb28FF7659481F1c8516f616A576339936AFB06",
			Description:         "test agent",
			Url:                 "https://xxx.com",
		})
		suite.Require().Error(err)
	})

	suite.Run("not admin", func() {
		_, err := msgServer.AddAgent(suite.ctx, &types.MsgAddAgent{
			Sender:              app.CreateTestAddrs(2)[1].String(),
			Name:                "agent1",
			BtcReceivingAddress: "tb1ptt9gnjnvje343y47z2wd7r8w6mnuylp8z0w74qftv7p5x323vxeq9jrn6f",
			EthAddr:             "0xBAb28FF7659481F1c8516f616A576339936AFB06",
			Description:         "test agent",
			Url:                 "https://xxx.com",
		})
		suite.Require().Error(err)
	})

	suite.Run("success", func() {
		msg := &types.MsgAddAgent{
			Sender:              testAdmin.String(),
			Name:                "agent2",
			BtcReceivingAddress: "tb1ptt9gnjnvje343y47z2wd7r8w6mnuylp8z0w74qftv7p5x323vxeq9jrn6f",
			EthAddr:             "0xBAb28FF7659481F1c8516f616A576339936AFB06",
			Description:         "test agent",
			Url:                 "https://xxx.com",
		}
		response, err := msgServer.AddAgent(suite.ctx, msg)
		suite.Require().NoError(err)
		suite.Require().Equal(uint64(2), response.Id, "wrong agent id")

		agent, has := suite.keeper.GetAgent(suite.ctx, response.Id)
		suite.Require().True(has)
		suite.Require().EqualValues(types.Agent{
			Id:                  response.Id,
			Name:                msg.Name,
			BtcReceivingAddress: msg.BtcReceivingAddress,
			EthAddr:             msg.EthAddr,
			Description:         msg.Description,
			Url:                 msg.Url,
		}, agent, "not match agent")
	})
}

func (suite *KeeperTestSuite) TestMsgServer_EditAgent() {
	msgServer := keeper.NewMsgServerImpl(suite.keeper)
	suite.Run("not admin", func() {
		_, err := msgServer.EditAgent(suite.ctx, &types.MsgEditAgent{
			Sender: app.CreateTestAddrs(2)[1].String(),
			Id:     1,
		})
		suite.Require().Error(err)
	})

	suite.Run("agent not exist", func() {
		_, err := msgServer.EditAgent(suite.ctx, &types.MsgEditAgent{
			Sender: testAdmin.String(),
			Id:     3,
		})
		suite.Require().Error(err)
	})

	suite.Run("edit name", func() {
		suite.SetupTest()
		msgServer = keeper.NewMsgServerImpl(suite.keeper)

		_, err := msgServer.EditAgent(suite.ctx, &types.MsgEditAgent{
			Sender:      testAdmin.String(),
			Id:          1,
			Name:        "agent1_test",
			Description: types.DoNotModifyDesc,
			Url:         types.DoNotModifyDesc,
		})
		suite.Require().NoError(err)

		agent, has := suite.keeper.GetAgent(suite.ctx, 1)
		suite.Require().True(has)
		suite.Require().EqualValues(types.Agent{
			Id:                  1,
			Name:                "agent1_test",
			BtcReceivingAddress: agents[0].BtcReceivingAddress,
			EthAddr:             agents[0].EthAddr,
			Description:         agents[0].Description,
			Url:                 agents[0].Url,
		}, agent, "not match agent")
	})

	suite.Run("edit description", func() {
		suite.SetupTest()
		msgServer = keeper.NewMsgServerImpl(suite.keeper)

		_, err := msgServer.EditAgent(suite.ctx, &types.MsgEditAgent{
			Sender:      testAdmin.String(),
			Id:          1,
			Name:        types.DoNotModifyDesc,
			Description: "xxxx",
			Url:         types.DoNotModifyDesc,
		})
		suite.Require().NoError(err)

		agent, has := suite.keeper.GetAgent(suite.ctx, 1)
		suite.Require().True(has)
		suite.Require().EqualValues(types.Agent{
			Id:                  1,
			Name:                agents[0].Name,
			BtcReceivingAddress: agents[0].BtcReceivingAddress,
			EthAddr:             agents[0].EthAddr,
			Description:         "xxxx",
			Url:                 agents[0].Url,
		}, agent, "not match agent")
	})

	suite.Run("edit url", func() {
		suite.SetupTest()
		msgServer = keeper.NewMsgServerImpl(suite.keeper)

		_, err := msgServer.EditAgent(suite.ctx, &types.MsgEditAgent{
			Sender:      testAdmin.String(),
			Id:          1,
			Name:        types.DoNotModifyDesc,
			Description: types.DoNotModifyDesc,
			Url:         "xxxx",
		})
		suite.Require().NoError(err)

		agent, has := suite.keeper.GetAgent(suite.ctx, 1)
		suite.Require().True(has)
		suite.Require().EqualValues(types.Agent{
			Id:                  1,
			Name:                agents[0].Name,
			BtcReceivingAddress: agents[0].BtcReceivingAddress,
			EthAddr:             agents[0].EthAddr,
			Description:         agents[0].Description,
			Url:                 "xxxx",
		}, agent, "not match agent")
	})
}

func (suite *KeeperTestSuite) TestMsgServer_RemoveAgent() {
	msgServer := keeper.NewMsgServerImpl(suite.keeper)
	suite.Run("not admin", func() {
		_, err := msgServer.RemoveAgent(suite.ctx, &types.MsgRemoveAgent{
			Sender: app.CreateTestAddrs(2)[1].String(),
			Id:     1,
		})
		suite.Require().Error(err)
	})

	suite.Run("agent not exist", func() {
		_, err := msgServer.RemoveAgent(suite.ctx, &types.MsgRemoveAgent{
			Sender: testAdmin.String(),
			Id:     2,
		})
		suite.Require().Error(err)
	})

	suite.Run("remove success", func() {
		suite.SetupTest()
		msgServer = keeper.NewMsgServerImpl(suite.keeper)

		_, err := msgServer.RemoveAgent(suite.ctx, &types.MsgRemoveAgent{
			Sender: testAdmin.String(),
			Id:     1,
		})
		suite.Require().NoError(err)

		_, has := suite.keeper.GetAgent(suite.ctx, 1)
		suite.Require().False(has)
	})
}
