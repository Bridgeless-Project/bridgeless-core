package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRemoveTokenInfoMetadata = "remove_token_info_metadata"

var _ sdk.Msg = &MsgRemoveTokenInfoMetadata{}

func NewMsgRemoveTokenInfoMetadata(creator, chain, address string) *MsgRemoveTokenInfoMetadata {
	return &MsgRemoveTokenInfoMetadata{
		Creator: creator,
		ChainId: chain,
		Address: address,
	}
}

func (msg *MsgRemoveTokenInfoMetadata) Route() string {
	return RouterKey
}

func (msg *MsgRemoveTokenInfoMetadata) Type() string {
	return TypeMsgRemoveTokenInfoMetadata
}

func (msg *MsgRemoveTokenInfoMetadata) GetSigners() []sdk.AccAddress {
	accAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(errorsmod.Wrapf(err, "failed to acc address from bech32 string, given string: %s", msg.Creator))
	}

	return []sdk.AccAddress{accAddress}
}

func (msg *MsgRemoveTokenInfoMetadata) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveTokenInfoMetadata) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address: %s", err)
	}

	if msg.ChainId == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "chain id cannot be empty")
	}

	if msg.Address == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "token address is empty")
	}

	return nil
}
