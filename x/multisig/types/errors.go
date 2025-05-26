package types

import errorsmod "cosmossdk.io/errors"

// DONTCOVER

// x/multisig module sentinel errors
var (
	ErrInvalidVoteOption = errorsmod.Register(ModuleName, 2001, "invalid vote option")
)
