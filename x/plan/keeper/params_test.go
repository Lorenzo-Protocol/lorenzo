package keeper_test

import (
	"reflect"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/Lorenzo-Protocol/lorenzo/v2/x/plan/types"
)

func (suite *KeeperTestSuite) TestParams() {
	testCases := []struct {
		name      string
		paramsFun func() interface{}
		getFun    func() interface{}
		expected  bool
	}{
		{
			name: "success - Checks if the admin is set correctly",
			paramsFun: func() interface{} {
				params := types.Params{
					AllowList: []string{
						testAdmin.String(),
					},
				}
				err := suite.lorenzoApp.PlanKeeper.SetParams(suite.ctx, params)
				suite.Require().NoError(err)
				return params.AllowList
			},
			getFun: func() interface{} {
				return suite.lorenzoApp.PlanKeeper.GetParams(suite.ctx).AllowList
			},
			expected: true,
		},
		{
			name: "success - Checks if the beacon is set correctly",
			paramsFun: func() interface{} {
				suite.Commit()
				contractAddr := crypto.CreateAddress(types.ModuleAddress, 1)
				return contractAddr.Hex()
			},
			getFun: func() interface{} {
				return suite.lorenzoApp.PlanKeeper.GetParams(suite.ctx).Beacon
			},
			expected: true,
		},
	}

	for _, tc := range testCases {
		suite.SetupTest()
		suite.Run(tc.name, func() {
			outcome := reflect.DeepEqual(tc.paramsFun(), tc.getFun())
			suite.Require().Equal(tc.expected, outcome)
		})
	}
}
