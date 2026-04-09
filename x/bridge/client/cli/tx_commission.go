package cli

import (
	"math/big"

	"cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func TxCommissionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "commission",
		Short:                      "Commission subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdSetCommission(),
		CmdUpdateCommission(),
		CmdRemoveCommission(),
	)
	// this line is used by starport scaffolding # 1

	return cmd
}

func CmdSetCommission() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-commission [token-id] [amount]",
		Short: "Set a commission",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			tokenId, ok := big.NewInt(0).SetString(args[0], 10)
			if !ok {
				return errors.Wrap(types.ErrInvalidDataType, "token-id must be a valid integer")
			}

			msg := types.NewMsgSetCommission(
				clientCtx.GetFromAddress().String(),
				tokenId.Uint64(),
				args[1],
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateCommission() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-commission [token-id] [amount]",
		Short: "Update a commission",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			tokenId, ok := big.NewInt(0).SetString(args[0], 10)
			if !ok {
				return errors.Wrap(types.ErrInvalidDataType, "token-id must be a valid integer")
			}

			msg := types.NewMsgUpdateCommission(
				clientCtx.GetFromAddress().String(),
				tokenId.Uint64(),
				args[1],
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdRemoveCommission() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-commission [token-id]",
		Short: "Remove a commission",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			tokenId, ok := big.NewInt(0).SetString(args[0], 10)
			if !ok {
				return errors.Wrap(types.ErrInvalidDataType, "token-id must be a valid integer")
			}

			msg := types.NewMsgRemoveCommission(
				clientCtx.GetFromAddress().String(),
				tokenId.Uint64(),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
