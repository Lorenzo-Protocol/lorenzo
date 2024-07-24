package keeper

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/Lorenzo-Protocol/lorenzo/v2/contracts/erc20"
	"github.com/Lorenzo-Protocol/lorenzo/v2/x/token/types"
)

// ConvertNativeCoinToVoucherERC20 converts token from bank to contract. Token source is sdk.Coin.
func (k Keeper) ConvertNativeCoinToVoucherERC20(
	ctx sdk.Context,
	pair types.TokenPair,
	msg *types.MsgConvertCoin,
	receiver common.Address,
	sender sdk.AccAddress,
) (*types.MsgConvertCoinResponse, error) {
	// 1. escrow native coins on module account
	coins := sdk.Coins{msg.Coin}
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, coins)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to escrow coins")
	}

	// 2. mint tokens on contract
	erc20ABI := erc20.ERC20MinterBurnerDecimalsContract.ABI // nolint: staticcheck
	erc20Addr := pair.GetERC20ContractAddress()
	balanceBefore := k.ERC20BalanceOf(ctx, erc20ABI, erc20Addr, receiver)
	if balanceBefore == nil {
		return nil, errorsmod.Wrapf(types.ErrEVMCall, "failed to retrieve balance")
	}

	_, err = k.CallEVM(ctx, erc20ABI, types.ModuleAddress, erc20Addr, true,
		"mint", receiver, msg.Coin.Amount.BigInt())
	if err != nil {
		return nil, err
	}

	// 3. check erc20 balances before and after
	balanceAfter := k.ERC20BalanceOf(ctx, erc20ABI, erc20Addr, receiver)
	if balanceAfter == nil {
		return nil, errorsmod.Wrapf(types.ErrEVMCall, "failed to retrieve balance")
	}
	balanceExpected := big.NewInt(0).Add(balanceBefore, msg.Coin.Amount.BigInt())
	if balanceAfter.Cmp(balanceExpected) != 0 {
		return nil, errorsmod.Wrapf(types.ErrBalanceInvariance,
			"invalid token balance - expected: %v, actual: %v", balanceExpected, balanceAfter)
	}

	ctx.EventManager().EmitEvents(
		sdk.Events{
			sdk.NewEvent(
				types.EventTypeConvertCoin,
				sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
				sdk.NewAttribute(types.AttributeKeyReceiver, msg.Receiver),
				sdk.NewAttribute(sdk.AttributeKeyAmount, msg.Coin.Amount.String()),
				sdk.NewAttribute(types.AttributeKeyCosmosCoin, msg.Coin.Denom),
				sdk.NewAttribute(types.AttributeKeyERC20Token, pair.ContractAddress),
			),
		},
	)

	return &types.MsgConvertCoinResponse{}, nil
}

// ConvertVoucherERC20ToNativeCoin converts token from contract to bank. Token destination is sdk.Coin.
func (k Keeper) ConvertVoucherERC20ToNativeCoin(
	ctx sdk.Context,
	pair types.TokenPair,
	msg *types.MsgConvertERC20,
	receiver sdk.AccAddress,
	sender common.Address,
) (*types.MsgConvertERC20Response, error) {
	erc20ABI := erc20.ERC20MinterBurnerDecimalsContract.ABI // nolint: staticcheck
	erc20Addr := pair.GetERC20ContractAddress()
	balanceCoinBefore := k.bankKeeper.GetBalance(ctx, receiver, pair.Denom)
	balanceTokenBefore := k.ERC20BalanceOf(ctx, erc20ABI, erc20Addr, sender)
	if balanceTokenBefore == nil {
		return nil, errorsmod.Wrap(types.ErrEVMCall, "failed to retrieve balance")
	}
	// 1. burn coins on contract
	_, err := k.CallEVM(ctx, erc20ABI, types.ModuleAddress, erc20Addr, true, "burnCoins", sender, msg.Amount.BigInt())
	if err != nil {
		return nil, err
	}

	// 2. un-escrow coins and sent to receiver
	coins := sdk.Coins{sdk.Coin{Denom: pair.Denom, Amount: msg.Amount}}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiver, coins)
	if err != nil {
		return nil, err
	}

	// 3. check balances before and after on both side.
	balanceCoinAfter := k.bankKeeper.GetBalance(ctx, receiver, pair.Denom)
	balanceCoinExpected := balanceCoinBefore.Add(coins[0])
	if !balanceCoinAfter.IsEqual(balanceCoinExpected) {
		return nil, errorsmod.Wrapf(
			types.ErrBalanceInvariance,
			"invalid coin balance - expected: %v, actual: %v", balanceCoinExpected, balanceCoinAfter,
		)
	}

	balanceTokenAfter := k.ERC20BalanceOf(ctx, erc20ABI, erc20Addr, sender)
	if balanceTokenAfter == nil {
		return nil, errorsmod.Wrap(types.ErrEVMCall, "failed to retrieve balance")
	}
	balanceTokenExpected := big.NewInt(0).Sub(balanceTokenBefore, coins[0].Amount.BigInt())
	if balanceTokenAfter.Cmp(balanceTokenExpected) != 0 {
		return nil, errorsmod.Wrapf(
			types.ErrBalanceInvariance,
			"invalid token balance - expected: %v, actual: %v", balanceTokenExpected, balanceTokenAfter,
		)
	}

	ctx.EventManager().EmitEvents(
		sdk.Events{
			sdk.NewEvent(
				types.EventTypeConvertERC20,
				sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
				sdk.NewAttribute(types.AttributeKeyReceiver, msg.Receiver),
				sdk.NewAttribute(sdk.AttributeKeyAmount, msg.Amount.String()),
				sdk.NewAttribute(types.AttributeKeyCosmosCoin, pair.Denom),
				sdk.NewAttribute(types.AttributeKeyERC20Token, msg.ContractAddress),
			),
		},
	)

	return &types.MsgConvertERC20Response{}, nil
}

// ConvertVoucherCoinToNativeERC20 converts token from bank to contract. Token source is ERC20 contract.
func (k Keeper) ConvertVoucherCoinToNativeERC20(
	ctx sdk.Context,
	pair types.TokenPair,
	msg *types.MsgConvertCoin,
	receiver common.Address,
	sender sdk.AccAddress,
) (*types.MsgConvertCoinResponse, error) {
	// 1. escrow coins on module account
	coins := sdk.Coins{msg.Coin}
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, coins); err != nil {
		return nil, errorsmod.Wrap(err, "failed to escrow coins")
	}

	// 2. un-escrow tokens from contract and sent to receiver
	erc20ABI := erc20.ERC20MinterBurnerDecimalsContract.ABI // nolint: staticcheck
	erc20Addr := pair.GetERC20ContractAddress()
	balanceBefore := k.ERC20BalanceOf(ctx, erc20ABI, erc20Addr, receiver)
	if balanceBefore == nil {
		return nil, errorsmod.Wrapf(types.ErrEVMCall, "failed to retrieve balance")
	}

	res, err := k.CallEVM(ctx, erc20ABI, types.ModuleAddress, erc20Addr, true,
		"transfer", receiver, msg.Coin.Amount.BigInt())
	if err != nil {
		return nil, err
	}

	// 3. check unpacked return
	var unpackedRet types.ERC20BoolResponse
	if err := erc20ABI.UnpackIntoInterface(&unpackedRet, "transfer", res.Ret); err != nil {
		return nil, err
	}
	if !unpackedRet.Value {
		return nil, errorsmod.Wrap(types.ErrEVMCall, "failed to transfer tokens")
	}

	// 4. check balances before and after
	balanceAfter := k.ERC20BalanceOf(ctx, erc20ABI, erc20Addr, receiver)
	if balanceAfter == nil {
		return nil, errorsmod.Wrap(types.ErrEVMCall, "failed to retrieve balance")
	}

	balanceExpected := big.NewInt(0).Add(balanceBefore, msg.Coin.Amount.BigInt())
	if balanceAfter.Cmp(balanceExpected) != 0 {
		return nil, errorsmod.Wrapf(
			types.ErrBalanceInvariance,
			"invalid token balance - expected: %v, actual: %v", balanceExpected, balanceAfter,
		)
	}

	// 5. burn escrowed coins from module account
	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, coins)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to burn coins")
	}

	// 6. monitor unexpected approval
	if err := k.assureNoApprovalEvent(res); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(
		sdk.Events{
			sdk.NewEvent(
				types.EventTypeConvertCoin,
				sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
				sdk.NewAttribute(types.AttributeKeyReceiver, msg.Receiver),
				sdk.NewAttribute(sdk.AttributeKeyAmount, msg.Coin.Amount.String()),
				sdk.NewAttribute(types.AttributeKeyCosmosCoin, msg.Coin.Denom),
				sdk.NewAttribute(types.AttributeKeyERC20Token, pair.ContractAddress),
			),
		},
	)
	return &types.MsgConvertCoinResponse{}, nil
}

// ConvertNativeERC20ToVoucherCoin converts token from contract to contract. Token destination is ERC20 contract.
func (k Keeper) ConvertNativeERC20ToVoucherCoin(
	ctx sdk.Context,
	pair types.TokenPair,
	msg *types.MsgConvertERC20,
	receiver sdk.AccAddress,
	sender common.Address,
) (*types.MsgConvertERC20Response, error) {
	erc20ABI := erc20.ERC20MinterBurnerDecimalsContract.ABI // nolint: staticcheck
	erc20Addr := pair.GetERC20ContractAddress()
	balanceCoinBefore := k.bankKeeper.GetBalance(ctx, receiver, pair.Denom)
	balanceTokenBefore := k.ERC20BalanceOf(ctx, erc20ABI, erc20Addr, types.ModuleAddress)
	if balanceTokenBefore == nil {
		return nil, errorsmod.Wrap(types.ErrEVMCall, "failed to retrieve balance")
	}

	// 1. escrow sender's token on module account
	res, err := k.CallEVM(ctx, erc20ABI, sender, erc20Addr, true, "transfer", types.ModuleAddress, msg.Amount.BigInt())
	if err != nil {
		return nil, err
	}

	var unpackedRet types.ERC20BoolResponse
	if err := erc20ABI.UnpackIntoInterface(&unpackedRet, "transfer", res.Ret); err != nil {
		return nil, err
	}
	if !unpackedRet.Value {
		return nil, errorsmod.Wrap(errortypes.ErrLogic, "failed to execute transfer")
	}

	// 2. check token balances before and after
	balanceTokenAfter := k.ERC20BalanceOf(ctx, erc20ABI, erc20Addr, types.ModuleAddress)
	if balanceTokenAfter == nil {
		return nil, errorsmod.Wrap(types.ErrEVMCall, "failed to retrieve balance")
	}
	balanceTokenExpected := big.NewInt(0).Add(balanceTokenBefore, msg.Amount.BigInt())
	if balanceTokenAfter.Cmp(balanceTokenExpected) != 0 {
		return nil, errorsmod.Wrapf(
			types.ErrBalanceInvariance,
			"invalid token balance - expected: %v, actual: %v", balanceTokenExpected, balanceTokenAfter,
		)
	}

	// 3. mint coins and send to receiver
	coins := sdk.Coins{sdk.Coin{Denom: pair.Denom, Amount: msg.Amount}}
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, coins); err != nil {
		return nil, err
	}
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiver, coins); err != nil {
		return nil, err
	}

	// 4. check coin balances before and after
	balanceCoinAfter := k.bankKeeper.GetBalance(ctx, receiver, pair.Denom)
	balanceCoinExpected := balanceCoinBefore.Add(coins[0])
	if !balanceCoinAfter.IsEqual(balanceCoinExpected) {
		return nil, errorsmod.Wrapf(
			types.ErrBalanceInvariance,
			"invalid coin balance - expected: %v, actual: %v", balanceCoinExpected, balanceCoinAfter,
		)
	}

	// 5. monitor unexpected approval
	if err := k.assureNoApprovalEvent(res); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(
		sdk.Events{
			sdk.NewEvent(
				types.EventTypeConvertERC20,
				sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
				sdk.NewAttribute(types.AttributeKeyReceiver, msg.Receiver),
				sdk.NewAttribute(sdk.AttributeKeyAmount, msg.Amount.String()),
				sdk.NewAttribute(types.AttributeKeyCosmosCoin, pair.Denom),
				sdk.NewAttribute(types.AttributeKeyERC20Token, msg.ContractAddress),
			),
		},
	)

	return &types.MsgConvertERC20Response{}, nil
}
