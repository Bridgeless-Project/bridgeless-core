package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRemoveReferral = "remove_referral"

var _ sdk.Msg = &MsgRemoveReferral{}

func NewMsgRemoveReferral(creator string, id uint32) *MsgRemoveReferral {
	return &MsgRemoveReferral{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgRemoveReferral) Route() string {
	return RouterKey
}

func (msg *MsgRemoveReferral) Type() string {
	return TypeMsgRemoveReferral
}

func (msg *MsgRemoveReferral) GetSigners() []sdk.AccAddress {
	accAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(errorsmod.Wrapf(err, "failed to acc address from bech32 string, given string: %s", msg.Creator))
	}

	return []sdk.AccAddress{accAddress}
}

func (msg *MsgRemoveReferral) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveReferral) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address: %s", err)
	}

	return nil
}
