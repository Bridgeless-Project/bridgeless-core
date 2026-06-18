package cli

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/swap/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdQuerySwapByID() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swap [chain-id] [tx_hash] [tx_nonce]",
		Short: "Query swap by its id",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			nonce, ok := big.NewInt(0).SetString(args[2], 10)
			if !ok {
				return errors.New(fmt.Sprintf("invalid nonce: %s", args[2]))
			}

			if nonce.Cmp(big.NewInt(0)) == -1 {
				return errors.New(fmt.Sprintf("negative nonce: %s", args[2]))
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.GetSwapById(context.Background(), &types.QueryGetSwapById{
				ChainId: args[0],
				TxHash:  args[1],
				TxNonce: nonce.Uint64(),
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQuerySwaps() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swaps",
		Short: "Query all swaps",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.AllSwaps(context.Background(), &types.QueryAllSwaps{Pagination: pageReq})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "swaps")

	return cmd
}
