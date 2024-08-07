package keeper_test

import (
	"github.com/Lorenzo-Protocol/lorenzo/v2/x/bnblightclient/testutil"
	"github.com/Lorenzo-Protocol/lorenzo/v2/x/bnblightclient/types"
)

func (suite *KeeperTestSuite) TestUploadHeaders() {
	headers := testutil.GetTestHeaders(suite.T())
	testCases := []struct {
		name         string
		headers      []*types.Header
		expectErr    bool
		expectErrMsg string
	}{
		{
			name:         "empty RawHeader",
			headers:      []*types.Header{{
				RawHeader:   []byte{},
				ParentHash:  headers[5].ParentHash,
				Hash:        headers[5].Hash,
				Number:      headers[5].Number,
				ReceiptRoot: headers[5].ReceiptRoot,
			}},
			expectErr:    true,
			expectErrMsg: "unmarshal header failed",
		},
		{
			name:         "wrong ParentHash",
			headers:      []*types.Header{{
				RawHeader:   headers[5].RawHeader,
				ParentHash:  headers[6].ParentHash,
				Hash:        headers[5].Hash,
				Number:      headers[5].Number,
				ReceiptRoot: headers[5].ReceiptRoot,
			}},
			expectErr:    true,
			expectErrMsg: "parentHash not equal",
		},
		{
			name:         "wrong Hash",
			headers:      []*types.Header{{
				RawHeader:   headers[5].RawHeader,
				ParentHash:  headers[5].ParentHash,
				Hash:        headers[6].Hash,
				Number:      headers[5].Number,
				ReceiptRoot: headers[5].ReceiptRoot,
			}},
			expectErr:    true,
			expectErrMsg: "hash not equal",
		},
		{
			name:         "wrong Number",
			headers:      []*types.Header{{
				RawHeader:   headers[5].RawHeader,
				ParentHash:  headers[5].ParentHash,
				Hash:        headers[5].Hash,
				Number:      headers[6].Number,
				ReceiptRoot: headers[5].ReceiptRoot,
			}},
			expectErr:    true,
			expectErrMsg: "number not equal",
		},
		{
			name:         "wrong ReceiptRoot",
			headers:      []*types.Header{{
				RawHeader:   headers[5].RawHeader,
				ParentHash:  headers[5].ParentHash,
				Hash:        headers[5].Hash,
				Number:      headers[5].Number,
				ReceiptRoot: headers[6].ReceiptRoot,
			}},
			expectErr:    true,
			expectErrMsg: "receiptHash not equal",
		},
		{
			name:         "discontinuous header",
			headers:      []*types.Header{headers[6]},
			expectErr:    true,
			expectErrMsg: "hash not equal",
		},
		{
			name:         "succcess",
			headers:      headers[5:],
			expectErr:    false,
		},
	}
	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			suite.SetupTest()

			_, err := suite.msgServer.UploadHeaders(suite.ctx, &types.MsgUploadHeaders{
				Headers: tc.headers,
				Signer:  testAdmin.String(),
			})

			if tc.expectErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expectErrMsg)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}
