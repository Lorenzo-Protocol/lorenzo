package keeper_test

import (
	"encoding/json"
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	lrz "github.com/Lorenzo-Protocol/lorenzo/v3/types"

	btclctypes "github.com/Lorenzo-Protocol/lorenzo/v3/x/btclightclient/types"

	"github.com/cosmos/cosmos-sdk/baseapp"

	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/Lorenzo-Protocol/lorenzo/v3/testutil"
	"github.com/evmos/ethermint/crypto/ethsecp256k1"

	plantypes "github.com/Lorenzo-Protocol/lorenzo/v3/x/plan/types"

	appparams "github.com/Lorenzo-Protocol/lorenzo/v3/app/params"

	evmtypes "github.com/evmos/ethermint/x/evm/types"
	feemarkettypes "github.com/evmos/ethermint/x/feemarket/types"

	"github.com/Lorenzo-Protocol/lorenzo/v3/app"
	"github.com/Lorenzo-Protocol/lorenzo/v3/x/btcstaking/keeper"
	"github.com/Lorenzo-Protocol/lorenzo/v3/x/btcstaking/types"

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
		planGenesis := &plantypes.GenesisState{
			Params: plantypes.Params{
				AllowList: []string{
					testAdmin.String(),
				},
			},
		}
		state[plantypes.ModuleName] = cdc.MustMarshalJSON(planGenesis)

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

		// btclightclient
		baseHeaderHex := "0100000000000000000000000000000000000000000000000000000000000000000000003ba3edfd7a7b12b27ac72c3e67768f617fc81bc3888a51323a9fb8aa4b1e5e4a45068653ffff7f2002000000"
		baseHeader, err := lrz.NewBTCHeaderBytesFromHex(baseHeaderHex)
		suite.Require().NoError(err)
		baseHeaderHashHex := "000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f"
		baseHeaderHash, err := lrz.NewBTCHeaderHashBytesFromHex(baseHeaderHashHex)
		suite.Require().NoError(err)
		genesisWork := sdkmath.NewUint(0)

		baseBtcHeader := btclctypes.NewBTCHeaderInfo(
			&baseHeader,
			&baseHeaderHash,
			0,
			&genesisWork,
		)

		btclcGenesisState := &btclctypes.GenesisState{
			Params: btclctypes.Params{
				InsertHeadersAllowList: []string{testAdmin.String()},
			},
			BaseBtcHeader: *baseBtcHeader,
		}
		state[btclctypes.ModuleName] = cdc.MustMarshalJSON(btclcGenesisState)

		// btcstaking
		params := &types.Params{
			Receivers:             nil,
			BtcConfirmationsDepth: 1,
			MinterAllowList:       []string{testAdmin.String()},
			BridgeAddr:            "0xb7C0817Dd23DE89E4204502dd2C2EF7F57d3A3B8",
			TxoutDustAmount:       546,
		}
		genesis := &types.GenesisState{
			Params: params,
		}

		state[types.ModuleName] = cdc.MustMarshalJSON(genesis)
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
	suite.keeper = &lorenzoApp.BTCStakingKeeper
	suite.lorenzoApp = lorenzoApp

	//
	err = testutil.FundModuleAccount(
		suite.ctx,
		suite.lorenzoApp.BankKeeper,
		types.ModuleName, sdk.NewCoins(sdk.NewCoin(types.NativeTokenDenom, sdk.NewInt(100000000000000))))
	suite.Require().NoError(err)

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
	types.RegisterQueryServer(queryHelper, keeper.NewQuerierImpl(&suite.lorenzoApp.BTCStakingKeeper))
	queryClient := types.NewQueryClient(queryHelper)
	suite.msgServer = keeper.NewMsgServerImpl(&suite.lorenzoApp.BTCStakingKeeper)
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
