package cli

import (
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
			tr, err = parseSubmitTx(args[1])
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
