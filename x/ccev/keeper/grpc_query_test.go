package keeper_test

import (
	"github.com/Lorenzo-Protocol/lorenzo/v3/x/ccev/testutil"
	"github.com/Lorenzo-Protocol/lorenzo/v3/x/ccev/types"
)

func (suite *KeeperTestSuite) TestClient() {
	headers := testutil.GetTestHeaders(suite.T())
	testCases := []struct {
		name         string
		msg          *types.QueryClientRequest
		setup        func()
		expect       *types.Client
		expectErr    bool
		expectErrMsg string
	}{
		{
			name: "client not found",
			msg: &types.QueryClientRequest{
				ChainId: 56,
			},
			setup:        func() {},
			expectErr:    true,
			expectErrMsg: "client not found",
		},
		{
			name: "client found",
			msg: &types.QueryClientRequest{
				ChainId: 56,
			},
			setup: func() {
				suite.CreateClient(56, "Binance Smart Chain", *headers[0])
			},
			expect: &types.Client{
				ChainId:      56,
				ChainName:    "Binance Smart Chain",
				InitialBlock: *headers[0],
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest()
			tc.setup()
			res, err := suite.queryClient.Client(suite.ctx, tc.msg)
			if tc.expectErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expectErrMsg)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expect, res.Client)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestClients() {
	headers := testutil.GetTestHeaders(suite.T())
	testCases := []struct {
		name         string
		setup        func()
		expect       []*types.Client
		expectErr    bool
		expectErrMsg string
	}{
		{
			name:         "success(no clients)",
			setup:        func() {},
		},
		{
			name: "success(one client)",
			setup: func() {
				suite.CreateClient(56, "Binance Smart Chain", *headers[0])
			},
			expect: []*types.Client{
				{
					ChainId:      56,
					ChainName:    "Binance Smart Chain",
					InitialBlock: *headers[0],
				},
			},
		},
		{
			name: "success(multiple client)",
			setup: func() {
				suite.CreateClient(56, "Binance Smart Chain", *headers[0])
				suite.CreateClient(1, "Ethereum Mainnet", *headers[1])
			},
			expect: []*types.Client{
				{
					ChainId:      1,
					ChainName:    "Ethereum Mainnet",
					InitialBlock: *headers[1],
				},
				{
					ChainId:      56,
					ChainName:    "Binance Smart Chain",
					InitialBlock: *headers[0],
				},
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest()
			tc.setup()
			res, err := suite.queryClient.Clients(suite.ctx, &types.QueryClientsRequest{})
			if tc.expectErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expectErrMsg)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expect, res.Clients)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestContract()     {
	headers := testutil.GetTestHeaders(suite.T())
	testCases := []struct {
		name         string
		msg          *types.QueryContractRequest
		setup        func()
		expect       *types.Contract
		expectErr    bool
		expectErrMsg string
	}{
		{
			name: "client not found",
			msg: &types.QueryContractRequest{
				ChainId: 56,
				Address: "0x8BCd1CCDA853677Ac865C882B60FBaF5030EeF50",
			},
			setup:        func() {},
			expectErr:    true,
			expectErrMsg: "contract not found",
		},
		{
			name: "contract not found",
			msg: &types.QueryContractRequest{
				ChainId: 56,
				Address: "0x8BCd1CCDA853677Ac865C882B60FBaF5030EeF50",
			},
			setup: func() {
				suite.CreateClient(56, "Binance Smart Chain", *headers[0])
			},
			expectErr:    true,
			expectErrMsg: "contract not found",
		},
		{
			name: "sunccess",
			msg: &types.QueryContractRequest{
				ChainId: 56,
				Address: "0x8BCd1CCDA853677Ac865C882B60FBaF5030EeF50",
			},
			setup: func() {
				suite.CreateClient(56, "Binance Smart Chain", *headers[0])
				suite.UploadContract(56, "0x8BCd1CCDA853677Ac865C882B60FBaF5030EeF50", "StakeBTC2JoinStakePlan", []byte{})
			},
			expect: &types.Contract{
				ChainId:   56,
				Address:   "0x8BCd1CCDA853677Ac865C882B60FBaF5030EeF50",
				EventName: "StakeBTC2JoinStakePlan",
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest()
			tc.setup()
			res, err := suite.queryClient.Contract(suite.ctx, tc.msg)
			if tc.expectErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expectErrMsg)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expect, res.Contract)
			}
		})
	}
}
func (suite *KeeperTestSuite) TestHeader() {
	headers := testutil.GetTestHeaders(suite.T())
	testCases := []struct {
		name         string
		msg          *types.QueryHeaderRequest
		setup        func()
		expect       *types.TinyHeader
		expectErr    bool
		expectErrMsg string
	}{
		{
			name: "client not found",
			msg: &types.QueryHeaderRequest{
				ChainId: 56,
			},
			setup:        func() {},
			expectErr:    true,
			expectErrMsg: "client not found",
		},
		{
			name: "header not found",
			msg: &types.QueryHeaderRequest{
				ChainId: 56,
			},
			setup: func() {
				suite.CreateClient(56, "Binance Smart Chain", *headers[0])
			},
			expectErr:    true,
			expectErrMsg: "header not found",
		},
		{
			name: "success",
			msg: &types.QueryHeaderRequest{
				ChainId: 56,
			},
			setup: func() {
				suite.CreateClient(56, "Binance Smart Chain", *headers[0])
				suite.UploadHeaders(56, []types.TinyHeader{*headers[1]})
			},
			expect: headers[1],
		},

	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest()
			tc.setup()
			res, err := suite.queryClient.Header(suite.ctx, tc.msg)
			if tc.expectErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expectErrMsg)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expect, res.Header)
			}
		})
	}
}
func (suite *KeeperTestSuite) TestHeaderByHash() {}
func (suite *KeeperTestSuite) TestLatestHeader() {}
func (suite *KeeperTestSuite) TestParams()       {}
