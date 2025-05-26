package cli

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
	"github.com/spf13/cobra"
)

func TxChainsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "chains",
		Short:                      "Chain transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdInsertChain(),
		CmdRemoveChain(),
	)
	// this line is used by starport scaffolding # 1

	return cmd
}

func CmdInsertChain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "insert [from_key_or_address] [path-to-chain-json]",
		Short: "Set a new chain",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Set(flags.FlagFrom, args[0])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			chain, err := parseInsertChain(args[1])
			if err != nil {
				return errorsmod.Wrap(err, "failed unmarshalling chain")
			}

			msg := types.NewMsgInsertChain(
				clientCtx.GetFromAddress().String(),
				*chain,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdRemoveChain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove [from_key_or_address] [chain_id]",
		Short: "Remove the chain",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Set(flags.FlagFrom, args[0])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteChain(
				clientCtx.GetFromAddress().String(),
				args[1],
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
