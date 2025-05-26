package cli

import (
	"cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
	"github.com/spf13/cobra"
	"strconv"
)

func TxTssThresholdCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "tss-threshold",
		Short:                      "Tss threshold transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdSetTssThreshold())
	// this line is used by starport scaffolding # 1
	return cmd
}

func CmdSetTssThreshold() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set [from_key_or_address] [threshold]",
		Short: "Set a new parties list",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Set(flags.FlagFrom, args[0])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return errors.Wrap(err, "cannot get client tx context")
			}

			threshold, err := strconv.ParseUint(args[1], 10, 32)
			if err != nil {
				return errors.Wrap(err, "cannot parse threshold")
			}

			msg := types.NewMsgSetTssThreshold(clientCtx.GetFromAddress().String(), uint32(threshold))
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	return cmd
}
