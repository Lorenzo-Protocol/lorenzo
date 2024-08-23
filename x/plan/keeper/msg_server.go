package keeper

import (
	"context"

	"github.com/ethereum/go-ethereum/common"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/Lorenzo-Protocol/lorenzo/v3/x/plan/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	// This should be a reference to Keeper
	k *Keeper
}

func (m msgServer) UpdateParams(goCtx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if m.k.authority != msg.Authority {
		return nil, errorsmod.Wrapf(
			govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", m.k.authority, msg.Authority)
	}

	if err := msg.Params.Validate(); err != nil {
		return nil, govtypes.ErrInvalidProposalMsg.Wrapf("invalid parameter: %v", err)
	}

	sdkCtx := sdk.UnwrapSDKContext(goCtx)

	if err := m.k.SetParams(sdkCtx, msg.Params); err != nil {
		return nil, err
	}

	return &types.MsgUpdateParamsResponse{}, nil
}

func (m msgServer) UpgradePlan(goCtx context.Context, msg *types.MsgUpgradePlan) (*types.MsgUpgradePlanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return nil, err
	}

	if m.k.authority != msg.Authority && !m.k.Authorized(ctx, sender) {
		return nil, errorsmod.Wrapf(
			sdkerrors.ErrUnauthorized,
			"unauthorized",
		)
	}
	// Check if the implementation address is a valid address
	if !common.IsHexAddress(msg.Implementation) {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid implementation address")
	}
	oldImplementation, err := m.k.GetPlanImplementationFromBeacon(ctx)
	if err != nil {
		return nil, err
	}
	implementation := common.HexToAddress(msg.Implementation)
	if err := m.k.UpgradeBeaconForPlan(ctx, implementation); err != nil {
		return nil, err
	}
	ctx.EventManager().EmitEvent(
		types.NewUpgradePlanEvent(sender, oldImplementation.Hex(), msg.Implementation),
	)
	return &types.MsgUpgradePlanResponse{}, nil
}

func (m msgServer) CreatePlan(goCtx context.Context, msg *types.MsgCreatePlan) (*types.MsgCreatePlanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}
	if !m.k.Authorized(ctx, sender) {
		return nil, errorsmod.Wrapf(types.ErrUnauthorized, "unauthorized")
	}
	// Check if the plan's start time greater than current block time
	if int64(msg.PlanStartTime) < ctx.BlockTime().Unix() {
		return nil, errorsmod.Wrapf(
			types.ErrInvalidPlanStartTime,
			"plan start time should be greater than current block time")
	}
	plan := types.Plan{
		Name:               msg.Name,
		PlanDescUri:        msg.PlanDescUri,
		AgentId:            msg.AgentId,
		PlanStartTime:      msg.PlanStartTime,
		PeriodTime:         msg.PeriodTime,
		YatContractAddress: msg.YatContractAddress,
	}
	planResult, err := m.k.AddPlan(ctx, plan)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		types.NewCreatePlanEvent(sender, planResult),
	)
	return &types.MsgCreatePlanResponse{Id: planResult.Id}, nil
}

func (m msgServer) SetMerkleRoot(goCtx context.Context, msg *types.MsgSetMerkleRoot) (*types.MsgSetMerkleRootResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}
	if !m.k.Authorized(ctx, sender) {
		return nil, errorsmod.Wrapf(types.ErrUnauthorized, "unauthorized")
	}

	plan, found := m.k.GetPlan(ctx, msg.PlanId)
	if !found {
		return nil, errorsmod.Wrapf(types.ErrPlanNotFound, "plan not found")
	}

	contractAddr := common.HexToAddress(plan.ContractAddress)

	if err := m.k.SetMerkleRoot(
		ctx,
		contractAddr,
		msg.RoundId.BigInt(),
		msg.MerkleRoot,
	); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		types.NewSetMerkleRootEvent(sender, msg.PlanId, msg.MerkleRoot),
	)

	return &types.MsgSetMerkleRootResponse{}, nil
}

func (m msgServer) Claims(goCtx context.Context, msg *types.MsgClaims) (*types.MsgClaimsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	plan, found := m.k.GetPlan(ctx, msg.PlanId)
	if !found {
		return nil, errorsmod.Wrapf(types.ErrPlanNotFound, "plan not found")
	}

	// check if the plan is disabled
	if plan.Enabled == types.PlanStatus_Pause {
		return nil, errorsmod.Wrapf(types.ErrPlanPaused, "plan is paused")
	}

	if err := m.k.Withdraw(ctx,
		msg.PlanId,
		msg.Receiver,
		msg.RoundId.BigInt(),
		msg.Amount.BigInt(),
		msg.MerkleProof,
	); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		types.NewClaimsEvent(
			sender,
			msg.PlanId,
			msg.Receiver,
			msg.RoundId.String(),
			msg.Amount.String(),
			msg.MerkleProof,
		),
	)

	return &types.MsgClaimsResponse{}, nil
}

func (m msgServer) CreateYAT(goCtx context.Context, msg *types.MsgCreateYAT) (*types.MsgCreateYATResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}
	if !m.k.Authorized(ctx, sender) {
		return nil, errorsmod.Wrapf(types.ErrUnauthorized, "unauthorized")
	}
	yatContract, err := m.k.DeployYATContract(ctx, msg.Name, msg.Symbol)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		types.NewCreateYATEvent(sender, yatContract.Hex(), msg.Name, msg.Symbol),
	)

	return &types.MsgCreateYATResponse{ContractAddress: yatContract.Hex()}, nil
}

func (m msgServer) UpdatePlanStatus(goCtx context.Context, msg *types.MsgUpdatePlanStatus) (*types.MsgUpdatePlanStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}
	if !m.k.Authorized(ctx, sender) {
		return nil, errorsmod.Wrapf(types.ErrUnauthorized, "unauthorized")
	}

	plan, found := m.k.GetPlan(ctx, msg.PlanId)
	if !found {
		return nil, errorsmod.Wrapf(types.ErrPlanNotFound, "plan not found")
	}
	if plan.Enabled == msg.Status {
		return nil, errorsmod.Wrapf(types.ErrInvalidPlanStatus, "plan already %s", msg.Status)
	}

	if err := m.k.UpdatePlanStatus(ctx, msg.PlanId, msg.Status); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		types.NewUpdatePlanStatusEvent(sender, msg.PlanId, plan.Enabled, msg.Status),
	)

	return &types.MsgUpdatePlanStatusResponse{}, nil
}

func (m msgServer) SetMinter(goCtx context.Context, msg *types.MsgSetMinter) (*types.MsgSetMinterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}
	if !m.k.Authorized(ctx, sender) {
		return nil, errorsmod.Wrapf(types.ErrUnauthorized, "unauthorized")
	}

	if err := m.k.UpdateMinter(ctx, msg.ContractAddress, msg.Minter, UpdateMinterTypeAdd); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		types.NewSetMinterEvent(sender, msg.Minter, msg.ContractAddress),
	)

	return &types.MsgSetMinterResponse{}, nil
}

func (m msgServer) RemoveMinter(goCtx context.Context, msg *types.MsgRemoveMinter) (*types.MsgRemoveMinterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}
	if !m.k.Authorized(ctx, sender) {
		return nil, errorsmod.Wrapf(types.ErrUnauthorized, "unauthorized")
	}

	if err := m.k.UpdateMinter(ctx, msg.ContractAddress, msg.Minter, UpdateMinterTypeRemove); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		types.NewRemoveMinterEvent(sender, msg.Minter, msg.ContractAddress),
	)

	return &types.MsgRemoveMinterResponse{}, nil
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper *Keeper) types.MsgServer {
	return &msgServer{keeper}
}
