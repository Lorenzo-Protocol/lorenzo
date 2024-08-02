package keeper_test

import (
	"encoding/json"
	"testing"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/Lorenzo-Protocol/lorenzo/v2/app"
	"github.com/Lorenzo-Protocol/lorenzo/v2/x/bnblightclient/keeper"
	"github.com/Lorenzo-Protocol/lorenzo/v2/x/bnblightclient/types"
)

const (
	bnbRPC = "https://bsc-testnet-dataseed.bnbchain.org"
	testHeight = 42158723
)

type KeeperTestSuite struct {
	suite.Suite

	ctx        sdk.Context
	keeper     keeper.Keeper

	msgServer   types.MsgServer
	queryClient types.QueryClient
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	merge := func(cdc codec.Codec, state map[string]json.RawMessage) {
		genesis := &types.GenesisState{}
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
