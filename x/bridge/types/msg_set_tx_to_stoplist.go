package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeAddTxToStopList = "add_tx_to_stop_list"

var _ sdk.Msg = &MsgAddTxToStopList{}

func NewMsgAddTxToStopList(creator string, tx *Transaction) *MsgAddTxToStopList {
	return &MsgAddTxToStopList{
		Creator:     creator,
		Transaction: tx,
	}
}

func (msg *MsgAddTxToStopList) Route() string {
	return RouterKey
}

func (msg *MsgAddTxToStopList) Type() string {
	return TypeAddTxToStopList
}

func (msg *MsgAddTxToStopList) GetSigners() []sdk.AccAddress {
	accAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(errorsmod.Wrapf(err, "failed to acc address from bech32 string, given string: %s", msg.Creator))
	}

	return []sdk.AccAddress{accAddress}
}

func (msg *MsgAddTxToStopList) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddTxToStopList) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address: %s", err)
	}

	err = validateTransaction(msg.Transaction)
	if err != nil {
		return errorsmod.Wrapf(err, "invalid transaction")
	}

	return nil
}
