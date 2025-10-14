package cli

import (
	"context"
	"strconv"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/multisig/types"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

func CmdQueryVotes() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "votes",
		Short: "Query multisig votes",
		Args:  cobra.NoArgs,
	}

	cmd.AddCommand(QueryAllVotes(), QueryVote(), QueryVoteByProposalAndProposer())

	return cmd
}

func QueryAllVotes() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-all",
		Short: "Lists all votes",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.VoteAll(context.Background(), &types.QueryAllVoteRequest{})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	return cmd
}

func QueryVote() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vote [prorosalId]",
		Short: "Query multisig votes for proposal",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)
			proposalId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			if err != nil {
				return err
			}
			res, err := queryClient.VotesByProposal(context.Background(), &types.QueryVotesByProposalRequest{
				ProposalId: proposalId,
			})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	return cmd
}

func QueryVoteByProposalAndProposer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vote-detailed [proposalId] [proposer]",
		Short: "Query multisig vote by proposalId and proposer",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)
			proposalId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			_, err = sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			res, err := queryClient.Vote(context.Background(), &types.QueryGetVoteRequest{
				ProposalId: proposalId,
				Voter:      args[1],
			})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	return cmd
}
