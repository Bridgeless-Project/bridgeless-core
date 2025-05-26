package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetTssThreshold = "set_tss_threshold"

var _ sdk.Msg = &MsgSetTssThreshold{}

func NewMsgSetTssThreshold(submitter string, threshold uint32) *MsgSetTssThreshold {
	return &MsgSetTssThreshold{
		Creator:   submitter,
		Threshold: threshold,
	}
}

func (msg *MsgSetTssThreshold) Route() string {
	return RouterKey
}

func (msg *MsgSetTssThreshold) Type() string {
	return TypeMsgSetTssThreshold
}

func (msg *MsgSetTssThreshold) GetSigners() []sdk.AccAddress {
	accAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(errorsmod.Wrapf(err, "failed to acc address from bech32 string, given string: %s", msg.Creator))
	}

	return []sdk.AccAddress{accAddress}
}

func (msg *MsgSetTssThreshold) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetTssThreshold) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid submitter address: %s", err)

	}
	return nil
}
