package types_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/Lorenzo-Protocol/lorenzo/v3/x/btcstaking/types"
	"github.com/stretchr/testify/suite"
)

type MsgsTestSuite struct {
	suite.Suite
}

func TestMsgsTestSuite(t *testing.T) {
	suite.Run(t, new(MsgsTestSuite))
}

func (suite *MsgsTestSuite) TestMsgUpdateParams() {
	testCases := []struct {
		name      string
		msgUpdate *types.MsgUpdateParams
		expPass   bool
	}{
		{
			"fail - invalid authority address",
			&types.MsgUpdateParams{
				Authority: "invalid",
				Params: types.Params{
					Receivers:             nil,
					BtcConfirmationsDepth: 1,
					MinterAllowList: []string{
						"cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
					},
					BridgeAddr:      "0xb7C0817Dd23DE89E4204502dd2C2EF7F57d3A3B8",
					TxoutDustAmount: 546,
				},
			},
			false,
		},
		{
			"fail - duplicate address",
			&types.MsgUpdateParams{
				Authority: authtypes.NewModuleAddress(govtypes.ModuleName).String(),
				Params: types.Params{
					Receivers:             nil,
					BtcConfirmationsDepth: 1,
					MinterAllowList: []string{
						"cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
						"cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
					},
					BridgeAddr:      "0xb7C0817Dd23DE89E4204502dd2C2EF7F57d3A3B8",
					TxoutDustAmount: 546,
				},
			},
			false,
		},
		{
			"pass - valid msg",
			&types.MsgUpdateParams{
				Authority: authtypes.NewModuleAddress(govtypes.ModuleName).String(),
				Params: types.Params{
					Receivers:             nil,
					BtcConfirmationsDepth: 1,
					MinterAllowList: []string{
						"cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
					},
					BridgeAddr:      "0xb7C0817Dd23DE89E4204502dd2C2EF7F57d3A3B8",
					TxoutDustAmount: 546,
				},
			},
			true,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			err := tc.msgUpdate.ValidateBasic()
			if tc.expPass {
				suite.NoError(err)
			} else {
				suite.Error(err)
			}
		})
	}
}

func (suite *MsgsTestSuite) TestMsgCreateBTCStaking() {
	testCases := []struct {
		name    string
		msg     *types.MsgCreateBTCStaking
		expPass bool
	}{
		{
			"fail - invalid authority address",
			&types.MsgCreateBTCStaking{
				Signer: "invalid",
				StakingTx: &types.TransactionInfo{
					Key: &types.TransactionKey{
						Index: 1,
						Hash:  nil,
					},
					Transaction: []byte("tx"),
					Proof:       []byte("tx"),
				},
				AgentId: 1,
			},
			false,
		},
		{
			"fail - staking tx is nil",
			&types.MsgCreateBTCStaking{
				Signer:    "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				StakingTx: nil,
				AgentId:   1,
			},
			false,
		},
		{
			"fail - agent id is zero",
			&types.MsgCreateBTCStaking{
				Signer: "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				StakingTx: &types.TransactionInfo{
					Key: &types.TransactionKey{
						Index: 1,
						Hash:  nil,
					},
					Transaction: []byte("tx"),
					Proof:       []byte("tx"),
				},
				AgentId: 0,
			},
			false,
		},
		{
			"pass - valid msg",
			&types.MsgCreateBTCStaking{
				Signer: "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				StakingTx: &types.TransactionInfo{
					Key: &types.TransactionKey{
						Index: 1,
						Hash:  nil,
					},
					Transaction: []byte("tx"),
					Proof:       []byte("tx"),
				},
				AgentId: 1,
			},
			true,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			err := tc.msg.ValidateBasic()
			if tc.expPass {
				suite.NoError(err)
			} else {
				suite.Error(err)
			}
		})
	}
}

func (suite *MsgsTestSuite) TestMsgBurnRequest() {
	testCases := []struct {
		name    string
		msg     *types.MsgBurnRequest
		expPass bool
	}{
		{
			"fail - invalid authority address",
			&types.MsgBurnRequest{
				Signer:           "invalid",
				BtcTargetAddress: "tb1ptt9gnjnvje343y47z2wd7r8w6mnuylp8z0w74qftv7p5x323vxeq9jrn6f",
				Amount:           sdkmath.NewInt(1000000),
			},
			false,
		},
		{
			"fail - btc target address is empty",
			&types.MsgBurnRequest{
				Signer:           "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				BtcTargetAddress: "",
				Amount:           sdkmath.NewInt(1000000),
			},
			false,
		},
		{
			"fail - amount is zero",
			&types.MsgBurnRequest{
				Signer:           "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				BtcTargetAddress: "tb1ptt9gnjnvje343y47z2wd7r8w6mnuylp8z0w74qftv7p5x323vxeq9jrn6f",
				Amount:           sdkmath.NewInt(0),
			},
			false,
		},
		{
			"pass - valid msg",
			&types.MsgBurnRequest{
				Signer:           "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				BtcTargetAddress: "tb1ptt9gnjnvje343y47z2wd7r8w6mnuylp8z0w74qftv7p5x323vxeq9jrn6f",
				Amount:           sdkmath.NewInt(1000000),
			},
			true,
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			err := tc.msg.ValidateBasic()
			if tc.expPass {
				suite.NoError(err)
			} else {
				suite.Error(err)
			}
		})
	}
}
