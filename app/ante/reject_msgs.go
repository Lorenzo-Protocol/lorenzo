package ante

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	vesting "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	"golang.org/x/exp/slices"

	evmtypes "github.com/evmos/ethermint/x/evm/types"

	btcstakingtypes "github.com/Lorenzo-Protocol/lorenzo/x/btcstaking/types"
)

// RejectMessagesDecorator prevents invalid msg types from being executed
type RejectMessagesDecorator struct {
	disabledMsgTypeURLs []string
	onlyonceMsgTypeURLs []string
}

var (
	_                  sdk.AnteDecorator = RejectMessagesDecorator{}
	authzExecsgTypeURL                   = sdk.MsgTypeURL(&authz.MsgExec{})
)

// NewRejectMessagesDecorator creates a decorator to block vesting messages from reaching the mempool
func NewRejectMessagesDecorator() RejectMessagesDecorator {
	return RejectMessagesDecorator{
		disabledMsgTypeURLs: []string{
			sdk.MsgTypeURL(&vesting.MsgCreateVestingAccount{}),
			sdk.MsgTypeURL(&vesting.MsgCreatePermanentLockedAccount{}),
			sdk.MsgTypeURL(&vesting.MsgCreatePeriodicVestingAccount{}),
		},
		onlyonceMsgTypeURLs: []string{
			sdk.MsgTypeURL(&btcstakingtypes.MsgCreateBTCStaking{}),
			sdk.MsgTypeURL(&btcstakingtypes.MsgBurnRequest{}),
		},
	}
}

// AnteHandle rejects messages that requires ethereum-specific authentication.
// For example `MsgEthereumTx` requires fee to be deducted in the antehandler in
// order to perform the refund.
func (rmd RejectMessagesDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	msgTypeCount := make(map[string]int)

	for _, msg := range tx.GetMsgs() {
		if _, ok := msg.(*evmtypes.MsgEthereumTx); ok {
			return ctx, errorsmod.Wrapf(
				sdkerrors.ErrInvalidType,
				"MsgEthereumTx needs to be contained within a tx with 'ExtensionOptionsEthereumTx' option",
			)
		}

		typeURL := sdk.MsgTypeURL(msg)

		// check if the msg type is disabled
		if slices.Contains(rmd.disabledMsgTypeURLs, typeURL) {
			return ctx, errorsmod.Wrapf(
				sdkerrors.ErrUnauthorized,
				"MsgTypeURL %s not supported",
				typeURL,
			)
		}

		
		if slices.Contains(rmd.onlyonceMsgTypeURLs, typeURL) {
			msgTypeCount[typeURL]++
		}

		// check if the msg type is authzExecsgTypeURL
		if typeURL == authzExecsgTypeURL {
			msgs, err := msg.(*authz.MsgExec).GetMessages()
			if err != nil {
				return ctx, err
			}

			for _, msg := range msgs {
				if slices.Contains(rmd.onlyonceMsgTypeURLs, sdk.MsgTypeURL(msg)) {
					msgTypeCount[typeURL]++
				}
			}
		}

		// check if the msg type is only once
		if msgTypeCount[typeURL] > 1 {
			return ctx, errorsmod.Wrapf(
				sdkerrors.ErrInvalidRequest,
				"a transaction can only contain one %s message",
				typeURL,
			)
		}
	}
	return next(ctx, tx, simulate)
}
