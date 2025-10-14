package cli

import (
	"context"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/multisig/types"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

func CmdQueryGroups() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "groups",
		Short: "Query multisig groups",
		Args:  cobra.NoArgs,
	}

	cmd.AddCommand(
		QueryAllGroups(), QueryGroup())

	return cmd
}

func QueryAllGroups() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-all",
		Short: "Lists all multisig groupss",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.GroupAll(context.Background(), &types.QueryAllGroupRequest{})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	return cmd
}

func QueryGroup() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "group [group_account_addr]",
		Short: "Query multisig group",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)
			groupAddr := args[0]
			_, err := sdk.AccAddressFromBech32(groupAddr)
			if err != nil {
				return err
			}
			res, err := queryClient.Group(context.Background(), &types.QueryGetGroupRequest{
				groupAddr,
			})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	return cmd
}
