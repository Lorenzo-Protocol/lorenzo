package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "agent"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	RouterKey = ModuleName

	DoNotModifyDesc = "[do-not-modify]"
)

var (
	ParamsKey     = []byte{0x01}
	AgentKey      = []byte{0x02}
	NextNumberKey = []byte{0x03}
)

func KeyAgent(id uint64) []byte {
	bz := sdk.Uint64ToBigEndian(id)
	return append(AgentKey, bz...)
}

func KeyNextNumber() []byte {
	return NextNumberKey
}
