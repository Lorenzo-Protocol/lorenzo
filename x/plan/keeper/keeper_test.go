package keeper_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"

	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/evmos/ethermint/crypto/ethsecp256k1"

	appparams "github.com/Lorenzo-Protocol/lorenzo/v3/app/params"
	evmtypes "github.com/evmos/ethermint/x/evm/types"

	feemarkettypes "github.com/evmos/ethermint/x/feemarket/types"

	"github.com/Lorenzo-Protocol/lorenzo/v3/testutil"

	"github.com/Lorenzo-Protocol/lorenzo/v3/app"

	"github.com/Lorenzo-Protocol/lorenzo/v3/x/plan/keeper"
	"github.com/Lorenzo-Protocol/lorenzo/v3/x/plan/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stretchr/testify/suite"
)

var testAdmin = app.CreateTestAddrs(1)[0]

type KeeperTestSuite struct {
	suite.Suite

	ctx        sdk.Context
	keeper     *keeper.Keeper
	lorenzoApp *app.LorenzoApp

	msgServer   types.MsgServer
	queryClient types.QueryClient
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	merge := func(cdc codec.Codec, state map[string]json.RawMessage) {
		genesis := &types.GenesisState{
			Params: types.Params{
				AllowList: []string{
					testAdmin.String(),
				},
			},
		}
		state[types.ModuleName] = cdc.MustMarshalJSON(genesis)

		feeMarketGenesis := &feemarkettypes.GenesisState{
			Params:   feemarkettypes.DefaultParams(),
			BlockGas: 0,
		}
		state[feemarkettypes.ModuleName] = cdc.MustMarshalJSON(feeMarketGenesis)

		evmParams := evmtypes.DefaultParams()
		evmParams.EvmDenom = appparams.BaseDenom
		evmGenesis := &evmtypes.GenesisState{
			Params: evmParams,
		}
		state[evmtypes.ModuleName] = cdc.MustMarshalJSON(evmGenesis)
	}

	lorenzoApp := app.SetupWithGenesisMergeFn(suite.T(), merge)

	// consensus key
	privCons, err := ethsecp256k1.GenerateKey()
	suite.Require().NoError(err)
	consAddress := sdk.ConsAddress(privCons.PubKey().Address())
	header := testutil.NewHeader(
		lorenzoApp.LastBlockHeight()+1, time.Now().UTC(), app.SimAppChainID, consAddress, nil, nil,
	)
	ctx := lorenzoApp.GetBaseApp().NewContext(false, header)

	suite.ctx = ctx
	suite.keeper = lorenzoApp.PlanKeeper
	suite.lorenzoApp = lorenzoApp

	// setup validators
	valAddr := sdk.ValAddress(privCons.PubKey().Address().Bytes())
	validator, err := stakingtypes.NewValidator(valAddr, privCons.PubKey(), stakingtypes.Description{})
	suite.Require().NoError(err)
	validator = stakingkeeper.TestingUpdateValidator(suite.lorenzoApp.StakingKeeper, suite.ctx, validator, true)
	err = suite.lorenzoApp.StakingKeeper.Hooks().AfterValidatorCreated(suite.ctx, validator.GetOperator())
	suite.Require().NoError(err)
	err = suite.lorenzoApp.StakingKeeper.SetValidatorByConsAddr(suite.ctx, validator)
	suite.Require().NoError(err)

	queryHelper := baseapp.NewQueryServerTestHelper(suite.ctx, suite.lorenzoApp.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, keeper.NewQuerierImpl(suite.lorenzoApp.PlanKeeper))
	queryClient := types.NewQueryClient(queryHelper)
	suite.msgServer = keeper.NewMsgServerImpl(suite.lorenzoApp.PlanKeeper)
	suite.queryClient = queryClient
}

// Commit commits and starts a new block with an updated context.
func (suite *KeeperTestSuite) Commit() {
	suite.CommitAfter(time.Second * 1)
}

// Commit commits a block at a given time.
func (suite *KeeperTestSuite) CommitAfter(t time.Duration) {
	var err error
	suite.ctx, err = testutil.Commit(suite.ctx, suite.lorenzoApp, t, nil)
	suite.Require().NoError(err)
}
