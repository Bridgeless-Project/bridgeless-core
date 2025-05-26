package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetParties = "submit_parties"

var _ sdk.Msg = &MsgSetParties{}

func NewMsgSetParties(creator string, partiesList []*Party) *MsgSetParties {
	return &MsgSetParties{
		Creator: creator,
		Parties: partiesList,
	}
}

func (msg *MsgSetParties) Route() string {
	return RouterKey
}

func (msg *MsgSetParties) Type() string {
	return TypeMsgSetParties
}

func (msg *MsgSetParties) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Creator)}
}

func (msg *MsgSetParties) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetParties) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if len(msg.Parties) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "no parties specified")
	}

	return nil
}
