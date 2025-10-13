package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRemoveReferralRewards = "remove_referral_rewards"

var _ sdk.Msg = &MsgRemoveReferralRewards{}

func NewMsgRemoveReferralRewards(creator string, referralId uint32, tokenId uint64) *MsgRemoveReferralRewards {
	return &MsgRemoveReferralRewards{
		Creator:    creator,
		ReferrerId: referralId,
		TokenId:    tokenId,
	}
}

func (msg *MsgRemoveReferralRewards) Route() string {
	return RouterKey
}

func (msg *MsgRemoveReferralRewards) Type() string {
	return TypeMsgRemoveReferralRewards
}

func (msg *MsgRemoveReferralRewards) GetSigners() []sdk.AccAddress {
	accAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(errorsmod.Wrapf(err, "failed to acc address from bech32 string, given string: %s", msg.Creator))
	}

	return []sdk.AccAddress{accAddress}
}

func (msg *MsgRemoveReferralRewards) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveReferralRewards) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address: %s", err)
	}

	if msg.ReferrerId == 0 {
		return errorsmod.Wrapf(ErrReferralIdMustBePositive, "referrer id cannot be 0")
	}

	if msg.TokenId == 0 {
		return errorsmod.Wrapf(ErrTokenIdMustBePositive, "token id cannot be 0")
	}

	return nil
}
