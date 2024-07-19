package keeper

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/Lorenzo-Protocol/lorenzo/x/btcstaking/types"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/ethereum/go-ethereum/common"
)

const (
	EthAddrLen        = 20
	ChainIDLen        = 4
	SatoshiToStBTCMul = 1e10
)

const (
	Dep0Amount = 4e5
	Dep1Amount = 2e6
	Dep2Amount = 1e7
	Dep3Amount = 5e7
)

type msgServer struct {
	*Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper *Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func NewBTCTxFromBytes(txBytes []byte) (*wire.MsgTx, error) {
	var msgTx wire.MsgTx
	rbuf := bytes.NewReader(txBytes)
	if err := msgTx.Deserialize(rbuf); err != nil {
		return nil, err
	}

	return &msgTx, nil
}

const maxOpReturnPkScriptSize = 83

func ExtractPaymentTo(tx *wire.MsgTx, addr btcutil.Address) (uint64, error) {
	payToAddrScript, err := txscript.PayToAddrScript(addr)
	if err != nil {
		return 0, fmt.Errorf("invalid address")
	}
	var amt uint64 = 0
	for _, out := range tx.TxOut {
		if bytes.Equal(out.PkScript, payToAddrScript) {
			amt += uint64(out.Value)
		}
	}
	return amt, nil
}

func ExtractPaymentToWithOpReturnIdAndDust(tx *wire.MsgTx, addr btcutil.Address, dustAmount int64) (uint64, []byte, error) {
	payToAddrScript, err := txscript.PayToAddrScript(addr)
	if err != nil {
		return 0, nil, fmt.Errorf("invalid address")
	}
	var amt uint64 = 0
	foundOpReturnId := false
	var opReturnId []byte
	for _, out := range tx.TxOut {
		if bytes.Equal(out.PkScript, payToAddrScript) && out.Value >= dustAmount {
			amt += uint64(out.Value)
		} else {
			pkScript := out.PkScript
			pkScriptLen := len(pkScript)
			// valid op return script will have at least 2 bytes
			// - fisrt byte should be OP_RETURN marker
			// - second byte should indicate how many bytes there are in opreturn script
			if pkScriptLen > 1 &&
				pkScriptLen <= maxOpReturnPkScriptSize &&
				pkScript[0] == txscript.OP_RETURN {

				// if this is OP_PUSHDATA1, we need to drop first 3 bytes as those are related
				// to script iteslf i.e OP_RETURN + OP_PUSHDATA1 + len of bytes
				if pkScript[1] == txscript.OP_PUSHDATA1 {
					opReturnId = pkScript[3:]
				} else if pkScript[1] == txscript.OP_PUSHDATA2 {
					opReturnId = pkScript[4:]
				} else if pkScript[1] == txscript.OP_PUSHDATA4 {
					opReturnId = pkScript[6:]
				} else {
					// this should be one of OP_DATAXX opcodes we drop first 2 bytes
					opReturnId = pkScript[2:]
				}
				foundOpReturnId = true
			}
		}
	}
	if !foundOpReturnId {
		return 0, nil, fmt.Errorf("expected op_return_id not found")
	}
	return amt, opReturnId, nil
}

func canPerformMint(signer sdk.AccAddress, p types.Params) bool {
	if len(p.MinterAllowList) == 0 {
		return true
	}
	for _, addr := range p.MinterAllowList {
		if sdk.MustAccAddressFromBech32(addr).Equals(signer) {
			return true
		}
	}
	return false
}

func checkBTCTxDepth(stakingTxDepth uint64, btcAmount uint64) error {
	if btcAmount < Dep0Amount { // no depth check required
	} else if btcAmount < Dep1Amount { // at least 1 depth
		if stakingTxDepth < 1 {
			return types.ErrBlkHdrNotConfirmed.Wrapf("not k-deep: k=1; depth=%d", stakingTxDepth)
		}
	} else if btcAmount < Dep2Amount {
		if stakingTxDepth < 2 {
			return types.ErrBlkHdrNotConfirmed.Wrapf("not k-deep: k=2; depth=%d", stakingTxDepth)
		}
	} else if btcAmount < Dep3Amount {
		if stakingTxDepth < 3 {
			return types.ErrBlkHdrNotConfirmed.Wrapf("not k-deep: k=3; depth=%d", stakingTxDepth)
		}
	} else if stakingTxDepth < 4 {
		return types.ErrBlkHdrNotConfirmed.Wrapf("not k-deep: k=4; depth=%d", stakingTxDepth)
	}
	return nil
}

func (ms msgServer) CreateBTCStaking(goCtx context.Context, req *types.MsgCreateBTCStaking) (*types.MsgCreateBTCStakingResponse, error) {
	stakingMsgTx, err := NewBTCTxFromBytes(req.StakingTx.Transaction)
	if err != nil {
		return nil, types.ErrParseBTCTx.Wrap(err.Error())
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	stakingTxHash := stakingMsgTx.TxHash()
	if ms.getBTCStakingRecord(ctx, stakingTxHash) != nil {
		return nil, types.ErrDupBTCTx
	}

	stakingTxHeader := ms.btclcKeeper.GetHeaderByHash(ctx, req.StakingTx.Key.Hash)
	if stakingTxHeader == nil {
		return nil, types.ErrBlkHdrNotFound
	}

	p := ms.GetParams(ctx)
	btcTip := ms.btclcKeeper.GetTipInfo(ctx)
	stakingTxDepth := btcTip.Height - stakingTxHeader.Height
	btclcParams := ms.btclcKeeper.GetBTCNet()
	if err := req.StakingTx.VerifyInclusion(stakingTxHeader.Header, btclcParams.PowLimit); err != nil {
		return nil, types.ErrBTCTxNotIncluded.Wrap(err.Error())
	}
	_, receiver := findReceiver(p.Receivers, req.Receiver)
	if receiver == nil {
		return nil, types.ErrInvalidReceivingAddr.Wrapf("Receiver(%s) not exists", req.Receiver)
	}

	btcReceivingAddr, err := btcutil.DecodeAddress(receiver.Addr, btclcParams)
	if err != nil {
		return nil, types.ErrInvalidReceivingAddr.Wrap(err.Error())
	}
	var mintToAddr []byte
	var receiverAddr []byte
	var btcAmount uint64
	lrzChainId := uint32(ms.evmKeeper.ChainID().Uint64())
	var chainId uint32 = lrzChainId
	if common.IsHexAddress(receiver.EthAddr) {
		signers := req.GetSigners()
		if len(signers) == 0 || !canPerformMint(req.GetSigners()[0], *p) {
			return nil, types.ErrNotInAllowList
		}
		mintToAddr = common.HexToAddress(receiver.EthAddr).Bytes()
		receiverAddr = mintToAddr
		btcAmount, err = ExtractPaymentTo(stakingMsgTx, btcReceivingAddr)
	} else {
		var opReturnMsg []byte
		btcAmount, opReturnMsg, err = ExtractPaymentToWithOpReturnIdAndDust(stakingMsgTx, btcReceivingAddr, p.TxoutDustAmount)
		if err != nil {
			return nil, types.ErrInvalidTransaction.Wrap(err.Error())
		}
		if len(opReturnMsg) == EthAddrLen {
			mintToAddr = opReturnMsg
			receiverAddr = mintToAddr
		} else if len(opReturnMsg) == EthAddrLen+ChainIDLen {
			receiverAddr = opReturnMsg[:EthAddrLen]
			chainId = binary.BigEndian.Uint32(opReturnMsg[EthAddrLen:])
			if chainId != lrzChainId {
				mintToAddr = common.HexToAddress(p.BridgeAddr).Bytes()
			} else {
				mintToAddr = receiverAddr
			}
		} else {
			return nil, types.ErrMintToAddr
		}
	}
	if err != nil || btcAmount == 0 {
		return nil, types.ErrInvalidTransaction
	}
	err = checkBTCTxDepth(stakingTxDepth, btcAmount)
	if err != nil {
		return nil, err
	}

	toMintAmount := sdkmath.NewIntFromUint64(btcAmount).Mul(sdkmath.NewIntFromUint64(SatoshiToStBTCMul))

	coins := []sdk.Coin{
		{
			Denom:  types.NativeTokenDenom,
			Amount: toMintAmount,
		},
	}
	err = ms.bankKeeper.MintCoins(ctx, types.ModuleName, coins)
	if err != nil {
		return nil, types.ErrMintToModule.Wrap(err.Error())
	}
	err = ms.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, mintToAddr, coins)
	if err != nil {
		return nil, types.ErrTransferToAddr.Wrap(err.Error())
	}
	stakingRecord := types.BTCStakingRecord{
		TxHash:       stakingTxHash[:],
		Amount:       btcAmount,
		ReceiverAddr: receiverAddr,
		AgentName:    receiver.Name,
		AgentBtcAddr: receiver.Addr,
		ChainId:      chainId,
	}
	err = ms.addBTCStakingRecord(ctx, &stakingRecord)
	if err != nil {
		return nil, types.ErrRecordStaking.Wrap(err.Error())
	}
	err = ctx.EventManager().EmitTypedEvent(types.NewEventBTCStakingCreated(&stakingRecord))
	if err != nil {
		panic(fmt.Errorf("fail to emit EventBTCStakingCreated : %w", err))
	}
	return &types.MsgCreateBTCStakingResponse{}, nil
}

func (ms msgServer) Burn(goCtx context.Context, req *types.MsgBurnRequest) (*types.MsgBurnResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	btcNetworkParams := ms.btclcKeeper.GetBTCNet()
	btcTargetAddress, err := btcutil.DecodeAddress(req.BtcTargetAddress, btcNetworkParams)
	if err != nil {
		return nil, types.ErrInvalidBurnBtcTargetAddress.Wrap(err.Error())
	}

	amount := sdk.NewCoin(types.NativeTokenDenom, req.Amount)

	signers := req.GetSigners()
	if len(signers) != 1 {
		return nil, types.ErrBurnInvalidSigner
	}
	signer := signers[0]
	balance := ms.bankKeeper.GetBalance(ctx, signer, types.NativeTokenDenom)
	if balance.IsLT(amount) {
		return nil, types.ErrBurnInsufficientBalance
	}

	err = ms.bankKeeper.SendCoinsFromAccountToModule(ctx, signer, types.ModuleName, []sdk.Coin{amount})
	if err != nil {
		return nil, types.ErrBurn.Wrap(err.Error())
	}
	err = ms.bankKeeper.BurnCoins(ctx, types.ModuleName, []sdk.Coin{amount})
	if err != nil {
		return nil, types.ErrBurn.Wrap(err.Error())
	}

	err = ctx.EventManager().EmitTypedEvent(types.NewEventBurnCreated(signer, btcTargetAddress, amount))
	if err != nil {
		return nil, types.ErrEmitEvent.Wrap(err.Error())
	}

	return &types.MsgBurnResponse{}, nil
}

func findReceiver(receivers []*types.Receiver, name string) (int, *types.Receiver) {
	var receiver *types.Receiver = nil
	idx := -1
	for i, r := range receivers {
		if r != nil && r.Name == name {
			idx = i
			receiver = r
			break
		}
	}
	return idx, receiver
}

func (ms msgServer) UpdateParams(goCtx context.Context, req *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if ms.authority != req.Authority {
		return nil, errorsmod.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", ms.authority, req.Authority)
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := req.Params.Validate(); err != nil {
		return nil, err
	}
	if err := ms.SetParams(ctx, &req.Params); err != nil {
		return nil, err
	}
	return &types.MsgUpdateParamsResponse{}, nil
}

func (ms msgServer) AddReceiver(goCtx context.Context, req *types.MsgAddReceiver) (*types.MsgAddReceiverResponse, error) {
	return nil, fmt.Errorf("deprecated, use UpdateParams instead")
}

func (ms msgServer) RemoveReceiver(goCtx context.Context, req *types.MsgRemoveReceiver) (*types.MsgRemoveReceiverResponse, error) {
	return nil, fmt.Errorf("deprecated, use UpdateParams instead")
}
