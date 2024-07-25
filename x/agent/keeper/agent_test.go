package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestSetAdmin() {
	testCases := []struct {
		name       string
		args       sdk.AccAddress
		expectErr  bool
		validation func(sdk.AccAddress)
	}{
		{
			name:      "fail - admin already set",
			args:      testAdmin,
			expectErr: true,
		},
		{
			name:      "success - set admin",
			args:      sdk.AccAddress("lrz1xa40j022h2rcmnte47gyjg8688grln94pp84lc"),
			expectErr: false,
			validation: func(args sdk.AccAddress) {
				admin := suite.keeper.GetAdmin(suite.ctx)
				suite.Require().Equal(args, admin)
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("KeeperSetAdmin - %s", tc.name), func() {
			suite.SetupTest()
			err := suite.keeper.SetAdmin(suite.ctx, tc.args)
			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
			if tc.validation != nil {
				tc.validation(tc.args)
			}
		})
	}
}
