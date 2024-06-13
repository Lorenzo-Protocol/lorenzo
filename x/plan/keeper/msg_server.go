package keeper

import (
	"context"

	"github.com/ethereum/go-ethereum/common"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/Lorenzo-Protocol/lorenzo/x/plan/types"
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

func (m msgServer) UpgradeYAT(goCtx context.Context, msg *types.MsgUpgradeYAT) (*types.MsgUpgradeYATResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if m.k.authority != msg.Authority {
		return nil, errorsmod.Wrapf(
			sdkerrors.ErrUnauthorized,
			"invalid authority; expected %s, got %s",
			m.k.authority,
			msg.Authority,
		)
	}
	implementation := common.HexToAddress(msg.Implementation)
	if err := m.k.UpgradeYAT(ctx, implementation); err != nil {
		return nil, err
	}

	return &types.MsgUpgradeYATResponse{}, nil
}

func (m msgServer) CreatePlan(goCtx context.Context, msg *types.MsgCreatePlan) (*types.MsgCreatePlanResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)
	sender := sdk.AccAddress(msg.Sender)
	if !m.k.Authorized(ctx, sender) {
		return nil, errorsmod.Wrapf(types.ErrUnauthorized, "unauthorized")
	}
	plan := types.Plan{
		Name:                  msg.Name,
		Symbol:                msg.Symbol,
		PlanDescUri:           msg.PlanDescUri,
		AgentId:               msg.AgentId,
		SubscriptionStartTime: msg.SubscriptionStartTime,
		SubscriptionEndTime:   msg.SubscriptionEndTime,
		EndTime:               msg.EndTime,
		MerkleRoot:            msg.MerkleRoot,
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

func (m msgServer) Claims(goCtx context.Context, msg *types.MsgClaims) (*types.MsgClaimsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sender := sdk.AccAddress(msg.Sender)
	if !m.k.Authorized(ctx, sender) {
		return nil, errorsmod.Wrapf(types.ErrUnauthorized, "unauthorized")
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	return &types.MsgClaimsResponse{}, nil
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper *Keeper) types.MsgServer {
	return &msgServer{keeper}
}
