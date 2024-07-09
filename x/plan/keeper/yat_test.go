package keeper_test

import (
	"fmt"
	"math/big"

	"github.com/Lorenzo-Protocol/lorenzo/x/plan/types"

	"github.com/ethereum/go-ethereum/common"
)

func (suite *KeeperTestSuite) TestYatMint() {
	testCases := []struct {
		name      string
		malleate  func() common.Address
		expectErr bool
	}{
		{
			name: "fail - invalid yat",
			malleate: func() common.Address {
				return common.Address{}
			},
			expectErr: true,
		},
		{
			name: "fail - sender not minter",
			malleate: func() common.Address {
				// deploy yat
				yatAddr, err := suite.lorenzoApp.PlanKeeper.DeployYATContract(
					suite.ctx, "lorenzo", "ALRZ")
				suite.Require().NoError(err)
				return yatAddr
			},
			expectErr: true,
		},
		{
			name: "success - mint",
			malleate: func() common.Address {
				// deploy yat
				yatAddr, err := suite.lorenzoApp.PlanKeeper.DeployYATContract(
					suite.ctx, "lorenzo", "ALRZ")
				suite.Require().NoError(err)
				// set minter
				err = suite.lorenzoApp.PlanKeeper.SetMinter(suite.ctx, yatAddr, types.ModuleAddress)
				suite.Require().NoError(err)
				return yatAddr
			},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("YatMint - %s", tc.name), func() {
			suite.SetupTest() // reset

			yatAddress := tc.malleate()

			testAdminEthAddr := common.BytesToAddress(testAdmin.Bytes())
			err := suite.lorenzoApp.PlanKeeper.Mint(
				suite.ctx, yatAddress,
				testAdminEthAddr,
				big.NewInt(100000),
			)
			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestYatHasRoleFromYAT() {
	testCases := []struct {
		name      string
		malleate  func() common.Address
		expectErr bool
		expectRes bool
	}{
		{
			name: "fail - invalid yat",
			malleate: func() common.Address {
				return common.Address{}
			},
			expectErr: true,
		},
		{
			name: "fail - sender not minter",
			malleate: func() common.Address {
				// deploy yat
				yatAddr, err := suite.lorenzoApp.PlanKeeper.DeployYATContract(
					suite.ctx, "lorenzo", "ALRZ")
				suite.Require().NoError(err)
				return yatAddr
			},
			expectErr: false,
			expectRes: false,
		},
		{
			name: "success - sender is minter",
			malleate: func() common.Address {
				// deploy yat
				yatAddr, err := suite.lorenzoApp.PlanKeeper.DeployYATContract(
					suite.ctx, "lorenzo", "ALRZ")
				suite.Require().NoError(err)
				// set minter
				err = suite.lorenzoApp.PlanKeeper.SetMinter(suite.ctx, yatAddr, types.ModuleAddress)
				suite.Require().NoError(err)
				return yatAddr
			},
			expectErr: false,
			expectRes: true,
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("YatHasRoleFromYAT - %s", tc.name), func() {
			suite.SetupTest() // reset

			yatAddress := tc.malleate()
			found, err := suite.lorenzoApp.PlanKeeper.HasRoleFromYAT(
				suite.ctx,
				yatAddress,
				"YAT_MINTER_ROLE",
				types.ModuleAddress,
			)
			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expectRes, found)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestYatSetMinter() {
	testCases := []struct {
		name       string
		malleate   func() common.Address
		validation func(common.Address)
	}{
		{
			name: "success - set minter",
			malleate: func() common.Address {
				// deploy yat
				yatAddr, err := suite.lorenzoApp.PlanKeeper.DeployYATContract(
					suite.ctx, "lorenzo", "ALRZ")
				suite.Require().NoError(err)
				return yatAddr
			},
			validation: func(yatAddr common.Address) {
				// get minter
				found, err := suite.lorenzoApp.PlanKeeper.HasRoleFromYAT(
					suite.ctx,
					yatAddr,
					"YAT_MINTER_ROLE",
					types.ModuleAddress)
				suite.Require().NoError(err)
				suite.Require().True(found)
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("YatSetMinter - %s", tc.name), func() {
			suite.SetupTest() // reset

			yatAddress := tc.malleate()
			err := suite.lorenzoApp.PlanKeeper.SetMinter(
				suite.ctx, yatAddress, types.ModuleAddress)
			suite.Require().NoError(err)
			tc.validation(yatAddress)
		})
	}
}

func (suite *KeeperTestSuite) TestYatRemoveMinter() {
	testCases := []struct {
		name      string
		malleate  func() common.Address
		expectErr bool
	}{
		{
			name: "success - remove minter",
			malleate: func() common.Address {
				// deploy yat
				yatAddr, err := suite.lorenzoApp.PlanKeeper.DeployYATContract(
					suite.ctx, "lorenzo", "ALRZ")
				suite.Require().NoError(err)
				// set minter
				err = suite.lorenzoApp.PlanKeeper.SetMinter(suite.ctx, yatAddr, types.ModuleAddress)
				suite.Require().NoError(err)
				return yatAddr
			},
			expectErr: false,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("YatRemoveMinter - %s", tc.name), func() {
			suite.SetupTest() // reset

			yatAddress := tc.malleate()
			err := suite.lorenzoApp.PlanKeeper.RemoveMinter(suite.ctx, yatAddress, types.ModuleAddress)
			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestYatGetOwner() {
	testCases := []struct {
		name     string
		malleate func() common.Address
	}{
		{
			name: "success - get owner",
			malleate: func() common.Address {
				// deploy yat
				yatAddr, err := suite.lorenzoApp.PlanKeeper.DeployYATContract(
					suite.ctx, "lorenzo", "ALRZ")
				suite.Require().NoError(err)
				return yatAddr
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("YatRemoveMinter - %s", tc.name), func() {
			suite.SetupTest() // reset
			yatAddress := tc.malleate()
			owner, err := suite.lorenzoApp.PlanKeeper.GetOwner(suite.ctx, yatAddress)
			suite.Require().NoError(err)
			suite.Require().Equal(owner, types.ModuleAddress)
		})
	}
}
