package cli

import (
	"context"
	"strconv"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdQueryChainById() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "chain [id]",
		Short: "shows the chain info by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)
			req := &types.QueryGetChainById{
				Id: args[0],
			}
			res, err := queryClient.GetChainById(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryChains() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "chains",
		Short: "Query all chains",
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

			req := &types.QueryGetChains{
				Pagination: pageReq,
			}

			res, err := queryClient.GetChains(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryChainsByType() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "chains-by-type [chain-type]",
		Short: "Query chains by type",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			chainType, err := strconv.ParseUint(args[0], 10, 32)
			if err != nil {
				return err
			}

			req := &types.QueryGetChainsByType{
				ChainType:  types.ChainType(chainType),
				Pagination: pageReq,
			}

			res, err := queryClient.GetChainsByType(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "chains-by-type")

	return cmd
}
