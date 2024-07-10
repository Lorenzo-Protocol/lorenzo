package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
	evmtypes "github.com/evmos/ethermint/x/evm/types"

	ibcgotesting "github.com/cosmos/ibc-go/v7/testing"

	"github.com/Lorenzo-Protocol/lorenzo/app"
	"github.com/Lorenzo-Protocol/lorenzo/x/token/types"
)

type KeeperTestSuite struct {
	suite.Suite

	// app testing
	ctx sdk.Context
	app *app.LorenzoApp

	queryClientEvm evmtypes.QueryClient
	queryClient    types.QueryClient

	// ibc-go testing
	LorenzoChain *ibcgotesting.TestChain
	CosmosChain  *ibcgotesting.TestChain
}

var (
	s *KeeperTestSuite
)

func TestKeeperTestSuite(t *testing.T) {
	s = new(KeeperTestSuite)
	suite.Run(t, s)
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.execSetupTest()
}

func (suite *KeeperTestSuite) execSetupTest() {

}
