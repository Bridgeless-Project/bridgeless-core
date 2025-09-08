package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRemoveTransaction = "remove_transaction"

var _ sdk.Msg = &MsgRemoveTransaction{}

func NewMsgRemoveTransaction(submitter string, depositChainId, depositTxHash string, depositTxIndex uint64) *MsgRemoveTransaction {
	return &MsgRemoveTransaction{
		Creator:        submitter,
		DepositChainId: depositChainId,
		DepositTxHash:  depositTxHash,
		DepositTxIndex: depositTxIndex,
	}
}

func (msg *MsgRemoveTransaction) Route() string {
	return RouterKey
}

func (msg *MsgRemoveTransaction) Type() string {
	return TypeMsgRemoveTransaction
}

func (msg *MsgRemoveTransaction) GetSigners() []sdk.AccAddress {
	accAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(errorsmod.Wrapf(err, "failed to acc address from bech32 string, given string: %s", msg.Creator))
	}

	return []sdk.AccAddress{accAddress}
}

func (msg *MsgRemoveTransaction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveTransaction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid submitter address: %s", err)
	}

	if len(msg.DepositChainId) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "deposit chain id and withdrawal id cannot be empty")
	}

	return nil
}
