package types

import (
	errorsmod "cosmossdk.io/errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSubmitTransactions = "submit_transactions"

var _ sdk.Msg = &MsgSubmitTransactions{}

func NewMsgSubmitTransactions(submitter string, transactions ...Transaction) *MsgSubmitTransactions {
	return &MsgSubmitTransactions{
		Submitter:    submitter,
		Transactions: transactions,
	}
}

func (msg *MsgSubmitTransactions) Route() string {
	return RouterKey
}

func (msg *MsgSubmitTransactions) Type() string {
	return TypeMsgSubmitTransactions
}

func (msg *MsgSubmitTransactions) GetSigners() []sdk.AccAddress {
	accAddress, err := sdk.AccAddressFromBech32(msg.Submitter)
	if err != nil {
		panic(errorsmod.Wrapf(err, "failed to acc address from bech32 string, given string: %s", msg.Submitter))
	}

	return []sdk.AccAddress{accAddress}
}

func (msg *MsgSubmitTransactions) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSubmitTransactions) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Submitter)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid submitter address: %s", err)
	}

	if len(msg.Transactions) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "transactions cannot be empty")
	}

	for _, tx := range msg.Transactions {
		if len(tx.DepositChainId) == 0 || len(tx.WithdrawalChainId) == 0 {
			return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "deposit chain id and withdrawal id cannot be empty")
		}
		if tx.WithdrawalChainId == tx.DepositChainId {
			return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "withdrawal chain id cannot be the same as deposit chain id ")
		}
		if err = validateTransaction(&tx); err != nil {
			return errorsmod.Wrap(sdkerrors.ErrInvalidRequest,
				fmt.Sprintf("invalid transaction %s: %s", TransactionId(&tx), err),
			)
		}
	}

	return nil
}
