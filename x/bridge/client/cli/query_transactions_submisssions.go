package cli

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
	"github.com/spf13/cobra"
)

func CmdQueryTransactionsSubmissions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transactions-submissions",
		Short: "Query bridge transactions submissions",
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
			res, err := queryClient.GetTxsSubmissions(cmd.Context(), &types.QueryGetTxsSubmissions{Pagination: pageReq})
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

func CmdQueryTransactionSubmissionsByHash() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transaction-submissions [tx_hash]",
		Short: "Query bridge transaction submissions by hash",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			_, err = hexutil.Decode(args[0])
			if err != nil {
				return errorsmod.Wrap(err, "invalid transaction hash")
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.GetTxSubmissionsByHash(cmd.Context(), &types.QueryGetTxSubmissionsByHash{
				TxHash: args[0],
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
