package keeper_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/suite"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Lorenzo-Protocol/lorenzo/app/helpers"
	"github.com/Lorenzo-Protocol/lorenzo/x/agent/keeper"
	"github.com/Lorenzo-Protocol/lorenzo/x/agent/types"
)

var (
	testAdmin  = helpers.CreateTestAddrs(1)[0]
	agents     = []types.Agent{
		{
			Id:                  1,
			Name:                "agent1",
			BtcReceivingAddress: "tb1ptt9gnjnvje343y47z2wd7r8w6mnuylp8z0w74qftv7p5x323vxeq9jrn6f",
			EthAddr:             "0xBAb28FF7659481F1c8516f616A576339936AFB06",
			Description:         "test agent",
			Url:                 "https://xxx.com",
		},
	}
)

type KeeperTestSuite struct {
	suite.Suite

	ctx    sdk.Context
	keeper keeper.Keeper
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	merge := func(cdc codec.Codec, state map[string]json.RawMessage) {
		genesis := &types.GenesisState{
			Agents:     agents,
			Admin:      testAdmin.String(),
		}
		state[types.ModuleName] = cdc.MustMarshalJSON(genesis)
	}

	app := helpers.SetupWithGenesisMergeFn(suite.T(), merge)
	suite.ctx = app.BaseApp.NewContext(false, tmproto.Header{})
	suite.keeper = app.AgentKeeper
}

func (suite *KeeperTestSuite) TestGetAgent() {
	suite.Run("not found", func() {
		_, has := suite.keeper.GetAgent(suite.ctx, 2)
		suite.False(has)
	})

	suite.Run("found", func() {
		agent, has := suite.keeper.GetAgent(suite.ctx, agents[0].Id)
		suite.True(has)
		suite.Equal(agent, agents[0])
	})
}

func (suite *KeeperTestSuite) TestGetNextNumber() {
	suite.SetupTest()

	nextNumber := suite.keeper.GetNextNumber(suite.ctx)
	suite.Equal(nextNumber, uint64(2))
}

func (suite *KeeperTestSuite) TestGetAdmin() {
	suite.SetupTest()
	
	admin := suite.keeper.GetAdmin(suite.ctx)
	suite.Equal(admin, testAdmin)
}
