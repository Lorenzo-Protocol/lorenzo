package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/ethereum/go-ethereum/common"

	"github.com/Lorenzo-Protocol/lorenzo/x/token/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	*Keeper
}

func NewMsgServer(k *Keeper) types.MsgServer {
	return &msgServer{k}
}

// ConvertCoin implements MsgServer.ConvertCoin
func (m msgServer) ConvertCoin(goCtx context.Context, msg *types.MsgConvertCoin) (*types.MsgConvertCoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	receiver := common.HexToAddress(msg.Receiver)
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	// check if the token pair is enabled for minting and converting.
	pair, err := m.MintEnabled(ctx, sender, receiver.Bytes(), msg.Coin.Denom)
	if err != nil {
		return nil, err
	}

	// NOTE: check if the contract is self-destructed
	erc20 := common.HexToAddress(pair.ContractAddress)
	acc := m.evmKeeper.GetAccountWithoutBalance(ctx, erc20)

	if acc == nil || !acc.IsContract() {
		m.DeleteTokenPair(ctx, pair)
		m.Logger(ctx).Debug(
			"deleting self-destructed token pair from state",
			"contract", pair.ContractAddress,
		)
		// NOTE: return nil error to persist the changes from the deletion
		return nil, nil
	}

	// execute as per token source type.
	switch {
	case pair.IsNativeCoin():
		return m.Keeper.ConvertCoinNativeCoin(ctx, pair, msg, receiver, sender)
	case pair.IsNativeERC20():
		return m.Keeper.ConvertCoinNativeERC20(ctx, pair, msg, receiver, sender)
	default:
		return nil, types.ErrUndefinedOwner
	}
}

// ConvertERC20 implements MsgServer.ConvertERC20
func (m msgServer) ConvertERC20(goCtx context.Context, msg *types.MsgConvertERC20) (*types.MsgConvertERC20Response, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender := common.HexToAddress(msg.Sender)
	receiver, err := sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return nil, err
	}

	// check if the token pair is enabled for minting and converting.
	pair, err := m.MintEnabled(ctx, sender.Bytes(), receiver, msg.ContractAddress)
	if err != nil {
		return nil, err
	}

	// NOTE: check if the contract is self-destructed
	erc20 := common.HexToAddress(pair.ContractAddress)
	acc := m.evmKeeper.GetAccountWithoutBalance(ctx, erc20)

	if acc == nil || !acc.IsContract() {
		m.DeleteTokenPair(ctx, pair)
		m.Logger(ctx).Debug(
			"deleting self-destructed token pair from state",
			"contract", pair.ContractAddress,
		)
		// NOTE: return nil error to persist the changes from the deletion
		return nil, nil
	}

	// execute as per token destination type.
	switch {
	case pair.IsNativeCoin():
		return m.ConvertERC20NativeCoin(ctx, pair, msg, receiver, sender) // case 1.2
	case pair.IsNativeERC20():
		return m.ConvertERC20NativeERC20(ctx, pair, msg, receiver, sender) // case 2.1
	default:
		return nil, types.ErrUndefinedOwner
	}
}

// UpdateParams implements MsgServer.UpdateParams
func (m msgServer) UpdateParams(goCtx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if m.authority.String() != msg.Authority {
		return nil, errorsmod.Wrapf(govtypes.ErrInvalidSigner,
			"invalid authority; expected %s, got %s", m.authority.String(), msg.Authority)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := m.SetParams(ctx, msg.Params); err != nil {
		return nil, err
	}

	return &types.MsgUpdateParamsResponse{}, nil
}
