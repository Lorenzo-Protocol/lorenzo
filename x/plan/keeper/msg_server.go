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
	// Check if the implementation address is a valid address
	if !common.IsHexAddress(msg.Implementation) {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid implementation address")
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
		Name:               msg.Name,
		PlanDescUri:        msg.PlanDescUri,
		AgentId:            msg.AgentId,
		PlanStartBlock:     msg.PlanStartBlock,
		PeriodBlocks:       msg.PeriodBlocks,
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

func (m msgServer) Claims(goCtx context.Context, msg *types.MsgClaims) (*types.MsgClaimsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sender := sdk.AccAddress(msg.Sender)
	if !m.k.Authorized(ctx, sender) {
		return nil, errorsmod.Wrapf(types.ErrUnauthorized, "unauthorized")
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
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
	sender := sdk.AccAddress(msg.Sender)
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

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper *Keeper) types.MsgServer {
	return &msgServer{keeper}
}
