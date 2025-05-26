package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	_ paramtypes.ParamSet = (*Params)(nil)
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair([]byte(ParamModuleAdminKey), &p.ModuleAdmin, validateModuleAdmin),
		paramtypes.NewParamSetPair([]byte(ParamModulePartiesKey), &p.Parties, validateModuleParties),
		paramtypes.NewParamSetPair([]byte(ParamTssThresholdKey), &p.TssThreshold, validateTssThreshold),
	}
}

// NewParams creates a new Params instance
func NewParams(moduleAdmin string, parties []*Party) Params {
	return Params{
		ModuleAdmin: moduleAdmin,
		Parties:     parties,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams("", []*Party{})
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateModuleAdmin(p.ModuleAdmin); err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid module admin address (%s)", err)
	}

	if err := validateModuleParties(p.Parties); err != nil {
		return errorsmod.Wrapf(ErrInvalidPartiesList, "invalid parties list (%s)", err)
	}

	return nil
}

func validateModuleAdmin(i interface{}) error {
	adm, ok := i.(string)
	if !ok {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidType, "invalid parameter type: %T", i)
	}

	_, err := sdk.AccAddressFromBech32(adm)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid module admin address: %s", err.Error())
	}

	return nil
}
func validateModuleParties(i interface{}) error {
	parties, ok := i.([]*Party)
	if !ok {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidType, "invalid parameter type: %T", i)
	}

	for _, party := range parties {
		_, err := sdk.AccAddressFromBech32(party.Address)
		if err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid party address (%s)", err.Error())
		}
	}

	return nil
}

func validateTssThreshold(i interface{}) error {
	_, ok := i.(uint32)
	if !ok {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidType, "invalid parameter type: %T", i)
	}

	return nil
}
