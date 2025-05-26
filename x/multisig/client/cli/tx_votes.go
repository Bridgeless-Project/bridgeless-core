package cli

import (
	errorsmod "cosmossdk.io/errors"

	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/hyle-team/bridgeless-core/v12/x/multisig/types"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

func TxVoteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "votes",
		Short:                      "proposals subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	cmd.AddCommand(
		CmdNewVote())
	return cmd
}

func CmdNewVote() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vote [voter_address] [proposal_id] [option]",
		Short: "Create a new vote",
		Long: strings.TrimSpace(fmt.Sprintf(`Create a vote for proposal
Example:
$ %s tx multisig votes vote bridge1... 1 0
0 - YES
1 - NO`, version.AppName)),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Set(flags.FlagFrom, args[0])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			proposalId, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			vote_option, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}

			if vote_option != 0 && vote_option != 1 {
				return errorsmod.Wrap(types.ErrInvalidVoteOption, "vote option must be 0 or 1")
			}

			opt := types.VoteOption_NO
			if vote_option == 0 {
				opt = types.VoteOption_YES
			}

			msg := types.NewMsgVote(clientCtx.GetFromAddress().String(), proposalId, opt)

			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
