package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRemovePool = "remove_pool"

var _ sdk.Msg = &MsgRemovePool{}

func NewMsgRemovePool(creator string, tokenID string) *MsgRemovePool {
	return &MsgRemovePool{
		Creator: creator,
		TokenId: tokenID,
	}
}

func (msg *MsgRemovePool) Route() string {
	return RouterKey
}

func (msg *MsgRemovePool) Type() string {
	return TypeMsgRemovePool
}

func (msg *MsgRemovePool) GetSigners() []sdk.AccAddress {
	accAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(errorsmod.Wrapf(err, "failed to acc address from bech32 string, given string: %s", msg.Creator))
	}

	return []sdk.AccAddress{accAddress}
}

func (msg *MsgRemovePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemovePool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address: %s", err)
	}

	if len(msg.TokenId) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "token id cannot be empty")
	}

	return nil
}
