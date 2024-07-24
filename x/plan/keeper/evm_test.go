package keeper_test

import (
	"fmt"

	contractsplan "github.com/Lorenzo-Protocol/lorenzo/v2/contracts/plan"
	utiltx "github.com/Lorenzo-Protocol/lorenzo/v2/testutil/tx"
	"github.com/Lorenzo-Protocol/lorenzo/v2/x/plan/types"
	"github.com/ethereum/go-ethereum/common"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
)

func (suite *KeeperTestSuite) TestCallEVM() {
	testCases := []struct {
		name    string
		method  string
		expPass bool
	}{
		{
			"unknown method",
			"",
			false,
		},
		{
			"pass",
			"balanceOf",
			true,
		},
	}
	for _, tc := range testCases {
		suite.SetupTest() // reset

		erc20 := contractsplan.YieldAccruingTokenContract.ABI
		contract, err := suite.lorenzoApp.PlanKeeper.DeployYATContract(
			suite.ctx, "lorenzo", "ALRZ")
		suite.Require().NoError(err)
		account := utiltx.GenerateAddress()

		res, err := suite.lorenzoApp.PlanKeeper.CallEVM(suite.ctx, erc20, types.ModuleAddress, contract, true, tc.method, account)
		if tc.expPass {
			suite.Require().IsTypef(&evmtypes.MsgEthereumTxResponse{}, res, tc.name)
			suite.Require().NoError(err)
		} else {
			suite.Require().Error(err)
		}
	}
}

func (suite *KeeperTestSuite) TestCallEVMWithData() {
	yat := contractsplan.YieldAccruingTokenContract.ABI
	testCases := []struct {
		name     string
		from     common.Address
		malleate func() ([]byte, *common.Address)
		expPass  bool
	}{
		{
			"unknown method",
			types.ModuleAddress,
			func() ([]byte, *common.Address) {
				contract, err := suite.lorenzoApp.PlanKeeper.DeployYATContract(
					suite.ctx, "lorenzo", "ALRZ")
				suite.Require().NoError(err)
				account := utiltx.GenerateAddress()
				data, _ := yat.Pack("", account)
				return data, &contract
			},
			false,
		},
		{
			"pass",
			types.ModuleAddress,
			func() ([]byte, *common.Address) {
				contract, err := suite.lorenzoApp.PlanKeeper.DeployYATContract(
					suite.ctx, "lorenzo", "ALRZ")
				suite.Require().NoError(err)
				account := utiltx.GenerateAddress()
				data, _ := yat.Pack("balanceOf", account)
				return data, &contract
			},
			true,
		},
		{
			"fail empty data",
			types.ModuleAddress,
			func() ([]byte, *common.Address) {
				contract, err := suite.lorenzoApp.PlanKeeper.DeployYATContract(
					suite.ctx, "lorenzo", "ALRZ")
				suite.Require().NoError(err)
				return []byte{}, &contract
			},
			false,
		},

		{
			"fail empty sender",
			common.Address{},
			func() ([]byte, *common.Address) {
				contract, err := suite.lorenzoApp.PlanKeeper.DeployYATContract(
					suite.ctx, "lorenzo", "ALRZ")
				suite.Require().NoError(err)
				return []byte{}, &contract
			},
			false,
		},
		{
			"deploy",
			types.ModuleAddress,
			func() ([]byte, *common.Address) {
				suite.Commit()
				ctorArgs, _ := contractsplan.YieldAccruingTokenContract.ABI.Pack("", "test", "test", types.ModuleAddress)
				data := append(contractsplan.YieldAccruingTokenContract.Bin, ctorArgs...) //nolint:gocritic
				return data, nil
			},
			true,
		},
		{
			"fail deploy",
			types.ModuleAddress,
			func() ([]byte, *common.Address) {
				suite.Commit()
				params := suite.lorenzoApp.EvmKeeper.GetParams(suite.ctx)
				params.EnableCreate = false
				_ = suite.lorenzoApp.EvmKeeper.SetParams(suite.ctx, params)
				ctorArgs, _ := contractsplan.YieldAccruingTokenContract.ABI.Pack("", "test", "test", types.ModuleAddress)
				data := append(contractsplan.YieldAccruingTokenContract.Bin, ctorArgs...) //nolint:gocritic
				return data, nil
			},
			false,
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.SetupTest() // reset

			data, contract := tc.malleate()

			res, err := suite.lorenzoApp.PlanKeeper.CallEVMWithData(suite.ctx, tc.from, contract, data, true)
			if tc.expPass {
				suite.Require().IsTypef(&evmtypes.MsgEthereumTxResponse{}, res, tc.name)
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}
