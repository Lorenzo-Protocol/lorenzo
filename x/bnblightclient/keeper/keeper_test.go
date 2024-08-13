package keeper_test

import (
	"encoding/json"
	"testing"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/Lorenzo-Protocol/lorenzo/v3/app"
	"github.com/Lorenzo-Protocol/lorenzo/v3/x/bnblightclient/keeper"
	"github.com/Lorenzo-Protocol/lorenzo/v3/x/bnblightclient/testutil"
	"github.com/Lorenzo-Protocol/lorenzo/v3/x/bnblightclient/types"
)

var testAdmin = app.CreateTestAddrs(1)[0]

type KeeperTestSuite struct {
	suite.Suite

	ctx    sdk.Context
	keeper keeper.Keeper

	msgServer   types.MsgServer
	queryClient types.QueryClient
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	merge := func(cdc codec.Codec, state map[string]json.RawMessage) {
		headers := testutil.GetTestHeaders(suite.T())
		genesis := &types.GenesisState{
			Params: &types.Params{
				StakePlanHubAddress: "0x9ADb675bc89d9EC5d829709e85562b7c99658D59",
				EventName:           "StakeBTC2JoinStakePlan",
				RetainedBlocks:      10,
				AllowList: []string{
					testAdmin.String(),
				},
				ChainId: 56,
			},
			Headers: headers[:5],
		}
		state[types.ModuleName] = cdc.MustMarshalJSON(genesis)
	}

	lorenzoApp := app.SetupWithGenesisMergeFn(suite.T(), merge)

	ctx := lorenzoApp.BaseApp.NewContext(false, tmproto.Header{})

	suite.ctx = ctx
	suite.keeper = lorenzoApp.BNBLightClientKeeper

	queryHelper := baseapp.NewQueryServerTestHelper(suite.ctx, lorenzoApp.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, keeper.NewQuerierImpl(suite.keeper))
	// types.RegisterMsgServer(queryHelper, keeper.NewMsgServerImpl(suite.keeper))

	queryClient := types.NewQueryClient(queryHelper)
	suite.msgServer = keeper.NewMsgServerImpl(suite.keeper)
	suite.queryClient = queryClient
}
