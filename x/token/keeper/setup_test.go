package keeper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/ethermint/crypto/ethsecp256k1"
	evmtypes "github.com/evmos/ethermint/x/evm/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/Lorenzo-Protocol/lorenzo/v2/app"
	"github.com/Lorenzo-Protocol/lorenzo/v2/testutil"

	utiltx "github.com/Lorenzo-Protocol/lorenzo/v2/testutil/tx"
	"github.com/Lorenzo-Protocol/lorenzo/v2/x/token/keeper"
	"github.com/Lorenzo-Protocol/lorenzo/v2/x/token/types"
)

type KeeperTestSuite struct {
	suite.Suite

	// app testing
	ctx sdk.Context
	app *app.LorenzoApp

	// account
	priv    cryptotypes.PrivKey
	address common.Address
	signer  keyring.Signer

	// services
	msgServer      types.MsgServer
	queryClient    types.QueryClient
	queryClientEvm evmtypes.QueryClient
}

var s *KeeperTestSuite

func TestKeeperTestSuite(t *testing.T) {
	s = new(KeeperTestSuite)
	suite.Run(t, s)
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.execSetupTest()
}

func (suite *KeeperTestSuite) execSetupTest() {
	// account key
	priv, err := ethsecp256k1.GenerateKey()
	suite.Require().NoError(err)
	suite.priv = priv
	suite.address = common.BytesToAddress(priv.PubKey().Address().Bytes())
	suite.signer = utiltx.NewSigner(priv)

	// consAddress
	privCons, err := ethsecp256k1.GenerateKey()
	suite.Require().NoError(err)
	consAddress := sdk.ConsAddress(privCons.PubKey().Address())

	// init app
	// TODO: setup with genesis merge fn need recheck, it's probably not errorless.
	suite.app = app.SetupWithGenesisMergeFn(suite.T(), nil)
	header := testutil.NewHeader(
		suite.app.LastBlockHeight()+1, time.Now().UTC(), app.SimAppChainID, consAddress, nil, nil,
	)
	suite.ctx = suite.app.GetBaseApp().NewContext(false, header)

	// set validator
	valAddr := sdk.ValAddress(privCons.PubKey().Address().Bytes())
	validator, err := stakingtypes.NewValidator(valAddr, privCons.PubKey(), stakingtypes.Description{})
	suite.Require().NoError(err)
	validator = stakingkeeper.TestingUpdateValidator(suite.app.StakingKeeper, suite.ctx, validator, true)
	err = suite.app.StakingKeeper.Hooks().AfterValidatorCreated(suite.ctx, validator.GetOperator())
	suite.Require().NoError(err)
	err = suite.app.StakingKeeper.SetValidatorByConsAddr(suite.ctx, validator)
	suite.Require().NoError(err)

	// services and query client
	suite.msgServer = keeper.NewMsgServerImpl(suite.app.TokenKeeper)

	queryHelper := baseapp.NewQueryServerTestHelper(suite.ctx, suite.app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, keeper.NewQuerier(suite.app.TokenKeeper))
	evmtypes.RegisterQueryServer(queryHelper, suite.app.EvmKeeper)
	suite.queryClient = types.NewQueryClient(queryHelper)
	suite.queryClientEvm = evmtypes.NewQueryClient(queryHelper)
}

// Commit commits and starts a new block with an updated context.
func (suite *KeeperTestSuite) Commit() {
	suite.CommitAfter(time.Second * 1)
}

// Commit commits a block at a given time.
func (suite *KeeperTestSuite) CommitAfter(t time.Duration) {
	var err error
	suite.ctx, err = testutil.Commit(suite.ctx, suite.app, t, nil)
	suite.Require().NoError(err)
}
