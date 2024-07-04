package types

const (
	EventTypeTokenLock             = "token_lock"
	EventTypeTokenUnlock           = "token_unlock"
	EventTypeMint                  = "mint"
	EventTypeConvertCoin           = "convert_coin"
	EventTypeConvertERC20          = "convert_erc20"
	EventTypeBurn                  = "burn"
	EventTypeRegisterCoin          = "register_coin"
	EventTypeRegisterERC20         = "register_erc20"
	EventTypeToggleTokenConversion = "toggle_token_conversion" // #nosec

	AttributeKeyCosmosCoin = "cosmos_coin"
	AttributeKeyERC20Token = "erc20_token" // #nosec
	AttributeKeyReceiver   = "receiver"

	// ERC20EventTransfer defines the transfer event for ERC20
	ERC20EventTransfer = "Transfer"
)
