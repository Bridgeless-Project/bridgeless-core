package types

import (
	"math/big"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateCommission = "update_commission"

var _ sdk.Msg = &MsgUpdateCommission{}

func NewMsgUpdateCommission(creator string, tokenId uint64, amount string) *MsgUpdateCommission {
	return &MsgUpdateCommission{
		Creator: creator,
		TokenId: tokenId,
		Amount:  amount,
	}
}

func (msg *MsgUpdateCommission) Route() string {
	return RouterKey
}

func (msg *MsgUpdateCommission) Type() string {
	return TypeMsgUpdateCommission
}

func (msg *MsgUpdateCommission) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateCommission) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateCommission) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	amount, ok := new(big.Int).SetString(msg.Amount, 10)
	if !ok || amount.Sign() < 0 {
		return errorsmod.Wrap(ErrInvalidCommission, "amount must be a non-negative 18-decimal integer")
	}
	return nil
}
