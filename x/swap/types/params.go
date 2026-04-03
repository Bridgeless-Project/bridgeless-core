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
func NewParams(moduleAdmin, uniswapRouterAddress, wrappedBridge string, swapDeadlineSeconds uint64) Params {
	return Params{
		ModuleAdmin:          moduleAdmin,
		UniswapRouterAddress: uniswapRouterAddress,
		WrappedBridge:        wrappedBridge,
		SwapDeadlineSeconds:  swapDeadlineSeconds,
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
		paramtypes.NewParamSetPair([]byte(ParamUniswapRouterAddressKey), &p.UniswapRouterAddress, validateUniswapRouterAddress),
		paramtypes.NewParamSetPair([]byte(ParamWrappedBridgeKey), &p.WrappedBridge, validateWrappedBridge),
		paramtypes.NewParamSetPair([]byte(ParamSwapDeadlineSecondsKey), &p.SwapDeadlineSeconds, validateSwapDeadlineSeconds),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateModuleAdmin(p.ModuleAdmin); err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid module admin address (%s)", err)
	}
	if err := validateUniswapRouterAddress(p.UniswapRouterAddress); err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid uniswap router address (%s)", err)
	}
	if err := validateWrappedBridge(p.WrappedBridge); err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid wrapped bridge address (%s)", err)
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

func validateUniswapRouterAddress(i interface{}) error {
	addr, ok := i.(string)
	if !ok {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidType, "invalid parameter type: %T", i)
	}

	if !common.IsHexAddress(addr) {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid uniswap router address: %s", addr)
	}

	return nil
}

func validateWrappedBridge(i interface{}) error {
	addr, ok := i.(string)
	if !ok {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidType, "invalid parameter type: %T", i)
	}

	if !common.IsHexAddress(addr) {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid wrapped bridge address: %s", addr)
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
