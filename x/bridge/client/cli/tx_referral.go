package cli

import (
	"strconv"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func TxReferralCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "referral",
		Short:                      "Referral subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdSetReferral(),
		CmdRemoveReferral(),
	)
	// this line is used by starport scaffolding # 1

	return cmd
}

func CmdSetReferral() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set [from_key_or_address] [referral-id] [referral-withdrawal-address] [commission-rate]",
		Short: "set a new referral",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Set(flags.FlagFrom, args[0])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			referralId, err := strconv.ParseUint(args[1], 10, 32)
			if err != nil {
				return err
			}

			commissionRate, err := strconv.ParseInt(args[4], 10, 32)
			if err != nil {
				return err
			}

			referral := types.Referral{
				Id:                uint32(referralId),
				WithdrawalAddress: args[3],
				CommissionRate:    int32(commissionRate),
			}

			msg := types.NewMsgSetReferral(
				clientCtx.GetFromAddress().String(),
				referral,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdRemoveReferral() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove [from_key_or_address] [id]",
		Short: "remove the referral",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Set(flags.FlagFrom, args[0])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			referralId, err := strconv.ParseUint(args[1], 10, 32)
			if err != nil {
				return err
			}

			msg := types.NewMsgRemoveReferral(
				clientCtx.GetFromAddress().String(),
				uint32(referralId),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
