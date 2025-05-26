package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeRemoveTokenInfo = "remove_token_info"

var _ sdk.Msg = &MsgRemoveTokenInfo{}

func NewMsgRemoveTokenInfo(creator string, tokenId uint64, chainId string) *MsgRemoveTokenInfo {
	return &MsgRemoveTokenInfo{
		Creator: creator,
		TokenId: tokenId,
		ChainId: chainId,
	}
}

func (msg *MsgRemoveTokenInfo) Route() string {
	return RouterKey
}

func (msg *MsgRemoveTokenInfo) Type() string {
	return TypeRemoveTokenInfo
}

func (msg *MsgRemoveTokenInfo) GetSigners() []sdk.AccAddress {
	accAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(errorsmod.Wrapf(err, "failed to acc address from bech32 string, given string: %s", msg.Creator))
	}

	return []sdk.AccAddress{accAddress}
}

func (msg *MsgRemoveTokenInfo) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveTokenInfo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address: %s", err)
	}

	if msg.TokenId == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidType, "token id cannot be zero")
	}

	if len(msg.ChainId) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidChainID, "chain id cannot be empty")
	}

	return nil
}
