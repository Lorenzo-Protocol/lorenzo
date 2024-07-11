package keeper_test

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/Lorenzo-Protocol/lorenzo/contracts/erc20"
	"github.com/Lorenzo-Protocol/lorenzo/x/token/types"
)

// EVM Test Helpers

func (suite *KeeperTestSuite) DeployERC20ContractWithCommit(name, symbol string, decimals uint8) (common.Address, error) {
	contractArgs, err := erc20.ERC20MinterBurnerDecimalsContract.ABI.Pack("", name, symbol, decimals)
	if err != nil {
		return common.Address{}, err
	}

	data := make([]byte, len(erc20.ERC20MinterBurnerDecimalsContract.Bin)+len(contractArgs))
	copy(data[:len(erc20.ERC20MinterBurnerDecimalsContract.Bin)], erc20.ERC20MinterBurnerDecimalsContract.Bin)
	copy(data[len(erc20.ERC20MinterBurnerDecimalsContract.Bin):], contractArgs)

	nonce, err := suite.app.AccountKeeper.GetSequence(suite.ctx, types.ModuleAddress.Bytes())
	if err != nil {
		return common.Address{}, err
	}

	contractAddr := crypto.CreateAddress(types.ModuleAddress, nonce)
	_, err = suite.app.TokenKeeper.CallEVMWithData(suite.ctx, types.ModuleAddress, nil, data, true)
	if err != nil {
		return common.Address{}, err
	}

	suite.Commit()
	return contractAddr, nil
}

// IBC Test Helpers

func (suite *KeeperTestSuite) SetupIBCTest() {
}
