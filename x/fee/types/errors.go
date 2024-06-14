package types

import (
	errorsmod "cosmossdk.io/errors"
)

var (
	// ErrDuplicateMsg is returned if duplicate msgs are provided
	ErrDuplicateMsg = errorsmod.Register(ModuleName, 2, "duplicate msg")
)
