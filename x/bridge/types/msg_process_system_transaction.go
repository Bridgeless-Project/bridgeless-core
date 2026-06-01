package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgProcessSystemWithdrawal = "process_system_transaction"

var _ sdk.Msg = &MsgProcessSystemWithdrawal{}

func NewMsgProcessSystemWithdrawal(creator string, withdrawals ...SystemWithdrawal) *MsgProcessSystemWithdrawal {
	return &MsgProcessSystemWithdrawal{
		Creator:    creator,
		Withdrawal: withdrawals,
	}
}

func (msg *MsgProcessSystemWithdrawal) Route() string {
	return RouterKey
}

func (msg *MsgProcessSystemWithdrawal) Type() string {
	return TypeMsgProcessSystemWithdrawal
}

func (msg *MsgProcessSystemWithdrawal) GetSigners() []sdk.AccAddress {
	accAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(errorsmod.Wrapf(err, "failed to acc address from bech32 string, given string: %s", msg.Creator))
	}

	return []sdk.AccAddress{accAddress}
}

func (msg *MsgProcessSystemWithdrawal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgProcessSystemWithdrawal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address: %s", err)
	}

	return nil
}
