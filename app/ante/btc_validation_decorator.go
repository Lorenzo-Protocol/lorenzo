package ante

import (
	bbn "github.com/Lorenzo-Protocol/lorenzo/types"
	btclightclient "github.com/Lorenzo-Protocol/lorenzo/x/btclightclient/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type BtcValidationDecorator struct {
	BtcCfg bbn.BtcConfig
}

func newBtcValidationDecorator(
	cfg bbn.BtcConfig,
) BtcValidationDecorator {
	return BtcValidationDecorator{
		BtcCfg: cfg,
	}
}

func (bvd BtcValidationDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {

	// only do this validation when handling mempool addition. During DeliverTx they
	// should be performed by btclightclient and btccheckpoint modules
	if ctx.IsCheckTx() || ctx.IsReCheckTx() {
		for _, m := range tx.GetMsgs() {
			switch msg := m.(type) {
			case *btclightclient.MsgInsertHeaders:
				powLimit := bvd.BtcCfg.PowLimit()
				err := msg.ValidateHeaders(&powLimit)
				if err != nil {
					return ctx, btclightclient.ErrInvalidProofOfWOrk
				}
			default:
				// NOOP in case of other messages
			}
		}
	}
	return next(ctx, tx, simulate)
}
