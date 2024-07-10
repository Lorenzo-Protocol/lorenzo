package keeper

import (
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	evmtypes "github.com/evmos/ethermint/x/evm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Lorenzo-Protocol/lorenzo/contracts/erc20"
	"github.com/Lorenzo-Protocol/lorenzo/x/token/types"
)

var _ evmtypes.EvmHooks = Hooks{}

type Hooks struct {
	keeper *Keeper
}

func (k Keeper) Hooks() Hooks {
	return Hooks{&k}
}

func (h Hooks) PostTxProcessing(ctx sdk.Context, msg core.Message, receipt *ethtypes.Receipt) error {
	return h.keeper.PostTxProcessing(ctx, msg, receipt)
}

func (k Keeper) PostTxProcessing(
	ctx sdk.Context,
	_ core.Message,
	receipt *ethtypes.Receipt,
) error {
	params := k.GetParams(ctx)
	if !params.EnableConvert || !params.EnableEVMHook {
		return nil
	}

	erc20 := erc20.ERC20MinterBurnerDecimalsContract.ABI

	for i, log := range receipt.Logs {
		// Note: the `Transfer` event contains 3 topics (id, from, to)
		if len(log.Topics) != 3 {
			continue
		}

		// Check if event is included in ERC20
		eventID := log.Topics[0]
		event, err := erc20.EventByID(eventID)
		if err != nil {
			continue
		}

		// Check if event is a `Transfer` event.
		if event.Name != types.ERC20EventTransfer {
			k.Logger(ctx).Info("emitted event", "name", event.Name, "signature", event.Sig)
			continue
		}

		transferEvent, err := erc20.Unpack(event.Name, log.Data)
		if err != nil {
			k.Logger(ctx).Error("failed to unpack transfer event", "error", err.Error())
			continue
		}

		if len(transferEvent) == 0 {
			continue
		}

		tokens, ok := transferEvent[0].(*big.Int)
		// safety check and ignore if amount not positive
		if !ok || tokens == nil || tokens.Sign() != 1 {
			continue
		}

		// Check that the contract is a registered token pair
		contractAddr := log.Address
		id := k.GetTokenPairIdByERC20(ctx, contractAddr)
		pair, found := k.GetTokenPair(ctx, id)
		if !found {
			continue
		}

		// Check if tokens are sent to module address
		to := common.BytesToAddress(log.Topics[2].Bytes())
		if !bytes.Equal(to.Bytes(), types.ModuleAddress.Bytes()) {
			continue
		}

		// Check that conversion for the pair is enabled. Fail
		if !pair.Enabled {
			// continue to allow transfers for the ERC20 in case the token pair is
			// disabled
			k.Logger(ctx).Debug(
				"ERC20 token -> Cosmos coin conversion is disabled for pair",
				"coin", pair.Denom, "contract", pair.ContractAddress,
			)
			continue
		}

		// create the corresponding sdk.Coin that is paired with ERC20
		coins := sdk.Coins{{Denom: pair.Denom, Amount: sdk.NewIntFromBigInt(tokens)}}

		switch pair.Ownership {
		case types.OWNER_MODULE:
			_, err = k.CallEVM(ctx, erc20, types.ModuleAddress, contractAddr, true, "burn", tokens)
		case types.OWNER_EXTERNAL:
			err = k.bankKeeper.MintCoins(ctx, types.ModuleName, coins)
		default:
			err = types.ErrUndefinedOwner
		}

		if err != nil {
			k.Logger(ctx).Debug(
				"failed to process EVM hook for ER20 -> coin conversion",
				"coin", pair.Denom, "contract", pair.ContractAddress, "error", err.Error(),
			)
			continue
		}

		// Only need last 20 bytes from log.topics
		from := common.BytesToAddress(log.Topics[1].Bytes())
		recipient := sdk.AccAddress(from.Bytes())

		// transfer the tokens from ModuleAccount to sender address
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recipient, coins); err != nil {
			k.Logger(ctx).Debug(
				"failed to process EVM hook for ER20 -> coin conversion",
				"tx-hash", receipt.TxHash.Hex(), "log-idx", i,
				"coin", pair.Denom, "contract", pair.ContractAddress, "error", err.Error(),
			)
			continue
		}
	}

	return nil
}
