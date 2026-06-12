package keeper

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/contracts"
	contractstypes "github.com/Bridgeless-Project/bridgeless-core/v12/contracts/types"
	"github.com/Bridgeless-Project/bridgeless-core/v12/utils"
	bridgetypes "github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	swaptypes "github.com/Bridgeless-Project/bridgeless-core/v12/x/swap/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

const contractEventSwappedAndRouted = "SwappedAndRouted"

// PostTxProcessing stores swaps initiated directly on the configured Swapper
// contract. It intentionally does not emit Cosmos SDK events.
func (k Keeper) PostTxProcessing(ctx sdk.Context, _ core.Message, receipt *ethtypes.Receipt) error {
	if receipt == nil || len(receipt.Logs) == 0 {
		k.Logger(ctx).Info("skipping swap EVM hook: receipt is nil or has no logs")
		return nil
	}

	swapperAddress := k.GetParams(ctx).SwapperAddress
	if !common.IsHexAddress(swapperAddress) {
		k.Logger(ctx).Error("skipping swap EVM hook: invalid swapper address", "address", swapperAddress)
		return nil
	}

	contractAddress := common.HexToAddress(swapperAddress)
	k.Logger(ctx).Info(
		"processing EVM receipt for swap events",
		"tx_hash", receipt.TxHash.Hex(),
		"log_count", len(receipt.Logs),
		"swapper_address", contractAddress.Hex(),
	)
	for _, evmLog := range receipt.Logs {
		if evmLog == nil {
			k.Logger(ctx).Error("skipping nil EVM log", "tx_hash", receipt.TxHash.Hex())
			continue
		}
		if evmLog.Address != contractAddress || len(evmLog.Topics) == 0 {
			continue
		}

		event, err := contracts.SwapperContract.ABI.EventByID(evmLog.Topics[0])
		if err != nil {
			k.Logger(ctx).Error(
				"failed to resolve swapper contract event",
				"tx_hash", evmLog.TxHash.Hex(),
				"log_index", evmLog.Index,
				"error", err,
			)
			continue
		}

		if event.Name != contractEventSwappedAndRouted {
			continue
		}

		k.Logger(ctx).Info(
			"matched SwappedAndRouted event",
			"tx_hash", evmLog.TxHash.Hex(),
			"log_index", evmLog.Index,
			"block", evmLog.BlockNumber,
		)

		eventBody := contractstypes.SwapperSwappedAndRouted{}
		if err = utils.UnpackLog(contracts.SwapperContract.ABI, &eventBody, event.Name, evmLog); err != nil {
			k.Logger(ctx).Error(
				errorsmod.Wrap(err, "failed to unpack SwappedAndRouted event").Error(),
				"tx_hash", evmLog.TxHash.Hex(),
				"log_index", evmLog.Index,
			)
			continue
		}

		if len(eventBody.SwapParams.Path) == 0 {
			k.Logger(ctx).Error(
				"skipping SwappedAndRouted event with empty swap path",
				"tx_hash", evmLog.TxHash.Hex(),
				"log_index", evmLog.Index,
			)
			continue
		}

		txHash := evmLog.TxHash.Hex()
		txIndex := uint64(evmLog.Index)
		chainID := utils.GetChainId(ctx)
		if _, found := k.GetSwap(ctx, txHash, txIndex, chainID); found {
			k.Logger(ctx).Info(
				"swap transaction already stored",
				"tx_hash", txHash,
				"log_index", txIndex,
				"chain_id", chainID,
			)
			continue
		}

		finalToken := eventBody.SwapParams.Path[len(eventBody.SwapParams.Path)-1]
		swap := swaptypes.SwapTransaction{
			Tx: bridgetypes.Transaction{
				DepositChainId:    chainID,
				DepositTxHash:     txHash,
				DepositTxIndex:    txIndex,
				DepositBlock:      evmLog.BlockNumber,
				DepositToken:      eventBody.SwapParams.Path[0].Hex(),
				DepositAmount:     eventBody.SwapParams.AmountIn.String(),
				Depositor:         eventBody.Sender.Hex(),
				Receiver:          swapperAddress,
				WithdrawalChainId: eventBody.DestinationDepositParams.Network,
				WithdrawalTxHash:  txHash, // set same tx hash
				WithdrawalToken:   finalToken.Hex(),
				IsWrapped:         eventBody.DestinationDepositParams.IsWrapped,
				WithdrawalAmount:  eventBody.SwapParams.MinDestinationAmount.String(),
				CommissionAmount:  "0", // we dont pay commission on this step ( user will pay it next step )
				ReferralId:        uint32(eventBody.DestinationDepositParams.ReferralId),
			},
			FinalReceiver:      eventBody.DestinationDepositParams.Receiver,
			FinalToken:         finalToken.Hex(),
			FinalChainId:       eventBody.DestinationDepositParams.Network,
			SwapDeadline:       eventBody.SwapParams.SwapDeadline.Uint64(),
			SwapOutAmount:      eventBody.SwapParams.MinDestinationAmount.String(),
			FinalDepositTxHash: txHash, // set same tx hash
		}
		k.SetSwap(ctx, swap)
		k.Logger(ctx).Info(
			"stored SwappedAndRouted transaction",
			"tx_hash", txHash,
			"log_index", txIndex,
			"chain_id", chainID,
			"sender", eventBody.Sender.Hex(),
			"destination_chain", eventBody.DestinationDepositParams.Network,
			"destination_receiver", eventBody.DestinationDepositParams.Receiver,
		)
	}

	return nil
}
