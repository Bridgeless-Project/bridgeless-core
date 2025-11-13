package cli

import (
	"strconv"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func TxTransactionsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "transactions",
		Short:                      "transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdSubmitTx(),
		CmdRemoveTx(),
		CmdUpdateTx(),
	)

	return cmd
}

func CmdSubmitTx() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit [from_key_or_address] [path-to-json-tx]",
		Short: "Submit a transaction to the bridge module",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Set(flags.FlagFrom, args[0])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var tr []types.Transaction
			tr, err = parseTxs(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgSubmitTransactions(clientCtx.GetFromAddress().String(), tr...)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateTx() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update [from_key_or_address] [path-to-json-tx]",
		Short: "Update the transaction on the bridge module",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Set(flags.FlagFrom, args[0])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			tr, err := parseTx(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateTransaction(clientCtx.GetFromAddress().String(), *tr)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdRemoveTx() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove [from_key_or_address] [deposit-tx-hash] [deposit-tx-index] [deposit-chain-id]",
		Short: "Remove the transaction from the bridge module",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Set(flags.FlagFrom, args[0])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			depositHash := args[1]
			depositIndex, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}
			depostChainId := args[3]

			msg := types.NewMsgRemoveTransaction(clientCtx.GetFromAddress().String(), depostChainId, depositHash, depositIndex)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
