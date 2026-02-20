package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetEpochSignature = "set_epoch_signature"

var _ sdk.Msg = &MsgSetEpochSignature{}

func NewMsgSetEpochSignature(creator string, signatures []EpochChainSignatures, addresses []EpochBridgeAddress) *MsgSetEpochSignature {
	return &MsgSetEpochSignature{
		Creator:              creator,
		EpochChainSignatures: signatures,
		Addresses:            addresses,
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

	if len(msg.EpochChainSignatures) == 0 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "no signatures provided")
	}
	if len(msg.Addresses) == 0 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "no addresses provided")
	}

	for _, chainSignature := range msg.EpochChainSignatures {
		if chainSignature.AddedSignature == nil && chainSignature.RemovedSignature == nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "added and removed signature cannot be nil")
		}
		if chainSignature.EpochId == 0 {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "epoch ID must be greater than zero")
		}
		if chainSignature.ChainType < 0 {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "epoch ID must be greater or equal than zero")
		}
	}

	for _, address := range msg.Addresses {
		if address.ChainId == "" {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "address must contain chain ID")
		}
		if address.Address == "" {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "address must contain address")
		}
	}

	return nil
}
