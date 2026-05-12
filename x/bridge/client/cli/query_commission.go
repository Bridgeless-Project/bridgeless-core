package cli

import (
	"context"
	"math/big"

	"cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdQueryGetCommissionByToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "commission [token-id]",
		Short: "shows the commission by tokenId",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			tokenId, ok := big.NewInt(0).SetString(args[0], 10)
			if !ok {
				return errors.Wrap(types.ErrInvalidDataType, "token-id must be a valid integer")
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.GetCommissionByToken(context.Background(), &types.QueryGetCommissionByToken{TokenId: tokenId.Uint64()})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
