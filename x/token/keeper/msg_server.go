package keeper

import (
	"context"

	"github.com/ethereum/go-ethereum/common"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/Lorenzo-Protocol/lorenzo/v3/x/token/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	*Keeper
}

// NewMsgServerImpl creates a new token MsgServer instance
func NewMsgServerImpl(k *Keeper) types.MsgServer {
	return &msgServer{Keeper: k}
}

// RegisterCoin implements MsgServer.RegisterCoin
func (m msgServer) RegisterCoin(goCtx context.Context, msg *types.MsgRegisterCoin) (*types.MsgRegisterCoinResponse, error) {
	if m.authority.String() != msg.Authority {
		return nil, errorsmod.Wrapf(govtypes.ErrInvalidSigner,
			"invalid authority; expected %s, got %s", m.authority.String(), msg.Authority)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if !m.IsConvertEnabled(ctx) {
		return nil, errorsmod.Wrap(
			types.ErrConvertDisabled, "register is not allowed",
		)
	}

	for _, metadata := range msg.Metadata {
		pair, err := m.Keeper.RegisterCoin(ctx, metadata)
		if err != nil {
			return nil, err
		}

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeRegisterCoin,
				sdk.NewAttribute(types.AttributeKeyCosmosCoin, pair.Denom),
				sdk.NewAttribute(types.AttributeKeyERC20Token, pair.ContractAddress),
			),
		)
	}

	return &types.MsgRegisterCoinResponse{}, nil
}

// RegisterERC20 implements MsgServer.RegisterERC20
func (m msgServer) RegisterERC20(goCtx context.Context, msg *types.MsgRegisterERC20) (*types.MsgRegisterERC20Response, error) {
	if m.authority.String() != msg.Authority {
		return nil, errorsmod.Wrapf(govtypes.ErrInvalidSigner,
			"invalid authority; expected %s, got %s", m.authority.String(), msg.Authority)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if !m.IsConvertEnabled(ctx) {
		return nil, errorsmod.Wrap(
			types.ErrConvertDisabled, "register is not allowed",
		)
	}

	for _, address := range msg.ContractAddresses {
		pair, err := m.Keeper.RegisterERC20(ctx, common.HexToAddress(address))
		if err != nil {
			return nil, err
		}

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeRegisterERC20,
				sdk.NewAttribute(types.AttributeKeyCosmosCoin, pair.Denom),
				sdk.NewAttribute(types.AttributeKeyERC20Token, pair.ContractAddress),
			),
		)
	}

	return &types.MsgRegisterERC20Response{}, nil
}

// ToggleConversion implements MsgServer.ToggleConversion
func (m msgServer) ToggleConversion(goCtx context.Context, msg *types.MsgToggleConversion) (*types.MsgToggleConversionResponse, error) {
	if m.authority.String() != msg.Authority {
		return nil, errorsmod.Wrapf(govtypes.ErrInvalidSigner,
			"invalid authority; expected %s, got %s", m.authority.String(), msg.Authority)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if !m.IsConvertEnabled(ctx) {
		return nil, errorsmod.Wrap(
			types.ErrConvertDisabled, "register is not allowed",
		)
	}

	pair, err := m.Keeper.ToggleConversion(ctx, msg.Token)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeToggleTokenConversion,
			sdk.NewAttribute(types.AttributeKeyCosmosCoin, pair.Denom),
			sdk.NewAttribute(types.AttributeKeyERC20Token, pair.ContractAddress),
		),
	)

	return &types.MsgToggleConversionResponse{}, nil
}

// UpdateParams implements MsgServer.UpdateParams
func (m msgServer) UpdateParams(goCtx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if m.authority.String() != msg.Authority {
		return nil, errorsmod.Wrapf(govtypes.ErrInvalidSigner,
			"invalid authority; expected %s, got %s", m.authority.String(), msg.Authority)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	m.SetParams(ctx, msg.Params)

	return &types.MsgUpdateParamsResponse{}, nil
}

// ConvertCoin implements MsgServer.ConvertCoin
// NOTE: impl it as keeper method as we need to export it as expected keeper for lorenzo ibc transfer wrapper.
func (k Keeper) ConvertCoin(goCtx context.Context, msg *types.MsgConvertCoin) (*types.MsgConvertCoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	receiver := common.HexToAddress(msg.Receiver)
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	// check if the token pair is enabled for minting and converting.
	pair, err := k.MintEnabled(ctx, sender, receiver.Bytes(), msg.Coin.Denom)
	if err != nil {
		return nil, err
	}

	// NOTE: check if the contract is self-destructed
	erc20 := common.HexToAddress(pair.ContractAddress)
	acc := k.evmKeeper.GetAccountWithoutBalance(ctx, erc20)

	if acc == nil || !acc.IsContract() {
		k.RemoveTokenPair(ctx, pair)
		k.Logger(ctx).Debug(
			"deleting self-destructed token pair from state",
			"contract", pair.ContractAddress,
		)
		// NOTE: return nil error to persist the changes from the deletion
		return nil, nil
	}

	// execute as per token source type.
	switch {
	case pair.IsNativeCoin():
		return k.ConvertNativeCoinToVoucherERC20(ctx, pair, msg, receiver, sender)
	case pair.IsNativeERC20():
		return k.ConvertVoucherCoinToNativeERC20(ctx, pair, msg, receiver, sender)
	default:
		return nil, types.ErrUndefinedOwner
	}
}

// ConvertERC20 implements MsgServer.ConvertERC20
// NOTE: impl it as keeper method as we need to export it as expected keeper for lorenzo ibc transfer wrapper.
func (k Keeper) ConvertERC20(goCtx context.Context, msg *types.MsgConvertERC20) (*types.MsgConvertERC20Response, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender := common.HexToAddress(msg.Sender)
	receiver, err := sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return nil, err
	}

	// check if the token pair is enabled for minting and converting.
	pair, err := k.MintEnabled(ctx, sender.Bytes(), receiver, msg.ContractAddress)
	if err != nil {
		return nil, err
	}

	// NOTE: check if the contract is self-destructed
	erc20 := common.HexToAddress(pair.ContractAddress)
	acc := k.evmKeeper.GetAccountWithoutBalance(ctx, erc20)

	if acc == nil || !acc.IsContract() {
		k.RemoveTokenPair(ctx, pair)
		k.Logger(ctx).Debug(
			"deleting self-destructed token pair from state",
			"contract", pair.ContractAddress,
		)
		// NOTE: return nil error to persist the changes from the deletion
		return nil, nil
	}

	// execute as per token destination type.
	switch {
	case pair.IsNativeCoin():
		return k.ConvertVoucherERC20ToNativeCoin(ctx, pair, msg, receiver, sender)
	case pair.IsNativeERC20():
		return k.ConvertNativeERC20ToVoucherCoin(ctx, pair, msg, receiver, sender)
	default:
		return nil, types.ErrUndefinedOwner
	}
}
