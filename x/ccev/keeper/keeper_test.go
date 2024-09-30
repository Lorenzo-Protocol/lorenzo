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
	"github.com/Lorenzo-Protocol/lorenzo/v3/x/ccev/keeper"
	"github.com/Lorenzo-Protocol/lorenzo/v3/x/ccev/types"
)

var testAccounts = app.CreateTestAddrs(2)

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
	suite.T().Log("setting up keeper test suite")
	merge := func(cdc codec.Codec, state map[string]json.RawMessage) {
		genesis := &types.GenesisState{
			Params: &types.Params{
				AllowList: []string{
					testAccounts[0].String(),
				},
			},
		}
		state[types.ModuleName] = cdc.MustMarshalJSON(genesis)
	}

	lorenzoApp := app.SetupWithGenesisMergeFn(suite.T(), merge)
	suite.ctx = lorenzoApp.BaseApp.NewContext(false, tmproto.Header{})
	suite.keeper = lorenzoApp.CCEVkeeper

	queryHelper := baseapp.NewQueryServerTestHelper(suite.ctx, lorenzoApp.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, keeper.NewQuerier(suite.keeper))
	// types.RegisterMsgServer(queryHelper, keeper.NewMsgServerImpl(suite.keeper))

	queryClient := types.NewQueryClient(queryHelper)
	suite.msgServer = keeper.NewMsgServerImpl(suite.keeper)
	suite.queryClient = queryClient
}

func (suite *KeeperTestSuite) CreateClient(chainID uint32, chainName string, initialBlock types.TinyHeader) {
	err := suite.keeper.CreateClient(suite.ctx, &types.Client{
		ChainId:      chainID,
		ChainName:    chainName,
		InitialBlock: initialBlock,
	})
	suite.Require().NoError(err)
}

func (suite *KeeperTestSuite) UploadContract(chainID uint32,address string,eventName string,abi []byte) {
	err := suite.keeper.UploadContract(suite.ctx, chainID, address, eventName, abi)
	suite.Require().NoError(err)
}

func (suite *KeeperTestSuite) UploadHeaders(chainID uint32, headers []types.TinyHeader) {
	err := suite.keeper.UploadHeaders(suite.ctx, chainID, headers)
	suite.Require().NoError(err)
}
