package cli

import (
	"context"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdQueryGetReferralRewardsByToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "referral-rewards [referral-id] [token-id]",
		Short: "Query bridge referral rewards by referral id and token id",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			referraid, err := cmd.Flags().GetUint32("referral-id")
			if err != nil {
				return err
			}

			tokenId, err := cmd.Flags().GetUint64("token-id")
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.GetReferralRewardsByToken(cmd.Context(), &types.QueryGetReferralRewardByToken{ReferralId: referraid, TokenId: tokenId})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "txs")

	return cmd
}

func CmdQueryGetReferralsRewardsById() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "referrals",
		Short: "Query referrals",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			req := &types.QueryGetReferralRewardsById{
				Pagination: pageReq,
			}

			res, err := queryClient.GetReferralsRewardsById(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
