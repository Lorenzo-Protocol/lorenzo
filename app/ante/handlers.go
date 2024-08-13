package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	ibcante "github.com/cosmos/ibc-go/v7/modules/core/ante"
	ethante "github.com/evmos/ethermint/app/ante"
	evmtypes "github.com/evmos/ethermint/x/evm/types"

	cosmosante "github.com/Lorenzo-Protocol/lorenzo/v3/app/ante/cosmos"
)

func newEthAnteHandler(options HandlerOptions) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		ethante.NewEthSetUpContextDecorator(options.EvmKeeper),

		// FIXME: need to allow universal fees for Eth as well
		ethante.NewEthMempoolFeeDecorator(options.EvmKeeper),                           // Check eth effective gas price against minimal-gas-prices
		ethante.NewEthMinGasPriceDecorator(options.FeeMarketKeeper, options.EvmKeeper), // Check eth effective gas price against the global MinGasPrice

		ethante.NewEthValidateBasicDecorator(options.EvmKeeper),
		ethante.NewEthSigVerificationDecorator(options.EvmKeeper),
		ethante.NewEthAccountVerificationDecorator(options.AccountKeeper, options.EvmKeeper),
		ethante.NewCanTransferDecorator(options.EvmKeeper),
		ethante.NewEthGasConsumeDecorator(options.EvmKeeper, options.MaxTxGasWanted),
		ethante.NewEthIncrementSenderSequenceDecorator(options.AccountKeeper), // innermost AnteDecorator.
		ethante.NewGasWantedDecorator(options.EvmKeeper, options.FeeMarketKeeper),
		ethante.NewEthEmitEventDecorator(options.EvmKeeper), // emit eth tx hash and index at the very last ante handler.
	)
}

// newLegacyCosmosAnteHandlerEip712 creates an AnteHandler to process legacy EIP-712
// transactions, as defined by the presence of an ExtensionOptionsWeb3Tx extension.
func newLegacyCosmosAnteHandlerEip712(options HandlerOptions) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		NewRejectMessagesDecorator(), // reject MsgEthereumTxs and vesting messages
		ethante.NewAuthzLimiterDecorator([]string{ // disable the Msg types that cannot be included on an authz.MsgExec msgs field
			sdk.MsgTypeURL(&evmtypes.MsgEthereumTx{}),
			sdk.MsgTypeURL(&vestingtypes.MsgCreateVestingAccount{}),
			sdk.MsgTypeURL(&vestingtypes.MsgCreatePeriodicVestingAccount{}),
			sdk.MsgTypeURL(&vestingtypes.MsgCreatePermanentLockedAccount{}),
		},
		),
		ante.NewSetUpContextDecorator(),
		ante.NewValidateBasicDecorator(),
		ante.NewTxTimeoutHeightDecorator(),
		// min gas price check
		cosmosante.NewMinGasPriceDecorator(options.FeeMarketKeeper, options.EvmKeeper, options.FeeKeeper),
		ante.NewValidateMemoDecorator(options.AccountKeeper),
		ante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		ante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper, options.TxFeeChecker),
		// SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewSetPubKeyDecorator(options.AccountKeeper),
		ante.NewValidateSigCountDecorator(options.AccountKeeper),
		ante.NewSigGasConsumeDecorator(options.AccountKeeper, ethante.DefaultSigVerificationGasConsumer),
		// Note: signature verification uses EIP instead of the cosmos signature validator
		NewLegacyEip712SigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
		ante.NewIncrementSequenceDecorator(options.AccountKeeper),
		ibcante.NewRedundantRelayDecorator(options.IBCKeeper),
		ethante.NewGasWantedDecorator(options.EvmKeeper, options.FeeMarketKeeper),
	)
}

func newCosmosAnteHandler(options HandlerOptions) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		NewRejectMessagesDecorator(), // reject MsgEthereumTxs and vesting msgs
		ethante.NewAuthzLimiterDecorator([]string{ // disable the Msg types that cannot be included on an authz.MsgExec msgs field
			sdk.MsgTypeURL(&evmtypes.MsgEthereumTx{}),
			sdk.MsgTypeURL(&vestingtypes.MsgCreateVestingAccount{}),
			sdk.MsgTypeURL(&vestingtypes.MsgCreatePeriodicVestingAccount{}),
			sdk.MsgTypeURL(&vestingtypes.MsgCreatePermanentLockedAccount{}),
		},
		),
		ante.NewSetUpContextDecorator(),
		ante.NewExtensionOptionsDecorator(options.ExtensionOptionChecker),
		// Use Mempool Fee Decorator from our txfees module instead of default one from auth
		ante.NewValidateBasicDecorator(),
		ante.NewTxTimeoutHeightDecorator(),
		// min gas price check
		cosmosante.NewMinGasPriceDecorator(options.FeeMarketKeeper, options.EvmKeeper, options.FeeKeeper),
		ante.NewValidateMemoDecorator(options.AccountKeeper),
		ante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		ante.NewSetPubKeyDecorator(options.AccountKeeper), // SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewValidateSigCountDecorator(options.AccountKeeper),
		ante.NewSigGasConsumeDecorator(options.AccountKeeper, ethante.DefaultSigVerificationGasConsumer),
		ante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper, options.TxFeeChecker),
		ante.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
		ante.NewIncrementSequenceDecorator(options.AccountKeeper),
		ibcante.NewRedundantRelayDecorator(options.IBCKeeper),
		ethante.NewGasWantedDecorator(options.EvmKeeper, options.FeeMarketKeeper),
		newBtcValidationDecorator(options.BtcConfig),
	)
}
