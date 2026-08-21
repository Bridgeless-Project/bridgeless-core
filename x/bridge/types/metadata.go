package types

import (
	"errors"

	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func validateTokenInfoMetadata(metadata *TokenInfoMetadata) error {
	if metadata == nil {
		return errors.New("token info metadata is nil")
	}
	if metadata.ChainId == "" {
		return errors.New("chain id cannot be empty")
	}
	if metadata.Address == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "token address is empty")
	}

	return nil
}
