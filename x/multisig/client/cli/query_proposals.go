package cli

import (
	"context"
	"strconv"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/multisig/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

func CmdQueryProposals() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proposals",
		Short: "Query multisig proposals",
		Args:  cobra.NoArgs,
	}

	cmd.AddCommand(QueryAllProposals(), QueryProposal())

	return cmd
}

func QueryAllProposals() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-all",
		Short: "Lists all proposals",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.ProposalAll(context.Background(), &types.QueryAllProposalRequest{})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	return cmd
}

func QueryProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proposal [proposal_id]",
		Short: "Query multisig proposal",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)
			proposalId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			res, err := queryClient.Proposal(context.Background(), &types.QueryGetProposalRequest{ProposalId: proposalId})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	return cmd
}
