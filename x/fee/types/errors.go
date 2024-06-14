package types

import (
	errorsmod "cosmossdk.io/errors"
)

// ErrDuplicateMsg is returned if duplicate msgs are provided
var ErrDuplicateMsg = errorsmod.Register(ModuleName, 2, "duplicate msg")
