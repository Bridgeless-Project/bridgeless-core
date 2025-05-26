package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeMsgChangeGroup = "change_group"
)

var _ sdk.Msg = &MsgChangeGroup{}

func NewMsgChangeGroup(
	creator string,
	group string,
	members []string,
	threshold uint64,
) *MsgChangeGroup {
	return &MsgChangeGroup{
		Creator:   creator,
		Group:     group,
		Members:   members,
		Threshold: threshold,
	}
}

func (msg *MsgChangeGroup) Route() string {
	return RouterKey
}

func (msg *MsgChangeGroup) Type() string {
	return TypeMsgChangeGroup
}

func (msg *MsgChangeGroup) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(errorsmod.Wrapf(err, "failed to acc address from bech32 string, given string: %s", msg.Creator))
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgChangeGroup) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgChangeGroup) ValidateBasic() error {
	return validateGroupMessage(msg)
}
