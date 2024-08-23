package keeper_test

import (
	"github.com/Lorenzo-Protocol/lorenzo/v3/x/bnblightclient/testutil"
)

func (suite *KeeperTestSuite) TestEndBlock() {
	headers := testutil.GetTestHeaders(suite.T())
	testCases := []struct {
		name           string
		retainedBlocks uint64
	}{
		{
			name:           "retained 1 blocks",
			retainedBlocks: 1,
		},
		{
			name:           "retained 2 blocks",
			retainedBlocks: 2,
		},
		{
			name:           "retained 3 blocks",
			retainedBlocks: 3,
		},
		{
			name:           "retained 4 blocks",
			retainedBlocks: 4,
		},
		{
			name:           "retained 5 blocks",
			retainedBlocks: 5,
		},
		{
			name:           "retained 6 blocks",
			retainedBlocks: 6,
		},
		{
			name:           "retained 10 blocks",
			retainedBlocks: 10,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			suite.SetupTest()
			suite.Require().NoError(suite.keeper.UploadHeaders(suite.ctx, headers[5:]), "upload headers failed")

			params := suite.keeper.GetParams(suite.ctx)
			params.RetainedBlocks = tc.retainedBlocks
			suite.Require().NoError(suite.keeper.SetParams(suite.ctx, params), "set params failed")

			suite.keeper.EndBlock(suite.ctx)

			// last header
			idx := len(headers) - int(tc.retainedBlocks)
			suite.Require().True(suite.keeper.HasHeader(suite.ctx, headers[idx].Number), "header not found")

			idx--
			if idx >= 0 {
				suite.Require().False(suite.keeper.HasHeader(suite.ctx, headers[idx].Number), "header found")
			}
		})
	}
}
