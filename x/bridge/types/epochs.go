package types

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func validateEpoch(epoch *Epoch) error {
	if epoch.StartBlock > epoch.EndBlock {
		return errorsmod.Wrap(sdkerrors.ErrInvalidHeight, fmt.Sprintf("invalid epoch startBlock: %d", epoch.StartBlock))
	}

	return nil
}
