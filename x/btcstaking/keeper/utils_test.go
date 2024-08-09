package keeper_test

import (
	"fmt"

	"github.com/Lorenzo-Protocol/lorenzo/v2/x/btcstaking/keeper"
)

func (suite *KeeperTestSuite) TestCheckBTCTxDepth() {
	testCases := []struct {
		name           string
		stakingTxDepth uint64
		btcAmount      uint64
		expectErr      bool
	}{
		{
			name:           "no depth check required",
			stakingTxDepth: 0,
			btcAmount:      3e5,
			expectErr:      false,
		},
		{
			name:           "at least 1 depth required",
			stakingTxDepth: 1,
			btcAmount:      1e6,
			expectErr:      false,
		},
		{
			name:           "at least 2 depth required",
			stakingTxDepth: 2,
			btcAmount:      1e6 + 1e5,
			expectErr:      false,
		},
		{
			name:           "at least 3 depth required",
			stakingTxDepth: 3,
			btcAmount:      4e7,
			expectErr:      false,
		},
		{
			name:           "not k-deep",
			stakingTxDepth: 3,
			btcAmount:      5e7,
			expectErr:      true,
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("MsgUpdateParams - %s", tc.name), func() {
			err := keeper.CheckBTCTxDepth(tc.stakingTxDepth, tc.btcAmount)
			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}
