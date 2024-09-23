package keeper

import (
	"fmt"
	"math/big"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"

	"github.com/Lorenzo-Protocol/lorenzo/v3/x/btcstaking/types"
	ccevtypes "github.com/Lorenzo-Protocol/lorenzo/v3/x/ccev/types"
)

// stakingxBTCEvent is a struct that contains the sender, plan id, BTC contract address, stake amount, and stBTC amount.
type stakingxBTCEvent struct {
	Contract           common.Address `json:"contract"`
	Identifier         uint64         `json:"identifier"`
	Sender             common.Address `json:"sender"`
	PlanID             uint64         `json:"plan_id"`
	BTCcontractAddress common.Address `json:"btc_contract_address"`
	StakeAmount        big.Int        `json:"stake_amount"`
	StBTCAmount        big.Int        `json:"st_btc_amount"`
}

var _ ccevtypes.EventHandler = (*eventHandler)(nil)

type eventHandler struct {
	keeper Keeper
}

// Execute implements types.EventHandler.
func (e *eventHandler) Execute(ctx sdk.Context, chainID uint32, events []*ccevtypes.Event) error {
	totalStBTCAmt := new(big.Int)
	for i := range events {
		event, err := e.parseEvent(events[i])
		if err != nil {
			return err
		}

		if e.keeper.hasBTCBStakingRecord(ctx, chainID, event.Contract.Bytes(), event.Identifier) {
			return types.ErrDuplicateStakingEvent.Wrapf("duplicate event,planID %d,stakingIdx %d,contract %s",
				event.PlanID,
				event.Identifier,
				event.Contract.String(),
			)
		}

		amount := new(big.Int).SetBytes(event.StBTCAmount.Bytes())
		result := ""
		totalStBTCAmt = totalStBTCAmt.Add(totalStBTCAmt, amount)

		btcbStakingRecord := &types.BTCBStakingRecord{
			StakingIdx:    event.Identifier,
			Contract:      event.Contract.Bytes(),
			ReceiverAddr:  event.Sender.String(),
			Amount:        math.NewIntFromBigInt(amount),
			ChainId:       chainID,
			MintYatResult: result,
			PlanId:        event.PlanID,
		}

		e.keeper.addBTCBStakingRecord(ctx, btcbStakingRecord)
		// emit an event
		ctx.EventManager().EmitTypedEvent(types.NewEventBTCBStakingCreated(btcbStakingRecord)) //nolint:errcheck,gosec
	}

	// mint stBTC to the bridgeAddr
	totalStBTC := sdk.NewCoins(sdk.NewCoin(types.NativeTokenDenom, sdk.NewIntFromBigInt(totalStBTCAmt)))
	if err := e.keeper.bankKeeper.MintCoins(ctx, types.ModuleName, totalStBTC); err != nil {
		return err
	}

	bridgeAddr := sdk.AccAddress(common.HexToAddress(e.keeper.GetParams(ctx).BridgeAddr).Bytes())
	if err := e.keeper.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, bridgeAddr, totalStBTC); err != nil {
		return err
	}
	return nil
}

// GetUniqueID implements types.EventHandler.
func (e *eventHandler) GetUniqueID(ctx sdk.Context, event *ccevtypes.Event) (string, error) {
	stakingEvent, err := e.parseEvent(event)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", stakingEvent.Identifier), nil
}

func (e *eventHandler) parseEvent(event *ccevtypes.Event) (*stakingxBTCEvent, error) {
	if len(event.Topics) != 4 {
		return nil, errorsmod.Wrapf(
			ccevtypes.ErrInvalidEvent,
			"event has wrong number of topics, expected 4, actual: %d",
			len(event.Topics),
		)
	}
	// stakeIndex
	identifier := new(big.Int).SetBytes(event.Topics[1].Bytes())
	// planId
	planID := new(big.Int).SetBytes(event.Topics[2].Bytes())
	// sender
	sender := common.BytesToAddress(event.Topics[3].Bytes())

	if len(event.Args) != 3 {
		return nil, errorsmod.Wrapf(
			ccevtypes.ErrInvalidEvent,
			"event has wrong number of parameters, expected 3, actual: %d",
			len(event.Args),
		)
	}

	// btcContractAddress
	btcContractAddress, ok := event.Args[0].(common.Address)
	if !ok {
		return nil, errorsmod.Wrap(
			ccevtypes.ErrInvalidEvent,
			"event `btcContractAddress` parameters is invalid, expected `common.Address`",
		)
	}

	// stakeAmount
	stakeAmount, ok := event.Args[1].(*big.Int)
	if !ok {
		return nil, errorsmod.Wrap(
			ccevtypes.ErrInvalidEvent,
			"event `stakeAmount` parameters is invalid, expected `*big.Int`",
		)
	}

	// stBTCAmount
	stBTCAmount, ok := event.Args[2].(*big.Int)
	if !ok {
		return nil, errorsmod.Wrap(
			ccevtypes.ErrInvalidEvent,
			"event `stBTCAmount` parameters is invalid, expected `*big.Int`",
		)
	}

	return &stakingxBTCEvent{
		Contract:           event.Address,
		Identifier:         identifier.Uint64(),
		Sender:             sender,
		PlanID:             planID.Uint64(),
		BTCcontractAddress: btcContractAddress,
		StakeAmount:        *stakeAmount,
		StBTCAmount:        *stBTCAmount,
	}, nil
}
