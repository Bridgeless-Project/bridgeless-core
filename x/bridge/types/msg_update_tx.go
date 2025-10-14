package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateTransaction = "update_transaction"

var _ sdk.Msg = &MsgUpdateTransaction{}

func NewMsgUpdateTransaction(submitter string, tx Transaction) *MsgUpdateTransaction {
	return &MsgUpdateTransaction{
		Submitter:   submitter,
		Transaction: tx,
	}
}

func (msg *MsgUpdateTransaction) Route() string {
	return RouterKey
}

func (msg *MsgUpdateTransaction) Type() string {
	return TypeMsgUpdateTransaction
}

func (msg *MsgUpdateTransaction) GetSigners() []sdk.AccAddress {
	accAddress, err := sdk.AccAddressFromBech32(msg.Submitter)
	if err != nil {
		panic(errorsmod.Wrapf(err, "failed to acc address from bech32 string, given string: %s", msg.Submitter))
	}

	return []sdk.AccAddress{accAddress}
}

func (msg *MsgUpdateTransaction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateTransaction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Submitter)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid submitter address: %s", err)
	}

	if msg.Transaction.Signature == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "missing signature")
	}

	if msg.Transaction.WithdrawalTxHash != "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "withdrawal tx hash must be empty")
	}

	return nil
}
