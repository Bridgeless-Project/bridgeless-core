package types

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetReferral = "set_referral"

var _ sdk.Msg = &MsgSetReferral{}

func NewMsgSetReferral(creator string, referral Referral) *MsgSetReferral {
	return &MsgSetReferral{
		Creator:  creator,
		Referral: referral,
	}
}

func (msg *MsgSetReferral) Route() string {
	return RouterKey
}

func (msg *MsgSetReferral) Type() string {
	return TypeMsgSetReferral
}

func (msg *MsgSetReferral) GetSigners() []sdk.AccAddress {
	accAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(errorsmod.Wrapf(err, "failed to acc address from bech32 string, given string: %s", msg.Creator))
	}

	return []sdk.AccAddress{accAddress}
}

func (msg *MsgSetReferral) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetReferral) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address: %s", err)
	}

	if err = validateReferral(&msg.Referral); err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest,
			fmt.Sprintf("invalid referral %d: %s", &msg.Referral.Id, err),
		)
	}

	return nil
}
