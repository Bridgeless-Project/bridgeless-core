package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetEpochSignature = "set_epoch_signature"

var _ sdk.Msg = &MsgSetEpochSignature{}

func NewMsgSetEpochSignature(creator string, data EpochChainSignatures) *MsgSetEpochSignature {
	return &MsgSetEpochSignature{
		Creator:              creator,
		EpochChainSignatures: data,
	}
}

func (msg *MsgSetEpochSignature) Route() string {
	return RouterKey
}

func (msg *MsgSetEpochSignature) Type() string {
	return TypeMsgSetEpochSignature
}

func (msg *MsgSetEpochSignature) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Creator)}
}

func (msg *MsgSetEpochSignature) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetEpochSignature) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.EpochChainSignatures.ChainId == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "chain ID must be greater than zero")
	}

	if msg.EpochChainSignatures.EpochId == 0 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "epoch ID must be greater than zero")
	}

	if msg.EpochChainSignatures.AddedSignature == nil && msg.EpochChainSignatures.RemovedSignature == nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "added and removed signature cannot be nil")
	}

	return nil
}
