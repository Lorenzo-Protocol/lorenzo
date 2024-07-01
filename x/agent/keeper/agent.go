package keeper

import (
	"github.com/Lorenzo-Protocol/lorenzo/x/agent/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetAgent retrieves an agent from the Keeper's store based on the given ID.
//
// Parameters:
// - ctx: the SDK context.
// - id: the ID of the agent.
//
// Returns:
// - agent: the agent with the given ID, or an empty agent if it does not exist.
// - bool: true if the agent exists, false otherwise.
func (k Keeper) GetAgent(ctx sdk.Context, id uint64) (types.Agent, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyAgent(id))
	if bz == nil {
		return types.Agent{}, false
	}

	var agent types.Agent
	k.cdc.MustUnmarshal(bz, &agent)
	return agent, true
}

// GetNextNumber retrieves the next number from the Keeper's store.
//
// Parameters:
// - ctx: the SDK context.
//
// Returns:
// - uint64: the next number, or 1 if it is not found in the store.
func (k Keeper) GetNextNumber(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyNextNumber())
	if bz == nil {
		return 1
	}

	return sdk.BigEndianToUint64(bz)
}

// GetAdmin retrieves the admin address from the Keeper's store.
//
// Parameters:
// - ctx: the SDK context.
//
// Returns:
// - sdk.AccAddress: the admin address, or nil if it is not found in the store.
func (k Keeper) GetAdmin(ctx sdk.Context) sdk.AccAddress {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyAdmin())
	if bz == nil {
		return nil
	}
	return sdk.AccAddress(bz)
}

// Allowed checks if the given address is authorized.
//
// Parameters:
// - ctx: the SDK context.
// - address: the address to check.
//
// Returns:
// - bool: true if the address is authorized, false otherwise.
func (k Keeper) Allowed(ctx sdk.Context, address sdk.AccAddress) bool {
	return k.GetAdmin(ctx).Equals(address)
}

// GetAgentsPrefixStore returns the store for the agents
func (k Keeper) GetAgentsPrefixStore(ctx sdk.Context) prefix.Store {
	store := ctx.KVStore(k.storeKey)
	return prefix.NewStore(store, types.AgentKey)
}

func (k Keeper) getAgents(ctx sdk.Context) (agents []types.Agent) {
	store := ctx.KVStore(k.storeKey)

	it := sdk.KVStorePrefixIterator(store, types.AgentKey)
	defer it.Close() //nolint:errcheck

	for ; it.Valid(); it.Next() {
		var agent types.Agent
		k.cdc.MustUnmarshal(it.Value(), &agent)
		agents = append(agents, agent)
	}
	return agents
}

// addAgent adds a new agent to the Keeper's store.
//
// Parameters:
// - ctx: the SDK context.
// - name: the name of the agent.
// - btcReceivingAddress: the Bitcoin receiving address of the agent.
// - ethAddr: the Ethereum address of the agent.
// - description: the description of the agent.
// - url: the URL of the agent.
//
// Returns:
// - uint64: the ID of the newly added agent.
func (k Keeper) addAgent(ctx sdk.Context, name, btcReceivingAddress, ethAddr, description, url string) uint64 {
	id := k.GetNextNumber(ctx)
	agent := types.Agent{
		Id:                  id,
		Name:                name,
		BtcReceivingAddress: btcReceivingAddress,
		EthAddr:             ethAddr,
		Description:         description,
		Url:                 url,
	}
	k.setAgent(ctx, agent)
	k.setNextNumber(ctx, id+1)
	return id
}

// setNextNumber sets the next number in the Keeper's store.
//
// ctx - the SDK context.
// number - the number to be set.
func (k Keeper) setNextNumber(ctx sdk.Context, number uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyNextNumber(), sdk.Uint64ToBigEndian(number))
}

func (k Keeper) setAgent(ctx sdk.Context, agent types.Agent) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&agent)
	store.Set(types.KeyAgent(agent.Id), bz)
}

func (k Keeper) removeAgent(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyAgent(id))
}

func (k Keeper) setAdmin(ctx sdk.Context, admin sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyAdmin(), admin.Bytes())
}
