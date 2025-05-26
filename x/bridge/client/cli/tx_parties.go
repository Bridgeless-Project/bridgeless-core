package cli

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
	"github.com/spf13/cobra"
	"strings"
)

func TxPartiesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "parties",
		Short:                      "Parties transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdSubmitParties())
	// this line is used by starport scaffolding # 1
	return cmd
}

func CmdSubmitParties() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set [from_key_or_address] [parties-list]",
		Short: "Set a new parties list",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Set(flags.FlagFrom, args[0])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return errorsmod.Wrap(err, "cannot get client tx context")
			}

			arg2 := args[1]
			parties := strings.Split(arg2, ",")
			for _, party := range parties {
				_, err = sdk.AccAddressFromBech32(party)
				if err != nil {
					return errorsmod.Wrap(err, "failed to set parties")
				}
			}

			var partiesList []*types.Party

			for _, party := range parties {
				partiesList = append(partiesList, &types.Party{
					Address: party,
				})
			}
			msg := types.NewMsgSetParties(clientCtx.GetFromAddress().String(), partiesList)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	return cmd
}
