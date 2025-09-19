package types

import (
	"errors"
	"fmt"
	"math/big"
)

func validateReferral(referral *Referral) error {
	if referral == nil {
		return fmt.Errorf("transaction is nil")
	}
	if referral.WithdrawalAddress == "" {
		return errors.New("withdrawal address cannot be empty")
	}

	if referral.Id <= 0 {
		return errors.New("referral id must be greater than zero")
	}

	commissionRete, ok := big.NewFloat(0).SetString(referral.CommissionRate)
	if !ok {
		return errors.New("invalid commission rate")
	}

	if commissionRete.Cmp(big.NewFloat(0)) == -1 || commissionRete.Cmp(big.NewFloat(1)) == 1 {
		return errors.New("commission rate must be in [0, 1] range")
	}
	return nil
}

func validateReferralRewards(rewards *ReferralRewards) error {
	if rewards == nil {
		return errors.New("rewards is nil")
	}

	totalCollectedAmount, ok := big.NewInt(0).SetString(rewards.TotalCollectedAmount, 10)
	if !ok {
		return errors.New("invalid total collected amount")
	}

	toClaim, ok := big.NewInt(0).SetString(rewards.ToClaim, 10)
	if !ok {
		return errors.New("invalid to claim amount")
	}

	if totalCollectedAmount.Cmp(toClaim) == -1 {
		return errors.New("total collected amount must be greater or equal to claim")
	}

	return nil
}
