package types

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/multisig/tx"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgSubmitProposal = "submit_proposal"
)

var _ sdk.Msg = &MsgSubmitProposal{}
var _ types.UnpackInterfacesMessage = &MsgSubmitProposal{}

func NewMsgSubmitProposal(
	creator string,
	group string,
	messages []sdk.Msg,
) (*MsgSubmitProposal, error) {
	m := &MsgSubmitProposal{
		Creator: creator,
		Group:   group,
	}

	err := m.SetMsgs(messages)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (msg *MsgSubmitProposal) Route() string {
	return RouterKey
}

func (msg *MsgSubmitProposal) Type() string {
	return TypeMsgSubmitProposal
}

func (msg *MsgSubmitProposal) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(errorsmod.Wrapf(err, "failed to acc address from bech32 string, given string: %s", msg.Creator))
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSubmitProposal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSubmitProposal) SetMsgs(msgs []sdk.Msg) error {
	anys, err := tx.SetMsgs(msgs)
	if err != nil {
		return err
	}
	msg.Messages = anys
	return nil
}

func (msg *MsgSubmitProposal) GetMsgs() ([]sdk.Msg, error) {
	return tx.GetMsgs(msg.Messages, "proposal")
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (msg *MsgSubmitProposal) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	return tx.UnpackInterfaces(unpacker, msg.Messages)
}

func (msg *MsgSubmitProposal) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if _, err := sdk.AccAddressFromBech32(msg.Group); err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid group address (%s)", err)
	}

	msgs, err := msg.GetMsgs()
	if err != nil {
		return err
	}

	if len(msgs) == 0 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "no messages to submit")
	}

	for i, msg := range msgs {
		if err := msg.ValidateBasic(); err != nil {
			return errorsmod.Wrapf(err, "msg %d", i)
		}
	}
	return nil
}
