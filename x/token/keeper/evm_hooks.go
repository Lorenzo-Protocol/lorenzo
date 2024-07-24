package keeper

import (
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	evmtypes "github.com/evmos/ethermint/x/evm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Lorenzo-Protocol/lorenzo/v2/contracts/erc20"
	"github.com/Lorenzo-Protocol/lorenzo/v2/x/token/types"
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

// PostTxProcessing searches for erc20 transfer event specific to module account, and continues conversion of sdk logic.
func (k Keeper) PostTxProcessing(
	ctx sdk.Context,
	_ core.Message,
	receipt *ethtypes.Receipt,
) error {
	params := k.GetParams(ctx)
	if !params.EnableConversion || !params.EnableEVMHook {
		return nil
	}

	erc20 := erc20.ERC20MinterBurnerDecimalsContract.ABI

	// if non target log and inner data found, skip
	for i, log := range receipt.Logs {
		// transfer event topics length equals 3
		if len(log.Topics) != 3 {
			continue
		}

		// check transfer event and unpack its data
		eventID := log.Topics[0]
		event, err := erc20.EventByID(eventID)
		if err != nil {
			continue
		}

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

		// check positive amount of tokens transferred
		tokens, ok := transferEvent[0].(*big.Int)
		if !ok || tokens == nil || tokens.Sign() != 1 {
			continue
		}

		// only handle contract that are registered.
		contractAddr := log.Address
		id := k.GetTokenPairIdByERC20(ctx, contractAddr)
		pair, found := k.GetTokenPair(ctx, id)
		if !found {
			continue
		}

		// token must be sent to the module account evm address
		to := common.BytesToAddress(log.Topics[2].Bytes())
		if !bytes.Equal(to.Bytes(), types.ModuleAddress.Bytes()) {
			continue
		}

		// NOTE: now that token module and hook is enabled, target event is logged,
		// and token pair if found, we shall continue process convert logic.
		if !pair.Enabled {
			// NOTE: if pair conversion is disabled, we shall return evm tx. otherwise
			// the token is send to module account.
			k.Logger(ctx).Debug(
				"ERC20 token -> Cosmos coin conversion is disabled for pair",
				"coin", pair.Denom, "contract", pair.ContractAddress,
			)
			return types.ErrTokenPairDisabled
		}

		// create the corresponding sdk.Coin that is paired with ERC20
		coins := sdk.Coins{{Denom: pair.Denom, Amount: sdk.NewIntFromBigInt(tokens)}}

		switch pair.Source {
		case types.OWNER_MODULE:
			_, err = k.CallEVM(ctx, erc20, types.ModuleAddress, contractAddr, true, "burn", tokens)
		case types.OWNER_CONTRACT:
			err = k.bankKeeper.MintCoins(ctx, types.ModuleName, coins)
		default:
			err = types.ErrUndefinedOwner
		}

		if err != nil {
			k.Logger(ctx).Debug(
				"failed to process EVM hook for ER20 -> coin conversion",
				"coin", pair.Denom, "contract", pair.ContractAddress, "error", err.Error(),
			)
			return err
		}

		// get sender address
		from := common.BytesToAddress(log.Topics[1].Bytes())
		recipient := sdk.AccAddress(from.Bytes())

		// transfer the tokens from ModuleAccount to sender address
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recipient, coins); err != nil {
			k.Logger(ctx).Debug(
				"failed to process EVM hook for ER20 -> coin conversion",
				"tx-hash", receipt.TxHash.Hex(), "log-idx", i,
				"coin", pair.Denom, "contract", pair.ContractAddress, "error", err.Error(),
			)
			return err
		}
	}

	return nil
}
