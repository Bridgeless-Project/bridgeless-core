package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/swap module sentinel errors
var (
	ErrPoolNotFound     = errorsmod.Register(ModuleName, 1100, "pool not found")
	ErrSwapNotFound     = errorsmod.Register(ModuleName, 1101, "swap not found")
	ErrPermissionDenied = errorsmod.Register(ModuleName, 1102, "permission denied")
	ErrAlreadyExists    = errorsmod.Register(ModuleName, 1103, "entity already exists")
)
