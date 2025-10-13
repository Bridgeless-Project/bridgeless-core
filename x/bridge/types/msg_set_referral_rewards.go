package types

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetReferralRewards = "set_referral_rewards"

var _ sdk.Msg = &MsgSetReferralRewards{}

func NewMsgSetReferralRewards(creator string, rewards ReferralRewards) *MsgSetReferralRewards {
	return &MsgSetReferralRewards{
		Creator: creator,
		Rewards: rewards,
	}
}

func (msg *MsgSetReferralRewards) Route() string {
	return RouterKey
}

func (msg *MsgSetReferralRewards) Type() string {
	return TypeMsgSetReferralRewards
}

func (msg *MsgSetReferralRewards) GetSigners() []sdk.AccAddress {
	accAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(errorsmod.Wrapf(err, "failed to acc address from bech32 string, given string: %s", msg.Creator))
	}

	return []sdk.AccAddress{accAddress}
}

func (msg *MsgSetReferralRewards) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetReferralRewards) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address: %s", err)
	}

	if err = validateReferralRewards(&msg.Rewards); err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest,
			fmt.Sprintf("invalid referral rewards %d: %s", &msg.Rewards.ReferralId, err),
		)
	}

	return nil
}
