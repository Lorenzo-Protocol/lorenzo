package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestAuthorized() {
	testCases := []struct {
		name         string
		args         sdk.AccAddress
		malleate     func()
		expectResult bool
	}{
		{
			name:         "invalid address - invalid bech32",
			args:         sdk.AccAddress("foobar"),
			expectResult: false,
		},
		{
			name:         "invalid address - empty",
			args:         sdk.AccAddress(nil),
			expectResult: false,
		},
		{
			name: "success - valid address",
			args: testAdmin,
			malleate: func() {
				suite.Commit()
			},
			expectResult: true,
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("KeeperAuthorized - %s", tc.name), func() {
			suite.SetupTest()
			if tc.malleate != nil {
				tc.malleate()
			}
			authorized := suite.keeper.Authorized(suite.ctx, tc.args)
			suite.Require().Equal(tc.expectResult, authorized)
		})
	}
}
