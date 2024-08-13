package keeper_test

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Lorenzo-Protocol/lorenzo/v3/contracts/erc20"
	utiltx "github.com/Lorenzo-Protocol/lorenzo/v3/testutil/tx"
	"github.com/Lorenzo-Protocol/lorenzo/v3/x/token/types"
)

func (suite *KeeperTestSuite) TestPostTxProcessing() {
	var (
		receipt      *ethtypes.Receipt
		pair         *types.TokenPair
		account      = utiltx.GenerateAddress()
		erc20ABI     = erc20.ERC20MinterBurnerDecimalsContract.ABI
		transferData = make([]byte, 32)
		msg          = ethtypes.NewMessage(
			types.ModuleAddress,
			&common.Address{},
			0,
			big.NewInt(0),
			uint64(0),
			big.NewInt(0),
			big.NewInt(0),
			big.NewInt(0),
			[]byte{},
			ethtypes.AccessList{},
			true,
		)
	)

	transferEvent := erc20ABI.Events["Transfer"]
	transferData[31] = uint8(10)

	testCases := []struct {
		name          string
		malleate      func()
		expConversion bool
		expectErr     bool
	}{
		{
			"empty logs",
			func() {
				log := ethtypes.Log{}
				receipt = &ethtypes.Receipt{
					Logs: []*ethtypes.Log{&log},
				}
			},
			false,
			false,
		},
		{
			"no data ",
			func() {
				topics := []common.Hash{transferEvent.ID, account.Hash(), types.ModuleAddress.Hash()}
				log := ethtypes.Log{
					Topics: topics,
				}
				receipt = &ethtypes.Receipt{
					Logs: []*ethtypes.Log{&log},
				}
			},
			false,
			false,
		},
		{
			"unknown topics in logs",
			func() {
				topics := []common.Hash{{}, account.Hash(), account.Hash()}
				log := ethtypes.Log{
					Topics: topics,
					Data:   transferData,
				}
				receipt = &ethtypes.Receipt{
					Logs: []*ethtypes.Log{&log},
				}
			},
			false,
			false,
		},
		{
			"no transfer event",
			func() {
				approvalEvent := erc20ABI.Events["Approval"]
				topics := []common.Hash{approvalEvent.ID, account.Hash(), account.Hash()}
				log := ethtypes.Log{
					Topics: topics,
					Data:   transferData,
				}
				receipt = &ethtypes.Receipt{
					Logs: []*ethtypes.Log{&log},
				}
			},
			false,
			false,
		},
		{
			"no contract address in logs",
			func() {
				topics := []common.Hash{transferEvent.ID, account.Hash(), types.ModuleAddress.Hash()}
				log := ethtypes.Log{
					Topics: topics,
					Data:   transferData,
				}
				receipt = &ethtypes.Receipt{
					Logs: []*ethtypes.Log{&log},
				}
			},
			false,
			false,
		},
		{
			"no full topics in logs",
			func() {
				topics := []common.Hash{transferEvent.ID}
				log := ethtypes.Log{
					Topics: topics,
					Data:   transferData,
				}
				receipt = &ethtypes.Receipt{
					Logs: []*ethtypes.Log{&log},
				}
			},
			false,
			false,
		},
		{
			"receiver not token module account",
			func() {
				contractAddr := suite.utilsERC20Deploy("coin", "token", erc20Decimals)
				_, err := suite.app.TokenKeeper.RegisterERC20(suite.ctx, contractAddr)
				suite.Require().NoError(err)

				topics := []common.Hash{transferEvent.ID, account.Hash(), account.Hash()}
				log := ethtypes.Log{
					Topics:  topics,
					Data:    transferData,
					Address: contractAddr,
				}
				receipt = &ethtypes.Receipt{
					Logs: []*ethtypes.Log{&log},
				}
			},
			false,
			false,
		},
		{
			"successfully burn",
			func() {
				resPair, contractAddr := suite.utilsFundAndRegisterERC20("coin", "token", erc20Decimals, account, 1000)
				pair = &resPair
				topics := []common.Hash{transferEvent.ID, account.Hash(), types.ModuleAddress.Hash()}
				log := ethtypes.Log{
					Topics:  topics,
					Data:    transferData,
					Address: contractAddr,
				}
				receipt = &ethtypes.Receipt{
					Logs: []*ethtypes.Log{&log},
				}
			},
			true,
			false,
		},
		{
			"undefined ownership",
			func() {
				resPair, contractAddr := suite.utilsFundAndRegisterERC20("coin", "token", erc20Decimals, account, 1000)
				pair = &resPair

				pair.Source = types.OWNER_UNDEFINED
				suite.app.TokenKeeper.SetTokenPair(suite.ctx, *pair)

				topics := []common.Hash{transferEvent.ID, account.Hash(), types.ModuleAddress.Hash()}
				log := ethtypes.Log{
					Topics:  topics,
					Data:    transferData,
					Address: contractAddr,
				}
				receipt = &ethtypes.Receipt{
					Logs: []*ethtypes.Log{&log},
				}
			},
			false,
			true,
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest()
			suite.utilsAssureAppSetEVMHooks()
			suite.Commit()

			tc.malleate()

			err := suite.app.TokenKeeper.Hooks().PostTxProcessing(suite.ctx, msg, receipt)
			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}

			if tc.expConversion {
				sender := sdk.AccAddress(account.Bytes())
				cosmosBalance := suite.app.BankKeeper.GetBalance(suite.ctx, sender, pair.Denom)

				transferEvent, err := erc20ABI.Unpack("Transfer", transferData)
				suite.Require().NoError(err)

				tokens, _ := transferEvent[0].(*big.Int)
				suite.Require().Equal(cosmosBalance.Amount.String(), tokens.String())
			}
		})
	}
}

func (suite *KeeperTestSuite) utilsAssureAppSetEVMHooks() {
	defer func() {
		err := recover()
		suite.Require().NotNil(err)
	}()
	suite.app.EvmKeeper.SetHooks(suite.app.TokenKeeper.Hooks())
}
