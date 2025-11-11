package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeRemoveTxStopList = "remove_tx_from_stop_list"

var _ sdk.Msg = &MsgRemoveTxFromStopList{}

func NewMsgRemoveTxFromStopList(creator string, nonce uint64, txHash, chainId string) *MsgRemoveTxFromStopList {
	return &MsgRemoveTxFromStopList{
		ChainId: chainId,
		TxHash:  txHash,
		TxNonce: nonce,
		Creator: creator,
	}
}

func (msg *MsgRemoveTxFromStopList) Route() string {
	return RouterKey
}

func (msg *MsgRemoveTxFromStopList) Type() string {
	return TypeRemoveTxStopList
}

func (msg *MsgRemoveTxFromStopList) GetSigners() []sdk.AccAddress {
	accAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(errorsmod.Wrapf(err, "failed to acc address from bech32 string, given string: %s", msg.Creator))
	}

	return []sdk.AccAddress{accAddress}
}

func (msg *MsgRemoveTxFromStopList) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveTxFromStopList) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address: %s", err)
	}

	return nil
}
