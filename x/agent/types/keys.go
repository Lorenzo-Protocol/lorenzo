package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "agent"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	DoNotModifyDesc = "[do-not-modify]"
)

var (
	AgentKey      = []byte{0x01}
	NextNumberKey = []byte{0x02}
	AdminKey      = []byte{0x03}
)

func KeyAgent(id uint64) []byte {
	bz := sdk.Uint64ToBigEndian(id)
	return append(AgentKey, bz...)
}

func KeyNextNumber() []byte {
	return NextNumberKey
}

func KeyAdmin() []byte {
	return AdminKey
}
