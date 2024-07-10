package keeper_test

import banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

const (
	// erc20 metadata
	erc20Name     = "Coin Token"
	erc20Symbol   = "COIN"
	erc20Decimals = uint8(18)

	// coin metadata
	coinBaseDenom    = "acoin"
	coinDisplay      = "coin"
	coinBaseExponent = uint32(0)
	coinMaxExponent  = uint32(18)
)

var (
	coinMetadata = banktypes.Metadata{
		Description: "",
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    coinBaseDenom,
				Exponent: coinBaseExponent,
			},
			{
				Denom:    coinDisplay,
				Exponent: coinMaxExponent,
			},
		},
		Base:    coinBaseDenom,
		Display: coinDisplay,
		Name:    erc20Name,
		Symbol:  erc20Symbol,
	}
)

func (suite *KeeperTestSuite) TestRegisterCoin() {
	testcases := []struct {
		name       string
		malleate   func()
		expectPass bool
	}{
		{
			name:       "fail: sender is not authority",
			expectPass: false,
		},
		{
			name:       "fail: conversion is disabled",
			expectPass: false,
		},
		{
			name:       "fail: coin already registered",
			expectPass: false,
		},
		{
			name:       "fail: coin has no supply",
			expectPass: false,
		},
		{
			name:       "fail: inconsistent metadata",
			expectPass: false,
		},
		{
			name:       "fail: force evm to fail",
			expectPass: false,
		},
		{
			name:       "success: register coin",
			expectPass: true,
		},
	}

	for _, tc := range testcases {
		suite.Run(tc.name, func() {
			suite.SetupTest() // reset before each case

			if tc.expectPass {

			} else {

			}
		})
	}
}

func (suite *KeeperTestSuite) TestRegisterERC20() {
	testcases := []struct {
		name       string
		malleate   func()
		expectPass bool
	}{
		{
			name:       "success: register erc20",
			expectPass: true,
		},
		{
			name:       "fail: erc20 already registered",
			expectPass: false,
		},
	}

	for _, tc := range testcases {
		suite.Run(tc.name, func() {
			suite.SetupTest() // reset before each case

			// deploy contract
			contractAddr, err := suite.DeployERC20Contract("name", "symbol", 18)

			if tc.expectPass {

			} else {

			}
		})
	}
}

func (suite *KeeperTestSuite) TestToggleConversion() {

}

func (suite *KeeperTestSuite) TestUpdateParams() {

}

func (suite *KeeperTestSuite) TestConvertCoin() {

}

func (suite *KeeperTestSuite) TestConvertERC20() {

}
