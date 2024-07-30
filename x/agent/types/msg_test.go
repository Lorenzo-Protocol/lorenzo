package types_test

import (
	"testing"

	"github.com/Lorenzo-Protocol/lorenzo/v2/x/agent/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

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
				Params: types.NewParams(
					[]string{"cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66"},
				),
			},
			false,
		},
		{
			"fail - duplicate address",
			&types.MsgUpdateParams{
				Authority: authtypes.NewModuleAddress(govtypes.ModuleName).String(),
				Params: types.NewParams(
					[]string{
						"cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
						"cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
					},
				),
			},
			false,
		},
		{
			"pass - valid msg",
			&types.MsgUpdateParams{
				Authority: authtypes.NewModuleAddress(govtypes.ModuleName).String(),
				Params: types.NewParams(
					[]string{"cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66"},
				),
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

func (suite *MsgsTestSuite) TestMsgAddAgent() {
	testCases := []struct {
		name      string
		msgUpdate *types.MsgAddAgent
		expPass   bool
	}{
		{
			"fail - invalid sender address",
			&types.MsgAddAgent{
				Name:                "sinohope4",
				BtcReceivingAddress: "tb1ptt9gnjnvje343y47z2wd7r8w6mnuylp8z0w74qftv7p5x323vxeq9jrn6f",
				EthAddr:             "",
				Description:         "test",
				Url:                 "https://sinohope.com",
				Sender:              "invalid",
			},
			false,
		},
		{
			"fail - empty BtcReceivingAddress",
			&types.MsgAddAgent{
				Name:                "sinohope4",
				BtcReceivingAddress: "",
				EthAddr:             "",
				Description:         "test",
				Url:                 "https://sinohope.com",
				Sender:              "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
			},
			false,
		},
		{
			"fail - empty EthAddr",
			&types.MsgAddAgent{
				Name:                "sinohope4",
				BtcReceivingAddress: "tb1ptt9gnjnvje343y47z2wd7r8w6mnuylp8z0w74qftv7p5x323vxeq9jrn6f",
				EthAddr:             "0x123456",
				Description:         "test",
				Url:                 "https://sinohope.com",
				Sender:              "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
			},
			false,
		},
		{
			"success - valid msg when EthAddr is empty",
			&types.MsgAddAgent{
				Name:                "sinohope4",
				BtcReceivingAddress: "tb1ptt9gnjnvje343y47z2wd7r8w6mnuylp8z0w74qftv7p5x323vxeq9jrn6f",
				EthAddr:             "",
				Description:         "test",
				Url:                 "https://sinohope.com",
				Sender:              "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
			},
			true,
		},
		{
			"success - valid msg when EthAddr not is empty",
			&types.MsgAddAgent{
				Name:                "sinohope4",
				BtcReceivingAddress: "tb1ptt9gnjnvje343y47z2wd7r8w6mnuylp8z0w74qftv7p5x323vxeq9jrn6f",
				EthAddr:             "0xbCC0CdF7683120a1965A343245FA602314C13b9A",
				Description:         "test",
				Url:                 "https://sinohope.com",
				Sender:              "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
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

func (suite *MsgsTestSuite) TestMsgRemoveAgent() {
	testCases := []struct {
		name      string
		msgUpdate *types.MsgRemoveAgent
		expPass   bool
	}{
		{
			"fail - invalid sender address",
			&types.MsgRemoveAgent{
				Id:     1,
				Sender: "invalid",
			},
			false,
		},
		{
			"fail - invalid id",
			&types.MsgRemoveAgent{
				Id:     0,
				Sender: "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
			},
			false,
		},
		{
			"succeed - valid msg",
			&types.MsgRemoveAgent{
				Id:     1,
				Sender: "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
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

func (suite *MsgsTestSuite) TestMsgEditAgent() {
	testCases := []struct {
		name      string
		msgUpdate *types.MsgEditAgent
		expPass   bool
	}{
		{
			"fail - invalid sender address",
			&types.MsgEditAgent{
				Id:          1,
				Sender:      "invalid",
				Name:        "sinohope4",
				Description: "test",
				Url:         "https://sinohope.com",
			},
			false,
		},
		{
			"fail - invalid id",
			&types.MsgEditAgent{
				Id:          0,
				Sender:      "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				Name:        "sinohope4",
				Description: "test",
				Url:         "https://sinohope.com",
			},
			false,
		},
		{
			"fail - invalid name",
			&types.MsgEditAgent{
				Id:          1,
				Name:        "",
				Description: "test",
				Url:         "https://sinohope.com",
				Sender:      "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
			},
			false,
		},
		{
			"succeed - valid msg",
			&types.MsgEditAgent{
				Id:          1,
				Name:        "sinohope4",
				Description: "test",
				Url:         "https://sinohope.com",
				Sender:      "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
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
