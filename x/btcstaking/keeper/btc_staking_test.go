package keeper_test

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"

	plantypes "github.com/Lorenzo-Protocol/lorenzo/v3/x/plan/types"

	"github.com/btcsuite/btcd/chaincfg/chainhash"

	sdkmath "cosmossdk.io/math"
	"github.com/Lorenzo-Protocol/lorenzo/v3/x/btcstaking/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Lorenzo-Protocol/lorenzo/v3/x/btcstaking/types"
)

func (suite *KeeperTestSuite) TestDelegate() {
	type tmpRequest struct {
		btcStakingRecord *types.BTCStakingRecord
		mintAddr         sdk.AccAddress
		receiverAddr     sdk.AccAddress
		btcAmount        uint64
		planId           uint64
		agentId          uint64
	}
	testCases := []struct {
		name       string
		request    *tmpRequest
		malleate   func(request *tmpRequest)
		validation func(request *tmpRequest)
		expectErr  bool
	}{
		{
			name: "fail - plan not found",
			request: &tmpRequest{
				btcStakingRecord: &types.BTCStakingRecord{
					Amount:       1e7,
					ReceiverAddr: testAdmin,
					AgentName:    "lorenzo",
					AgentBtcAddr: "tb1p97g0dpmsm2fxkmkw9w7mpasmxprsye3k0v49qknwmclwxj78rfjqu6nacq",
					ChainId:      8329,
				},
				mintAddr:     testAdmin,
				receiverAddr: testAdmin,
				btcAmount:    1e7,
				planId:       0,
				agentId:      1,
			},
			malleate: func(request *tmpRequest) {
				txHash, err := hex.DecodeString("84b6addf9aca33604ab20453b71a5da8ab2bd9b773a12a9aecb9944a1123639a")
				suite.Require().NoError(err)
				request.btcStakingRecord.TxHash = txHash
			},
			validation: func(request *tmpRequest) {
				balance := suite.lorenzoApp.BankKeeper.GetBalance(
					suite.ctx, request.receiverAddr, types.NativeTokenDenom)
				toMintAmount := sdkmath.NewIntFromUint64(request.btcAmount).Mul(sdkmath.NewIntFromUint64(keeper.SatoshiToStBTCMul))
				suite.Require().Equal(balance.Amount, toMintAmount)

				// check btc staking record

				txHash, err := chainhash.NewHash(request.btcStakingRecord.TxHash)
				suite.Require().NoError(err)
				btcStakingRecord := suite.keeper.GetBTCStakingRecord(suite.ctx, *txHash)
				suite.Require().NotNil(btcStakingRecord)
				suite.Require().Equal(btcStakingRecord.MintYatResult, "")
			},
			expectErr: false,
		},
		{
			name: "fail - agentID not match",
			request: &tmpRequest{
				btcStakingRecord: &types.BTCStakingRecord{
					Amount:       1e7,
					ReceiverAddr: testAdmin,
					AgentName:    "lorenzo",
					AgentBtcAddr: "tb1p97g0dpmsm2fxkmkw9w7mpasmxprsye3k0v49qknwmclwxj78rfjqu6nacq",
					ChainId:      8329,
				},
				mintAddr:     testAdmin,
				receiverAddr: testAdmin,
				btcAmount:    1e7,
				planId:       1,
				agentId:      2,
			},
			malleate: func(request *tmpRequest) {
				suite.Commit()
				txHash, err := hex.DecodeString("84b6addf9aca33604ab20453b71a5da8ab2bd9b773a12a9aecb9944a1123639a")
				suite.Require().NoError(err)
				request.btcStakingRecord.TxHash = txHash

				// create agent
				name := "lorenzo"
				btcReceivingAddress := "tb1p97g0dpmsm2fxkmkw9w7mpasmxprsye3k0v49qknwmclwxj78rfjqu6nacq"
				ethAddr := ""
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
				planReq := plantypes.Plan{
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
			validation: func(request *tmpRequest) {
				balance := suite.lorenzoApp.BankKeeper.GetBalance(
					suite.ctx, request.receiverAddr, types.NativeTokenDenom)
				toMintAmount := sdkmath.NewIntFromUint64(request.btcAmount).Mul(sdkmath.NewIntFromUint64(keeper.SatoshiToStBTCMul))
				suite.Require().Equal(balance.Amount, toMintAmount)

				// check btc staking record

				txHash, err := chainhash.NewHash(request.btcStakingRecord.TxHash)
				suite.Require().NoError(err)
				btcStakingRecord := suite.keeper.GetBTCStakingRecord(suite.ctx, *txHash)
				suite.Require().NotNil(btcStakingRecord)
				suite.Require().Equal(btcStakingRecord.MintYatResult, keeper.AgentIdNotMatch)
			},
			expectErr: false,
		},
		{
			name: "fail - mint yat error",
			request: &tmpRequest{
				btcStakingRecord: &types.BTCStakingRecord{
					Amount:       1e7,
					ReceiverAddr: testAdmin,
					AgentName:    "lorenzo",
					AgentBtcAddr: "tb1p97g0dpmsm2fxkmkw9w7mpasmxprsye3k0v49qknwmclwxj78rfjqu6nacq",
					ChainId:      8329,
				},
				mintAddr:     testAdmin,
				receiverAddr: testAdmin,
				btcAmount:    1e7,
				planId:       1,
				agentId:      1,
			},
			malleate: func(request *tmpRequest) {
				suite.Commit()
				txHash, err := hex.DecodeString("84b6addf9aca33604ab20453b71a5da8ab2bd9b773a12a9aecb9944a1123639a")
				suite.Require().NoError(err)
				request.btcStakingRecord.TxHash = txHash

				// create agent
				name := "lorenzo"
				btcReceivingAddress := "tb1p97g0dpmsm2fxkmkw9w7mpasmxprsye3k0v49qknwmclwxj78rfjqu6nacq"
				ethAddr := ""
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
				planReq := plantypes.Plan{
					Name:               "lorenzo-stake-plan",
					PlanDescUri:        "https://lorenzo-protocol.io/lorenzo-stake-plan",
					AgentId:            uint64(1),
					PlanStartTime:      uint64(time.Now().UTC().Unix()) + 100000,
					PeriodTime:         1000,
					YatContractAddress: yatAddr.Hex(),
				}

				_, err = suite.lorenzoApp.PlanKeeper.AddPlan(suite.ctx, planReq)
				suite.Require().NoError(err)
			},
			validation: func(request *tmpRequest) {
				balance := suite.lorenzoApp.BankKeeper.GetBalance(
					suite.ctx, request.receiverAddr, types.NativeTokenDenom)
				toMintAmount := sdkmath.NewIntFromUint64(request.btcAmount).Mul(sdkmath.NewIntFromUint64(keeper.SatoshiToStBTCMul))
				suite.Require().Equal(balance.Amount, toMintAmount)
				plan, planFound := suite.lorenzoApp.PlanKeeper.GetPlan(suite.ctx, request.planId)
				suite.Require().True(planFound)
				yatContractAddress := common.HexToAddress(plan.YatContractAddress)
				receiverAddr := common.BytesToAddress(request.receiverAddr)
				yatAmount, err := suite.lorenzoApp.PlanKeeper.BalanceOfFromYAT(
					suite.ctx,
					yatContractAddress,
					receiverAddr,
				)
				suite.Require().NoError(err)
				suite.Require().Equal(yatAmount.Uint64(), uint64(0))

				// check btc staking record
				txHash, err := chainhash.NewHash(request.btcStakingRecord.TxHash)
				suite.Require().NoError(err)
				btcStakingRecord := suite.keeper.GetBTCStakingRecord(suite.ctx, *txHash)
				suite.Require().NotNil(btcStakingRecord)
				suite.Require().Contains(btcStakingRecord.MintYatResult, "execution reverted")
			},
			expectErr: false,
		},
		{
			name: "success",
			request: &tmpRequest{
				btcStakingRecord: &types.BTCStakingRecord{
					Amount:       1e7,
					ReceiverAddr: testAdmin,
					AgentName:    "lorenzo",
					AgentBtcAddr: "tb1p97g0dpmsm2fxkmkw9w7mpasmxprsye3k0v49qknwmclwxj78rfjqu6nacq",
					ChainId:      8329,
				},
				mintAddr:     testAdmin,
				receiverAddr: testAdmin,
				btcAmount:    1e7,
				planId:       1,
				agentId:      1,
			},
			malleate: func(request *tmpRequest) {
				suite.Commit()
				txHash, err := hex.DecodeString("84b6addf9aca33604ab20453b71a5da8ab2bd9b773a12a9aecb9944a1123639a")
				suite.Require().NoError(err)
				request.btcStakingRecord.TxHash = txHash

				// create agent
				name := "lorenzo"
				btcReceivingAddress := "tb1p97g0dpmsm2fxkmkw9w7mpasmxprsye3k0v49qknwmclwxj78rfjqu6nacq"
				ethAddr := ""
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
				planReq := plantypes.Plan{
					Name:               "lorenzo-stake-plan",
					PlanDescUri:        "https://lorenzo-protocol.io/lorenzo-stake-plan",
					AgentId:            uint64(1),
					PlanStartTime:      uint64(time.Now().UTC().Unix()) - 100000,
					PeriodTime:         1000,
					YatContractAddress: yatAddr.Hex(),
				}

				_, err = suite.lorenzoApp.PlanKeeper.AddPlan(suite.ctx, planReq)
				suite.Require().NoError(err)
			},
			validation: func(request *tmpRequest) {
				balance := suite.lorenzoApp.BankKeeper.GetBalance(
					suite.ctx, request.receiverAddr, types.NativeTokenDenom)
				toMintAmount := sdkmath.NewIntFromUint64(request.btcAmount).Mul(sdkmath.NewIntFromUint64(keeper.SatoshiToStBTCMul))
				suite.Require().Equal(balance.Amount, toMintAmount)
				plan, planFound := suite.lorenzoApp.PlanKeeper.GetPlan(suite.ctx, request.planId)
				suite.Require().True(planFound)
				yatContractAddress := common.HexToAddress(plan.YatContractAddress)
				receiverAddr := common.BytesToAddress(request.receiverAddr)
				yatAmount, err := suite.lorenzoApp.PlanKeeper.BalanceOfFromYAT(
					suite.ctx,
					yatContractAddress,
					receiverAddr,
				)
				suite.Require().NoError(err)
				suite.Require().Equal(yatAmount, toMintAmount.BigInt())

				// check btc staking record
				txHash, err := chainhash.NewHash(request.btcStakingRecord.TxHash)
				suite.Require().NoError(err)
				btcStakingRecord := suite.keeper.GetBTCStakingRecord(suite.ctx, *txHash)
				suite.Require().NotNil(btcStakingRecord)
				suite.Require().Equal(btcStakingRecord.MintYatResult, keeper.Success)

				events := suite.ctx.EventManager().Events()
				abciEvents := events.ToABCIEvents()
				for _, abciEvent := range abciEvents {
					if abciEvent.Type == "mint_yat" {
						eventAttribute := abciEvent.GetAttributes()

						suite.Require().Equal(eventAttribute[0].Key, "plan_id")
						suite.Require().Equal(eventAttribute[0].Value, fmt.Sprintf("%d", request.planId))
						suite.Require().Equal(eventAttribute[1].Key, "account")
						suite.Require().Equal(eventAttribute[1].Value, common.BytesToAddress(request.receiverAddr).Hex())
						suite.Require().Equal(eventAttribute[2].Key, "amount")
						suite.Require().Equal(eventAttribute[2].Value, toMintAmount.String())

					}
				}

			},
			expectErr: false,
		},
		{
			name: "success - mintAddr not equal to receiverAddr",
			request: &tmpRequest{
				btcStakingRecord: &types.BTCStakingRecord{
					Amount:       1e7,
					ReceiverAddr: testAdmin,
					AgentName:    "lorenzo",
					AgentBtcAddr: "tb1p97g0dpmsm2fxkmkw9w7mpasmxprsye3k0v49qknwmclwxj78rfjqu6nacq",
					ChainId:      8329,
				},
				mintAddr:     testAdmin,
				receiverAddr: sdk.AccAddress("lrz1cpldpp5960ed8s63w4v9fml84w875wv0emcda5"),
				btcAmount:    1e7,
				planId:       1,
				agentId:      1,
			},
			malleate: func(request *tmpRequest) {
				suite.Commit()
				txHash, err := hex.DecodeString("84b6addf9aca33604ab20453b71a5da8ab2bd9b773a12a9aecb9944a1123639a")
				suite.Require().NoError(err)
				request.btcStakingRecord.TxHash = txHash

				// create agent
				name := "lorenzo"
				btcReceivingAddress := "tb1p97g0dpmsm2fxkmkw9w7mpasmxprsye3k0v49qknwmclwxj78rfjqu6nacq"
				ethAddr := ""
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
				planReq := plantypes.Plan{
					Name:               "lorenzo-stake-plan",
					PlanDescUri:        "https://lorenzo-protocol.io/lorenzo-stake-plan",
					AgentId:            uint64(1),
					PlanStartTime:      uint64(time.Now().UTC().Unix()) - 100000,
					PeriodTime:         1000,
					YatContractAddress: yatAddr.Hex(),
				}

				_, err = suite.lorenzoApp.PlanKeeper.AddPlan(suite.ctx, planReq)
				suite.Require().NoError(err)
			},
			validation: func(request *tmpRequest) {
				balance := suite.lorenzoApp.BankKeeper.GetBalance(
					suite.ctx, request.mintAddr, types.NativeTokenDenom)
				toMintAmount := sdkmath.NewIntFromUint64(request.btcAmount).Mul(sdkmath.NewIntFromUint64(keeper.SatoshiToStBTCMul))
				suite.Require().Equal(balance.Amount, toMintAmount)
				plan, planFound := suite.lorenzoApp.PlanKeeper.GetPlan(suite.ctx, request.planId)
				suite.Require().True(planFound)
				yatContractAddress := common.HexToAddress(plan.YatContractAddress)
				receiverAddr := common.BytesToAddress(request.receiverAddr)
				yatAmount, err := suite.lorenzoApp.PlanKeeper.BalanceOfFromYAT(
					suite.ctx,
					yatContractAddress,
					receiverAddr,
				)
				suite.Require().NoError(err)
				suite.Require().Equal(yatAmount, toMintAmount.BigInt())

				// check btc staking record
				txHash, err := chainhash.NewHash(request.btcStakingRecord.TxHash)
				suite.Require().NoError(err)
				btcStakingRecord := suite.keeper.GetBTCStakingRecord(suite.ctx, *txHash)
				suite.Require().NotNil(btcStakingRecord)
				suite.Require().Equal(btcStakingRecord.MintYatResult, keeper.Success)

				events := suite.ctx.EventManager().Events()
				abciEvents := events.ToABCIEvents()
				for _, abciEvent := range abciEvents {
					if abciEvent.Type == "mint_yat" {
						eventAttribute := abciEvent.GetAttributes()

						suite.Require().Equal(eventAttribute[0].Key, "plan_id")
						suite.Require().Equal(eventAttribute[0].Value, fmt.Sprintf("%d", request.planId))
						suite.Require().Equal(eventAttribute[1].Key, "account")
						suite.Require().Equal(eventAttribute[1].Value, common.BytesToAddress(request.receiverAddr).Hex())
						suite.Require().Equal(eventAttribute[2].Key, "amount")
						suite.Require().Equal(eventAttribute[2].Value, toMintAmount.String())

					}
				}

			},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("KeeperDelegate - %s", tc.name), func() {
			suite.SetupTest()
			if tc.malleate != nil {
				tc.malleate(tc.request)
			}
			err := suite.keeper.Delegate(
				suite.ctx,
				tc.request.btcStakingRecord,
				tc.request.mintAddr,
				tc.request.receiverAddr,
				tc.request.btcAmount,
				tc.request.planId,
				tc.request.agentId,
			)
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
