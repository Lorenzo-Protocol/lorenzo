package keeper_test

import (
	"github.com/Lorenzo-Protocol/lorenzo/v3/x/ccev/testutil"
	"github.com/Lorenzo-Protocol/lorenzo/v3/x/ccev/types"
)

func (suite *KeeperTestSuite) TestCreateClient() {
	headers := testutil.GetTestHeaders(suite.T())
	testCases := []struct {
		name         string
		msg          *types.MsgCreateClient
		expectErr    bool
		expectErrMsg string
	}{
		{
			name: "unauthorized sender",
			msg: &types.MsgCreateClient{
				Sender: testAccounts[1].String(),
				Client: types.Client{
					ChainId:      56,
					ChainName:    "Binance Smart Chain",
					InitialBlock: *headers[0],
				},
			},
			expectErr:    true,
			expectErrMsg: "tx unauthorized",
		},
		{
			name: "successful create client",
			msg: &types.MsgCreateClient{
				Sender: testAccounts[0].String(),
				Client: types.Client{
					ChainId:      56,
					ChainName:    "Binance Smart Chain",
					InitialBlock: *headers[0],
				},
			},
			expectErr: false,
		},
		{
			name: "duplicate client",
			msg: &types.MsgCreateClient{
				Sender: testAccounts[0].String(),
				Client: types.Client{
					ChainId:      56,
					ChainName:    "Binance Smart Chain",
					InitialBlock: *headers[0],
				},
			},
			expectErr:    true,
			expectErrMsg: "duplicate client",
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			_, err := suite.msgServer.CreateClient(suite.ctx, tc.msg)
			if tc.expectErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expectErrMsg)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestUploadContract() {
	headers := testutil.GetTestHeaders(suite.T())
	testCases := []struct {
		name         string
		msg          *types.MsgUploadContract
		setup        func()
		expectErr    bool
		expectErrMsg string
	}{
		{
			name: "unauthorized sender",
			setup: func() {
				suite.CreateClient(56, "Binance Smart Chain", *headers[0])
			},
			msg: &types.MsgUploadContract{
				Sender:    testAccounts[1].String(),
				ChainId:   56,
				Address:   "0x7f0d2a9f0f41b9f1e9f8f9a3fe7e0d9a3fe7e0d9",
				EventName: "StakeBTC2JoinStakePlan",
				Abi:       []byte{},
			},
			expectErr:    true,
			expectErrMsg: "tx unauthorized",
		},
		{
			name:  "client not found",
			setup: func() {},
			msg: &types.MsgUploadContract{
				Sender:    testAccounts[0].String(),
				ChainId:   56,
				Address:   "0x7f0d2a9f0f41b9f1e9f8f9a3fe7e0d9a3fe7e0d9",
				EventName: "StakeBTC2JoinStakePlan",
				Abi:       []byte{},
			},
			expectErr:    true,
			expectErrMsg: "client not found",
		},
		{
			name: "successful upload contract",
			setup: func() {
				suite.CreateClient(56, "Binance Smart Chain", *headers[0])
			},
			msg: &types.MsgUploadContract{
				Sender:    testAccounts[0].String(),
				ChainId:   56,
				Address:   "0x7f0d2a9f0f41b9f1e9f8f9a3fe7e0d9a3fe7e0d9",
				EventName: "StakeBTC2JoinStakePlan",
				Abi:       []byte{},
			},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		suite.SetupTest()
		suite.Run(tc.name, func() {
			tc.setup()
			_, err := suite.msgServer.UploadContract(suite.ctx, tc.msg)
			if tc.expectErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expectErrMsg)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestUploadHeaders() {
	headers := testutil.GetTestHeaders(suite.T())
	testCases := []struct {
		name         string
		msg          *types.MsgUploadHeaders
		setup        func()
		expectErr    bool
		expectErrMsg string
	}{
		{
			name: "empty headers",
			setup: func() {
				suite.CreateClient(56, "Binance Smart Chain", *headers[0])
			},
			msg: &types.MsgUploadHeaders{
				Sender:  testAccounts[0].String(),
				ChainId: 56,
				Headers: []types.TinyHeader{},
			},
			expectErr:    true,
			expectErrMsg: "header is empty",
		},
		{
			name: "unauthorized sender",
			setup: func() {
				suite.CreateClient(56, "Binance Smart Chain", *headers[0])
			},
			msg: &types.MsgUploadHeaders{
				Sender:  testAccounts[1].String(),
				ChainId: 56,
				Headers: []types.TinyHeader{*headers[1]},
			},
			expectErr:    true,
			expectErrMsg: "tx unauthorized",
		},
		{
			name:  "client not found",
			setup: func() {},
			msg: &types.MsgUploadHeaders{
				Sender:  testAccounts[0].String(),
				ChainId: 56,
				Headers: []types.TinyHeader{*headers[1]},
			},
			expectErr:    true,
			expectErrMsg: "client not found",
		},
		{
			name: "header number less than initial height",
			setup: func() {
				suite.CreateClient(56, "Binance Smart Chain", *headers[1])
			},
			msg: &types.MsgUploadHeaders{
				Sender:  testAccounts[0].String(),
				ChainId: 56,
				Headers: []types.TinyHeader{*headers[0]},
			},
			expectErr:    true,
			expectErrMsg: "invalid header",
		},
		{
			name: "duplicate header",
			setup: func() {
				suite.CreateClient(56, "Binance Smart Chain", *headers[0])
			},
			msg: &types.MsgUploadHeaders{
				Sender:  testAccounts[0].String(),
				ChainId: 56,
				Headers: []types.TinyHeader{*headers[0]},
			},
			expectErr:    true,
			expectErrMsg: "duplicate header",
		},
		{
			name: "successful upload headers(single)",
			setup: func() {
				suite.CreateClient(56, "Binance Smart Chain", *headers[0])
			},
			msg: &types.MsgUploadHeaders{
				Sender:  testAccounts[0].String(),
				ChainId: 56,
				Headers: []types.TinyHeader{*headers[1]},
			},
			expectErr: false,
		},
		{
			name: "successful upload headers(multiple)",
			setup: func() {
				suite.CreateClient(56, "Binance Smart Chain", *headers[0])
			},
			msg: &types.MsgUploadHeaders{
				Sender:  testAccounts[0].String(),
				ChainId: 56,
				Headers: []types.TinyHeader{*headers[1], *headers[2]},
			},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		suite.SetupTest()
		suite.Run(tc.name, func() {
			tc.setup()
			_, err := suite.msgServer.UploadHeaders(suite.ctx, tc.msg)
			if tc.expectErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expectErrMsg)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestUpdateParams() {
	testCases := []struct {
		name         string
		msg          *types.MsgUpdateParams
		expectErr    bool
		expectErrMsg string
	}{
		{
			name: "unauthorized sender",
			msg: &types.MsgUpdateParams{
				Authority: testAccounts[1].String(),
				Params: types.Params{
					AllowList: []string{testAccounts[0].String()},
				},
			},
			expectErr:    true,
			expectErrMsg: "invalid authority",
		},
		{
			name: "duplicate address",
			msg: &types.MsgUpdateParams{
				Authority: "lrz10d07y265gmmuvt4z0w9aw880jnsr700jq84749",
				Params: types.Params{
					AllowList: []string{testAccounts[0].String(), testAccounts[0].String()},
				},
			},
			expectErr:    true,
			expectErrMsg: "duplicate address",
		},
		{
			name: "invalid address",
			msg: &types.MsgUpdateParams{
				Authority: "lrz10d07y265gmmuvt4z0w9aw880jnsr700jq84749",
				Params: types.Params{
					AllowList: []string{"lrz10d07y265gmmuvt4z0w9aw880jnsr700jq8474"},
				},
			},
			expectErr:    true,
			expectErrMsg: "invalid address",
		},
		{
			name: "successful update params",
			msg: &types.MsgUpdateParams{
				Authority: "lrz10d07y265gmmuvt4z0w9aw880jnsr700jq84749",
				Params: types.Params{
					AllowList: []string{testAccounts[0].String()},
				},
			},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		suite.SetupTest()
		suite.Run(tc.name, func() {
			_, err := suite.msgServer.UpdateParams(suite.ctx, tc.msg)
			if tc.expectErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expectErrMsg)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestUpdateHeader() {
	headers := testutil.GetTestHeaders(suite.T())
	testCases := []struct {
		name         string
		msg          *types.MsgUpdateHeader
		setup        func()
		expectErr    bool
		expectErrMsg string
	}{
		{
			name: "unauthorized sender",
			setup: func() {
				suite.CreateClient(56, "Binance Smart Chain", *headers[1])
			},
			msg: &types.MsgUpdateHeader{
				Sender:  testAccounts[1].String(),
				ChainId: 56,
				Header:  *headers[1],
			},
			expectErr:    true,
			expectErrMsg: "tx unauthorized",
		},
		{
			name:  "client not found",
			setup: func() {},
			msg: &types.MsgUpdateHeader{
				Sender:  testAccounts[0].String(),
				ChainId: 56,
				Header:  *headers[1],
			},
			expectErr:    true,
			expectErrMsg: "client not found",
		},
		{
			name: "invalid header",
			setup: func() {
				suite.CreateClient(56, "Binance Smart Chain", *headers[1])
			},
			msg: &types.MsgUpdateHeader{
				Sender:  testAccounts[0].String(),
				ChainId: 56,
				Header:  *headers[0],
			},
			expectErr:    true,
			expectErrMsg: "invalid header",
		},
		{
			name: "header not found",
			setup: func() {
				suite.CreateClient(56, "Binance Smart Chain", *headers[1])
			},
			msg: &types.MsgUpdateHeader{
				Sender:  testAccounts[0].String(),
				ChainId: 56,
				Header:  *headers[2],
			},
			expectErr:    true,
			expectErrMsg: "header not found",
		},
		{
			name: "header not found",
			setup: func() {
				suite.CreateClient(56, "Binance Smart Chain", *headers[1])
			},
			msg: &types.MsgUpdateHeader{
				Sender:  testAccounts[0].String(),
				ChainId: 56,
				Header:  *headers[1],
			},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		suite.SetupTest()
		suite.Run(tc.name, func() {
			tc.setup()
			_, err := suite.msgServer.UpdateHeader(suite.ctx, tc.msg)
			if tc.expectErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expectErrMsg)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}
