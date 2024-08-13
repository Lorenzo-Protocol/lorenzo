package keeper_test

import (
	"encoding/json"
	"fmt"
	"time"

	evmtypes "github.com/evmos/ethermint/x/evm/types"

	"github.com/ethereum/go-ethereum/common"

	sdkmath "cosmossdk.io/math"

	"github.com/Lorenzo-Protocol/lorenzo/v3/x/plan/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func (suite *KeeperTestSuite) TestUpdateParams() {
	suite.SetupTest()
	testCases := []struct {
		name      string
		request   *types.MsgUpdateParams
		expectErr bool
	}{
		{
			name:      "fail - invalid authority",
			request:   &types.MsgUpdateParams{Authority: "foobar"},
			expectErr: true,
		},
		{
			name: "fail - beacon is not a valid beacon address",
			request: &types.MsgUpdateParams{
				Authority: authtypes.NewModuleAddress(govtypes.ModuleName).String(),
				Params:    types.Params{Beacon: "0x123456"},
			},
			expectErr: true,
		},
		{
			name: "fail - beacon is not a valid allow_list address",
			request: &types.MsgUpdateParams{
				Authority: authtypes.NewModuleAddress(govtypes.ModuleName).String(),
				Params: types.Params{
					Beacon: "0xbCC0CdF7683120a1965A343245FA602314C13b9A",
					AllowList: []string{
						"cosmos1qperwt9wrnkg5k9e5gzfgjppzpqqu8t3q4yjx9",
						"lrz1tffj9qp3wpdnuds443c86wffrac4jkapkjmmcy",
					},
				},
			},
			expectErr: true,
		},
		{
			name: "pass - valid Update msg",
			request: &types.MsgUpdateParams{
				Authority: authtypes.NewModuleAddress(govtypes.ModuleName).String(),
				Params: types.Params{
					Beacon: "0xbCC0CdF7683120a1965A343245FA602314C13b9A",
					AllowList: []string{
						"lrz1tffj9qp3wpdnuds443c86wffrac4jkapkjmmcy",
					},
				},
			},
			expectErr: false,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("MsgUpdateParams - %s", tc.name), func() {
			_, err := suite.msgServer.UpdateParams(suite.ctx, tc.request)
			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestCreatePlan() {
	testCases := []struct {
		name       string
		request    *types.MsgCreatePlan
		malleate   func(request *types.MsgCreatePlan)
		validation func(request *types.MsgCreatePlan)
		expectErr  bool
	}{
		{
			name:      "fail - invalid sender",
			request:   &types.MsgCreatePlan{Sender: "foobar"},
			expectErr: true,
		},
		{
			name:      "fail - sender not authorized",
			request:   &types.MsgCreatePlan{Sender: "lrz1tffj9qp3wpdnuds443c86wffrac4jkapkjmmcy"},
			expectErr: true,
		},
		{
			name: "fail - invalid PlanStartTime",
			request: &types.MsgCreatePlan{
				Name:               "lorenzo-stake-plan",
				PlanDescUri:        "https://lorenzo-protocol.io/lorenzo-stake-plan",
				AgentId:            uint64(1),
				PlanStartTime:      1000,
				PeriodTime:         1000,
				YatContractAddress: "0xbCC0CdF7683120a1965A343245FA602314C13b9A",
				Sender:             testAdmin.String(),
			},
			malleate: func(request *types.MsgCreatePlan) {
				// create agent
				suite.Commit()
			},
			expectErr: true,
		},
		{
			name: "fail - invalid yat contract address",
			request: &types.MsgCreatePlan{
				Name:               "lorenzo-stake-plan",
				PlanDescUri:        "https://lorenzo-protocol.io/lorenzo-stake-plan",
				AgentId:            uint64(1),
				PlanStartTime:      uint64(time.Now().UTC().Unix()) + 1000,
				PeriodTime:         1000,
				YatContractAddress: "0xbCC0CdF7683120a1965A343245FA602314C13b9A",
				Sender:             testAdmin.String(),
			},
			malleate: func(request *types.MsgCreatePlan) {
				// create agent
				name := "sinohope4"
				btcReceivingAddress := "3C7VPws9fMW3kcwRJvMkSVdqMs4SAhQCqq"
				ethAddr := "0x6508d68f4e5931f93fadc3b7afac5092e195b80f"
				description := "lorenzo"
				url := "https://lorenzo-protocol.io"
				agentId := suite.lorenzoApp.AgentKeeper.AddAgent(
					suite.ctx,
					name, btcReceivingAddress, ethAddr, description, url)
				suite.Require().NotEqual(agentId, 0)
				suite.Require().Equal(agentId, uint64(1))
			},
			expectErr: true,
		},
		{
			name: "success - valid create plan",
			request: &types.MsgCreatePlan{
				Name:          "lorenzo-stake-plan",
				PlanDescUri:   "https://lorenzo-protocol.io/lorenzo-stake-plan",
				AgentId:       uint64(1),
				PlanStartTime: uint64(time.Now().UTC().Unix()) + 1000,
				PeriodTime:    1000,
				Sender:        testAdmin.String(),
			},
			malleate: func(request *types.MsgCreatePlan) {
				suite.Commit()
				// create agent
				name := "sinohope4"
				btcReceivingAddress := "3C7VPws9fMW3kcwRJvMkSVdqMs4SAhQCqq"
				ethAddr := "0x6508d68f4e5931f93fadc3b7afac5092e195b80f"
				description := "lorenzo"
				url := "https://lorenzo-protocol.io"
				agentId := suite.lorenzoApp.AgentKeeper.AddAgent(
					suite.ctx,
					name, btcReceivingAddress, ethAddr, description, url)
				suite.Require().NotEqual(agentId, 0)
				suite.Require().Equal(agentId, uint64(1))
				yatAddr, err := suite.lorenzoApp.PlanKeeper.DeployYATContract(
					suite.ctx, "lorenzo", "ALRZ")
				suite.Require().NoError(err)

				request.YatContractAddress = yatAddr.Hex()
			},
			validation: func(request *types.MsgCreatePlan) {
				plan, found := suite.lorenzoApp.PlanKeeper.GetPlan(suite.ctx, uint64(1))
				suite.Require().True(found)
				suite.Require().Equal(plan.Name, request.Name)
				suite.Require().Equal(plan.PlanDescUri, request.PlanDescUri)
				suite.Require().Equal(plan.AgentId, request.AgentId)
				suite.Require().Equal(plan.PlanStartTime, request.PlanStartTime)
				suite.Require().Equal(plan.PeriodTime, request.PeriodTime)

				planContractAddress := common.HexToAddress(plan.ContractAddress)

				// YatContractAddress
				yatContractAddress, err := suite.lorenzoApp.PlanKeeper.YatContractAddress(suite.ctx, planContractAddress)
				suite.Require().NoError(err)
				suite.Require().Equal(yatContractAddress, common.HexToAddress(plan.YatContractAddress))
				// StakePlanName
				stakePlanName, err := suite.lorenzoApp.PlanKeeper.StakePlanName(suite.ctx, planContractAddress)
				suite.Require().NoError(err)
				suite.Require().Equal(stakePlanName, plan.Name)

				agentId, err := suite.lorenzoApp.PlanKeeper.AgentId(suite.ctx, planContractAddress)
				suite.Require().NoError(err)
				suite.Require().Equal(agentId, plan.AgentId)

				planId, err := suite.lorenzoApp.PlanKeeper.PlanId(suite.ctx, planContractAddress)
				suite.Require().NoError(err)
				suite.Require().Equal(planId, plan.Id)

				// PlanDesc
				planDesc, err := suite.lorenzoApp.PlanKeeper.PlanDesc(suite.ctx, planContractAddress)
				suite.Require().NoError(err)
				suite.Require().Equal(planDesc, plan.PlanDescUri)

				// PlanStartTime
				PlanStartTime, err := suite.lorenzoApp.PlanKeeper.PlanStartTime(suite.ctx, planContractAddress)
				suite.Require().NoError(err)
				suite.Require().Equal(PlanStartTime, plan.PlanStartTime)

				// PeriodTime
				PeriodTime, err := suite.lorenzoApp.PlanKeeper.PeriodTime(suite.ctx, planContractAddress)
				suite.Require().NoError(err)
				suite.Require().Equal(PeriodTime, plan.PeriodTime)

				// NextRewardReceiveBlock
				nextRewardReceiveBlock, err := suite.lorenzoApp.PlanKeeper.NextRewardReceiveTime(suite.ctx, planContractAddress)
				suite.Require().NoError(err)
				suite.Require().Equal(nextRewardReceiveBlock, plan.PlanStartTime+plan.PeriodTime)

				// ClaimRoundId
				claimRoundId, err := suite.lorenzoApp.PlanKeeper.ClaimRoundId(suite.ctx, planContractAddress)
				suite.Require().NoError(err)
				suite.Require().Equal(claimRoundId, uint64(0))
			},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("MsgCreatPlan - %s", tc.name), func() {
			suite.SetupTest()
			if tc.malleate != nil {
				tc.malleate(tc.request)
			}
			_, err := suite.msgServer.CreatePlan(suite.ctx, tc.request)
			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
			if tc.validation != nil {
				tc.validation(tc.request)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestUpgradePlan() {
	testCases := []struct {
		name       string
		request    *types.MsgUpgradePlan
		malleate   func(request *types.MsgUpgradePlan)
		validation func(request *types.MsgUpgradePlan)
		expectErr  bool
	}{
		{
			name:      "fail - invalid sender",
			request:   &types.MsgUpgradePlan{Authority: "foobar"},
			expectErr: true,
		},
		{
			name:      "fail - sender not authorized",
			request:   &types.MsgUpgradePlan{Authority: "lrz1tffj9qp3wpdnuds443c86wffrac4jkapkjmmcy"},
			expectErr: true,
		},
		{
			name: "fail - invalid implementation address",
			request: &types.MsgUpgradePlan{
				Authority:      testAdmin.String(),
				Implementation: "0x123456",
			},
			expectErr: true,
		},
		{
			name: "fail - old implementation address is empty",
			request: &types.MsgUpgradePlan{
				Authority: testAdmin.String(),
			},
			malleate: func(request *types.MsgUpgradePlan) {
				newImplementation, err := suite.lorenzoApp.PlanKeeper.DeployStakePlanLogicContract(suite.ctx)
				suite.Require().NoError(err)
				request.Implementation = newImplementation.Hex()
			},
			expectErr: true,
		},
		{
			name: "success - valid upgrade plan",
			request: &types.MsgUpgradePlan{
				Authority: testAdmin.String(),
			},
			malleate: func(request *types.MsgUpgradePlan) {
				// 1. deploy old implementation
				var stakePlanContract evmtypes.CompiledContract
				stakePlanJSON := []byte(`{
  "abi":"[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyMerkleRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMerkleProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimYATToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"MerkleRootSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"planId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MintYAT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"adminPauseBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminUnpauseBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"agentId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId_\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"leafNode_\",\"type\":\"bytes32\"}],\"name\":\"claimLeafNode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimRoundId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"roundId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof_\",\"type\":\"bytes32[]\"}],\"name\":\"claimYATToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"planName_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"planDescUri_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"planId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"agentId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"planStartTime_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"periodTime_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"yatContractAddress_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId_\",\"type\":\"uint256\"}],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextRewardReceiveTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"planDesc\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"planId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"planStartTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId_\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"newMerkleRoot_\",\"type\":\"bytes32\"}],\"name\":\"setMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newPlanDescUri\",\"type\":\"string\"}],\"name\":\"setPlanDesc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakePlanName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"yatContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
  "bin": "6080604052348015600f57600080fd5b506016601a565b60ca565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161560695760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b039081161460c75780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6111f8806100d96000396000f3fe608060405234801561001057600080fd5b50600436106101425760003560e01c80638dbc9fdf116100b8578063e84f43b71161007c578063e84f43b71461028e578063e97f239014610296578063efe19ede146102a9578063f20a4f33146102b1578063f2fde38b146102df578063f7dd6d44146102f257600080fd5b80638dbc9fdf1461024e578063a6e57d9914610263578063bd30008914610276578063c76059db1461027e578063cd57ccce1461028657600080fd5b806341f837241161010a57806341f83724146101be57806343eab934146101c657806359e1656a146101d95780635c975abb146101fe578063715018a6146102165780638da5cb5b1461021e57600080fd5b806318712c21146101475780631d31fac01461015c5780632ddfb152146101735780633c70b3571461017b57806340c10f191461019b575b600080fd5b61015a610155366004610cdb565b6102fa565b005b6005545b6040519081526020015b60405180910390f35b600654610160565b610160610189366004610cfd565b6000908152600a602052604090205490565b6101ae6101a9366004610d32565b6103ca565b604051901515815260200161016a565b600854610160565b6101ae6101d4366004610d5c565b6104ee565b6007546001600160a01b03165b6040516001600160a01b03909116815260200161016a565b6000805160206111a38339815191525460ff166101ae565b61015a610707565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b03166101e6565b61025661071b565b60405161016a9190610df5565b61015a610271366004610ee8565b6107ad565b600254610160565b61025661095d565b61015a61096c565b600354610160565b61015a6102a4366004610f88565b61097c565b61015a610994565b6101ae6102bf366004610cdb565b600091825260096020908152604080842092845291905290205460ff1690565b61015a6102ed366004610fc5565b6109a4565b600454610160565b6103026109e7565b80610320576040516385ac2b9960e01b815260040160405180910390fd5b816008540361034257604051635435b28960e11b815260040160405180910390fd5b6006544311156103605760055460065461035c9190610ff6565b6006555b600880546000918261037183611009565b909155506000818152600a6020526040908190208490555190915081907fb04b7d6145a7588fdcf339a22877d5965f861c171204fc37688058c5f6c06d3b906103bd9085815260200190565b60405180910390a2505050565b60006103d46109e7565b6103dc610a42565b8115806103ea575060045442105b1561040857604051635435b28960e11b815260040160405180910390fd5b600654431115610426576005546006546104229190610ff6565b6006555b6007546040516340c10f1960e01b81526001600160a01b03858116600483015260248201859052909116906340c10f19906044016020604051808303816000875af1158015610479573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061049d9190611022565b50826001600160a01b03166002547fe9cc7b20f3ef886661af843b84de3d24a481f25ae98d81285497a17e2fe77186846040516104dc91815260200190565b60405180910390a35060015b92915050565b6000848152600a602052604081205461051a576040516385ac2b9960e01b815260040160405180910390fd5b6040516bffffffffffffffffffffffff19606088901b1660208201526034810185905260009060540160408051601f19818403018152918152815160209283012060008981526009845282812082825290935291205490915060ff161561059457604051630c8d9eab60e31b815260040160405180910390fd5b6105df84848080602002602001604051908101604052809392919081815260200183836020028082843760009201829052508b8152600a60205260409020549250859150610a739050565b6105fc5760405163582f497d60e11b815260040160405180910390fd5b60065442111561061a576005546006546106169190610ff6565b6006555b600086815260096020908152604080832084845290915290819020805460ff1916600117905560075490516340c10f1960e01b81526001600160a01b03898116600483015260248201889052909116906340c10f19906044016020604051808303816000875af1158015610692573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106b69190611022565b50866001600160a01b03167f3fba8f41c65d7cd646bc82e68db6aaff16e62b565dc52f036224fe8e79e09f25866040516106f291815260200190565b60405180910390a25060019695505050505050565b61070f6109e7565b6107196000610a89565b565b60606001805461072a90611044565b80601f016020809104026020016040519081016040528092919081815260200182805461075690611044565b80156107a35780601f10610778576101008083540402835291602001916107a3565b820191906000526020600020905b81548152906001019060200180831161078657829003601f168201915b5050505050905090565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff166000811580156107f35750825b905060008267ffffffffffffffff1660011480156108105750303b155b90508115801561081e575080155b1561083c5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561086657845460ff60401b1916600160401b1785555b87431180610872575086155b8061088457506001600160a01b038616155b156108a257604051635435b28960e11b815260040160405180910390fd5b6108ab33610afa565b6108b3610b0b565b60006108bf8d826110cd565b5060028a905560016108d18c826110cd565b506003899055600488905560058790556108eb8789610ff6565b600655600780546001600160a01b0319166001600160a01b038816179055831561094f57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050505050565b60606000805461072a90611044565b6109746109e7565b610719610b1b565b6109846109e7565b600161099082826110cd565b5050565b61099c6109e7565b610719610b7e565b6109ac6109e7565b6001600160a01b0381166109db57604051631e4fbdf760e01b8152600060048201526024015b60405180910390fd5b6109e481610a89565b50565b33610a197f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146107195760405163118cdaa760e01b81523360048201526024016109d2565b6000805160206111a38339815191525460ff16156107195760405163d93c066560e01b815260040160405180910390fd5b600082610a808584610bc4565b14949350505050565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b610b02610c07565b6109e481610c50565b610b13610c07565b610719610c58565b610b23610a42565b6000805160206111a3833981519152805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258335b6040516001600160a01b03909116815260200160405180910390a150565b610b86610c79565b6000805160206111a3833981519152805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa33610b60565b600081815b8451811015610bff57610bf582868381518110610be857610be861118c565b6020026020010151610ca9565b9150600101610bc9565b509392505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661071957604051631afcd79f60e31b815260040160405180910390fd5b6109ac610c07565b610c60610c07565b6000805160206111a3833981519152805460ff19169055565b6000805160206111a38339815191525460ff1661071957604051638dfc202b60e01b815260040160405180910390fd5b6000818310610cc5576000828152602084905260409020610cd4565b60008381526020839052604090205b9392505050565b60008060408385031215610cee57600080fd5b50508035926020909101359150565b600060208284031215610d0f57600080fd5b5035919050565b80356001600160a01b0381168114610d2d57600080fd5b919050565b60008060408385031215610d4557600080fd5b610d4e83610d16565b946020939093013593505050565b600080600080600060808688031215610d7457600080fd5b610d7d86610d16565b94506020860135935060408601359250606086013567ffffffffffffffff811115610da757600080fd5b8601601f81018813610db857600080fd5b803567ffffffffffffffff811115610dcf57600080fd5b8860208260051b8401011115610de457600080fd5b959894975092955050506020019190565b602081526000825180602084015260005b81811015610e235760208186018101516040868401015201610e06565b506000604082850101526040601f19601f83011684010191505092915050565b634e487b7160e01b600052604160045260246000fd5b600082601f830112610e6a57600080fd5b813567ffffffffffffffff811115610e8457610e84610e43565b604051601f8201601f19908116603f0116810167ffffffffffffffff81118282101715610eb357610eb3610e43565b604052818152838201602001851015610ecb57600080fd5b816020850160208301376000918101602001919091529392505050565b600080600080600080600060e0888a031215610f0357600080fd5b873567ffffffffffffffff811115610f1a57600080fd5b610f268a828b01610e59565b975050602088013567ffffffffffffffff811115610f4357600080fd5b610f4f8a828b01610e59565b96505060408801359450606088013593506080880135925060a08801359150610f7a60c08901610d16565b905092959891949750929550565b600060208284031215610f9a57600080fd5b813567ffffffffffffffff811115610fb157600080fd5b610fbd84828501610e59565b949350505050565b600060208284031215610fd757600080fd5b610cd482610d16565b634e487b7160e01b600052601160045260246000fd5b808201808211156104e8576104e8610fe0565b60006001820161101b5761101b610fe0565b5060010190565b60006020828403121561103457600080fd5b81518015158114610cd457600080fd5b600181811c9082168061105857607f821691505b60208210810361107857634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156110c857806000526020600020601f840160051c810160208510156110a55750805b601f840160051c820191505b818110156110c557600081556001016110b1565b50505b505050565b815167ffffffffffffffff8111156110e7576110e7610e43565b6110fb816110f58454611044565b8461107e565b6020601f82116001811461112f57600083156111175750848201515b600019600385901b1c1916600184901b1784556110c5565b600084815260208120601f198516915b8281101561115f578785015182556020948501946001909201910161113f565b508482101561117d5786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b634e487b7160e01b600052603260045260246000fdfecd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300a2646970667358221220527f8539f7a30bb5d7e1df18925462556c5d169db619c0ada5080b289f216a0d64736f6c634300081a0033",
  "contractName": "StakePlan"
}`)
				err := json.Unmarshal(stakePlanJSON, &stakePlanContract)
				suite.Require().NoError(err)
				suite.Require().NotEqual(len(stakePlanContract.Bin), 0)

				logicAddr, err := suite.lorenzoApp.PlanKeeper.DeployContract(suite.ctx, stakePlanContract)
				suite.Require().NoError(err)

				// 2. deploy a new plan beacon contract
				beaconAddr, err := suite.lorenzoApp.PlanKeeper.DeployBeaconForPlan(suite.ctx, logicAddr)
				suite.Require().NoError(err)

				// set beacon address in params
				params := suite.lorenzoApp.PlanKeeper.GetParams(suite.ctx)
				params.Beacon = beaconAddr.Hex()
				err = suite.lorenzoApp.PlanKeeper.SetParams(suite.ctx, params)
				suite.Require().NoError(err)

				// suite.Commit()

				// set data
				// create agent
				name := "sinohope4"
				btcReceivingAddress := "3C7VPws9fMW3kcwRJvMkSVdqMs4SAhQCqq"
				ethAddr := "0x6508d68f4e5931f93fadc3b7afac5092e195b80f"
				description := "lorenzo"
				url := "https://lorenzo-protocol.io"
				agentId := suite.lorenzoApp.AgentKeeper.AddAgent(
					suite.ctx,
					name, btcReceivingAddress, ethAddr, description, url)
				suite.Require().NotEqual(agentId, 0)
				suite.Require().Equal(agentId, uint64(1))

				yatAddr, err := suite.lorenzoApp.PlanKeeper.DeployYATContract(
					suite.ctx, "lorenzo", "ALRZ")
				suite.Require().NoError(err)
				// create plan
				planReq := types.Plan{
					Name:               "lorenzo-stake-plan",
					PlanDescUri:        "https://lorenzo-protocol.io/lorenzo-stake-plan",
					AgentId:            uint64(1),
					PlanStartTime:      uint64(time.Now().UTC().Unix()) + 1000,
					PeriodTime:         1000,
					YatContractAddress: yatAddr.Hex(),
				}

				planResult, err := suite.lorenzoApp.PlanKeeper.AddPlan(suite.ctx, planReq)
				suite.Require().NoError(err)

				// set merkle root
				merkelRoot := "0x39c19150c14c397b133682e95742b651babde3418edaaa4375a3197604159346"
				err = suite.lorenzoApp.PlanKeeper.SetMerkleRoot(
					suite.ctx,
					common.HexToAddress(planResult.ContractAddress),
					sdkmath.NewInt(1).BigInt(),
					merkelRoot)
				suite.Require().NoError(err)

				// 3. deploy new implementation
				newImplementation, err := suite.lorenzoApp.PlanKeeper.DeployStakePlanLogicContract(suite.ctx)
				suite.Require().NoError(err)
				request.Implementation = newImplementation.Hex()
			},
			expectErr: false,
			validation: func(request *types.MsgUpgradePlan) {
				planImplementation, err := suite.lorenzoApp.PlanKeeper.GetPlanImplementationFromBeacon(suite.ctx)
				suite.Require().NoError(err)
				suite.Require().Equal(planImplementation.Hex(), request.Implementation)

				// check data
				plan, found := suite.lorenzoApp.PlanKeeper.GetPlan(suite.ctx, 1)
				suite.Require().True(found)

				// get merkle root
				contractAddr := common.HexToAddress(plan.ContractAddress)
				roundId, err := suite.lorenzoApp.PlanKeeper.ClaimRoundId(suite.ctx, contractAddr)
				suite.Require().NoError(err)
				suite.Require().Equal(roundId, uint64(1))
				merkleRoot, err := suite.lorenzoApp.PlanKeeper.MerkleRoot(
					suite.ctx, contractAddr, sdkmath.NewIntFromUint64(roundId-1).BigInt())
				suite.Require().NoError(err)
				expectMerkelRoot := "0x39c19150c14c397b133682e95742b651babde3418edaaa4375a3197604159346"
				suite.Require().Equal(expectMerkelRoot, merkleRoot)
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("MsgUpgradePlan - %s", tc.name), func() {
			suite.SetupTest()
			if tc.malleate != nil {
				tc.malleate(tc.request)
			}
			_, err := suite.msgServer.UpgradePlan(suite.ctx, tc.request)
			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
			if tc.validation != nil {
				tc.validation(tc.request)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestUpdatePlanStatus() {
	testCases := []struct {
		name       string
		request    *types.MsgUpdatePlanStatus
		malleate   func()
		validation func()
		expectErr  bool
	}{
		{
			name:      "fail - invalid sender",
			request:   &types.MsgUpdatePlanStatus{Sender: "foobar"},
			expectErr: true,
		},
		{
			name:      "fail - sender not authorized",
			request:   &types.MsgUpdatePlanStatus{Sender: "lrz1tffj9qp3wpdnuds443c86wffrac4jkapkjmmcy"},
			expectErr: true,
		},
		{
			name: "fail - plan not found",
			request: &types.MsgUpdatePlanStatus{
				Sender: testAdmin.String(),
				PlanId: 1,
				Status: types.PlanStatus_Pause,
			},
			expectErr: true,
		},
		{
			name: "fail - plan status equals to current status",
			request: &types.MsgUpdatePlanStatus{
				Sender: testAdmin.String(),
				PlanId: 1,
				Status: types.PlanStatus_Unpause,
			},
			malleate: func() {
				suite.Commit()
				// create agent
				name := "sinohope4"
				btcReceivingAddress := "3C7VPws9fMW3kcwRJvMkSVdqMs4SAhQCqq"
				ethAddr := "0x6508d68f4e5931f93fadc3b7afac5092e195b80f"
				description := "lorenzo"
				url := "https://lorenzo-protocol.io"
				agentId := suite.lorenzoApp.AgentKeeper.AddAgent(
					suite.ctx,
					name, btcReceivingAddress, ethAddr, description, url)
				suite.Require().NotEqual(agentId, 0)
				suite.Require().Equal(agentId, uint64(1))

				yatAddr, err := suite.lorenzoApp.PlanKeeper.DeployYATContract(
					suite.ctx, "lorenzo", "ALRZ")
				suite.Require().NoError(err)
				// create plan
				planReq := types.Plan{
					Name:               "lorenzo-stake-plan",
					PlanDescUri:        "https://lorenzo-protocol.io/lorenzo-stake-plan",
					AgentId:            uint64(1),
					PlanStartTime:      uint64(time.Now().UTC().Unix()) + 1000,
					PeriodTime:         1000,
					YatContractAddress: yatAddr.Hex(),
				}

				_, err = suite.lorenzoApp.PlanKeeper.AddPlan(suite.ctx, planReq)
				suite.Require().NoError(err)
			},
			expectErr: true,
		},
		{
			name: "success - valid update plan status, plan pause",
			request: &types.MsgUpdatePlanStatus{
				Sender: testAdmin.String(),
				PlanId: 1,
				Status: types.PlanStatus_Pause,
			},
			malleate: func() {
				suite.Commit()
				// create agent
				name := "sinohope4"
				btcReceivingAddress := "3C7VPws9fMW3kcwRJvMkSVdqMs4SAhQCqq"
				ethAddr := "0x6508d68f4e5931f93fadc3b7afac5092e195b80f"
				description := "lorenzo"
				url := "https://lorenzo-protocol.io"
				agentId := suite.lorenzoApp.AgentKeeper.AddAgent(
					suite.ctx,
					name, btcReceivingAddress, ethAddr, description, url)
				suite.Require().NotEqual(agentId, 0)
				suite.Require().Equal(agentId, uint64(1))

				yatAddr, err := suite.lorenzoApp.PlanKeeper.DeployYATContract(
					suite.ctx, "lorenzo", "ALRZ")
				suite.Require().NoError(err)
				// create plan
				planReq := types.Plan{
					Name:               "lorenzo-stake-plan",
					PlanDescUri:        "https://lorenzo-protocol.io/lorenzo-stake-plan",
					AgentId:            uint64(1),
					PlanStartTime:      uint64(time.Now().UTC().Unix()) + 1000,
					PeriodTime:         1000,
					YatContractAddress: yatAddr.Hex(),
				}

				_, err = suite.lorenzoApp.PlanKeeper.AddPlan(suite.ctx, planReq)
				suite.Require().NoError(err)
			},
			expectErr: false,
			validation: func() {
				plan, found := suite.lorenzoApp.PlanKeeper.GetPlan(suite.ctx, uint64(1))
				suite.Require().True(found)
				suite.Require().Equal(plan.Enabled, types.PlanStatus_Pause)
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("MsgUpdatePlanStatus - %s", tc.name), func() {
			suite.SetupTest()
			if tc.malleate != nil {
				tc.malleate()
			}
			_, err := suite.msgServer.UpdatePlanStatus(suite.ctx, tc.request)
			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
			if tc.validation != nil {
				tc.validation()
			}
		})
	}
}

func (suite *KeeperTestSuite) TestSetMerkleRoot() {
	testCases := []struct {
		name       string
		request    *types.MsgSetMerkleRoot
		malleate   func() string
		validation func(string, *types.MsgSetMerkleRoot)
		expectErr  bool
	}{
		{
			name:      "fail - invalid sender",
			request:   &types.MsgSetMerkleRoot{Sender: "foobar"},
			expectErr: true,
		},
		{
			name:      "fail - sender not authorized",
			request:   &types.MsgSetMerkleRoot{Sender: "lrz1tffj9qp3wpdnuds443c86wffrac4jkapkjmmcy"},
			expectErr: true,
		},
		{
			name: "fail - plan not found",
			request: &types.MsgSetMerkleRoot{
				Sender:     testAdmin.String(),
				PlanId:     1,
				MerkleRoot: "0x34337eb06160f22cfc735517076cb8d69f60afae27700d20e918cfb41f9faca7",
			},
			expectErr: true,
		},
		{
			name: "success - valid set merkle root",
			request: &types.MsgSetMerkleRoot{
				Sender:     testAdmin.String(),
				PlanId:     1,
				RoundId:    sdkmath.NewInt(0),
				MerkleRoot: "0x34337eb06160f22cfc735517076cb8d69f60afae27700d20e918cfb41f9faca7",
			},
			malleate: func() string {
				suite.Commit()
				// create agent
				name := "sinohope4"
				btcReceivingAddress := "3C7VPws9fMW3kcwRJvMkSVdqMs4SAhQCqq"
				ethAddr := "0x6508d68f4e5931f93fadc3b7afac5092e195b80f"
				description := "lorenzo"
				url := "https://lorenzo-protocol.io"
				agentId := suite.lorenzoApp.AgentKeeper.AddAgent(
					suite.ctx,
					name, btcReceivingAddress, ethAddr, description, url)
				suite.Require().NotEqual(agentId, 0)
				suite.Require().Equal(agentId, uint64(1))

				yatAddr, err := suite.lorenzoApp.PlanKeeper.DeployYATContract(
					suite.ctx, "lorenzo", "ALRZ")
				suite.Require().NoError(err)
				// create plan
				planReq := types.Plan{
					Name:               "lorenzo-stake-plan",
					PlanDescUri:        "https://lorenzo-protocol.io/lorenzo-stake-plan",
					AgentId:            uint64(1),
					PlanStartTime:      uint64(time.Now().UTC().Unix()) + 1000,
					PeriodTime:         1000,
					YatContractAddress: yatAddr.Hex(),
				}

				planResult, err := suite.lorenzoApp.PlanKeeper.AddPlan(suite.ctx, planReq)
				suite.Require().NoError(err)
				return planResult.ContractAddress
			},
			validation: func(contractAddrHex string, request *types.MsgSetMerkleRoot) {
				contractAddr := common.HexToAddress(contractAddrHex)
				roundId, err := suite.lorenzoApp.PlanKeeper.ClaimRoundId(suite.ctx, contractAddr)
				suite.Require().NoError(err)
				suite.Require().Equal(roundId, uint64(1))
				merkleRoot, err := suite.lorenzoApp.PlanKeeper.MerkleRoot(
					suite.ctx, contractAddr, sdkmath.NewIntFromUint64(roundId-1).BigInt())
				suite.Require().NoError(err)
				suite.Require().Equal(merkleRoot, request.MerkleRoot)
			},
			expectErr: false,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("MsgSetMerkleRoot - %s", tc.name), func() {
			suite.SetupTest()
			var contractAddrHex string
			if tc.malleate != nil {
				contractAddrHex = tc.malleate()
			}
			_, err := suite.msgServer.SetMerkleRoot(suite.ctx, tc.request)
			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
			if tc.validation != nil {
				tc.validation(contractAddrHex, tc.request)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestClaims() {
	testCases := []struct {
		name       string
		request    *types.MsgClaims
		malleate   func()
		validation func()
		expectErr  bool
	}{
		{
			name:      "fail - invalid sender",
			request:   &types.MsgClaims{Sender: "foobar"},
			expectErr: true,
		},
		{
			name: "fail - plan not found",
			request: &types.MsgClaims{
				Sender:      "lrz1cpldpp5960ed8s63w4v9fml84w875wv0emcda5",
				PlanId:      1,
				Receiver:    "lrz1cpldpp5960ed8s63w4v9fml84w875wv0emcda5",
				RoundId:     sdkmath.NewInt(0),
				Amount:      sdkmath.NewInt(100),
				MerkleProof: "0x1764cb495e1c2565f6d033e298a2d46a527c93a5a48c8b318fa05e9b07489b33",
			},
			expectErr: true,
		},
		{
			name: "fail - plan not set merkle root",
			request: &types.MsgClaims{
				Sender:      "lrz1cpldpp5960ed8s63w4v9fml84w875wv0emcda5",
				PlanId:      1,
				Receiver:    "0xc07ed08685d3F2D3c351755854EFE7ab8fEa398F",
				RoundId:     sdkmath.NewInt(0),
				Amount:      sdkmath.NewInt(100),
				MerkleProof: "0x365cc96c249dc95f3f2e4934371b55ee1c5ef9e6f6da6407b1ec26aa6cd12109",
			},
			malleate: func() {
				suite.Commit()
				// create agent
				name := "sinohope4"
				btcReceivingAddress := "3C7VPws9fMW3kcwRJvMkSVdqMs4SAhQCqq"
				ethAddr := "0x6508d68f4e5931f93fadc3b7afac5092e195b80f"
				description := "lorenzo"
				url := "https://lorenzo-protocol.io"
				agentId := suite.lorenzoApp.AgentKeeper.AddAgent(
					suite.ctx,
					name, btcReceivingAddress, ethAddr, description, url)
				suite.Require().NotEqual(agentId, 0)
				suite.Require().Equal(agentId, uint64(1))

				yatAddr, err := suite.lorenzoApp.PlanKeeper.DeployYATContract(
					suite.ctx, "lorenzo", "ALRZ")
				suite.Require().NoError(err)
				// create plan
				planReq := types.Plan{
					Name:               "lorenzo-stake-plan",
					PlanDescUri:        "https://lorenzo-protocol.io/lorenzo-stake-plan",
					AgentId:            uint64(1),
					PlanStartTime:      uint64(time.Now().UTC().Unix()) + 1000,
					PeriodTime:         1000,
					YatContractAddress: yatAddr.Hex(),
				}

				planResult, err := suite.lorenzoApp.PlanKeeper.AddPlan(suite.ctx, planReq)
				suite.Require().NoError(err)

				// set minter
				err = suite.lorenzoApp.PlanKeeper.SetMinter(suite.ctx, yatAddr, common.HexToAddress(planResult.ContractAddress))
				suite.Require().NoError(err)
			},
			expectErr: true,
		},
		{
			name: "fail - plan is disabled",
			request: &types.MsgClaims{
				Sender:      "lrz1cpldpp5960ed8s63w4v9fml84w875wv0emcda5",
				PlanId:      1,
				Receiver:    "0xc07ed08685d3F2D3c351755854EFE7ab8fEa398F",
				RoundId:     sdkmath.NewInt(0),
				Amount:      sdkmath.NewInt(100),
				MerkleProof: "0x365cc96c249dc95f3f2e4934371b55ee1c5ef9e6f6da6407b1ec26aa6cd12109",
			},
			malleate: func() {
				suite.Commit()
				// create agent
				name := "sinohope4"
				btcReceivingAddress := "3C7VPws9fMW3kcwRJvMkSVdqMs4SAhQCqq"
				ethAddr := "0x6508d68f4e5931f93fadc3b7afac5092e195b80f"
				description := "lorenzo"
				url := "https://lorenzo-protocol.io"
				agentId := suite.lorenzoApp.AgentKeeper.AddAgent(
					suite.ctx,
					name, btcReceivingAddress, ethAddr, description, url)
				suite.Require().NotEqual(agentId, 0)
				suite.Require().Equal(agentId, uint64(1))

				yatAddr, err := suite.lorenzoApp.PlanKeeper.DeployYATContract(
					suite.ctx, "lorenzo", "ALRZ")
				suite.Require().NoError(err)
				// create plan
				planReq := types.Plan{
					Name:               "lorenzo-stake-plan",
					PlanDescUri:        "https://lorenzo-protocol.io/lorenzo-stake-plan",
					AgentId:            uint64(1),
					PlanStartTime:      uint64(time.Now().UTC().Unix()) + 1000,
					PeriodTime:         1000,
					YatContractAddress: yatAddr.Hex(),
				}

				planResult, err := suite.lorenzoApp.PlanKeeper.AddPlan(suite.ctx, planReq)
				suite.Require().NoError(err)

				// set merkle root
				merkelRoot := "0x39c19150c14c397b133682e95742b651babde3418edaaa4375a3197604159346"
				err = suite.lorenzoApp.PlanKeeper.SetMerkleRoot(
					suite.ctx,
					common.HexToAddress(planResult.ContractAddress),
					sdkmath.NewInt(0).BigInt(),
					merkelRoot)
				suite.Require().NoError(err)

				// set minter
				err = suite.lorenzoApp.PlanKeeper.SetMinter(suite.ctx, yatAddr, common.HexToAddress(planResult.ContractAddress))
				suite.Require().NoError(err)

				// plan disable
				err = suite.lorenzoApp.PlanKeeper.UpdatePlanStatus(suite.ctx, planResult.Id, types.PlanStatus_Pause)
				suite.Require().NoError(err)
			},
			expectErr: true,
		},
		{
			name: "success - valid claims",
			request: &types.MsgClaims{
				Sender:      "lrz1cpldpp5960ed8s63w4v9fml84w875wv0emcda5",
				PlanId:      1,
				Receiver:    "0xc07ed08685d3F2D3c351755854EFE7ab8fEa398F",
				RoundId:     sdkmath.NewInt(0),
				Amount:      sdkmath.NewInt(100),
				MerkleProof: "0x365cc96c249dc95f3f2e4934371b55ee1c5ef9e6f6da6407b1ec26aa6cd12109",
			},
			malleate: func() {
				suite.Commit()
				// create agent
				name := "sinohope4"
				btcReceivingAddress := "3C7VPws9fMW3kcwRJvMkSVdqMs4SAhQCqq"
				ethAddr := "0x6508d68f4e5931f93fadc3b7afac5092e195b80f"
				description := "lorenzo"
				url := "https://lorenzo-protocol.io"
				agentId := suite.lorenzoApp.AgentKeeper.AddAgent(
					suite.ctx,
					name, btcReceivingAddress, ethAddr, description, url)
				suite.Require().NotEqual(agentId, 0)
				suite.Require().Equal(agentId, uint64(1))

				yatAddr, err := suite.lorenzoApp.PlanKeeper.DeployYATContract(
					suite.ctx, "lorenzo", "ALRZ")
				suite.Require().NoError(err)
				// create plan
				planReq := types.Plan{
					Name:               "lorenzo-stake-plan",
					PlanDescUri:        "https://lorenzo-protocol.io/lorenzo-stake-plan",
					AgentId:            uint64(1),
					PlanStartTime:      uint64(time.Now().UTC().Unix()) + 1000,
					PeriodTime:         1000,
					YatContractAddress: yatAddr.Hex(),
				}

				planResult, err := suite.lorenzoApp.PlanKeeper.AddPlan(suite.ctx, planReq)
				suite.Require().NoError(err)

				// set merkle root
				merkelRoot := "0x39c19150c14c397b133682e95742b651babde3418edaaa4375a3197604159346"
				err = suite.lorenzoApp.PlanKeeper.SetMerkleRoot(
					suite.ctx,
					common.HexToAddress(planResult.ContractAddress),
					sdkmath.NewInt(0).BigInt(),
					merkelRoot)
				suite.Require().NoError(err)

				// set minter
				err = suite.lorenzoApp.PlanKeeper.SetMinter(suite.ctx, yatAddr, common.HexToAddress(planResult.ContractAddress))
				suite.Require().NoError(err)
			},
			validation: func() {
				plan, found := suite.lorenzoApp.PlanKeeper.GetPlan(suite.ctx, uint64(1))
				suite.Require().True(found)
				amount, err := suite.lorenzoApp.PlanKeeper.BalanceOfFromYAT(
					suite.ctx, common.HexToAddress(plan.YatContractAddress),
					common.HexToAddress("0xc07ed08685d3F2D3c351755854EFE7ab8fEa398F"))
				suite.Require().NoError(err)
				suite.Require().Equal(amount, sdkmath.NewInt(100).BigInt())
			},
			expectErr: false,
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("MsgClaims - %s", tc.name), func() {
			suite.SetupTest()
			if tc.malleate != nil {
				tc.malleate()
			}
			_, err := suite.msgServer.Claims(suite.ctx, tc.request)
			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
			if tc.validation != nil {
				tc.validation()
			}
		})
	}
}

func (suite *KeeperTestSuite) TestCreateYAT() {
	testCases := []struct {
		name       string
		request    *types.MsgCreateYAT
		malleate   func()
		validation func(string)
		expectErr  bool
	}{
		{
			name:      "fail - invalid sender",
			request:   &types.MsgCreateYAT{Sender: "foobar"},
			expectErr: true,
		},
		{
			name:      "fail - sender not authorized",
			request:   &types.MsgCreateYAT{Sender: "lrz1tffj9qp3wpdnuds443c86wffrac4jkapkjmmcy"},
			expectErr: true,
		},
		{
			name: "success - valid create yat",
			request: &types.MsgCreateYAT{
				Sender: testAdmin.String(),
				Name:   "lorenzo",
				Symbol: "ALRZ",
			},
			expectErr: false,
			validation: func(yatAddressHex string) {
				yatAddress := common.HexToAddress(yatAddressHex)
				owner, err := suite.lorenzoApp.PlanKeeper.GetOwner(suite.ctx, yatAddress)
				suite.Require().NoError(err)
				suite.Require().Equal(owner, types.ModuleAddress)
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("MsgCreateYAT - %s", tc.name), func() {
			suite.SetupTest()
			if tc.malleate != nil {
				tc.malleate()
			}
			result, err := suite.msgServer.CreateYAT(suite.ctx, tc.request)
			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
			if tc.validation != nil {
				tc.validation(result.ContractAddress)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestSetMinter() {
	testCases := []struct {
		name       string
		request    *types.MsgSetMinter
		malleate   func(*types.MsgSetMinter)
		validation func(*types.MsgSetMinter)
		expectErr  bool
	}{
		{
			name:      "fail - invalid sender",
			request:   &types.MsgSetMinter{Sender: "foobar"},
			expectErr: true,
		},
		{
			name:      "fail - sender not authorized",
			request:   &types.MsgSetMinter{Sender: "lrz1tffj9qp3wpdnuds443c86wffrac4jkapkjmmcy"},
			expectErr: true,
		},
		{
			name: "fail - invalid minter address",
			request: &types.MsgSetMinter{
				Sender:          testAdmin.String(),
				Minter:          "0x123456",
				ContractAddress: "0xbCC0CdF7683120a1965A343245FA602314C13b9A",
			},
			expectErr: true,
		},
		{
			name: "fail - invalid yat contract address",
			request: &types.MsgSetMinter{
				Sender:          testAdmin.String(),
				Minter:          types.ModuleAddress.Hex(),
				ContractAddress: "0x123456",
			},
			expectErr: true,
		},
		{
			name: "fail - yat not exists",
			request: &types.MsgSetMinter{
				Sender:          testAdmin.String(),
				ContractAddress: "0xbCC0CdF7683120a1965A343245FA602314C13b9A",
			},
			expectErr: true,
		},
		{
			name: "fail - minter not exists",
			request: &types.MsgSetMinter{
				Sender: testAdmin.String(),
				Minter: "0xbCC0CdF7683120a1965A343245FA602314C13b9A",
			},
			expectErr: true,
			malleate: func(request *types.MsgSetMinter) {
				yatAddr, err := suite.lorenzoApp.PlanKeeper.DeployYATContract(
					suite.ctx, "lorenzo", "ALRZ")
				suite.Require().NoError(err)
				request.ContractAddress = yatAddr.Hex()
			},
		},
		{
			name: "success - valid set minter",
			request: &types.MsgSetMinter{
				Sender: testAdmin.String(),
			},
			malleate: func(request *types.MsgSetMinter) {
				suite.Commit()
				// create agent
				name := "sinohope4"
				btcReceivingAddress := "3C7VPws9fMW3kcwRJvMkSVdqMs4SAhQCqq"
				ethAddr := "0x6508d68f4e5931f93fadc3b7afac5092e195b80f"
				description := "lorenzo"
				url := "https://lorenzo-protocol.io"
				agentId := suite.lorenzoApp.AgentKeeper.AddAgent(
					suite.ctx,
					name, btcReceivingAddress, ethAddr, description, url)
				suite.Require().NotEqual(agentId, 0)
				suite.Require().Equal(agentId, uint64(1))

				// deploy yat contract
				yatAddr, err := suite.lorenzoApp.PlanKeeper.DeployYATContract(
					suite.ctx, "lorenzo", "ALRZ")
				suite.Require().NoError(err)

				// create plan
				planReq := types.Plan{
					Name:               "lorenzo-stake-plan",
					PlanDescUri:        "https://lorenzo-protocol.io/lorenzo-stake-plan",
					AgentId:            uint64(1),
					PlanStartTime:      uint64(time.Now().UTC().Unix()) + 1000,
					PeriodTime:         1000,
					YatContractAddress: yatAddr.Hex(),
				}

				planResult, err := suite.lorenzoApp.PlanKeeper.AddPlan(suite.ctx, planReq)
				suite.Require().NoError(err)
				request.ContractAddress = yatAddr.Hex()
				request.Minter = planResult.ContractAddress
			},
			expectErr: false,
			validation: func(request *types.MsgSetMinter) {
				yatAddress := common.HexToAddress(request.ContractAddress)
				minterAddress := common.HexToAddress(request.Minter)
				found, err := suite.lorenzoApp.PlanKeeper.HasRoleFromYAT(
					suite.ctx,
					yatAddress,
					"YAT_MINTER_ROLE",
					minterAddress,
				)
				suite.Require().NoError(err)
				suite.Require().True(found)
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("MsgSetMinter - %s", tc.name), func() {
			suite.SetupTest()
			if tc.malleate != nil {
				tc.malleate(tc.request)
			}
			_, err := suite.msgServer.SetMinter(suite.ctx, tc.request)
			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
			if tc.validation != nil {
				tc.validation(tc.request)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestRemoveMinter() {
	testCases := []struct {
		name       string
		request    *types.MsgRemoveMinter
		malleate   func(*types.MsgRemoveMinter)
		validation func(*types.MsgRemoveMinter)
		expectErr  bool
	}{
		{
			name:      "fail - invalid sender",
			request:   &types.MsgRemoveMinter{Sender: "foobar"},
			expectErr: true,
		},
		{
			name:      "fail - sender not authorized",
			request:   &types.MsgRemoveMinter{Sender: "lrz1tffj9qp3wpdnuds443c86wffrac4jkapkjmmcy"},
			expectErr: true,
		},
		{
			name: "fail - invalid minter address",
			request: &types.MsgRemoveMinter{
				Sender:          testAdmin.String(),
				Minter:          "0x123456",
				ContractAddress: "0xbCC0CdF7683120a1965A343245FA602314C13b9A",
			},
			expectErr: true,
		},
		{
			name: "fail - invalid yat contract address",
			request: &types.MsgRemoveMinter{
				Sender:          testAdmin.String(),
				Minter:          types.ModuleAddress.Hex(),
				ContractAddress: "0x123456",
			},
			expectErr: true,
		},
		{
			name: "fail - yat not exists",
			request: &types.MsgRemoveMinter{
				Sender:          testAdmin.String(),
				ContractAddress: "0xbCC0CdF7683120a1965A343245FA602314C13b9A",
			},
			expectErr: true,
		},
		{
			name: "fail - minter not exists",
			request: &types.MsgRemoveMinter{
				Sender: testAdmin.String(),
				Minter: "0xbCC0CdF7683120a1965A343245FA602314C13b9A",
			},
			expectErr: true,
			malleate: func(request *types.MsgRemoveMinter) {
				yatAddr, err := suite.lorenzoApp.PlanKeeper.DeployYATContract(
					suite.ctx, "lorenzo", "ALRZ")
				suite.Require().NoError(err)
				request.ContractAddress = yatAddr.Hex()
			},
		},
		{
			name: "success - valid remove minter",
			request: &types.MsgRemoveMinter{
				Sender: testAdmin.String(),
			},
			malleate: func(request *types.MsgRemoveMinter) {
				suite.Commit()
				// create agent
				name := "sinohope4"
				btcReceivingAddress := "3C7VPws9fMW3kcwRJvMkSVdqMs4SAhQCqq"
				ethAddr := "0x6508d68f4e5931f93fadc3b7afac5092e195b80f"
				description := "lorenzo"
				url := "https://lorenzo-protocol.io"
				agentId := suite.lorenzoApp.AgentKeeper.AddAgent(
					suite.ctx,
					name, btcReceivingAddress, ethAddr, description, url)
				suite.Require().NotEqual(agentId, 0)
				suite.Require().Equal(agentId, uint64(1))

				// deploy yat contract
				yatAddr, err := suite.lorenzoApp.PlanKeeper.DeployYATContract(
					suite.ctx, "lorenzo", "ALRZ")
				suite.Require().NoError(err)

				// create plan
				planReq := types.Plan{
					Name:               "lorenzo-stake-plan",
					PlanDescUri:        "https://lorenzo-protocol.io/lorenzo-stake-plan",
					AgentId:            uint64(1),
					PlanStartTime:      uint64(time.Now().UTC().Unix()) + 1000,
					PeriodTime:         1000,
					YatContractAddress: yatAddr.Hex(),
				}

				planResult, err := suite.lorenzoApp.PlanKeeper.AddPlan(suite.ctx, planReq)
				suite.Require().NoError(err)
				// set minter
				err = suite.lorenzoApp.PlanKeeper.SetMinter(
					suite.ctx,
					yatAddr,
					common.HexToAddress(planResult.ContractAddress),
				)
				suite.Require().NoError(err)
				request.ContractAddress = yatAddr.Hex()
				request.Minter = planResult.ContractAddress
			},
			expectErr: false,
			validation: func(request *types.MsgRemoveMinter) {
				yatAddress := common.HexToAddress(request.ContractAddress)
				minterAddress := common.HexToAddress(request.Minter)
				found, err := suite.lorenzoApp.PlanKeeper.HasRoleFromYAT(
					suite.ctx,
					yatAddress,
					"YAT_MINTER_ROLE",
					minterAddress,
				)
				suite.Require().NoError(err)
				suite.Require().False(found)
			},
		},
	}
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("MsgRemoveMinter - %s", tc.name), func() {
			suite.SetupTest()
			if tc.malleate != nil {
				tc.malleate(tc.request)
			}
			_, err := suite.msgServer.RemoveMinter(suite.ctx, tc.request)
			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
			if tc.validation != nil {
				tc.validation(tc.request)
			}
		})
	}
}
