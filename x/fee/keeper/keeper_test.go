package keeper_test

import (
	"testing"

	"github.com/Lorenzo-Protocol/lorenzo/v2/app"
	"github.com/stretchr/testify/suite"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Lorenzo-Protocol/lorenzo/v2/x/fee/keeper"
	"github.com/Lorenzo-Protocol/lorenzo/v2/x/fee/types"
)

type KeeperTestSuite struct {
	suite.Suite

	cdc    *codec.LegacyAmino
	ifr    codectypes.InterfaceRegistry
	ctx    sdk.Context
	keeper *keeper.Keeper
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	app := app.Setup(suite.T())

	suite.cdc = app.LegacyAmino()
	suite.ifr = app.InterfaceRegistry()
	suite.ctx = app.BaseApp.NewContext(false, tmproto.Header{})
	suite.keeper = app.FeeKeeper
}

func (suite *KeeperTestSuite) TestGetSetParams() {
	params := types.Params{
		NonFeeMsgs: []string{
			sdk.MsgTypeURL(&types.MsgUpdateParams{}),
		},
	}
	suite.NoError(suite.keeper.SetParams(suite.ctx, params), "set params failed")

	expectedParams := suite.keeper.GetParams(suite.ctx)
	suite.True(params.Equal(expectedParams), "not expected params")
}
