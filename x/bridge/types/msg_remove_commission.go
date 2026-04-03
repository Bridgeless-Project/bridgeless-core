package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRemoveCommission = "remove_commission"

var _ sdk.Msg = &MsgRemoveCommission{}

func NewMsgRemoveCommission(creator string, tokenId uint64, amount string) *MsgRemoveCommission {
	return &MsgRemoveCommission{
		Creator: creator,
		TokenId: tokenId,
		Amount:  amount,
	}
}

func (msg *MsgRemoveCommission) Route() string {
	return RouterKey
}

func (msg *MsgRemoveCommission) Type() string {
	return TypeMsgRemoveCommission
}

func (msg *MsgRemoveCommission) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemoveCommission) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveCommission) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
