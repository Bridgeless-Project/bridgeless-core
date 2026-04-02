package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSubmitSwapTx = "submit_swap_tx"

var _ sdk.Msg = &MsgSubmitSwapTx{}

func NewMsgSubmitSwapTx(creator string, tx *SwapTransaction, isBridgeTx bool) *MsgSubmitSwapTx {
	return &MsgSubmitSwapTx{
		Creator:    creator,
		Tx:         tx,
		IsBridgeTx: isBridgeTx,
	}
}

func (msg *MsgSubmitSwapTx) Route() string {
	return RouterKey
}

func (msg *MsgSubmitSwapTx) Type() string {
	return TypeMsgSubmitSwapTx
}

func (msg *MsgSubmitSwapTx) GetSigners() []sdk.AccAddress {
	accAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(errorsmod.Wrapf(err, "failed to acc address from bech32 string, given string: %s", msg.Creator))
	}

	return []sdk.AccAddress{accAddress}
}

func (msg *MsgSubmitSwapTx) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSubmitSwapTx) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address: %s", err)
	}

	return nil
}
