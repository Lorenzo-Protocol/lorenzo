package keeper

import (
	"fmt"

	bbn "github.com/Lorenzo-Protocol/lorenzo/types"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/wire"
	"github.com/cometbft/cometbft/libs/log"

	"github.com/Lorenzo-Protocol/lorenzo/x/btclightclient/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Keeper struct {
		cdc       codec.BinaryCodec
		storeKey  storetypes.StoreKey
		hooks     types.BTCLightClientHooks
		btcConfig bbn.BtcConfig
		bl        *types.BtcLightClient
		authority string
	}
)

var _ types.BtcChainReadStore = (*headersState)(nil)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	btcConfig bbn.BtcConfig,
	authority string,
) Keeper {
	bl := types.NewBtcLightClientFromParams(btcConfig.NetParams())

	return Keeper{
		cdc:       cdc,
		storeKey:  storeKey,
		hooks:     nil,
		btcConfig: btcConfig,
		bl:        bl,
		authority: authority,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// SetHooks sets the btclightclient hooks
func (k *Keeper) SetHooks(bh types.BTCLightClientHooks) *Keeper {
	if k.hooks != nil {
		panic("cannot set btclightclient hooks twice")
	}
	k.hooks = bh

	return k
}

func (k Keeper) insertHeaders(
	ctx sdk.Context,
	headers []*wire.BlockHeader,
) error {

	headerState := k.headersState(ctx)

	result, err := k.bl.InsertHeaders(
		headerState,
		headers,
	)

	if err != nil {
		return err
	}

	// if we have rollback, first delete all headers up to the rollback point
	if result.RollbackInfo != nil {
		// roll back to the height
		headerState.rollBackHeadersUpTo(result.RollbackInfo.HeaderToRollbackTo.Height)
		// trigger rollback event
		k.triggerRollBack(ctx, result.RollbackInfo.HeaderToRollbackTo)
	}

	for _, header := range result.HeadersToInsert {
		h := header
		headerState.insertHeader(h)
		k.triggerHeaderInserted(ctx, h)
		k.triggerRollForward(ctx, h)
	}
	return nil
}

func (k Keeper) InsertHeaders(ctx sdk.Context, headers []bbn.BTCHeaderBytes) error {
	if len(headers) == 0 {
		return types.ErrEmptyMessage
	}

	blockHeaders := make([]*wire.BlockHeader, len(headers))
	for i, header := range headers {
		blockHeaders[i] = header.ToBlockHeader()
	}

	return k.insertHeaders(ctx, blockHeaders)
}

func (k Keeper) UpdateFeeRate(ctx sdk.Context, feeRate uint64) error {
	headerState := k.headersState(ctx)
	headerState.updateFeeRate(feeRate)
	k.triggerFeeRateUpdated(ctx, feeRate)
	return nil
}

// BlockHeight returns the height of the provided header
func (k Keeper) BlockHeight(ctx sdk.Context, headerHash *bbn.BTCHeaderHashBytes) (uint64, error) {
	if headerHash == nil {
		return 0, types.ErrEmptyMessage
	}

	headerInfo, err := k.headersState(ctx).GetHeaderByHash(headerHash)

	if err != nil {
		return 0, err
	}

	return headerInfo.Height, nil
}

// MainChainDepth returns the depth of the header in the main chain, or error if it does not exists
func (k Keeper) MainChainDepth(ctx sdk.Context, headerHashBytes *bbn.BTCHeaderHashBytes) (uint64, error) {
	if headerHashBytes == nil {
		return 0, types.ErrEmptyMessage
	}
	// Retrieve the header. If it does not exist, return an error
	headerInfo, err := k.headersState(ctx).GetHeaderByHash(headerHashBytes)
	if err != nil {
		return 0, err
	}
	// Retrieve the tip
	tipInfo := k.headersState(ctx).GetTip()

	// sanity check, to avoid silent error if something is wrong.
	if tipInfo.Height < headerInfo.Height {
		// panic, as tip should always be higher than the header than every header
		panic("tip height is less than header height")
	}

	headerDepth := tipInfo.Height - headerInfo.Height
	return headerDepth, nil
}

func (k Keeper) GetTipInfo(ctx sdk.Context) *types.BTCHeaderInfo {
	return k.headersState(ctx).GetTip()
}

// GetHeaderByHash returns header with given hash, if it does not exists returns nil
func (k Keeper) GetHeaderByHash(ctx sdk.Context, hash *bbn.BTCHeaderHashBytes) *types.BTCHeaderInfo {
	info, err := k.headersState(ctx).GetHeaderByHash(hash)

	if err != nil {
		return nil
	}

	return info
}

// GetHeaderByHeight returns header with given height from main chain, returns nil if such header is not found
func (k Keeper) GetHeaderByHeight(ctx sdk.Context, height uint64) *types.BTCHeaderInfo {
	header, err := k.headersState(ctx).GetHeaderByHeight(height)

	if err != nil {
		return nil
	}

	return header
}

// GetMainChainFrom returns the current canonical chain from the given height up to the tip
// If the height is higher than the tip, it returns an empty slice
// If startHeight is 0, it returns the entire main chain
func (k Keeper) GetMainChainFrom(ctx sdk.Context, startHeight uint64) []*types.BTCHeaderInfo {
	headers := make([]*types.BTCHeaderInfo, 0)
	accHeaderFn := func(header *types.BTCHeaderInfo) bool {
		headers = append(headers, header)
		return false
	}
	k.headersState(ctx).IterateForwardHeaders(startHeight, accHeaderFn)
	return headers
}

// GetMainChainUpTo returns the current canonical chain as a collection of block headers
// starting from the tip and ending on the header that has `depth` distance from it.
func (k Keeper) GetMainChainUpTo(ctx sdk.Context, depth uint64) []*types.BTCHeaderInfo {
	headers := make([]*types.BTCHeaderInfo, 0)

	var currentDepth = uint64(0)
	accHeaderFn := func(header *types.BTCHeaderInfo) bool {
		// header header is at depth 0.
		if currentDepth > depth {
			return true
		}

		headers = append(headers, header)
		currentDepth++
		return false
	}

	k.headersState(ctx).IterateReverseHeaders(accHeaderFn)

	return headers
}

// GetMainChainReverse Retrieves whole header chain in reverse order
func (k Keeper) GetMainChainReverse(ctx sdk.Context) []*types.BTCHeaderInfo {
	headers := make([]*types.BTCHeaderInfo, 0)
	accHeaderFn := func(header *types.BTCHeaderInfo) bool {
		headers = append(headers, header)
		return false
	}
	k.headersState(ctx).IterateReverseHeaders(accHeaderFn)
	return headers
}

func (k Keeper) GetBTCNet() *chaincfg.Params {
	return k.btcConfig.NetParams()
}
