package keeper

import (
	"github.com/Lorenzo-Protocol/lorenzo/v3/x/ccev/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

// UploadContract uploads a cross chain contract and saves it to the store.
// If the contract already exists, it will be overwritten.
func (k Keeper) UploadContract(
	ctx sdk.Context,
	chainID uint32,
	address string,
	eventName string,
	abi []byte,
) {
	contract := &types.Contract{
		ChainId:   chainID,
		Address:   address,
		EventName: eventName,
		Abi:       abi,
	}
	k.setContract(ctx, contract)
}

func (k Keeper) setContract(ctx sdk.Context, contract *types.Contract) {
	store := k.clientStore(ctx, contract.ChainId)
	store.Set(
		types.KeyCrossChainContract(common.HexToAddress(contract.Address)),
		k.cdc.MustMarshal(contract),
	)
}

func (k Keeper) getContract(
	ctx sdk.Context,
	chainID uint32,
	address common.Address,
) *types.Contract {
	store := k.clientStore(ctx, chainID)
	bz := store.Get(types.KeyCrossChainContract(address))
	if bz == nil {
		return nil
	}
	var contract types.Contract
	k.cdc.MustUnmarshal(bz, &contract)
	return &contract
}

func (k Keeper) getAllContracts(ctx sdk.Context,chainID uint32) (contracts []*types.Contract) {
	store := k.clientStore(ctx, chainID)

	it := sdk.KVStorePrefixIterator(store, types.KeyPrefixCrossChainContract)
	defer it.Close() //nolint:errcheck

	for ; it.Valid(); it.Next() {
		var contract types.Contract
		k.cdc.MustUnmarshal(it.Value(), &contract)
		contracts = append(contracts, &contract)
	}
	return
}
