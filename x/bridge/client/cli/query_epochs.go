package cli

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

// CmdQueryEpochTransactions queries transactions for a specific epoch.
func CmdQueryEpochTransactions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "epoch-transactions [epoch-id]",
		Short: "Query transactions for a specific epoch",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			epochID, err := strconv.ParseUint(args[0], 10, 32)
			if err != nil {
				return fmt.Errorf("invalid epoch ID: %w", err)
			}

			req := &types.QueryGetEpochTransactions{
				EpochId: uint32(epochID),
			}

			res, err := queryClient.GetEpochTransactions(context.Background(), req)
			if err != nil {
				return fmt.Errorf("failed to query epoch transactions: %w", err)
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQueryEpochByID queries details of a specific epoch by ID.
func CmdQueryEpochByID() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "epoch [epoch-id]",
		Short: "Query details of a specific epoch by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			epochID, err := strconv.ParseUint(args[0], 10, 32)
			if err != nil {
				return fmt.Errorf("invalid epoch ID: %w", err)
			}

			req := &types.QueryGetEpoch{
				EpochId: uint32(epochID),
			}

			res, err := queryClient.GetEpochById(context.Background(), req)
			if err != nil {
				return fmt.Errorf("failed to query epoch: %w", err)
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
