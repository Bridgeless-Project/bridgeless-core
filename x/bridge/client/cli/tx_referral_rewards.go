package cli

import (
	"math/big"
	"strconv"

	"cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
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

			_, ok := big.NewInt(0).SetString(args[3], 10)
			if !ok {
				return errors.Wrap(types.ErrInvalidDataType, "to-claim must be a valid integer")
			}

			_, ok = big.NewInt(0).SetString(args[4], 10)
			if !ok {
				return errors.Wrap(types.ErrInvalidDataType, "total-collected-amount must be a valid integer")
			}

			referral := types.ReferralRewards{
				ReferralId:           uint32(referralId),
				TokenId:              tokenId,
				ToClaim:              args[3],
				TotalCollectedAmount: args[4],
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
