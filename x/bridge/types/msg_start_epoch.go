package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgStartEpoch = "start_epoch"

var _ sdk.Msg = &MsgStartEpoch{}

func NewMsgStartEpoch(creator string, epochId uint32, info []TSSInfo) *MsgStartEpoch {
	return &MsgStartEpoch{
		Creator: creator,
		EpochId: epochId,
		Info:    info,
	}
}

func (msg *MsgStartEpoch) Route() string {
	return RouterKey
}

func (msg *MsgStartEpoch) Type() string {
	return TypeMsgStartEpoch
}

func (msg *MsgStartEpoch) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Creator)}
}

func (msg *MsgStartEpoch) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgStartEpoch) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if len(msg.Info) == 0 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "TSS info cannot be empty")
	}

	if msg.EpochId == 0 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "epoch ID must be greater than zero")
	}

	return nil
}
