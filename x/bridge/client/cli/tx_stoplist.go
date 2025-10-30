package cli

import (
	"strconv"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func TxStopListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "stop-list",
		Short:                      "stop-list subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdAddTxToSL(),
		CmdRemoveTxFromSL(),
	)

	return cmd
}

func CmdAddTxToSL() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-tx [from-key_or_address] [path-to-json-tx]",
		Short: "Submit a transaction to the bridge module",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Set(flags.FlagFrom, args[0])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var tr *types.Transaction
			tr, err = parseTx(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgAddTxToStopList(clientCtx.GetFromAddress().String(), tr)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdRemoveTxFromSL() *cobra.Command {
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

			msg := types.NewMsgRemoveTxFromStopList(clientCtx.GetFromAddress().String(), depositIndex, depositHash, depostChainId)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
