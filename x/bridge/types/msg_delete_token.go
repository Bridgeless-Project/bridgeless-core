package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeleteToken = "delete_token"

var _ sdk.Msg = &MsgDeleteToken{}

func NewMsgDeleteToken(creator string, tokenId uint64) *MsgDeleteToken {
	return &MsgDeleteToken{
		Creator: creator,
		TokenId: tokenId,
	}
}

func (msg *MsgDeleteToken) Route() string {
	return RouterKey
}

func (msg *MsgDeleteToken) Type() string {
	return TypeMsgDeleteToken
}

func (msg *MsgDeleteToken) GetSigners() []sdk.AccAddress {
	accAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{accAddress}
}

func (msg *MsgDeleteToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address: %s", err)
	}

	return nil
}
