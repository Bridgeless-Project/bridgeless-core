package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetTokenInfoMetadata = "set_token_info_metadata"

var _ sdk.Msg = &MsgSetTokenInfoMetadata{}

func NewMsgSetTokenInfoMetadata(creator, chain, address string, metadata TokenInfoMetadata) *MsgSetTokenInfoMetadata {
	return &MsgSetTokenInfoMetadata{
		Creator:  creator,
		ChainId:  chain,
		Address:  address,
		Metadata: metadata,
	}
}

func (msg *MsgSetTokenInfoMetadata) Route() string {
	return RouterKey
}

func (msg *MsgSetTokenInfoMetadata) Type() string {
	return TypeMsgSetTokenInfoMetadata
}

func (msg *MsgSetTokenInfoMetadata) GetSigners() []sdk.AccAddress {
	accAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(errorsmod.Wrapf(err, "failed to acc address from bech32 string, given string: %s", msg.Creator))
	}

	return []sdk.AccAddress{accAddress}
}

func (msg *MsgSetTokenInfoMetadata) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetTokenInfoMetadata) ValidateBasic() error {
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
