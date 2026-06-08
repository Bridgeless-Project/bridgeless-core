package types

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgProcessSystemWithdrawal = "process_system_transaction"

var _ sdk.Msg = &MsgProcessSystemWithdrawal{}

func NewMsgProcessSystemWithdrawal(creator string, epochId uint32, withdrawals ...SystemWithdrawal) *MsgProcessSystemWithdrawal {
	return &MsgProcessSystemWithdrawal{
		Creator:    creator,
		Withdrawal: withdrawals,
		EpochId:    epochId,
	}
}

func (msg *MsgProcessSystemWithdrawal) Route() string {
	return RouterKey
}

func (msg *MsgProcessSystemWithdrawal) Type() string {
	return TypeMsgProcessSystemWithdrawal
}

func (msg *MsgProcessSystemWithdrawal) GetSigners() []sdk.AccAddress {
	accAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(errorsmod.Wrapf(err, "failed to acc address from bech32 string, given string: %s", msg.Creator))
	}

	return []sdk.AccAddress{accAddress}
}

func (msg *MsgProcessSystemWithdrawal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgProcessSystemWithdrawal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address: %s", err)
	}

	if len(msg.Withdrawal) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "withdrawals cannot be empty")
	}

	for i, withdrawal := range msg.Withdrawal {
		if err = validateSystemWithdrawal(&withdrawal); err != nil {
			return errorsmod.Wrap(sdkerrors.ErrInvalidRequest,
				fmt.Sprintf("invalid system withdrawal at index %d: %s", i, err),
			)
		}
	}

	return nil
}
