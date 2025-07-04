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

package cli

import (
	errorsmod "cosmossdk.io/errors"
	"encoding/json"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"

	evmostypes "github.com/Bridgeless-Project/bridgeless-core/v12/types"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/revenue/v1/types"
)

// NewTxCmd returns a root CLI command handler for certain modules/revenue
// transaction commands.
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "revenue subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		NewRegisterRevenue(),
		NewCancelRevenue(),
		NewUpdateRevenue(),
	)
	return txCmd
}

// NewRegisterRevenue returns a CLI command handler for registering a
// contract for fee distribution
func NewRegisterRevenue() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register CONTRACT_HEX NONCE... [WITHDRAWER_BECH32]",
		Short: "Register a contract for fee distribution. **NOTE** Please ensure, that the deployer of the contract (or the factory that deployes the contract) is an account that is owned by your project, to avoid that an individual deployer who leaves your project becomes malicious.",
		Long:  "Register a contract for fee distribution.\nOnly the contract deployer can register a contract.\nProvide the account nonce(s) used to derive the contract address. E.g.: you have an account nonce of 4 when you send a deployment transaction for a contract A; you use this contract as a factory, to create another contract B. If you register A, the nonces value is \"4\". If you register B, the nonces value is \"4,1\" (B is the first contract created by A). \nThe withdrawer address defaults to the deployer address if not provided.",
		Args:  cobra.RangeArgs(2, 3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var withdrawer string
			deployer := cliCtx.GetFromAddress()

			contract := args[0]
			if err := evmostypes.ValidateNonZeroAddress(contract); err != nil {
				return errorsmod.Wrap(err, "invalid contract hex address")
			}

			var nonces []uint64
			if err = json.Unmarshal([]byte("["+args[1]+"]"), &nonces); err != nil {
				return errorsmod.Wrap(err, "invalid nonces")
			}

			if len(args) == 3 {
				withdrawer = args[2]
				if _, err := sdk.AccAddressFromBech32(withdrawer); err != nil {
					return errorsmod.Wrap(err, "invalid withdrawer bech32 address")
				}
			}

			// If withdraw address is the same as contract deployer, remove the
			// field for avoiding storage bloat
			if deployer.String() == withdrawer {
				withdrawer = ""
			}

			msg := &types.MsgRegisterRevenue{
				ContractAddress:   contract,
				DeployerAddress:   deployer.String(),
				WithdrawerAddress: withdrawer,
				Nonces:            nonces,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewCancelRevenue returns a CLI command handler for canceling a
// contract for fee distribution
func NewCancelRevenue() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel CONTRACT_HEX",
		Short: "Cancel a contract from fee distribution",
		Long:  "Cancel a contract from fee distribution. The deployer will no longer receive fees from users interacting with the contract. \nOnly the contract deployer can cancel a contract.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			deployer := cliCtx.GetFromAddress()

			contract := args[0]
			if err := evmostypes.ValidateNonZeroAddress(contract); err != nil {
				return errorsmod.Wrap(err, "invalid contract hex address")
			}

			msg := &types.MsgCancelRevenue{
				ContractAddress: contract,
				DeployerAddress: deployer.String(),
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewUpdateRevenue returns a CLI command handler for updating the withdraw
// address of a contract for fee distribution
func NewUpdateRevenue() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update CONTRACT_HEX WITHDRAWER_BECH32",
		Short: "Update withdrawer address for a contract registered for fee distribution.",
		Long:  "Update withdrawer address for a contract registered for fee distribution. \nOnly the contract deployer can update the withdrawer address.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			deployer := cliCtx.GetFromAddress()

			contract := args[0]
			if err := evmostypes.ValidateNonZeroAddress(contract); err != nil {
				return errorsmod.Wrap(err, "invalid contract hex address")
			}

			withdrawer := args[1]
			if _, err := sdk.AccAddressFromBech32(withdrawer); err != nil {
				return errorsmod.Wrap(err, "invalid withdrawer bech32 address")
			}

			msg := &types.MsgUpdateRevenue{
				ContractAddress:   contract,
				DeployerAddress:   deployer.String(),
				WithdrawerAddress: withdrawer,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
