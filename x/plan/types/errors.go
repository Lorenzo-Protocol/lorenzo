package types

import errorsmod "cosmossdk.io/errors"

var (
	ErrABIPack      = errorsmod.Register(ModuleName, 2, "contract ABI pack failed")
	ErrABIUnpack    = errorsmod.Register(ModuleName, 3, "contract ABI unpack failed")
	ErrUnauthorized = errorsmod.Register(ModuleName, 4, "unauthorized address")
)
