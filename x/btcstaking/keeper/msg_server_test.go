package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkmath "cosmossdk.io/math"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/Lorenzo-Protocol/lorenzo/v3/x/btcstaking/types"
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
			name: "fail -  allow list contains invalid address",
			request: &types.MsgUpdateParams{
				Authority: authtypes.NewModuleAddress(govtypes.ModuleName).String(),
				Params: types.Params{
					Receivers:             nil,
					BtcConfirmationsDepth: 1,
					MinterAllowList: []string{
						"cosmos1qperwt9wrnkg5k9e5gzfgjppzpqqu8t3q4yjx9",
						"lrz1tffj9qp3wpdnuds443c86wffrac4jkapkjmmcy",
					},
					BridgeAddr:      "0xb7C0817Dd23DE89E4204502dd2C2EF7F57d3A3B8",
					TxoutDustAmount: 546,
				},
			},
			expectErr: true,
		},
		{
			name: "fail -  allow list contains duplicate address",
			request: &types.MsgUpdateParams{
				Authority: authtypes.NewModuleAddress(govtypes.ModuleName).String(),
				Params: types.Params{
					Receivers:             nil,
					BtcConfirmationsDepth: 1,
					MinterAllowList: []string{
						"lrz1tffj9qp3wpdnuds443c86wffrac4jkapkjmmcy",
						"lrz1tffj9qp3wpdnuds443c86wffrac4jkapkjmmcy",
					},
					BridgeAddr:      "0xb7C0817Dd23DE89E4204502dd2C2EF7F57d3A3B8",
					TxoutDustAmount: 546,
				},
			},
			expectErr: true,
		},
		{
			name: "fail - bridge address is invalid",
			request: &types.MsgUpdateParams{
				Authority: authtypes.NewModuleAddress(govtypes.ModuleName).String(),
				Params: types.Params{
					Receivers:             nil,
					BtcConfirmationsDepth: 1,
					MinterAllowList: []string{
						"cosmos1qperwt9wrnkg5k9e5gzfgjppzpqqu8t3q4yjx9",
						"lrz1tffj9qp3wpdnuds443c86wffrac4jkapkjmmcy",
					},
					BridgeAddr:      "0x123456",
					TxoutDustAmount: 546,
				},
			},
			expectErr: true,
		},
		{
			name: "pass - valid Update msg",
			request: &types.MsgUpdateParams{
				Authority: authtypes.NewModuleAddress(govtypes.ModuleName).String(),
				Params: types.Params{
					Receivers:             nil,
					BtcConfirmationsDepth: 1,
					MinterAllowList: []string{
						"lrz1tffj9qp3wpdnuds443c86wffrac4jkapkjmmcy",
					},
					BridgeAddr:      "0xb7C0817Dd23DE89E4204502dd2C2EF7F57d3A3B8",
					TxoutDustAmount: 546,
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

func (suite *KeeperTestSuite) TestCreateBTCStaking() {
	testCases := []struct {
		name       string
		request    *types.MsgCreateBTCStaking
		malleate   func(request *types.MsgCreateBTCStaking)
		validation func(request *types.MsgCreateBTCStaking)
		expectErr  bool
	}{
		{
			name: "fail - invalid sender",
			request: &types.MsgCreateBTCStaking{
				Signer: "foobar",
			},
			expectErr: true,
		},
		//{
		//	name: "success - valid create request when eth address is set",
		//	request: &types.MsgCreateBTCStaking{
		//		Signer: testAdmin.String(),
		//	},
		//	malleate: func(request *types.MsgCreateBTCStaking) {
		//		suite.Commit()
		//		// create transaction
		//		txBytes, _ := hex.DecodeString("0200000000010274171503dbc24539663844c6f8e6c290947cdf021ddadcd3fdecc3ee049ea0ee0000000000fdffffff4e9c069460c3386eaac1c88d0c07cf89ee1399bcf1af12c9b39439ba9f44313b0200000000fdffffff034a01000000000000225120b3312a2c36383e101f2fa446eb16f00f1d2edef0eb7b839889e13f4592d80ecf00000000000000004a6a4800000a580000000000000000000000001c779b46ca5ffcf144f53aa09c17a5c3372de33e41494e4e000000000000000000000000000000000000000000000001158e460913d00000a7860000000000002251205aca89ca6c96635892be129cdf0ceed6e7c27c2713ddea812b678343455161b201406305bf2c1a3df628c581642f32e85dffa07b3dac3da4bca6b735c4b6c9dcca1e79959a589df01260464d41f8a7cd451a4f932b3fa50a76396503c0e87e380f530140e9ab6a00fe4a00f01cb61f013331efddb441dc807904f0a46c0f6d8da9eab950d5729db413d5fda9c3f3b0dcec33f8fdef8f6ad6776eb1baf9a08e9d6d0c826f00000000")
		//		proof, _ := hex.DecodeString("00000020657d06debfee161b0de46492a1cc776d6b56ee063c862ada1200000000000000dd80a6d8798a78c6671aa8af826404d2d8b796eef3784a99808d8b85f90ba8ecbcaa0366ffff001d04055dacf90100000a01b7ec9c4c909a9e65859935cc8d542f43b50e215473df6342b0b6373fe943ab932ffcb2edc6484d093a3a14513913a5e7ac4f741488fb9a1021361357c929b1070b29d5f13fe5ae24f115eece7c7f73d8e1a5eb6700781794435c2ef66861b21a27f8dda8f1f4597775c64dadfe30f9e03f65b21c1c483ebefc629b345a59d04e9c069460c3386eaac1c88d0c07cf89ee1399bcf1af12c9b39439ba9f44313b45e627c5780e608bdd72d2614088cb02d72e129889ec9faeb902f9ebae95adf9ef9c0d15192bd3e47cc08efef4f2366b14d3ec8075fcb842d541fbf4418ce0d1b1196114246b0cab7cbcf5633bb0ec3ab1d8a02cadcf1aecc853a8641bf3332fcf1e609853bd5c868240f7535ab6e988c53d3b16ea424960c0a453b9dced7256f1db8093992557847d26df562bf630ea2d6606aab2d44446fa40a7cace6d1246035d5b00")
		//		//tx, _ := types.ParseTransaction(txBytes)
		//		merkleBlk, _ := keeper.ParseMerkleBlock(proof)
		//		txIndex, proofBytes, _ := keeper.ParseBTCProof(merkleBlk)
		//		blkHdr := &merkleBlk.Header
		//
		//		var blkHdrHashBytes lrz.BTCHeaderHashBytes
		//		tmp := blkHdr.BlockHash()
		//		blkHdrHashBytes.FromChainhash(&tmp)
		//		txInfo := types.TransactionInfo{
		//			Key: &types.TransactionKey{
		//				Index: txIndex,
		//				Hash:  &blkHdrHashBytes,
		//			},
		//			Transaction: txBytes,
		//			Proof:       proofBytes,
		//		}
		//		request.StakingTx = &txInfo
		//
		//		// create agent
		//		name := "sinohope4"
		//		btcReceivingAddress := "3C7VPws9fMW3kcwRJvMkSVdqMs4SAhQCqq"
		//		ethAddr := "0x6508d68f4e5931f93fadc3b7afac5092e195b80f"
		//		description := "sinohope"
		//		url := "https://sinohope.io"
		//		agentId := suite.lorenzoApp.AgentKeeper.AddAgent(
		//			suite.ctx,
		//			name, btcReceivingAddress, ethAddr, description, url)
		//		suite.Require().NotEqual(agentId, 0)
		//		suite.Require().Equal(agentId, uint64(1))
		//		yatAddr, err := suite.lorenzoApp.PlanKeeper.DeployYATContract(
		//			suite.ctx, "sinohope", "SHOPE")
		//		suite.Require().NoError(err)
		//		// create plan
		//		planReq := plantypes.Plan{
		//			Name:               "sinohope-stake-plan",
		//			PlanDescUri:        "https://sinohope.io/sinohope-stake-plan",
		//			AgentId:            uint64(1),
		//			PlanStartTime:      uint64(time.Now().UTC().Unix()) + 1000,
		//			PeriodTime:         1000,
		//			YatContractAddress: yatAddr.Hex(),
		//		}
		//
		//		_, err = suite.lorenzoApp.PlanKeeper.AddPlan(suite.ctx, planReq)
		//		suite.Require().NoError(err)
		//
		//		request.AgentId = agentId
		//	},
		//	expectErr: false,
		//},
		//{
		//	name: "success - valid create request when eth address is empty",
		//	request: &types.MsgCreateBTCStaking{
		//		Signer: testAdmin.String(),
		//	},
		//	malleate: func(request *types.MsgCreateBTCStaking) {
		//		suite.Commit()
		//		// create transaction
		//		txBytes, _ := hex.DecodeString("0200000000010274171503dbc24539663844c6f8e6c290947cdf021ddadcd3fdecc3ee049ea0ee0000000000fdffffff4e9c069460c3386eaac1c88d0c07cf89ee1399bcf1af12c9b39439ba9f44313b0200000000fdffffff034a01000000000000225120b3312a2c36383e101f2fa446eb16f00f1d2edef0eb7b839889e13f4592d80ecf00000000000000004a6a4800000a580000000000000000000000001c779b46ca5ffcf144f53aa09c17a5c3372de33e41494e4e000000000000000000000000000000000000000000000001158e460913d00000a7860000000000002251205aca89ca6c96635892be129cdf0ceed6e7c27c2713ddea812b678343455161b201406305bf2c1a3df628c581642f32e85dffa07b3dac3da4bca6b735c4b6c9dcca1e79959a589df01260464d41f8a7cd451a4f932b3fa50a76396503c0e87e380f530140e9ab6a00fe4a00f01cb61f013331efddb441dc807904f0a46c0f6d8da9eab950d5729db413d5fda9c3f3b0dcec33f8fdef8f6ad6776eb1baf9a08e9d6d0c826f00000000")
		//		proof, _ := hex.DecodeString("00000020657d06debfee161b0de46492a1cc776d6b56ee063c862ada1200000000000000dd80a6d8798a78c6671aa8af826404d2d8b796eef3784a99808d8b85f90ba8ecbcaa0366ffff001d04055dacf90100000a01b7ec9c4c909a9e65859935cc8d542f43b50e215473df6342b0b6373fe943ab932ffcb2edc6484d093a3a14513913a5e7ac4f741488fb9a1021361357c929b1070b29d5f13fe5ae24f115eece7c7f73d8e1a5eb6700781794435c2ef66861b21a27f8dda8f1f4597775c64dadfe30f9e03f65b21c1c483ebefc629b345a59d04e9c069460c3386eaac1c88d0c07cf89ee1399bcf1af12c9b39439ba9f44313b45e627c5780e608bdd72d2614088cb02d72e129889ec9faeb902f9ebae95adf9ef9c0d15192bd3e47cc08efef4f2366b14d3ec8075fcb842d541fbf4418ce0d1b1196114246b0cab7cbcf5633bb0ec3ab1d8a02cadcf1aecc853a8641bf3332fcf1e609853bd5c868240f7535ab6e988c53d3b16ea424960c0a453b9dced7256f1db8093992557847d26df562bf630ea2d6606aab2d44446fa40a7cace6d1246035d5b00")
		//		//tx, _ := types.ParseTransaction(txBytes)
		//		merkleBlk, _ := keeper.ParseMerkleBlock(proof)
		//		txIndex, proofBytes, _ := keeper.ParseBTCProof(merkleBlk)
		//		blkHdr := &merkleBlk.Header
		//
		//		var blkHdrHashBytes lrz.BTCHeaderHashBytes
		//		tmp := blkHdr.BlockHash()
		//		blkHdrHashBytes.FromChainhash(&tmp)
		//		txInfo := types.TransactionInfo{
		//			Key: &types.TransactionKey{
		//				Index: txIndex,
		//				Hash:  &blkHdrHashBytes,
		//			},
		//			Transaction: txBytes,
		//			Proof:       proofBytes,
		//		}
		//		request.StakingTx = &txInfo
		//
		//		// create agent
		//		name := "lorenzo"
		//		btcReceivingAddress := "tb1p97g0dpmsm2fxkmkw9w7mpasmxprsye3k0v49qknwmclwxj78rfjqu6nacq"
		//		ethAddr := ""
		//		description := "lorenzo"
		//		url := "https://lorenzo-protocol.io"
		//		agentId := suite.lorenzoApp.AgentKeeper.AddAgent(
		//			suite.ctx,
		//			name, btcReceivingAddress, ethAddr, description, url)
		//		suite.Require().NotEqual(agentId, 0)
		//		suite.Require().Equal(agentId, uint64(1))
		//		yatAddr, err := suite.lorenzoApp.PlanKeeper.DeployYATContract(
		//			suite.ctx, "lorenzo", "ALRZ")
		//		suite.Require().NoError(err)
		//		// create plan
		//		planReq := plantypes.Plan{
		//			Name:               "lorenzo-stake-plan",
		//			PlanDescUri:        "https://lorenzo-protocol.io/lorenzo-stake-plan",
		//			AgentId:            uint64(1),
		//			PlanStartTime:      uint64(time.Now().UTC().Unix()) + 1000,
		//			PeriodTime:         1000,
		//			YatContractAddress: yatAddr.Hex(),
		//		}
		//
		//		_, err = suite.lorenzoApp.PlanKeeper.AddPlan(suite.ctx, planReq)
		//		suite.Require().NoError(err)
		//
		//		request.AgentId = agentId
		//
		//		// set btc header
		//		headerHex := ""
		//		header, err := lrz.NewBTCHeaderBytesFromHex(headerHex)
		//		headers := []lrz.BTCHeaderBytes{header}
		//		err = suite.lorenzoApp.BTCLightClientKeeper.InsertHeaders(suite.ctx, headers)
		//		suite.Require().NoError(err)
		//
		//	},
		//	expectErr: false,
		//},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("CreateBTCStaking - %s", tc.name), func() {
			suite.SetupTest()
			if tc.malleate != nil {
				tc.malleate(tc.request)
			}
			_, err := suite.msgServer.CreateBTCStaking(suite.ctx, tc.request)
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

func (suite *KeeperTestSuite) TestBurn() {
	testCases := []struct {
		name       string
		request    *types.MsgBurnRequest
		malleate   func(request *types.MsgBurnRequest)
		validation func(request *types.MsgBurnRequest)
		expectErr  bool
	}{
		{
			name:      "fail - invalid sender",
			request:   &types.MsgBurnRequest{Signer: "foobar"},
			expectErr: true,
		},
		{
			name: "fail - invalid btc target address",
			request: &types.MsgBurnRequest{
				Signer:           testAdmin.String(),
				BtcTargetAddress: "0xb7C0817Dd23DE89E4204502dd2C2EF7F57d3A3B8",
				Amount:           sdkmath.NewInt(100000000000),
			},
			expectErr: true,
		},
		{
			name: "fail - amount lt balance",
			request: &types.MsgBurnRequest{
				Signer:           testAdmin.String(),
				BtcTargetAddress: "tb1ptt9gnjnvje343y47z2wd7r8w6mnuylp8z0w74qftv7p5x323vxeq9jrn6f",
				Amount:           sdkmath.NewInt(200000000000),
			},
			malleate: func(request *types.MsgBurnRequest) {
				coins := []sdk.Coin{
					{
						Denom:  types.NativeTokenDenom,
						Amount: sdkmath.NewInt(100000000000),
					},
				}
				err := suite.lorenzoApp.BankKeeper.MintCoins(suite.ctx, types.ModuleName, coins)
				suite.Require().NoError(err)

				receiver, err := sdk.AccAddressFromBech32(request.Signer)
				suite.Require().NoError(err)

				err = suite.lorenzoApp.BankKeeper.SendCoinsFromModuleToAccount(
					suite.ctx, types.ModuleName, receiver, coins)
				suite.Require().NoError(err)
			},
			expectErr: true,
		},
		{
			name: "success - valid burn request",
			request: &types.MsgBurnRequest{
				Signer:           testAdmin.String(),
				BtcTargetAddress: "tb1ptt9gnjnvje343y47z2wd7r8w6mnuylp8z0w74qftv7p5x323vxeq9jrn6f",
				Amount:           sdkmath.NewInt(100000000000),
			},
			malleate: func(request *types.MsgBurnRequest) {
				coins := []sdk.Coin{
					{
						Denom:  types.NativeTokenDenom,
						Amount: sdkmath.NewInt(200000000000),
					},
				}
				err := suite.lorenzoApp.BankKeeper.MintCoins(suite.ctx, types.ModuleName, coins)
				suite.Require().NoError(err)

				receiver, err := sdk.AccAddressFromBech32(request.Signer)
				suite.Require().NoError(err)

				err = suite.lorenzoApp.BankKeeper.SendCoinsFromModuleToAccount(
					suite.ctx, types.ModuleName, receiver, coins)
				suite.Require().NoError(err)
			},
			validation: func(request *types.MsgBurnRequest) {
				balance := suite.lorenzoApp.BankKeeper.GetBalance(suite.ctx, testAdmin, types.NativeTokenDenom)
				suite.Require().Equal(sdkmath.NewInt(100000000000), balance.Amount)
			},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Burn - %s", tc.name), func() {
			suite.SetupTest()
			if tc.malleate != nil {
				tc.malleate(tc.request)
			}
			_, err := suite.msgServer.Burn(suite.ctx, tc.request)
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

func (suite *KeeperTestSuite) TestRepairStaking() {
	testCases := []struct {
		name       string
		request    *types.MsgRepairStaking
		malleate   func(request *types.MsgRepairStaking)
		validation func(request *types.MsgRepairStaking)
		expectErr  bool
	}{
		{
			name:      "fail - invalid authority",
			request:   &types.MsgRepairStaking{Authority: "foobar"},
			expectErr: true,
		},
		{
			name: "success - valid repair request",
			request: &types.MsgRepairStaking{
				Authority: authtypes.NewModuleAddress(govtypes.ModuleName).String(),
				ReceiverInfos: []*types.ReceiverInfo{
					{
						Address: "lrz1tffj9qp3wpdnuds443c86wffrac4jkapkjmmcy",
						Amount:  sdkmath.NewInt(100000000000),
					},
				},
			},
			validation: func(request *types.MsgRepairStaking) {
				// check balance
				for _, receiverInfo := range request.ReceiverInfos {
					receiver, err := sdk.AccAddressFromBech32(receiverInfo.Address)
					suite.Require().NoError(err)
					balance := suite.lorenzoApp.BankKeeper.GetBalance(suite.ctx, receiver, types.NativeTokenDenom)
					suite.Require().Equal(receiverInfo.Amount, balance.Amount)
				}

				// check event
				events := suite.ctx.EventManager().Events()
				abciEvents := events.ToABCIEvents()
				for _, abciEvent := range abciEvents {
					if abciEvent.Type == "lorenzo.btcstaking.v1.EventMintStBTC" {
						eventAttribute := abciEvent.GetAttributes()

						suite.Require().Equal(eventAttribute[0].Key, "amount")
						suite.Require().Equal(eventAttribute[1].Key, "cosmos_address")
						suite.Require().Equal(eventAttribute[2].Key, "eth_address")

					}
				}
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("MsgUpdateParams - %s", tc.name), func() {
			suite.SetupTest()
			if tc.malleate != nil {
				tc.malleate(tc.request)
			}
			_, err := suite.msgServer.RepairStaking(suite.ctx, tc.request)
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
