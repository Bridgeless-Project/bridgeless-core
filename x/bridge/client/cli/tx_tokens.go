package cli

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
	"github.com/spf13/cobra"
	"strconv"
)

func TxTokensCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "tokens",
		Short:                      "Tokens transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdInsertToken(),
		CmdRemoveToken(),
		CmdUpdateToken(),
		CmdAddTokenInfo(),
		CmdRemoveTokenInfo(),
	)
	// this line is used by starport scaffolding # 1

	return cmd
}

func CmdInsertToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "insert [from_key_or_address] [path-to-token-json] ",
		Short: "Insert a new token",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Set(flags.FlagFrom, args[0])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			token, err := parseInsertToken(args[1])
			if err != nil {
				return errorsmod.Wrap(err, "failed to parse token")
			}

			msg := types.NewMsgInsertToken(
				clientCtx.GetFromAddress().String(),
				*token,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdRemoveToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove [from_key_or_address] [token_id]",
		Short: "Remove the token",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			cmd.Flags().Set(flags.FlagFrom, args[0])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteToken(
				clientCtx.GetFromAddress().String(),
				id,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdUpdateToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update [from_key_or_address] [token-json]",
		Short: "Update the token metadata",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {

			cmd.Flags().Set(flags.FlagFrom, args[0])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			token := types.Token{}
			if err = types.ModuleCdc.UnmarshalJSON([]byte(args[1]), &token); err != nil {
				return err
			}

			msg := types.NewMsgUpdateToken(
				clientCtx.GetFromAddress().String(),
				token.Id, token.Metadata,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdAddTokenInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-info [from_key_or_address] [token-info-json]",
		Short: "Add token info to existing token",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			cmd.Flags().Set(flags.FlagFrom, args[0])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			tokenInfo := types.TokenInfo{}
			if err = types.ModuleCdc.UnmarshalJSON([]byte(args[1]), &tokenInfo); err != nil {
				return err
			}

			msg := types.NewMsgAddTokenInfo(
				clientCtx.GetFromAddress().String(),
				tokenInfo,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdRemoveTokenInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-info [from_key_or_address] [token-id] [chain-id]",
		Short: "Remove token info from existing token",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {

			cmd.Flags().Set(flags.FlagFrom, args[0])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgRemoveTokenInfo(
				clientCtx.GetFromAddress().String(),
				id, args[2],
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
