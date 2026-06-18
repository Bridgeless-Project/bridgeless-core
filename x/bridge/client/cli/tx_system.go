package cli

import (
	"strconv"

	errorsmod "cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func CmdDistributeFees() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "distribute-fees [from_key_or_address] [epoch_id]",
		Short: "Distribute bridge fees for an epoch",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Set(flags.FlagFrom, args[0])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			epochId, err := strconv.ParseUint(args[1], 10, 32)
			if err != nil {
				return errorsmod.Wrap(err, "failed to parse epoch ID")
			}

			msg := types.NewMsgDistributeFees(clientCtx.GetFromAddress().String(), uint32(epochId))

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdProcessSystemWithdrawal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "process-system-withdrawal [from_key_or_address] [path-to-system-withdrawals-json]",
		Short: "Process system withdrawals from a JSON file",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Set(flags.FlagFrom, args[0])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			withdrawals, err := parseSystemWithdrawals(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgProcessSystemWithdrawal(clientCtx.GetFromAddress().String(), withdrawals...)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
