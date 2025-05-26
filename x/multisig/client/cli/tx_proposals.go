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
	"strings"
)

func TxProposalsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "proposals",
		Short:                      "Proposals subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdSubmitProposal())

	return cmd
}

func CmdSubmitProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-proposal [path/to/proposal.json]",
		Short: "Submit a proposal along with some messages and metadata",
		Long: strings.TrimSpace(fmt.Sprintf(`Submit a proposal to group.
Example:
$ %s tx multisig proposals submit-proposal /path/to/proposal.json --from mykey

Where proposal.json looks like:

{
  "creator": "bridge1...",
  "group": "bridge1...",
  "messages": [{
    "@type": "/core.multisig.MsgChangeGroup",
    "creator": "bridge1...", //creator and group must be equal
    "group": "bridge1...",
    "members": ["bridge1...",
      "bridge1...",
      "bridge1..."],
    "threshold": 3
  }]
}`, version.AppName)),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			creator, group, msgs, err := parseSubmitProposal(clientCtx.Codec, args[0])
			if err != nil {
				return err
			}

			msg, err := types.NewMsgSubmitProposal(creator, group, msgs)
			if err != nil {
				return errorsmod.Wrap(err, "invalid message")
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
