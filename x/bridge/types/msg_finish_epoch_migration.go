package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgFinishEpochMigration = "set_finish_epoch_migration"

var _ sdk.Msg = &MsgFinishEpochMigration{}

func NewMsgFinishEpochMigration(creator string, epochId uint32) *MsgFinishEpochMigration {
	return &MsgFinishEpochMigration{
		Creator: creator,
		EpochId: epochId,
	}
}

func (msg *MsgFinishEpochMigration) Route() string {
	return RouterKey
}

func (msg *MsgFinishEpochMigration) Type() string {
	return TypeMsgFinishEpochMigration
}

func (msg *MsgFinishEpochMigration) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Creator)}
}

func (msg *MsgFinishEpochMigration) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgFinishEpochMigration) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.EpochId == 0 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "epoch ID must be greater than zero")
	}

	return nil
}
