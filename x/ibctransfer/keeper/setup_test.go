package keeper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/evmos/ethermint/crypto/ethsecp256k1"
	evmtypes "github.com/evmos/ethermint/x/evm/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"

	"github.com/Lorenzo-Protocol/lorenzo/v3/app"
	"github.com/Lorenzo-Protocol/lorenzo/v3/contracts/erc20"
	"github.com/Lorenzo-Protocol/lorenzo/v3/testutil"
	utiltx "github.com/Lorenzo-Protocol/lorenzo/v3/testutil/tx"
	tokentypes "github.com/Lorenzo-Protocol/lorenzo/v3/x/token/types"
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

	// query client
	queryHelper := baseapp.NewQueryServerTestHelper(suite.ctx, suite.app.InterfaceRegistry())
	evmtypes.RegisterQueryServer(queryHelper, suite.app.EvmKeeper)
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

var timeoutHeight = clienttypes.NewHeight(1000, 1000)

// utilsERC20Deploy deploys an ERC20 contract owned by module account.
func (suite *KeeperTestSuite) utilsERC20Deploy(name, symbol string, decimals uint8) common.Address {
	contractArgs, err := erc20.ERC20MinterBurnerDecimalsContract.ABI.Pack("", name, symbol, decimals)
	suite.Require().NoError(err)

	data := make([]byte, len(erc20.ERC20MinterBurnerDecimalsContract.Bin)+len(contractArgs))
	copy(data[:len(erc20.ERC20MinterBurnerDecimalsContract.Bin)], erc20.ERC20MinterBurnerDecimalsContract.Bin)
	copy(data[len(erc20.ERC20MinterBurnerDecimalsContract.Bin):], contractArgs)

	nonce, err := suite.app.AccountKeeper.GetSequence(suite.ctx, tokentypes.ModuleAddress.Bytes())
	suite.Require().NoError(err)
	contractAddr := crypto.CreateAddress(tokentypes.ModuleAddress, nonce)
	_, err = suite.app.TokenKeeper.CallEVMWithData(suite.ctx, tokentypes.ModuleAddress, nil, data, true)
	suite.Require().NoError(err)

	return contractAddr
}

func (suite *KeeperTestSuite) utilsERC20Mint(contract, from, to common.Address, amount int64) {
	_, err := suite.app.TokenKeeper.CallEVM(
		suite.ctx, erc20.ERC20MinterBurnerDecimalsContract.ABI,
		from, contract, true, "mint", to, sdk.NewInt(amount).BigInt())
	suite.Require().NoError(err)
}
