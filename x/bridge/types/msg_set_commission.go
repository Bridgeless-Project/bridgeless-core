package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetCommission = "set_commission"

var _ sdk.Msg = &MsgSetCommission{}

func NewMsgSetCommission(creator string, tokenId uint64, amount string) *MsgSetCommission {
	return &MsgSetCommission{
		Creator: creator,
		TokenId: tokenId,
		Amount:  amount,
	}
}

func (msg *MsgSetCommission) Route() string {
	return RouterKey
}

func (msg *MsgSetCommission) Type() string {
	return TypeMsgSetCommission
}

func (msg *MsgSetCommission) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetCommission) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetCommission) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
