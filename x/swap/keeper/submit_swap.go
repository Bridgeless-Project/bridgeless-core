package keeper

import (
	"math/big"
	"strconv"

	errorsmod "cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/contracts"
	"github.com/Bridgeless-Project/bridgeless-core/v12/utils"
	bridgetypes "github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	swaptypes "github.com/Bridgeless-Project/bridgeless-core/v12/x/swap/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	getAmountsOutMethod        = "getAmountsOut"
	withdrawSwapAndRouteMethod = "withdrawSwapAndRoute"
)

func (k Keeper) executeSwap(ctx sdk.Context, msg *swaptypes.MsgSubmitSwapTx) (*swaptypes.SwapTransaction, error) {
	params := k.GetParams(ctx)
	if !common.IsHexAddress(params.SwapperAddress) {
		return nil, errorsmod.Wrap(swaptypes.ErrInvalidConfig, "swapper address is not configured")
	}
	if !common.IsHexAddress(params.WrappedBridge) {
		return nil, errorsmod.Wrap(swaptypes.ErrInvalidConfig, "wrapped bridge address is not configured")
	}

	// its already bridgeless networks
	if _, found := k.bridge.GetChain(ctx, msg.Tx.Tx.WithdrawalChainId); !found {
		return nil, errorsmod.Wrapf(bridgetypes.ErrChainNotFound, "withdrawal chain not found: %d", msg.Tx.Tx.WithdrawalChainId)
	}

	//WithdrawalToken is the representation of deposited by user token
	finalDestinationTokenInfo, found := k.bridge.GetTokenInfo(ctx, msg.Tx.FinalChainId, msg.Tx.FinalToken)
	if !found {
		return nil, errorsmod.Wrapf(bridgetypes.ErrTokenInfoNotFound, "token info not found for %s on chain %d", msg.Tx.FinalToken, msg.Tx.FinalChainId)
	}

	// prepare the swap params
	// There we build the path (WithdrawalToken -> WrappedBridge -> FinalTokenOnBridgeless) and the swap params for the swapper contract call
	// if one of WithdrawalToken or FinalTokenOnBridgeless is WrappedBridge, the final path consist of 2 addresses only
	path, err := k.buildSwapPath(ctx, msg.Tx.Tx.WithdrawalToken, msg.Tx.FinalToken, msg.Tx.FinalChainId)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to build swap path")
	}

	amountIn, err := utils.ParseUintString(msg.Tx.Tx.WithdrawalAmount)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to parse deposit amount")
	}

	amountOutMin, err := utils.ParseUintString(msg.Tx.SwapOutAmount)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to parse amount_out_min")
	}

	signatureBytes, err := hexutil.Decode(msg.Tx.Tx.Signature)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to decode signature")
	}

	txResp, err := k.erc20.CallEVMAsTx(
		ctx,
		contracts.SwapperContract.ABI,
		common.HexToAddress(params.SwapperCallerAddress), // the address which calls swapper contract
		common.HexToAddress(params.SwapperAddress),
		true,
		withdrawSwapAndRouteMethod,
		swaptypes.SwapperWithdrawParams{
			Token:      common.HexToAddress(msg.Tx.Tx.WithdrawalToken),
			Amount:     amountIn,
			TxHash:     utils.TxHashToBytes32(msg.Tx.Tx.DepositTxHash),
			TxNonce:    new(big.Int).SetUint64(msg.Tx.Tx.DepositTxIndex),
			IsWrapped:  msg.Tx.Tx.IsWrapped,
			Signatures: [][]byte{signatureBytes},
		},
		swaptypes.SwapperSwapParams{
			AmountIn:                 amountIn,
			MinDestinationAmount:     amountOutMin,
			SwapDeadline:             new(big.Int).SetUint64(msg.Tx.SwapDeadline),
			Path:                     path,
			IsDestinationTokenNative: utils.IsZeroAddress(finalDestinationTokenInfo.Address) && finalDestinationTokenInfo.ChainId == utils.GetChainId(ctx),
		},
		swaptypes.SwapperDepositParams{
			Receiver:   msg.Tx.FinalReceiver,
			Network:    strconv.FormatUint(uint64(msg.Tx.FinalChainId), 10),
			IsWrapped:  finalDestinationTokenInfo.IsWrapped,
			ReferralId: uint16(msg.Tx.Tx.ReferralId),
		},
		swaptypes.SwapperDepositParams{
			Receiver:   msg.Tx.Tx.Depositor,
			Network:    strconv.FormatUint(uint64(msg.Tx.Tx.DepositChainId), 10),
			IsWrapped:  msg.Tx.Tx.IsWrapped,
			ReferralId: uint16(msg.Tx.Tx.ReferralId),
		},
	)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to execute swapper withdraw swap and route")
	}

	emitSubmitEvent(ctx, msg.Tx.Tx)
	k.Logger(ctx).Info("swap executed successfully", "txHash", txResp.Hash)
	msg.Tx.FinalDepositTxHash = txResp.Hash
	return msg.Tx, nil
}

func (k Keeper) buildSwapPath(ctx sdk.Context, sourceToken string, destinationToken string, destinationChain uint32) ([]common.Address, error) {
	params := k.GetParams(ctx)
	if !common.IsHexAddress(sourceToken) {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid source token address: %s", sourceToken)
	}
	if !common.IsHexAddress(params.WrappedBridge) {
		return nil, errorsmod.Wrap(swaptypes.ErrInvalidConfig, "wrapped bridge address is not configured")
	}

	dstToken, found := k.bridge.GetDstToken(ctx, destinationToken, destinationChain, utils.GetChainId(ctx))
	if !found {
		return nil, errorsmod.Wrapf(bridgetypes.ErrTokenInfoNotFound, "no token info found for destination token %s on chain %d", destinationToken, utils.GetChainId(ctx))
	}
	if !common.IsHexAddress(dstToken.Address) {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid bridgeless token address: %s", dstToken.Address)
	}

	sourceAddr := common.HexToAddress(sourceToken)
	wrappedAddr := common.HexToAddress(params.WrappedBridge)
	dstAddr := common.HexToAddress(dstToken.Address)
	// if one of tokens is WrappedBridge, we can skip it in the path and
	// swap directly between the other token and WrappedBridge

	if sourceAddr == wrappedAddr || sourceAddr == (common.Address{}) {
		return []common.Address{sourceAddr, dstAddr}, nil
	}

	if dstAddr == wrappedAddr || dstAddr == (common.Address{}) {
		return []common.Address{sourceAddr, dstAddr}, nil
	}

	return []common.Address{sourceAddr, wrappedAddr, dstAddr}, nil
}
