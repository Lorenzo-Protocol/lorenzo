package helpers

import (
	"bytes"
	"encoding/json"
	"strconv"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/require"

	dbm "github.com/cometbft/cometbft-db"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	tmtypes "github.com/cometbft/cometbft/types"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/testutil/mock"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/Lorenzo-Protocol/lorenzo/app"
)

// SimAppChainID hardcoded chainID for simulation
const (
	SimAppChainID = "lorenzo_9001-1"
)

// DefaultConsensusParams defines the default Tendermint consensus params used
// in LorenzoApp testing.
var DefaultConsensusParams = &tmproto.ConsensusParams{
	Block: &tmproto.BlockParams{
		MaxBytes: 200000,
		MaxGas:   2000000,
	},
	Evidence: &tmproto.EvidenceParams{
		MaxAgeNumBlocks: 302400,
		MaxAgeDuration:  504 * time.Hour, // 3 weeks is the max duration
		MaxBytes:        10000,
	},
	Validator: &tmproto.ValidatorParams{
		PubKeyTypes: []string{
			tmtypes.ABCIPubKeyTypeEd25519,
		},
	},
}

// PV implements PrivValidator interface
type PV struct {
	PrivKey cryptotypes.PrivKey
}

// EmptyAppOptions is a stub implementing AppOptions
type EmptyAppOptions struct{}

// Get implements AppOptions
func (EmptyAppOptions) Get(_ string) interface{} { return nil }

// Setup initializes and returns a new instance of app.LorenzoApp for testing purposes.
//
// Parameter:
// - t: A testing.T instance for testing purposes.
// Returns:
// - *app.LorenzoApp: The initialized instance of app.LorenzoApp.
func Setup(t *testing.T) *app.LorenzoApp {
	t.Helper()

	privVal := mock.NewPV()
	pubKey, err := privVal.GetPubKey()
	require.NoError(t, err)
	// create validator set with single validator
	validator := tmtypes.NewValidator(pubKey, 1)
	valSet := tmtypes.NewValidatorSet([]*tmtypes.Validator{validator})

	// generate genesis account
	senderPrivKey := mock.NewPV()
	senderPubKey := senderPrivKey.PrivKey.PubKey()

	acc := authtypes.NewBaseAccount(senderPubKey.Address().Bytes(), senderPubKey, 0, 0)
	balance := banktypes.Balance{
		Address: acc.GetAddress().String(),
		Coins:   sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100000000000000))),
	}
	genesisAccounts := []authtypes.GenesisAccount{acc}
	return SetupWithGenesisValSet(t, valSet, genesisAccounts, balance)
}

// SetupWithGenesisMergeFn initializes a new LorenzoApp with override module genesis data.
//
// Parameters:
// - t: A testing.T instance for testing purposes.
// - merge: A function to merge override module genesis data with the default genesis data.
// Returns:
// - *app.LorenzoApp: The initialized instance of app.LorenzoApp.
func SetupWithGenesisMergeFn(t *testing.T, merge func(cdc codec.Codec, state map[string]json.RawMessage)) *app.LorenzoApp {
	t.Helper()

	privVal := mock.NewPV()
	pubKey, err := privVal.GetPubKey()
	require.NoError(t, err)
	// create validator set with single validator
	validator := tmtypes.NewValidator(pubKey, 1)
	valSet := tmtypes.NewValidatorSet([]*tmtypes.Validator{validator})

	// generate genesis account
	senderPrivKey := mock.NewPV()
	senderPubKey := senderPrivKey.PrivKey.PubKey()

	acc := authtypes.NewBaseAccount(senderPubKey.Address().Bytes(), senderPubKey, 0, 0)
	balance := banktypes.Balance{
		Address: acc.GetAddress().String(),
		Coins:   sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100000000000000))),
	}
	genAccs := []authtypes.GenesisAccount{acc}

	lorenzoApp, genesisState := setup()
	if merge != nil {
		merge(lorenzoApp.AppCodec(), genesisState)
	}
	genesisState = genesisStateWithValSet(t, lorenzoApp, genesisState, valSet, genAccs, balance)

	stateBytes, err := json.MarshalIndent(genesisState, "", " ")
	require.NoError(t, err)

	// init chain will set the validator set and initialize the genesis accounts
	lorenzoApp.InitChain(
		abci.RequestInitChain{
			ChainId:         SimAppChainID,
			Validators:      []abci.ValidatorUpdate{},
			ConsensusParams: DefaultConsensusParams,
			AppStateBytes:   stateBytes,
		},
	)

	// commit genesis changes
	lorenzoApp.Commit()
	lorenzoApp.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{
		ChainID:            SimAppChainID,
		Height:             lorenzoApp.LastBlockHeight() + 1,
		AppHash:            lorenzoApp.LastCommitID().Hash,
		ValidatorsHash:     valSet.Hash(),
		NextValidatorsHash: valSet.Hash(),
	}})

	return lorenzoApp
}

// SetupWithGenesisValSet initializes a new LorenzoApp with a validator set and genesis accounts
// that also act as delegators. For simplicity, each validator is bonded with a delegation
// of one consensus engine unit in the default token of the LorenzoApp from first genesis
// account. A Nop logger is set in LorenzoApp.
func SetupWithGenesisValSet(t *testing.T, valSet *tmtypes.ValidatorSet, genAccs []authtypes.GenesisAccount, balances ...banktypes.Balance) *app.LorenzoApp {
	t.Helper()

	lorenzoApp, genesisState := setup()
	genesisState = genesisStateWithValSet(t, lorenzoApp, genesisState, valSet, genAccs, balances...)

	stateBytes, err := json.MarshalIndent(genesisState, "", " ")
	require.NoError(t, err)

	// init chain will set the validator set and initialize the genesis accounts
	lorenzoApp.InitChain(
		abci.RequestInitChain{
			ChainId:         SimAppChainID,
			Validators:      []abci.ValidatorUpdate{},
			ConsensusParams: DefaultConsensusParams,
			AppStateBytes:   stateBytes,
		},
	)

	// commit genesis changes
	lorenzoApp.Commit()
	lorenzoApp.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{
		ChainID:            SimAppChainID,
		Height:             lorenzoApp.LastBlockHeight() + 1,
		AppHash:            lorenzoApp.LastCommitID().Hash,
		ValidatorsHash:     valSet.Hash(),
		NextValidatorsHash: valSet.Hash(),
	}})

	return lorenzoApp
}

// CreateTestAddrs creates test addresses
func CreateTestAddrs(numAddrs int) []sdk.AccAddress {
	var addresses []sdk.AccAddress
	var buffer bytes.Buffer

	// start at 100 so we can make up to 999 test addresses with valid test addresses
	for i := 100; i < (numAddrs + 100); i++ {
		numString := strconv.Itoa(i)
		buffer.WriteString("A58856F0FD53BF058B4909A21AEC019107BA6") //base address string

		buffer.WriteString(numString) //adding on final two digits to make addresses unique
		res, _ := sdk.AccAddressFromHexUnsafe(buffer.String())
		bech := res.String()
		addresses = append(addresses, testAddr(buffer.String(), bech))
		buffer.Reset()
	}
	return addresses
}

func setup() (*app.LorenzoApp, app.GenesisState) {
	db := dbm.NewMemDB()
	appOptions := make(simtestutil.AppOptionsMap)
	appOptions[server.FlagInvCheckPeriod] = 5
	appOptions[server.FlagMinGasPrices] = "10alrz"
	appOptions["btc-config.network"] = "simnet"

	invCheckPeriod := cast.ToUint(appOptions.Get(server.FlagInvCheckPeriod))
	encConfig := app.MakeEncodingConfig()
	LorenzoApp := app.NewLorenzoApp(
		log.NewNopLogger(),
		db,
		nil,
		true,
		map[int64]bool{},
		app.DefaultNodeHome,
		invCheckPeriod,
		encConfig,
		appOptions,
		baseapp.SetChainID(SimAppChainID),
	)
	return LorenzoApp, app.NewDefaultGenesisState(encConfig.Codec)
}

func genesisStateWithValSet(t *testing.T,
	app *app.LorenzoApp, genesisState app.GenesisState,
	valSet *tmtypes.ValidatorSet, genAccs []authtypes.GenesisAccount,
	balances ...banktypes.Balance,
) app.GenesisState {
	t.Helper()
	// set genesis accounts
	authGenesis := authtypes.NewGenesisState(authtypes.DefaultParams(), genAccs)
	genesisState[authtypes.ModuleName] = app.AppCodec().MustMarshalJSON(authGenesis)

	validators := make([]stakingtypes.Validator, 0, len(valSet.Validators))
	delegations := make([]stakingtypes.Delegation, 0, len(valSet.Validators))

	bondAmt := sdk.DefaultPowerReduction

	for _, val := range valSet.Validators {
		pk, err := cryptocodec.FromTmPubKeyInterface(val.PubKey)
		require.NoError(t, err)
		pkAny, err := codectypes.NewAnyWithValue(pk)
		require.NoError(t, err)
		validator := stakingtypes.Validator{
			OperatorAddress: sdk.ValAddress(val.Address).String(),
			ConsensusPubkey: pkAny,
			Jailed:          false,
			Status:          stakingtypes.Bonded,
			Tokens:          bondAmt,
			DelegatorShares: sdk.OneDec(),
			Description:     stakingtypes.Description{},
			UnbondingHeight: int64(0),
			UnbondingTime:   time.Unix(0, 0).UTC(),
			Commission:      stakingtypes.NewCommission(sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec()),
		}
		validators = append(validators, validator)
		delegations = append(delegations, stakingtypes.NewDelegation(genAccs[0].GetAddress(), val.Address.Bytes(), sdk.OneDec()))

	}
	// set validators and delegations
	stakingGenesis := stakingtypes.NewGenesisState(stakingtypes.DefaultParams(), validators, delegations)
	genesisState[stakingtypes.ModuleName] = app.AppCodec().MustMarshalJSON(stakingGenesis)

	totalSupply := sdk.NewCoins()
	for _, b := range balances {
		// add genesis acc tokens to total supply
		totalSupply = totalSupply.Add(b.Coins...)
	}

	for range delegations {
		// add delegated tokens to total supply
		totalSupply = totalSupply.Add(sdk.NewCoin(sdk.DefaultBondDenom, bondAmt))
	}

	// add bonded amount to bonded pool module account
	balances = append(balances, banktypes.Balance{
		Address: authtypes.NewModuleAddress(stakingtypes.BondedPoolName).String(),
		Coins:   sdk.Coins{sdk.NewCoin(sdk.DefaultBondDenom, bondAmt)},
	})

	// update total supply
	bankGenesis := banktypes.NewGenesisState(banktypes.DefaultGenesisState().Params, balances, totalSupply, []banktypes.Metadata{}, []banktypes.SendEnabled{})
	genesisState[banktypes.ModuleName] = app.AppCodec().MustMarshalJSON(bankGenesis)

	return genesisState
}

func testAddr(addr string, bech string) sdk.AccAddress {
	res, err := sdk.AccAddressFromHexUnsafe(addr)
	if err != nil {
		panic(err)
	}
	bechexpected := res.String()
	if bech != bechexpected {
		panic("Bech encoding doesn't match reference")
	}

	bechres, err := sdk.AccAddressFromBech32(bech)
	if err != nil {
		panic(err)
	}
	if !bytes.Equal(bechres, res) {
		panic("Bech decode and hex decode don't match")
	}

	return res
}
