package keeper

import (
	"golang.org/x/exp/slices"

	"github.com/Lorenzo-Protocol/lorenzo/v3/x/ccev/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	if err := genState.Validate(); err != nil {
		panic(err)
	}

	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}

	for _, chain := range genState.ChainStates {
		slices.SortFunc(chain.Headers, func(a, b *types.TinyHeader) bool {
			return a.Number < b.Number
		})

		k.setClient(ctx, chain.Client)
		for i, header := range chain.Headers {
			k.setHeader(ctx, chain.Client.ChainId, header)
			if i == len(chain.Headers)-1 {
				k.setLatestNumber(ctx, chain.Client.ChainId, chain.Headers[i].Number)
			}
		}

		for _, contract := range chain.Contracts {
			k.setContract(ctx, contract)
		}
	}
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	var chainStates []*types.ChainState

	clients := k.getAllClients(ctx)
	for _, client := range clients {
		var contractState []*types.Contract

		contracts := k.getAllContracts(ctx, client.ChainId)
		for _, contract := range contracts {
			contractState = append(contractState, contract)
		}

		chainState := &types.ChainState{
			Client:    client,
			Headers:   k.GetAllHeaders(ctx, client.ChainId),
			Contracts: contractState,
		}
		chainStates = append(chainStates, chainState)
	}
	return &types.GenesisState{
		Params:      k.GetParams(ctx),
		ChainStates: chainStates,
	}
}
