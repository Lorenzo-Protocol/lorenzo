package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/Lorenzo-Protocol/lorenzo/v3/x/btcstaking/types"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/ethereum/go-ethereum/common"
)

type msgServer struct {
	k *Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper *Keeper) types.MsgServer {
	return &msgServer{k: keeper}
}

var _ types.MsgServer = msgServer{}

func (ms msgServer) CreateBTCStaking(goCtx context.Context, msg *types.MsgCreateBTCStaking) (*types.MsgCreateBTCStakingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, err
	}

	stakingMsgTx, err := NewBTCTxFromBytes(msg.StakingTx.Transaction)
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrParseBTCTx, "failed to parse btc tx: %v", err)
	}

	// Check if the staking transaction already exists
	stakingTxHash := stakingMsgTx.TxHash()
	if ms.k.GetBTCStakingRecord(ctx, stakingTxHash) != nil {
		return nil, types.ErrDupBTCTx
	}

	stakingTxHeader := ms.k.btclcKeeper.GetHeaderByHash(ctx, msg.StakingTx.Key.Hash)
	if stakingTxHeader == nil {
		return nil, types.ErrBlkHdrNotFound
	}

	params := ms.k.GetParams(ctx)

	btcTip := ms.k.btclcKeeper.GetTipInfo(ctx)
	stakingTxDepth := btcTip.Height - stakingTxHeader.Height

	// Verify the staking transaction
	btcLightClientParams := ms.k.btclcKeeper.GetBTCNet()
	if err := msg.StakingTx.VerifyInclusion(stakingTxHeader.Header, btcLightClientParams.PowLimit); err != nil {
		return nil, errorsmod.Wrapf(types.ErrBTCTxNotIncluded, "failed to verify inclusion: %v", err)
	}

	// Check if the agent exists
	agent, foundAgent := ms.k.agentKeeper.GetAgent(ctx, msg.AgentId)
	if !foundAgent {
		return nil, errorsmod.Wrapf(types.ErrInvalidReceivingAddr, "agent not found, id: %d", msg.AgentId)
	}

	// Check if the receiving address is valid
	btcReceivingAddr, err := btcutil.DecodeAddress(agent.BtcReceivingAddress, btcLightClientParams)
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrInvalidReceivingAddr, "failed to decode btc receiving address: %v", err)
	}

	var (
		mintToAddr, receiverAddr []byte
		btcAmount                uint64
		planId                   = uint64(0)
		lrzChainId               = uint32(ms.k.evmKeeper.ChainID().Uint64())
		chainId                  = lrzChainId
	)

	// parse opReturnMsg
	if common.IsHexAddress(agent.EthAddr) {
		// Check if the sender is authorized
		// when the agent's eth address is a valid hex address

		if !ms.k.Authorized(ctx, sender) {
			return nil, errorsmod.Wrapf(types.ErrNotInAllowList, "unauthorized")
		}
		mintToAddr = common.HexToAddress(agent.EthAddr).Bytes()
		receiverAddr = mintToAddr
		btcAmount, err = ExtractPaymentTo(stakingMsgTx, btcReceivingAddr)
		if err != nil {
			return nil, errorsmod.Wrapf(types.ErrInvalidTransaction, "failed to extract payment: %v", err)
		}
		if btcAmount == 0 {
			return nil, errorsmod.Wrapf(types.ErrMintAmount, "invalid mint amount: %d", btcAmount)
		}
	} else {
		var opReturnMsg []byte
		btcAmount, opReturnMsg, err = ExtractPaymentToWithOpReturnIdAndDust(stakingMsgTx, btcReceivingAddr, params.TxoutDustAmount)
		if err != nil {
			return nil, errorsmod.Wrapf(types.ErrInvalidTransaction, "failed to extract payment: %v", err)
		}
		if btcAmount == 0 {
			return nil, errorsmod.Wrapf(types.ErrMintAmount, "invalid mint amount: %d", btcAmount)
		}

		// Check if OpReturn length is valid
		if !opReturnMsgLenCheck(opReturnMsg) {
			return nil, errorsmod.Wrapf(types.ErrOpReturnLength, "invalid opReturnMsg length: %d", len(opReturnMsg))
		}

		receiverAddr = opReturnMsgGetEthAddr(opReturnMsg)

		// Check if OpReturn contains ChainID
		if opReturnMsgContainsChainId(opReturnMsg) {
			chainId = opReturnMsgGetChainId(opReturnMsg)
		}

		// mint to bridge address if chainId is not lrzChainId
		if chainId != lrzChainId {
			mintToAddr = common.HexToAddress(params.BridgeAddr).Bytes()
		} else {
			mintToAddr = receiverAddr
		}

		// Check if OpReturn contains PlanID
		if opReturnMsgContainsPlanId(opReturnMsg) {
			planId = opReturnMsgGetPlanId(opReturnMsg)
		}
	}

	if err := CheckBTCTxDepth(stakingTxDepth, btcAmount); err != nil {
		return nil, err
	}

	stakingRecord := &types.BTCStakingRecord{
		TxHash:       stakingTxHash[:],
		Amount:       btcAmount,
		ReceiverAddr: receiverAddr,
		AgentName:    agent.Name,
		AgentBtcAddr: agent.BtcReceivingAddress,
		ChainId:      chainId,
	}

	// mint stBTC to mintToAddr and record the staking
	if err := ms.k.Delegate(ctx,
		stakingRecord, mintToAddr, receiverAddr, btcAmount, planId, msg.AgentId); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitTypedEvent(types.NewEventBTCStakingCreated(stakingRecord)) //nolint:errcheck,gosec

	return &types.MsgCreateBTCStakingResponse{}, nil
}

func (ms msgServer) Burn(goCtx context.Context, msg *types.MsgBurnRequest) (*types.MsgBurnResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sender, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, err
	}

	btcNetworkParams := ms.k.btclcKeeper.GetBTCNet()
	btcTargetAddress, err := btcutil.DecodeAddress(msg.BtcTargetAddress, btcNetworkParams)
	if err != nil {
		return nil, errorsmod.Wrapf(
			types.ErrInvalidBurnBtcTargetAddress, "failed to decode btc target address: %v", err)
	}

	amount := sdk.NewCoin(types.NativeTokenDenom, msg.Amount)

	if err := ms.k.Undelegate(ctx, sender, amount); err != nil {
		return nil, err
	}

	err = ctx.EventManager().EmitTypedEvent(types.NewEventBurnCreated(sender, btcTargetAddress, amount))
	if err != nil {
		return nil, types.ErrEmitEvent.Wrap(err.Error())
	}

	return &types.MsgBurnResponse{}, nil
}

// CreateBTCStakingFromBNB implements types.MsgServer.
func (ms msgServer) CreateBTCBStaking(goctx context.Context, req *types.MsgCreateBTCBStaking) (*types.MsgCreateBTCBStakingResponse, error) {
	depositor, err := sdk.AccAddressFromBech32(req.Signer)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goctx)
	if err = ms.k.DepositBTCB(ctx, depositor, req.Number, req.Receipt, req.Proof); err != nil {
		return nil, err
	}
	return &types.MsgCreateBTCBStakingResponse{}, nil
}

// CreatexBTCStaking implements types.MsgServer.
func (ms msgServer) CreatexBTCStaking(goctx context.Context, req *types.MsgCreatexBTCStaking) (*types.MsgCreatexBTCStakingResponse, error) {
	depositor, err := sdk.AccAddressFromBech32(req.Signer)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goctx)
	if err = ms.k.DepositxBTC(ctx, depositor, req.ChainId, req.Number, req.Receipt, req.Proof); err != nil {
		return nil, err
	}
	return &types.MsgCreatexBTCStakingResponse{}, nil
}

func (ms msgServer) UpdateParams(goCtx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if ms.k.authority != msg.Authority {
		return nil, errorsmod.Wrapf(
			govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", ms.k.authority, msg.Authority)
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := msg.Params.Validate(); err != nil {
		return nil, err
	}
	if err := ms.k.SetParams(ctx, &msg.Params); err != nil {
		return nil, err
	}
	return &types.MsgUpdateParamsResponse{}, nil
}

func (ms msgServer) AddReceiver(goCtx context.Context, req *types.MsgAddReceiver) (*types.MsgAddReceiverResponse, error) {
	return nil, fmt.Errorf("deprecated, use UpdateParams instead")
}

func (ms msgServer) RemoveReceiver(goCtx context.Context, req *types.MsgRemoveReceiver) (*types.MsgRemoveReceiverResponse, error) {
	return nil, fmt.Errorf("deprecated, use UpdateParams instead")
}
