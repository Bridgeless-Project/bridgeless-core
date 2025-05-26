package types

import (
	errorsmod "cosmossdk.io/errors"
	"errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var ErrInvalidPeriod = errors.New("period cannot be 0")

const (
	DefaultPrunePeriod  = 240
	DefaultVotingPeriod = 120
)

// NewParams creates a new Params instance
func NewParams() Params {
	return Params{}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return Params{
		PrunePeriod:  DefaultPrunePeriod,
		VotingPeriod: DefaultVotingPeriod,
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validatePeriod(p.VotingPeriod); err != nil {
		return err
	}

	return validatePeriod(p.PrunePeriod)
}

func validatePeriod(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidType, "invalid parameter type: %T", i)
	}

	if v == 0 {
		return errorsmod.Wrapf(ErrInvalidPeriod, "voting period must be positive: %d", v)
	}

	return nil
}
