package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/suite"

	"github.com/Lorenzo-Protocol/lorenzo/x/token/types"
)

type MsgTestSuite struct {
	suite.Suite
}

func TestMsgTestSuite(t *testing.T) {
	suite.Run(t, new(MsgTestSuite))
}

func (suite *MsgTestSuite) TestMsgRegisterCoin() {
	coinMetadataMap := map[string]banktypes.Metadata{
		"coin": {
			DenomUnits: []*banktypes.DenomUnit{
				{
					Denom:    "acoin",
					Exponent: 0,
				},
				{
					Denom:    "coin",
					Exponent: 18,
				},
			},
			Base:    "acoin",
			Display: "coin",
			Name:    "Coin Token",
			Symbol:  "COIN",
		},
		"0coin": {
			Base: "0coin",
		},
		"erc20denom": {
			DenomUnits: []*banktypes.DenomUnit{
				{
					Denom:    "erc20/0x1cd55261EC1e167bf7c7201EE79517B7334F575c",
					Exponent: 0,
				},
				{
					Denom:    "erc20denom",
					Exponent: 18,
				},
			},
			Base:    "erc20/0x1cd55261EC1e167bf7c7201EE79517B7334F575c",
			Display: "erc20denom",
			Name:    "Coin Token",
			Symbol:  "COIN",
		},
	}

	testCases := []struct {
		name    string
		msg     *types.MsgRegisterCoin
		expPass bool
	}{
		{
			name:    "fail: invalid authority",
			expPass: false,
			msg: &types.MsgRegisterCoin{
				Authority: "invalid authority",
			},
		},
		{
			name:    "fail: invalid metadata",
			expPass: false,
			msg: &types.MsgRegisterCoin{
				Authority: "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				Metadata:  []banktypes.Metadata{coinMetadataMap["0coin"]},
			},
		},
		{
			name:    "fail: unexpected erc20 denom",
			expPass: false,
			msg: &types.MsgRegisterCoin{
				Authority: "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				Metadata:  []banktypes.Metadata{coinMetadataMap["erc20denom"]},
			},
		},
		{
			name:    "success",
			expPass: true,
			msg: &types.MsgRegisterCoin{
				Authority: "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				Metadata:  []banktypes.Metadata{coinMetadataMap["coin"]},
			},
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

func (suite *MsgTestSuite) TestMsgRegisterERC20() {
	testCases := []struct {
		name    string
		msg     *types.MsgRegisterERC20
		expPass bool
	}{
		{
			name: "fail: invalid authority",
			msg: &types.MsgRegisterERC20{
				Authority: "invalid authority",
			},
			expPass: false,
		},
		{
			name: "fail: invalid contract address",
			msg: &types.MsgRegisterERC20{
				Authority: "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				ContractAddresses: []string{
					"0x1cd55261EC1e167bf7c7201EE79517B7334F575c",
					"0x1cd55261EC1e167bf7c",
				},
			},
			expPass: false,
		},
		{
			name: "success",
			msg: &types.MsgRegisterERC20{
				Authority:         "invalid authority",
				ContractAddresses: []string{"0x1cd55261EC1e167bf7c7201EE79517B7334F575c"},
			},
			expPass: false,
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

func (suite *MsgTestSuite) TestMsgToggleConversion() {
	testCases := []struct {
		name    string
		msg     *types.MsgToggleConversion
		expPass bool
	}{
		{
			name:    "fail: invalid authority address",
			expPass: false,
			msg:     &types.MsgToggleConversion{Authority: "invalid"},
		},
		{
			name:    "fail: invalid denom",
			expPass: false,
			msg: &types.MsgToggleConversion{
				Authority: "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				Token:     "012token",
			},
		},
		{
			name:    "fail: invalid hex address",
			expPass: false,
			msg: &types.MsgToggleConversion{
				Authority: "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				Token:     "0x1cd55261EC1e167bf7c",
			},
		},
		{
			name:    "success: valid denom",
			expPass: true,
			msg: &types.MsgToggleConversion{
				Authority: "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				Token:     "coincoin",
			},
		},
		{
			name:    "success: hex address",
			expPass: true,
			msg: &types.MsgToggleConversion{
				Authority: "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				Token:     "0x1cd55261EC1e167bf7c7201EE79517B7334F575c",
			},
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

func (suite *MsgTestSuite) TestMsgConvertCoin() {
	testCases := []struct {
		name    string
		msg     *types.MsgConvertCoin
		expPass bool
	}{
		{
			name:    "fail: not valid denom",
			expPass: false,
			msg: &types.MsgConvertCoin{
				Sender:   "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				Receiver: "0x1cd55261EC1e167bf7c7201EE79517B7334F575c",
				Coin: sdk.Coin{
					Denom:  "0abc",
					Amount: sdk.NewInt(100),
				},
			},
		},
		{
			name:    "fail: amount not positive",
			expPass: false,
			msg: &types.MsgConvertCoin{
				Sender:   "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				Receiver: "0x1cd55261EC1e167bf7c7201EE79517B7334F575c",
				Coin: sdk.Coin{
					Denom:  "coincoin",
					Amount: sdk.NewInt(-100),
				},
			},
		},
		{
			name:    "fail: invalid sender",
			expPass: false,
			msg: &types.MsgConvertCoin{
				Sender:   "invalid sender",
				Receiver: "0x1cd55261EC1e167bf7c7201EE79517B7334F575c",
				Coin: sdk.Coin{
					Denom:  "coincoin",
					Amount: sdk.NewInt(100),
				},
			},
		},
		{
			name:    "fail: invalid receiver",
			expPass: false,
			msg: &types.MsgConvertCoin{
				Sender:   "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				Receiver: "0x1cd55261EC1e1",
				Coin: sdk.Coin{
					Denom:  "coincoin",
					Amount: sdk.NewInt(100),
				},
			},
		},
		{
			name:    "success",
			expPass: true,
			msg: &types.MsgConvertCoin{
				Sender:   "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				Receiver: "0x1cd55261EC1e167bf7c7201EE79517B7334F575c",
				Coin: sdk.Coin{
					Denom:  "coincoin",
					Amount: sdk.NewInt(100),
				},
			},
		},
		{
			name:    "success",
			expPass: true,
			msg: &types.MsgConvertCoin{
				Sender:   "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				Receiver: "0x1cd55261EC1e167bf7c7201EE79517B7334F575c",
				Coin: sdk.Coin{
					Denom:  "erc20/0x1cd55261EC1e167bf7c7201EE79517B7334F575c",
					Amount: sdk.NewInt(100),
				},
			},
		},
		{
			name:    "success",
			expPass: true,
			msg: &types.MsgConvertCoin{
				Sender:   "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				Receiver: "0x1cd55261EC1e167bf7c7201EE79517B7334F575c",
				Coin: sdk.Coin{
					Denom:  "ibc/7F1D3FCF4AE79E1554D670D1AD949A9BA4E4A3C76C63093E17E446A46061A7A2",
					Amount: sdk.NewInt(100),
				},
			},
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

func (suite *MsgTestSuite) TestMsgConcertERC20() {
	testCases := []struct {
		name    string
		msg     *types.MsgConvertERC20
		expPass bool
	}{
		{
			name:    "fail: invalid contract address",
			expPass: false,
			msg: &types.MsgConvertERC20{
				Sender:          "0x1cd55261EC1e167bf7c7201EE79517B7334F575c",
				Receiver:        "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				ContractAddress: "0x1cd55261EC1e16",
				Amount:          sdk.NewInt(100),
			},
		},
		{
			name:    "fail: invalid sender",
			expPass: false,
			msg: &types.MsgConvertERC20{
				Sender:          "0x1cd55261EC1e167bf",
				Receiver:        "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				ContractAddress: "0x1cd55261EC1e167bf7c7201EE79517B7334F575c",
				Amount:          sdk.NewInt(100),
			},
		},
		{
			name:    "fail: invalid receiver",
			expPass: false,
			msg: &types.MsgConvertERC20{
				Sender:          "0x1cd55261EC1e167bf7c7201EE79517B7334F575c",
				Receiver:        "invalid receiver",
				ContractAddress: "0x1cd55261EC1e167bf7c7201EE79517B7334F575c",
				Amount:          sdk.NewInt(100),
			},
		},
		{
			name:    "success",
			expPass: true,
			msg: &types.MsgConvertERC20{
				Sender:          "0x1cd55261EC1e167bf7c7201EE79517B7334F575c",
				Receiver:        "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				ContractAddress: "0x1cd55261EC1e167bf7c7201EE79517B7334F575c",
				Amount:          sdk.NewInt(100),
			},
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

func (suite *MsgTestSuite) TestMsgUpdateParams() {
	testCases := []struct {
		name    string
		msg     *types.MsgUpdateParams
		expPass bool
	}{
		{
			name: "fail: invalid authority update",
			msg: &types.MsgUpdateParams{
				Authority: "invalid authority",
				Params: types.Params{
					EnableConvert: false,
					EnableEVMHook: false,
				},
			},
			expPass: false,
		},
		{
			name: "valid update",
			msg: &types.MsgUpdateParams{
				Authority: "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				Params: types.Params{
					EnableConvert: false,
					EnableEVMHook: false,
				},
			},
			expPass: true,
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
