package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddPool = "add_pool"

var _ sdk.Msg = &MsgAddPool{}

func NewMsgAddPool(creator string, pool *SwapPool) *MsgAddPool {
	return &MsgAddPool{
		Creator: creator,
		Pool:    pool,
	}
}

func (msg *MsgAddPool) Route() string {
	return RouterKey
}

func (msg *MsgAddPool) Type() string {
	return TypeMsgAddPool
}

func (msg *MsgAddPool) GetSigners() []sdk.AccAddress {
	accAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(errorsmod.Wrapf(err, "failed to acc address from bech32 string, given string: %s", msg.Creator))
	}

	return []sdk.AccAddress{accAddress}
}

func (msg *MsgAddPool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddPool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address: %s", err)
	}

	if err = validatePool(msg.Pool); err != nil {
		return err
	}

	return nil
}
