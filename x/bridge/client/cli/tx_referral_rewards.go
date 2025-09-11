package cli

import (
	"strconv"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

func TxReferralRewardsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "referral-rewards",
		Short:                      "Referral rewards subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdSetReferralRewards(),
		CmdRemoveReferralRewards(),
	)
	// this line is used by starport scaffolding # 1

	return cmd
}

func CmdSetReferralRewards() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set [from_key_or_address] [referral-id] [token-id] [to-claim] [total-collected-amount]",
		Short: "set a new new referral rewards",
		Args:  cobra.ExactArgs(5),
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

			tokenId, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}

			allowedCoins, err := sdk.ParseCoinNormalized(args[3])
			if err != nil {
				return err
			}

			totalCollectedAmount, err := sdk.ParseCoinNormalized(args[4])
			if err != nil {
				return err
			}

			referral := types.ReferralRewards{
				ReferralId:           uint32(referralId),
				TokenId:              tokenId,
				ToClaim:              allowedCoins,
				TotalCollectedAmount: totalCollectedAmount,
			}

			msg := types.NewMsgSetReferralRewards(
				clientCtx.GetFromAddress().String(),
				referral,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdRemoveReferralRewards() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove [from_key_or_address] [referral-id] [token-id]",
		Short: "remove the referral rewards",
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
			tokenId, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgRemoveReferralRewards(
				clientCtx.GetFromAddress().String(),
				uint32(referralId),
				tokenId,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
