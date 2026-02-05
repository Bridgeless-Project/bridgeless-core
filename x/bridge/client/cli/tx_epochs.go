package cli

import (
	"strconv"

	errorsmod "cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func TxEpochsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "epochs",
		Short:                      "Epoch transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdStartEpoch(),
		CmdSetEpochSignature(),
		CmdSetEpochPubkey(),
	)

	return cmd
}

// CmdStartEpoch creates a CLI command for starting a new epoch
// The input is a JSON file containing:
//
//	{
//	  "epoch_id": 1,
//	  "tss_info": [
//	    {
//	      "certificate": "cert1",
//	      "domen": "domain1.com",
//	      "address": "0x...",
//	      "active": true
//	    }
//	  ],
//	  "tss_threshold": 2
//	}
func CmdStartEpoch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start [from_key_or_address] [path-to-epoch-json]",
		Short: "Start a new epoch with TSS info from JSON file",
		Long: `Start a new epoch by providing TSS information in a JSON file.

Example JSON file format:
{
  "epoch_id": 1,
  "tss_info": [
    {
      "certificate": "cert_pem_content",
      "domen": "tss-node1.example.com",
      "address": "0x1234567890abcdef",
      "active": true
    },
    {
      "certificate": "cert_pem_content",
      "domen": "tss-node2.example.com", 
      "address": "0xabcdef1234567890",
      "active": true
    }
  ],
  "tss_threshold": 2
}`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Set(flags.FlagFrom, args[0])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			input, err := parseStartEpochInput(args[1])
			if err != nil {
				return errorsmod.Wrap(err, "failed to parse start epoch input")
			}

			msg := types.NewMsgStartEpoch(
				clientCtx.GetFromAddress().String(),
				input.EpochID,
				input.TSSInfo,
				input.TSSThreshold,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// CmdSetEpochSignature creates a CLI command for setting epoch signatures
// The input is a JSON file containing an array of EpochChainSignatures:
//
//	[
//	  {
//	    "epoch_id": 1,
//	    "chain_type": 1,
//	    "added_signature": {...},
//	    "removed_signature": {...},
//	    "address": "0x...",
//	    "submittions": 0
//	  }
//	]
func CmdSetEpochSignature() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-signature [from_key_or_address] [path-to-signatures-json]",
		Short: "Set epoch signatures from JSON file",
		Long: `Set epoch chain signatures by providing signature data in a JSON file.

Example JSON file format:
[
  {
    "epoch_id": 1,
    "chain_type": 1,
    "added_signature": {
      "mod": 1,
      "epoch_id": 1,
      "signature": "0xsignature...",
      "data": {
        "new_signer": "0xnewsigner...",
        "start_time": 1234567890,
        "end_time": 1234567899,
        "nonce": "abc123"
      }
    },
    "removed_signature": null,
    "address": "0xcontractaddress...",
    "submittions": 0
  }
]`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Set(flags.FlagFrom, args[0])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			signatures, err := parseEpochChainSignatures(args[1])
			if err != nil {
				return errorsmod.Wrap(err, "failed to parse epoch chain signatures")
			}

			msg := types.NewMsgSetEpochSignature(
				clientCtx.GetFromAddress().String(),
				signatures,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// CmdSetEpochPubkey creates a CLI command for setting epoch pubkey
func CmdSetEpochPubkey() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-pubkey [from_key_or_address] [pubkey] [epoch_id]",
		Short: "Set the pubkey for an epoch",
		Long:  `Set the public key for a specific epoch. Only authorized parties can set the pubkey.`,
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Set(flags.FlagFrom, args[0])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			pubkey := args[1]
			epochId, err := strconv.ParseUint(args[2], 10, 32)
			if err != nil {
				return errorsmod.Wrap(err, "failed to parse epoch ID")
			}

			msg := types.NewMsgSetEpochPubkey(
				clientCtx.GetFromAddress().String(),
				pubkey,
				uint32(epochId),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
