package ibc

import (
	"encoding/json"
	"testing"
	"time"

	feemarkettypes "github.com/evmos/ethermint/x/feemarket/types"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/require"

	dbm "github.com/cometbft/cometbft-db"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	tmprotoversion "github.com/cometbft/cometbft/proto/tendermint/version"
	tmtypes "github.com/cometbft/cometbft/types"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/server"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/cosmos/ibc-go/v7/modules/core/keeper"
	ibctestingtypes "github.com/cosmos/ibc-go/v7/testing/types"

	"github.com/Lorenzo-Protocol/lorenzo/v3/app"
	appparams "github.com/Lorenzo-Protocol/lorenzo/v3/app/params"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
)

var DefaultTestingAppInit = SetupTestingApp

type TestingApp interface {
	abci.Application

	// ibc-go additions
	GetBaseApp() *baseapp.BaseApp
	GetStakingKeeper() ibctestingtypes.StakingKeeper
	GetIBCKeeper() *keeper.Keeper
	GetScopedIBCKeeper() capabilitykeeper.ScopedKeeper
	GetTxConfig() client.TxConfig

	// Implemented by SimApp
	AppCodec() codec.Codec

	// Implemented by BaseApp
	LastCommitID() storetypes.CommitID
	LastBlockHeight() int64
}

// SetupTestingApp is a test helper function to initialize a new Lorenzo app.
func SetupTestingApp() (TestingApp, map[string]json.RawMessage) {
	db := dbm.NewMemDB()
	appOptions := make(simtestutil.AppOptionsMap)
	appOptions[server.FlagInvCheckPeriod] = 5
	appOptions[server.FlagMinGasPrices] = "10alrz"
	appOptions["btc-config.network"] = "simnet"

	invCheckPeriod := cast.ToUint(appOptions.Get(server.FlagInvCheckPeriod))
	encodingConfig := app.MakeEncodingConfig()
	lorenzoApp := app.NewLorenzoApp(
		log.NewNopLogger(),
		db,
		nil,
		true,
		map[int64]bool{},
		app.DefaultNodeHome,
		invCheckPeriod,
		encodingConfig,
		appOptions,
	)

	return lorenzoApp, app.NewDefaultGenesisState(encodingConfig.Codec)
}

// SetupWithGenesisValSet initializes a new LorenzoApp with a validator set and genesis accounts
// that also act as delegators.
func SetupWithGenesisValSet(t *testing.T, valSet *tmtypes.ValidatorSet, genAccs []authtypes.GenesisAccount, chainID string, powerReduction math.Int, balances ...banktypes.Balance) TestingApp {
	return SetupWithGenesisValSetAndConsensusParams(t, app.DefaultConsensusParams, valSet, genAccs, chainID, powerReduction, balances...)
}

func SetupWithGenesisValSetAndConsensusParams(t *testing.T, consensusParams *tmproto.ConsensusParams, valSet *tmtypes.ValidatorSet, genAccs []authtypes.GenesisAccount, chainID string, powerReduction math.Int, balances ...banktypes.Balance) TestingApp {
	app, genesisState := DefaultTestingAppInit()

	// ensure baseapp has a chain-id set before running InitChain
	baseapp.SetChainID(chainID)(app.GetBaseApp())

	// set genesis accounts
	authGenesis := authtypes.NewGenesisState(authtypes.DefaultParams(), genAccs)
	genesisState[authtypes.ModuleName] = app.AppCodec().MustMarshalJSON(authGenesis)

	validators := make([]stakingtypes.Validator, 0, len(valSet.Validators))
	delegations := make([]stakingtypes.Delegation, 0, len(valSet.Validators))

	bondAmt := sdk.TokensFromConsensusPower(1, powerReduction)

	for _, val := range valSet.Validators {
		pk, err := cryptocodec.FromTmPubKeyInterface(val.PubKey)
		require.NoError(t, err)
		pkAny, err := codectypes.NewAnyWithValue(pk)
		require.NoError(t, err)
		validator := stakingtypes.Validator{
			OperatorAddress:   sdk.ValAddress(val.Address).String(),
			ConsensusPubkey:   pkAny,
			Jailed:            false,
			Status:            stakingtypes.Bonded,
			Tokens:            bondAmt,
			DelegatorShares:   sdk.OneDec(),
			Description:       stakingtypes.Description{},
			UnbondingHeight:   int64(0),
			UnbondingTime:     time.Unix(0, 0).UTC(),
			Commission:        stakingtypes.NewCommission(sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec()),
			MinSelfDelegation: sdk.ZeroInt(),
		}

		validators = append(validators, validator)
		delegations = append(delegations, stakingtypes.NewDelegation(genAccs[0].GetAddress(), val.Address.Bytes(), sdk.OneDec()))
	}

	// set validators and delegations
	var stakingGenesis stakingtypes.GenesisState
	app.AppCodec().MustUnmarshalJSON(genesisState[stakingtypes.ModuleName], &stakingGenesis)

	bondDenom := stakingGenesis.Params.BondDenom

	// add bonded amount to bonded pool module account
	balances = append(balances, banktypes.Balance{
		Address: authtypes.NewModuleAddress(stakingtypes.BondedPoolName).String(),
		Coins:   sdk.Coins{sdk.NewCoin(bondDenom, bondAmt.Mul(sdk.NewInt(int64(len(valSet.Validators)))))},
	})

	// set validators and delegations
	stakingGenesis = *stakingtypes.NewGenesisState(stakingGenesis.Params, validators, delegations)
	genesisState[stakingtypes.ModuleName] = app.AppCodec().MustMarshalJSON(&stakingGenesis)

	// update total supply
	bankGenesis := banktypes.NewGenesisState(banktypes.DefaultGenesisState().Params, balances, sdk.NewCoins(), []banktypes.Metadata{}, []banktypes.SendEnabled{})
	genesisState[banktypes.ModuleName] = app.AppCodec().MustMarshalJSON(bankGenesis)

	// NOTE: add lorenzo custom modules genesis here.

	// set evm genesis: use lrz as evm denom.
	evmParams := evmtypes.DefaultParams()
	evmParams.EvmDenom = appparams.BaseDenom
	evmGenesis := &evmtypes.GenesisState{Params: evmParams}
	genesisState[evmtypes.ModuleName] = app.AppCodec().MustMarshalJSON(evmGenesis)

	// set feemarket: disable fee market
	feemarketGenesis := feemarkettypes.NewGenesisState(feemarkettypes.DefaultParams(), 0)
	feemarketGenesis.Params.NoBaseFee = true
	genesisState[feemarkettypes.ModuleName] = app.AppCodec().MustMarshalJSON(feemarketGenesis)

	stateBytes, err := json.MarshalIndent(genesisState, "", " ")
	require.NoError(t, err)

	// init chain will set the validator set and initialize the genesis accounts
	app.InitChain(
		abci.RequestInitChain{
			ChainId:       chainID,
			Validators:    []abci.ValidatorUpdate{},
			AppStateBytes: stateBytes,
		},
	)

	// commit genesis changes
	app.Commit()
	app.BeginBlock(
		abci.RequestBeginBlock{
			Header: tmproto.Header{
				Version: tmprotoversion.Consensus{
					App: consensusParams.Version.App,
				},
				ChainID:            chainID,
				Height:             app.LastBlockHeight() + 1,
				AppHash:            app.LastCommitID().Hash,
				ValidatorsHash:     valSet.Hash(),
				NextValidatorsHash: valSet.Hash(),
				ProposerAddress:    valSet.Proposer.Address.Bytes(),
			},
		},
	)

	return app
}
