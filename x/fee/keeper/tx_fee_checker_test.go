package keeper_test

import (
	sdkmath "cosmossdk.io/math"
	"github.com/pkg/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/Lorenzo-Protocol/lorenzo/v3/x/fee/keeper"
	"github.com/Lorenzo-Protocol/lorenzo/v3/x/fee/types"
)

var (
	_ sdk.Tx    = MockTx{}
	_ sdk.FeeTx = MockTx{}
)

func (suite *KeeperTestSuite) TestTxFeeChecker() {
	params := types.Params{}
	suite.NoError(suite.keeper.SetParams(suite.ctx, params), "set params failed")

	ctx := suite.ctx.WithIsCheckTx(true)
	ctx = ctx.WithMinGasPrices(sdk.NewDecCoins(sdk.NewDecCoin("alrz", sdkmath.NewInt(20))))

	feeChecker := keeper.TxFeeChecker(suite.keeper)

	suite.Run("Single msg and no set non-fee list", func() {
		// requiredFee, priority, err := feeChecker(ctx, tx)
		tx := MockTx{
			gas: 20000,
			fee: sdk.NewCoins(sdk.NewInt64Coin("alrz", 0)),
			msgs: []sdk.Msg{
				&banktypes.MsgSend{},
			},
		}
		_, _, err := feeChecker(ctx, tx)
		suite.True(errors.Is(err, sdkerrors.ErrInsufficientFee))
	})

	suite.Run("Multiple msg and no set non-fee list", func() {
		// requiredFee, priority, err := feeChecker(ctx, tx)
		tx := MockTx{
			gas: 20000,
			fee: sdk.NewCoins(sdk.NewInt64Coin("alrz", 0)),
			msgs: []sdk.Msg{
				&banktypes.MsgSend{},
				&banktypes.MsgUpdateParams{},
			},
		}
		_, _, err := feeChecker(ctx, tx)
		suite.True(errors.Is(err, sdkerrors.ErrInsufficientFee))
	})

	params = types.Params{
		NonFeeMsgs: []string{
			sdk.MsgTypeURL(&banktypes.MsgSend{}),
			sdk.MsgTypeURL(&banktypes.MsgUpdateParams{}),
		},
	}
	suite.NoError(suite.keeper.SetParams(suite.ctx, params), "set params failed")

	suite.Run("Single msg and set non-fee list", func() {
		// requiredFee, priority, err := feeChecker(ctx, tx)
		tx := MockTx{
			gas: 20000,
			fee: sdk.NewCoins(sdk.NewInt64Coin("alrz", 0)),
			msgs: []sdk.Msg{
				&banktypes.MsgSend{},
			},
		}
		requiredFee, priority, err := feeChecker(ctx, tx)
		suite.NoError(err)
		suite.Len(requiredFee, 0)
		suite.Equal(priority, int64(0))
	})

	suite.Run("Multiple msg and set non-fee list", func() {
		// requiredFee, priority, err := feeChecker(ctx, tx)
		tx := MockTx{
			gas: 20000,
			fee: sdk.NewCoins(sdk.NewInt64Coin("alrz", 0)),
			msgs: []sdk.Msg{
				&banktypes.MsgSend{},
				&banktypes.MsgUpdateParams{},
			},
		}
		requiredFee, priority, err := feeChecker(ctx, tx)
		suite.NoError(err)
		suite.Len(requiredFee, 0)
		suite.Equal(priority, int64(0))
	})

	suite.Run("Mix of paid msg and non-paid msg", func() {
		// requiredFee, priority, err := feeChecker(ctx, tx)
		tx := MockTx{
			gas: 20000,
			fee: sdk.NewCoins(sdk.NewInt64Coin("alrz", 0)),
			msgs: []sdk.Msg{
				&banktypes.MsgSend{},
				&types.MsgUpdateParams{},
			},
		}
		_, _, err := feeChecker(ctx, tx)
		suite.True(errors.Is(err, sdkerrors.ErrInsufficientFee))
	})
}

type MockTx struct {
	gas     uint64
	fee     sdk.Coins
	payer   sdk.AccAddress
	granter sdk.AccAddress
	msgs    []sdk.Msg
}

func (m MockTx) GetGas() uint64 {
	return m.gas
}

func (m MockTx) GetFee() sdk.Coins {
	return m.fee
}

func (m MockTx) FeePayer() sdk.AccAddress {
	return m.payer
}

func (m MockTx) FeeGranter() sdk.AccAddress {
	return m.granter
}

func (m MockTx) GetMsgs() []sdk.Msg {
	return m.msgs
}

func (m MockTx) ValidateBasic() error {
	return nil
}
