package types

import (
	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func validatePool(pool *SwapPool) error {
	if pool == nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "pool cannot be nil")
	}

	if len(pool.TokenId) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "pool token id cannot be empty")
	}

	if len(pool.Address) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "pool address cannot be empty")
	}

	return nil
}
