package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	bridgeTypes "github.com/hyle-team/bridgeless-core/v12/types"
)

type groupMsg interface {
	GetCreator() string
	GetMembers() []string
	GetThreshold() uint64
}

func validateGroupMessage(msg groupMsg) error {
	if _, err := sdk.AccAddressFromBech32(msg.GetCreator()); err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	members := msg.GetMembers()
	index := make(map[string]struct{}, len(members))

	for _, member := range members {
		if _, err := sdk.AccAddressFromBech32(member); err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid member address (%s)", err)
		}
		if _, exists := index[member]; exists {
			return errorsmod.Wrapf(bridgeTypes.ErrDuplicatedValue, "duplicate member address: %s", member)
		}
		index[member] = struct{}{}
	}

	threshold := msg.GetThreshold()
	if threshold == 0 || threshold > uint64(len(members)) {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "threshold must be greater than 0 and less or equal to the number of members")
	}

	return nil
}
