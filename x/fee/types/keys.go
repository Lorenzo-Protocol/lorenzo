package types

const (
	// ModuleName defines the module name
	ModuleName = "fee"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName
)

var (
	// ParamsKey defines the key to store the Params object
	ParamsKey           = []byte{0x01} // key for params
)