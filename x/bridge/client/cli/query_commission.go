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
		Use:   "commission [token-id] [epoch-id]",
		Short: "shows the commission by tokenId and epochId",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			tokenId, ok := big.NewInt(0).SetString(args[0], 10)
			if !ok {
				return errors.Wrap(types.ErrInvalidDataType, "token-id must be a valid integer")
			}

			epochId, ok := big.NewInt(0).SetString(args[1], 10)
			if !ok {
				return errors.Wrap(types.ErrInvalidDataType, "epoch-id must be a valid integer")
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.GetCommissionByToken(context.Background(), &types.QueryGetCommissionByToken{
				TokenId: tokenId.Uint64(),
				EpochId: uint32(epochId.Uint64()),
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
