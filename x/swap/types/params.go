package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/ethereum/go-ethereum/common"
)

var _ paramtypes.ParamSet = (*Params)(nil)

const DefaultSwapDeadlineSeconds uint64 = 15 * 60

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(moduleAdmin, wrappedBridge, swapperAddress string, swapDeadlineSeconds uint64) Params {
	return Params{
		ModuleAdmin:         moduleAdmin,
		WrappedBridge:       wrappedBridge,
		SwapperAddress:      swapperAddress,
		SwapDeadlineSeconds: swapDeadlineSeconds,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams("", "", "", DefaultSwapDeadlineSeconds)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair([]byte(ParamModuleAdminKey), &p.ModuleAdmin, validateModuleAdmin),
		paramtypes.NewParamSetPair([]byte(ParamWrappedBridgeKey), &p.WrappedBridge, validateWrappedBridge),
		paramtypes.NewParamSetPair([]byte(ParamSwapperAddressKey), &p.SwapperAddress, validateSwapperAddress),
		paramtypes.NewParamSetPair([]byte(ParamSwapDeadlineSecondsKey), &p.SwapDeadlineSeconds, validateSwapDeadlineSeconds),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateModuleAdmin(p.ModuleAdmin); err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid module admin address (%s)", err)
	}
	if err := validateWrappedBridge(p.WrappedBridge); err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid wrapped bridge address (%s)", err)
	}
	if err := validateSwapperAddress(p.SwapperAddress); err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid swapper address (%s)", err)
	}
	if err := validateSwapDeadlineSeconds(p.SwapDeadlineSeconds); err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid swap deadline seconds (%s)", err)
	}

	return nil
}

func validateModuleAdmin(i interface{}) error {
	adm, ok := i.(string)
	if !ok {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidType, "invalid parameter type: %T", i)
	}
	if adm == "" {
		return nil
	}

	_, err := sdk.AccAddressFromBech32(adm)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid module admin address: %s", err.Error())
	}

	return nil
}

func validateWrappedBridge(i interface{}) error {
	addr, ok := i.(string)
	if !ok {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidType, "invalid parameter type: %T", i)
	}
	if addr == "" {
		return nil
	}

	if !common.IsHexAddress(addr) {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid wrapped bridge address: %s", addr)
	}

	return nil
}

func validateSwapperAddress(i interface{}) error {
	addr, ok := i.(string)
	if !ok {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidType, "invalid parameter type: %T", i)
	}
	if addr == "" {
		return nil
	}

	if !common.IsHexAddress(addr) {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid swapper address: %s", addr)
	}

	return nil
}

func validateSwapDeadlineSeconds(i interface{}) error {
	value, ok := i.(uint64)
	if !ok {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidType, "invalid parameter type: %T", i)
	}
	if value == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "swap deadline seconds must be greater than zero")
	}

	return nil
}
