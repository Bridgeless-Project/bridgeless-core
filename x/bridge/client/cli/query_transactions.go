package cli

import (
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
	"github.com/spf13/cobra"
	"math/big"
)

func CmdQueryTransactions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transactions",
		Short: "Query bridge transactions",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Transactions(cmd.Context(), &types.QueryTransactionsRequest{Pagination: pageReq})
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

func CmdQueryTransactionById() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transaction [chain-id] [tx_hash] [tx_nonce]",
		Short: "Query bridge transaction by its id",
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
			res, err := queryClient.TransactionById(cmd.Context(), &types.QueryTransactionByIdRequest{
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
