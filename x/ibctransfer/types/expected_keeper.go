package types

import (
	"context"

	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"

	tokentypes "github.com/Lorenzo-Protocol/lorenzo/v2/x/token/types"
)

// TokenKeeper defines the expected interface needed to convert erc20 to native coin.
type TokenKeeper interface {
	GetTokenPairId(ctx sdk.Context, token string) []byte
	GetTokenPair(ctx sdk.Context, id []byte) (tokentypes.TokenPair, bool)
	IsConvertEnabled(ctx sdk.Context) bool
	IsRegisteredByERC20(ctx sdk.Context, erc20Addr common.Address) bool
	ConvertERC20(goCtx context.Context, msg *tokentypes.MsgConvertERC20) (*tokentypes.MsgConvertERC20Response, error)
}

// BankKeeper defines the expected interface needed to check balances and send coins.
type BankKeeper interface {
	transfertypes.BankKeeper
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
}

// AccountKeeper defines the expected interface needed to retrieve account info.
type AccountKeeper interface {
	transfertypes.AccountKeeper
	GetAccount(sdk.Context, sdk.AccAddress) authtypes.AccountI
}
