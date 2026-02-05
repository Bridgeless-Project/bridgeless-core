package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetEpochPubkey = "set_epoch_pubkey"

var _ sdk.Msg = &MsgSetEpochPubkey{}

func NewMsgSetEpochPubkey(creator string, pubkey string, epochId uint32) *MsgSetEpochPubkey {
	return &MsgSetEpochPubkey{
		Creator: creator,
		Pubkey:  pubkey,
		EpochId: epochId,
	}
}

func (msg *MsgSetEpochPubkey) Route() string {
	return RouterKey
}

func (msg *MsgSetEpochPubkey) Type() string {
	return TypeMsgSetEpochPubkey
}

func (msg *MsgSetEpochPubkey) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Creator)}
}

func (msg *MsgSetEpochPubkey) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetEpochPubkey) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.Pubkey == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "pubkey cannot be empty")
	}

	if msg.EpochId == 0 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "epoch ID must be greater than zero")
	}

	return nil
}
