package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDistributeFees = "distribute_fees"

var _ sdk.Msg = &MsgDistributeFees{}

func NewMsgDistributeFees(creator string, epochId uint32) *MsgDistributeFees {
	return &MsgDistributeFees{
		Creator: creator,
		EpochId: epochId,
	}
}

func (msg *MsgDistributeFees) Route() string {
	return RouterKey
}

func (msg *MsgDistributeFees) Type() string {
	return TypeMsgDistributeFees
}

func (msg *MsgDistributeFees) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Creator)}
}

func (msg *MsgDistributeFees) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDistributeFees) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	
	return nil
}
