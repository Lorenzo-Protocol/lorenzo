package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	// ModuleName is the name of the module
	ModuleName = "plan"

	// StoreKey is the string store representation
	StoreKey = ModuleName

	// RouterKey is the msg router key for the module
	RouterKey = ModuleName
)

var (
	// PlanKey is the key to store the plan in the store
	PlanKey       = []byte{0x01}
	NextNumberKey = []byte{0x02}
)

func KeyPlan(id uint64) []byte {
	bz := sdk.Uint64ToBigEndian(id)
	return append(PlanKey, bz...)
}

func KeyNextNumber() []byte {
	return NextNumberKey
}
