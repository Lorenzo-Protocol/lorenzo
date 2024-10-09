package keeper_test

import (
	"github.com/Lorenzo-Protocol/lorenzo/v3/x/ccev/testutil"
	"github.com/Lorenzo-Protocol/lorenzo/v3/x/ccev/types"
	evmtypes "github.com/ethereum/go-ethereum/core/types"
)

func (suite *KeeperTestSuite) TestVerifyReceipt() {
	type params struct {
		chainID uint32
		number  uint64
		receipt *evmtypes.Receipt
		proof   *types.Proof
	}

	data := testutil.GetTestProvedReceipts(suite.T())
	pool := testutil.GetTestHeaderPool(suite.T())
	testCases := []struct {
		name         string
		request      params
		setup        func()
		expectErr    bool
		expectErrMsg string
	}{
		{
			name:         "header not found(client not created)",
			request:      params{chainID: 56, number: 42768118, receipt: data.Receipt, proof: data.Proof},
			expectErr:    true,
			setup:        func() {},
			expectErrMsg: "header not found",
		},
		{
			name:      "header not found",
			request:   params{chainID: 56, number: 42768118, receipt: data.Receipt, proof: data.Proof},
			expectErr: true,
			setup: func() {
				suite.CreateClient(56, "Binance Smart Chain", *pool.GetTestHeader(42768117))
			},
			expectErrMsg: "header not found",
		},
		{
			name:    "success",
			request: params{chainID: 56, number: 42768118, receipt: data.Receipt, proof: data.Proof},
			setup: func() {
				suite.CreateClient(56, "Binance Smart Chain", *pool.GetTestHeader(42768118))
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest()
			tc.setup()

			err := suite.keeper.VerifyReceipt(suite.ctx, tc.request.chainID, tc.request.number, tc.request.receipt, tc.request.proof)
			if tc.expectErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expectErrMsg)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}
