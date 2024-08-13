package keeper_test

import (
	"fmt"

	"github.com/Lorenzo-Protocol/lorenzo/v3/x/btcstaking/types"
)

func (suite *KeeperTestSuite) TestParamsQuery() {
	testCases := []struct {
		name       string
		request    *types.QueryParamsRequest
		malleate   func()
		validation func()
		expectErr  bool
	}{
		{
			name:      "success",
			request:   &types.QueryParamsRequest{},
			expectErr: false,
			malleate: func() {
				suite.Commit()
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("ParamsQuery - %s", tc.name), func() {
			suite.SetupTest()
			if tc.malleate != nil {
				tc.malleate()
			}
			_, err := suite.queryClient.Params(suite.ctx, tc.request)
			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
			if tc.validation != nil {
				tc.validation()
			}
		})
	}
}
