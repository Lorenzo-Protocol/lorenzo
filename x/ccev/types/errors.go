package types

import (
	errorsmod "cosmossdk.io/errors"
)

var (
	// ErrHeaderNotFound is returned when header is not found
	ErrHeaderNotFound = errorsmod.Register(ModuleName, 2, "header not found")

	// ErrInvalidHeader is returned when header is not valid
	ErrInvalidHeader = errorsmod.Register(ModuleName, 3, "invalid header")
	// ErrUnauthorized is returned when tx is not authorized
	ErrUnauthorized = errorsmod.Register(ModuleName, 4, "tx unauthorized")
	// ErrInvalidProof is returned when proof is not valid
	ErrInvalidProof = errorsmod.Register(ModuleName, 5, "invalid proof")
	// ErrInvalidTransaction is returned when transaction is not valid
	ErrInvalidTransaction = errorsmod.Register(ModuleName, 6, "invalid transaction")
	// ErrInvalidEvent is returned when event is not valid
	ErrInvalidEvent = errorsmod.Register(ModuleName, 7, "invalid event")
	// ErrDuplicateClient is returned when client is already registered
	ErrDuplicateClient = errorsmod.Register(ModuleName, 8, "duplicate client")
	// ErrNotFoundClient is returned when client is not found
	ErrNotFoundClient = errorsmod.Register(ModuleName, 9, "client not found")
	// ErrNotFoundContract is returned when contract is not found
	ErrNotFoundContract = errorsmod.Register(ModuleName, 10, "contract not found")
	// ErrInvalidClient is returned when client is not valid
	ErrInvalidClient = errorsmod.Register(ModuleName, 11, "invalid client")
	// ErrInvalidContract is returned when contract is not valid
	ErrInvalidContract = errorsmod.Register(ModuleName, 12, "invalid contract")
	// ErrDuplicateHeader is returned when header is already registered
	ErrDuplicateHeader = errorsmod.Register(ModuleName, 13, "duplicate header")
	// ErrDuplicateAddress is returned when address is already registered
	ErrDuplicateAddress = errorsmod.Register(ModuleName, 14, "duplicate address")
)
