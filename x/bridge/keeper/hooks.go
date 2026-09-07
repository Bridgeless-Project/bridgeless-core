// Copyright 2022 Evmos Foundation
// This file is part of the Evmos Network packages.
//
// Evmos is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Evmos packages are distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Evmos packages. If not, see https://github.com/evmos/evmos/blob/main/LICENSE

package keeper

import (
	evmtypes "github.com/Bridgeless-Project/bridgeless-core/v12/x/evm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/Zano-Execution-Layer/go-ethereum/core"
	ethtypes "github.com/Zano-Execution-Layer/go-ethereum/core/types"
)

var (
	_ evmtypes.EvmHooks         = Hooks{}
	_ govtypes.GovHooks         = Hooks{}
	_ stakingtypes.StakingHooks = Hooks{}
)

// Hooks wrapper struct for the claim keeper
type Hooks struct {
	k Keeper
}

func (h Hooks) ProcessStakingRewardsWithdrawal(_ sdk.Context, _ sdk.AccAddress, _ sdk.ValAddress) (sdk.Coins, error) {
	return nil, nil
}

func (k Keeper) Hooks() Hooks {
	return Hooks{k}
}

func (h Hooks) AfterProposalVote(ctx sdk.Context, proposalID uint64, voterAddr sdk.AccAddress) {}

func (h Hooks) AfterDelegationModified(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	return nil
}

// hooks for evm module to listen for EVM events and create a new position with the data from that event
func (h Hooks) PostTxProcessing(ctx sdk.Context, msg core.Message, receipt *ethtypes.Receipt) error {
	return h.k.PostTxProcessing(ctx, msg, receipt)
}

// ________________________________________________________________________________________

// governance hooks
func (h Hooks) AfterProposalFailedMinDeposit(_ sdk.Context, _ uint64) {
}

func (h Hooks) AfterProposalVotingPeriodEnded(_ sdk.Context, _ uint64) {
}

func (h Hooks) AfterProposalSubmission(_ sdk.Context, _ uint64) {}

func (h Hooks) AfterProposalDeposit(_ sdk.Context, _ uint64, _ sdk.AccAddress) {
}

func (h Hooks) AfterProposalInactive(_ sdk.Context, _ uint64) {}

func (h Hooks) AfterProposalActive(_ sdk.Context, _ uint64) {}

// staking hooks
func (h Hooks) AfterValidatorCreated(_ sdk.Context, _ sdk.ValAddress) error {
	return nil
}

func (h Hooks) BeforeValidatorModified(_ sdk.Context, _ sdk.ValAddress) error {
	return nil
}

func (h Hooks) AfterValidatorRemoved(_ sdk.Context, _ sdk.ConsAddress, _ sdk.ValAddress) error {
	return nil
}

func (h Hooks) AfterValidatorBonded(_ sdk.Context, _ sdk.ConsAddress, _ sdk.ValAddress) error {
	return nil
}

func (h Hooks) AfterValidatorBeginUnbonding(_ sdk.Context, _ sdk.ConsAddress, _ sdk.ValAddress) error {
	return nil
}

func (h Hooks) BeforeDelegationCreated(_ sdk.Context, _ sdk.AccAddress, _ sdk.ValAddress) error {
	return nil
}

func (h Hooks) BeforeDelegationSharesModified(_ sdk.Context, _ sdk.AccAddress, _ sdk.ValAddress) error {
	return nil
}

func (h Hooks) BeforeDelegationRemoved(_ sdk.Context, _ sdk.AccAddress, _ sdk.ValAddress) error {
	return nil
}

func (h Hooks) BeforeValidatorSlashed(_ sdk.Context, _ sdk.ValAddress, _ sdk.Dec) error {
	return nil
}

func (h Hooks) FundCommunityPoolFromModule(_ sdk.Context, _ sdk.Coins, _ string) error { return nil }
