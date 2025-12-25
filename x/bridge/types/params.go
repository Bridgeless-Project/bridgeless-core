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
		paramtypes.NewParamSetPair([]byte(ParamRelayerAccounts), &p.RelayerAccounts, validateRelayerAccounts),
		paramtypes.NewParamSetPair([]byte(ParamEpochSequence), &p.EpochSequence, validateEpochSequence),
	}
}

// NewParams creates a new Params instance
func NewParams(moduleAdmin string, parties []*Party, tssThreshold uint32, relayerAccounts []string, epochSequence uint64) Params {
	return Params{
		ModuleAdmin:     moduleAdmin,
		Parties:         parties,
		TssThreshold:    tssThreshold,
		RelayerAccounts: relayerAccounts,
		EpochSequence:   epochSequence,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams("", []*Party{}, 0, []string{}, 0)
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateModuleAdmin(p.ModuleAdmin); err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid module admin address (%s)", err)
	}

	if err := validateModuleParties(p.Parties); err != nil {
		return errorsmod.Wrapf(ErrInvalidPartiesList, "invalid parties list (%s)", err)
	}

	if err := validateTssThreshold(p.TssThreshold); err != nil {
		return errorsmod.Wrapf(ErrInvalidTssThreshold, "invalid TssThreshold (%s)", err)
	}

	if err := validateRelayerAccounts(p.RelayerAccounts); err != nil {
		return errorsmod.Wrap(err, "invalid relayer accounts")
	}

	if err := validateEpochSequence(p.EpochSequence); err != nil {
		return errorsmod.Wrapf(err, "invalid epoch sequence (%d)", p.EpochSequence)
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

func validateRelayerAccounts(i interface{}) error {
	relayerAccs, ok := i.([]string)
	if !ok {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidType, "invalid parameter type: %T", i)
	}

	for _, account := range relayerAccs {
		_, err := sdk.AccAddressFromBech32(account)
		if err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid relayer account address: %s", err.Error())
		}
	}

	return nil
}

func validateEpochSequence(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidType, "invalid parameter type: %T", i)
	}

	return nil
}
