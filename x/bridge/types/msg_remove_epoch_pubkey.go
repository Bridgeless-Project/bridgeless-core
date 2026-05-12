package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRemoveEpochPubKey = "remove_epoch_pubkey"

var _ sdk.Msg = &MsgRemoveEpochPubKey{}

func NewMsgRemoveEpochPubKey(creator string, epochId uint32) *MsgRemoveEpochPubKey {
	return &MsgRemoveEpochPubKey{
		Creator: creator,
		EpochId: epochId,
	}
}

func (msg *MsgRemoveEpochPubKey) Route() string {
	return RouterKey
}

func (msg *MsgRemoveEpochPubKey) Type() string {
	return TypeMsgRemoveEpochPubKey
}

func (msg *MsgRemoveEpochPubKey) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Creator)}
}

func (msg *MsgRemoveEpochPubKey) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveEpochPubKey) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.EpochId == 0 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "epoch ID must be greater than zero")
	}

	return nil
}
