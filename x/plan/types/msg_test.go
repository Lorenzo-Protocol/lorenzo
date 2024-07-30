package types_test

import (
	"testing"
	"time"

	"github.com/Lorenzo-Protocol/lorenzo/v2/x/plan/types"
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
					"0x1cd55261EC1e167bf7c7201EE79517B7334F575c",
				),
			},
			false,
		},
		{
			"fail - invalid beacon address",
			&types.MsgUpdateParams{
				Authority: authtypes.NewModuleAddress(govtypes.ModuleName).String(),
				Params: types.NewParams(
					[]string{"cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66"},
					"0x123456",
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
					"0x1cd55261EC1e167bf7c7201EE79517B7334F575c",
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
					"0x1cd55261EC1e167bf7c7201EE79517B7334F575c",
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

func (suite *MsgsTestSuite) TestMsgUpgradePlan() {
	testCases := []struct {
		name      string
		msgUpdate *types.MsgUpgradePlan
		expPass   bool
	}{
		{
			"fail - invalid authority address",
			&types.MsgUpgradePlan{
				Authority:      "invalid",
				Implementation: "0x6508d68f4e5931f93fadc3b7afac5092e195b80f",
			},
			false,
		},
		{
			"fail - invalid implementation address",
			&types.MsgUpgradePlan{
				Authority:      "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				Implementation: "0x123456",
			},
			false,
		},
		{
			"pass - valid msg",
			&types.MsgUpgradePlan{
				Authority:      "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				Implementation: "0x6508d68f4e5931f93fadc3b7afac5092e195b80f",
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

func (suite *MsgsTestSuite) TestMsgCreatePlan() {
	testCases := []struct {
		name      string
		msgUpdate *types.MsgCreatePlan
		expPass   bool
	}{
		{
			"fail - invalid sender address",
			&types.MsgCreatePlan{
				Name:               "lorenzo-stake-plan",
				PlanDescUri:        "https://lorenzo-protocol.io/lorenzo-stake-plan",
				AgentId:            uint64(1),
				PlanStartTime:      uint64(time.Now().UTC().Unix()) + 1000,
				PeriodTime:         1000,
				YatContractAddress: "0xbCC0CdF7683120a1965A343245FA602314C13b9A",
				Sender:             "invalid",
			},
			false,
		},
		{
			"fail - invalid yat contract address",
			&types.MsgCreatePlan{
				Name:               "lorenzo-stake-plan",
				PlanDescUri:        "https://lorenzo-protocol.io/lorenzo-stake-plan",
				AgentId:            uint64(1),
				PlanStartTime:      uint64(time.Now().UTC().Unix()) + 1000,
				PeriodTime:         1000,
				YatContractAddress: "0x123456",
				Sender:             "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
			},
			false,
		},
		{
			"fail - plan name cannot be empty",
			&types.MsgCreatePlan{
				Name:               "",
				PlanDescUri:        "https://lorenzo-protocol.io/lorenzo-stake-plan",
				AgentId:            uint64(1),
				PlanStartTime:      uint64(time.Now().UTC().Unix()) + 1000,
				PeriodTime:         1000,
				YatContractAddress: "0xbCC0CdF7683120a1965A343245FA602314C13b9A",
				Sender:             "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
			},
			false,
		},
		{
			"fail - plan name is empty space",
			&types.MsgCreatePlan{
				Name:               "  ",
				PlanDescUri:        "https://lorenzo-protocol.io/lorenzo-stake-plan",
				AgentId:            uint64(1),
				PlanStartTime:      uint64(time.Now().UTC().Unix()) + 1000,
				PeriodTime:         1000,
				YatContractAddress: "0xbCC0CdF7683120a1965A343245FA602314C13b9A",
				Sender:             "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
			},
			false,
		},
		{
			"fail - agent id cannot be zero",
			&types.MsgCreatePlan{
				Name:               "lorenzo-stake-plan",
				PlanDescUri:        "https://lorenzo-protocol.io/lorenzo-stake-plan",
				AgentId:            uint64(0),
				PlanStartTime:      uint64(time.Now().UTC().Unix()) + 1000,
				PeriodTime:         1000,
				YatContractAddress: "0xbCC0CdF7683120a1965A343245FA602314C13b9A",
				Sender:             "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
			},
			false,
		},
		{
			"fail - plan start time cannot be zero",
			&types.MsgCreatePlan{
				Name:               "lorenzo-stake-plan",
				PlanDescUri:        "https://lorenzo-protocol.io/lorenzo-stake-plan",
				AgentId:            uint64(1),
				PlanStartTime:      0,
				PeriodTime:         1000,
				YatContractAddress: "0xbCC0CdF7683120a1965A343245FA602314C13b9A",
				Sender:             "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
			},
			false,
		},
		{
			"fail - period time cannot be zero",
			&types.MsgCreatePlan{
				Name:               "lorenzo-stake-plan",
				PlanDescUri:        "https://lorenzo-protocol.io/lorenzo-stake-plan",
				AgentId:            uint64(1),
				PlanStartTime:      uint64(time.Now().UTC().Unix()) + 1000,
				PeriodTime:         0,
				YatContractAddress: "0xbCC0CdF7683120a1965A343245FA602314C13b9A",
				Sender:             "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
			},
			false,
		},
		{
			"pass - valid msg",
			&types.MsgCreatePlan{
				Name:               "lorenzo-stake-plan",
				PlanDescUri:        "https://lorenzo-protocol.io/lorenzo-stake-plan",
				AgentId:            uint64(1),
				PlanStartTime:      uint64(time.Now().UTC().Unix()) + 1000,
				PeriodTime:         1000,
				YatContractAddress: "0xbCC0CdF7683120a1965A343245FA602314C13b9A",
				Sender:             "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
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

func (suite *MsgsTestSuite) TestMsgClaims() {
	testCases := []struct {
		name      string
		msgUpdate *types.MsgClaims
		expPass   bool
	}{
		{
			"fail - invalid sender address",
			&types.MsgClaims{
				Receiver:    "0xbCC0CdF7683120a1965A343245FA602314C13b9A",
				MerkleProof: "0x9c22ff5f21f0b81b113e63f7db6da94fedef11b2119b4088b89664fb9a3cb658",
				Sender:      "invalid",
			},
			false,
		},
		{
			"fail - invalid receiver address",
			&types.MsgClaims{
				Receiver:    "0x123456",
				MerkleProof: "0x9c22ff5f21f0b81b113e63f7db6da94fedef11b2119b4088b89664fb9a3cb658",
				Sender:      "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
			},
			false,
		},
		{
			"fail - invalid merkle proof",
			&types.MsgClaims{
				Receiver:    "0xbCC0CdF7683120a1965A343245FA602314C13b9A",
				MerkleProof: "0x123456",
				Sender:      "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
			},
			false,
		},
		{
			"pass - valid msg",
			&types.MsgClaims{
				Receiver:    "0xbCC0CdF7683120a1965A343245FA602314C13b9A",
				MerkleProof: "0x9c22ff5f21f0b81b113e63f7db6da94fedef11b2119b4088b89664fb9a3cb658",
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

func (suite *MsgsTestSuite) TestMsgCreateYAT() {
	testCases := []struct {
		name      string
		msgUpdate *types.MsgCreateYAT
		expPass   bool
	}{
		{
			"fail - invalid sender address",
			&types.MsgCreateYAT{
				Name:   "lorenzo-yat",
				Symbol: "LYAT",
				Sender: "invalid",
			},
			false,
		},
		{
			"fail - yat name cannot be empty",
			&types.MsgCreateYAT{
				Name:   "",
				Symbol: "LYAT",
				Sender: "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
			},
			false,
		},
		{
			"fail - yat name is empty space",
			&types.MsgCreateYAT{
				Name:   "  ",
				Symbol: "LYAT",
				Sender: "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
			},
			false,
		},
		{
			"fail - yat symbol cannot be empty",
			&types.MsgCreateYAT{
				Name:   "lorenzo-yat",
				Symbol: "",
				Sender: "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
			},
			false,
		},
		{
			"fail - yat symbol is empty space",
			&types.MsgCreateYAT{
				Name:   "lorenzo-yat",
				Symbol: "  ",
				Sender: "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
			},
			false,
		},
		{
			"pass - valid msg",
			&types.MsgCreateYAT{
				Name:   "lorenzo-yat",
				Symbol: "LYAT",
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

func (suite *MsgsTestSuite) TestMsgUpdatePlanStatus() {
	testCases := []struct {
		name      string
		msgUpdate *types.MsgUpdatePlanStatus
		expPass   bool
	}{
		{
			"fail - invalid sender address",
			&types.MsgUpdatePlanStatus{
				Sender: "invalid",
				PlanId: 1,
				Status: types.PlanStatus_Pause,
			},
			false,
		},
		{
			"fail - plan id cannot be zero",
			&types.MsgUpdatePlanStatus{
				Sender: "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				PlanId: 0,
				Status: types.PlanStatus_Pause,
			},
			false,
		},
		{
			"fail - status cannot be greater than 1",
			&types.MsgUpdatePlanStatus{
				Sender: "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				PlanId: 1,
				Status: 2,
			},
			false,
		},
		{
			"fail - status cannot be less than 0",
			&types.MsgUpdatePlanStatus{
				Sender: "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				PlanId: 1,
				Status: -1,
			},
			false,
		},
		{
			"pass - valid msg",
			&types.MsgUpdatePlanStatus{
				Sender: "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				PlanId: 1,
				Status: types.PlanStatus_Pause,
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

func (suite *MsgsTestSuite) TestMsgSetMinter() {
	testCases := []struct {
		name      string
		msgUpdate *types.MsgSetMinter
		expPass   bool
	}{
		{
			"fail - invalid sender address",
			&types.MsgSetMinter{
				Sender:          "invalid",
				Minter:          "0xbCC0CdF7683120a1965A343245FA602314C13b9A",
				ContractAddress: "0x6508d68f4e5931f93fadc3b7afac5092e195b80f",
			},
			false,
		},
		{
			"fail - invalid minter address",
			&types.MsgSetMinter{
				Sender:          "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				Minter:          "0x123456",
				ContractAddress: "0x6508d68f4e5931f93fadc3b7afac5092e195b80f",
			},
			false,
		},
		{
			"fail - invalid yat contract address",
			&types.MsgSetMinter{
				Sender:          "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				Minter:          "0xbCC0CdF7683120a1965A343245FA602314C13b9A",
				ContractAddress: "0x123456",
			},
			false,
		},
		{
			"pass - valid msg",
			&types.MsgSetMinter{
				Sender:          "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				Minter:          "0xbCC0CdF7683120a1965A343245FA602314C13b9A",
				ContractAddress: "0x6508d68f4e5931f93fadc3b7afac5092e195b80f",
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

func (suite *MsgsTestSuite) TestMsgRemoveMinter() {
	testCases := []struct {
		name      string
		msgUpdate *types.MsgRemoveMinter
		expPass   bool
	}{
		{
			"fail - invalid sender address",
			&types.MsgRemoveMinter{
				Sender:          "invalid",
				Minter:          "0xbCC0CdF7683120a1965A343245FA602314C13b9A",
				ContractAddress: "0x6508d68f4e5931f93fadc3b7afac5092e195b80f",
			},
			false,
		},
		{
			"fail - invalid minter address",
			&types.MsgRemoveMinter{
				Sender:          "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				Minter:          "0x123456",
				ContractAddress: "0x6508d68f4e5931f93fadc3b7afac5092e195b80f",
			},
			false,
		},
		{
			"fail - invalid yat contract address",
			&types.MsgRemoveMinter{
				Sender:          "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				Minter:          "0xbCC0CdF7683120a1965A343245FA602314C13b9A",
				ContractAddress: "0x123456",
			},
			false,
		},
		{
			"pass - valid msg",
			&types.MsgRemoveMinter{
				Sender:          "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				Minter:          "0xbCC0CdF7683120a1965A343245FA602314C13b9A",
				ContractAddress: "0x6508d68f4e5931f93fadc3b7afac5092e195b80f",
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

func (suite *MsgsTestSuite) TestMsgSetMerkleRoot() {
	testCases := []struct {
		name      string
		msgUpdate *types.MsgSetMerkleRoot
		expPass   bool
	}{
		{
			"fail - invalid sender address",
			&types.MsgSetMerkleRoot{
				Sender:     "invalid",
				PlanId:     1,
				MerkleRoot: "0x9c22ff5f21f0b81b113e63f7db6da94fedef11b2119b4088b89664fb9a3cb658",
			},
			false,
		},
		{
			"fail - plan id cannot be zero",
			&types.MsgSetMerkleRoot{
				Sender:     "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				PlanId:     0,
				MerkleRoot: "0x9c22ff5f21f0b81b113e63f7db6da94fedef11b2119b4088b89664fb9a3cb658",
			},
			false,
		},

		{
			"fail - invalid merkle root",
			&types.MsgSetMerkleRoot{
				Sender:     "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				PlanId:     1,
				MerkleRoot: "0x123456",
			},
			false,
		},
		{
			"pass - valid msg",
			&types.MsgSetMerkleRoot{
				Sender:     "cosmos1qperwt9wrnkg5k9e5gzfgjppzpqhyav5j24d66",
				PlanId:     1,
				MerkleRoot: "0x9c22ff5f21f0b81b113e63f7db6da94fedef11b2119b4088b89664fb9a3cb658",
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
