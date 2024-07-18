package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/Lorenzo-Protocol/lorenzo/contracts/erc20"
	"github.com/Lorenzo-Protocol/lorenzo/x/token/types"
)

// utilsFundAndRegisterCoin funds receiver amount of coin and register a token pair for that coin.
func (suite *KeeperTestSuite) utilsFundAndRegisterCoin(
	metadata banktypes.Metadata,
	receiver sdk.AccAddress,
	amount int64,
) types.TokenPair {
	coins := sdk.NewCoins(sdk.NewInt64Coin(metadata.Base, amount))

	err := suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, sdk.NewCoins(coins...))
	suite.Require().NoError(err)
	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, receiver, sdk.NewCoins(coins...))
	suite.Require().NoError(err)

	_, err = suite.msgServer.RegisterCoin(suite.ctx, &types.MsgRegisterCoin{
		Authority: authority.String(),
		Metadata:  []banktypes.Metadata{metadata},
	})
	suite.Require().NoError(err)

	id := suite.app.TokenKeeper.GetTokenPairId(suite.ctx, metadata.Base)
	pair, found := suite.app.TokenKeeper.GetTokenPair(suite.ctx, id)
	suite.Require().True(found)
	return pair
}

// utilsFundAndRegisterERC20 funds receiver amount of token and register a token pair for that contract.
func (suite *KeeperTestSuite) utilsFundAndRegisterERC20(
	name, symbol string, decimals uint8,
	receiver common.Address, amount int64,
) (types.TokenPair, common.Address) {
	erc20Addr := suite.utilsERC20Deploy(name, symbol, decimals)
	suite.utilsERC20Mint(erc20Addr, types.ModuleAddress, receiver, amount)

	_, err := suite.msgServer.RegisterERC20(suite.ctx, &types.MsgRegisterERC20{
		Authority:         authority.String(),
		ContractAddresses: []string{erc20Addr.String()},
	})
	suite.Require().NoError(err)

	id := suite.app.TokenKeeper.GetTokenPairId(suite.ctx, erc20Addr.String())
	pair, found := suite.app.TokenKeeper.GetTokenPair(suite.ctx, id)
	suite.Require().True(found)
	return pair, erc20Addr
}

func (suite *KeeperTestSuite) utilsERC20BalanceOf(contract, address common.Address) int64 {
	erc20ABI := erc20.ERC20MinterBurnerDecimalsContract.ABI
	balance := suite.app.TokenKeeper.ERC20BalanceOf(suite.ctx, erc20ABI, contract, address)
	return balance.Int64()
}

// utilsERC20Deploy deploys an ERC20 contract owned by module account.
func (suite *KeeperTestSuite) utilsERC20Deploy(name, symbol string, decimals uint8) common.Address {
	contractArgs, err := erc20.ERC20MinterBurnerDecimalsContract.ABI.Pack("", name, symbol, decimals)
	suite.Require().NoError(err)

	data := make([]byte, len(erc20.ERC20MinterBurnerDecimalsContract.Bin)+len(contractArgs))
	copy(data[:len(erc20.ERC20MinterBurnerDecimalsContract.Bin)], erc20.ERC20MinterBurnerDecimalsContract.Bin)
	copy(data[len(erc20.ERC20MinterBurnerDecimalsContract.Bin):], contractArgs)

	nonce, err := suite.app.AccountKeeper.GetSequence(suite.ctx, types.ModuleAddress.Bytes())
	suite.Require().NoError(err)
	contractAddr := crypto.CreateAddress(types.ModuleAddress, nonce)
	_, err = suite.app.TokenKeeper.CallEVMWithData(suite.ctx, types.ModuleAddress, nil, data, true)
	suite.Require().NoError(err)

	return contractAddr
}

func (suite *KeeperTestSuite) utilsERC20Mint(contract, from, to common.Address, amount int64) {
	_, err := suite.app.TokenKeeper.CallEVM(
		suite.ctx, erc20.ERC20MinterBurnerDecimalsContract.ABI,
		from, contract, true, "mint", to, sdk.NewInt(amount).BigInt())
	suite.Require().NoError(err)
}
