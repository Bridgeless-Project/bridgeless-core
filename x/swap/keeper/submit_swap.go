package keeper

import (
	"math/big"

	errorsmod "cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/contracts"
	bridgetypes "github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	swaptypes "github.com/Bridgeless-Project/bridgeless-core/v12/x/swap/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const withdrawSwapAndRouteMethod = "withdrawSwapAndRoute"

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
		return nil, errorsmod.Wrapf(bridgetypes.ErrChainNotFound, "withdrawal chain not found: %s", msg.Tx.Tx.WithdrawalChainId)
	}

	//WithdrawalToken is the representation of deposited by user token
	finalDestinationTokenInfo, found := k.bridge.GetTokenInfo(ctx, msg.Tx.FinalChainId, msg.Tx.FinalToken)
	if !found {
		return nil, errorsmod.Wrapf(bridgetypes.ErrTokenInfoNotFound, "token info not found for %s on chain %s", msg.Tx.FinalToken, msg.Tx.FinalChainId)
	}

	// prepare the swap params
	// There we build the path (WithdrawalToken -> WrappedBridge -> FinalTokenOnBridgeless) and the swap params for the swapper contract call
	// if one of WithdrawalToken or FinalTokenOnBridgeless is WrappedBridge, the final path consist of 2 addresses only
	path, err := k.buildSwapPath(ctx, msg.Tx.Tx.WithdrawalToken, msg.Tx.FinalToken, msg.Tx.FinalChainId)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to build swap path")
	}

	amountIn, err := parseUintString(msg.Tx.Tx.DepositAmount)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to parse deposit amount")
	}

	amountOutMin, err := parseUintString(msg.Tx.SwapOutAmount)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to parse amount_out_min")
	}

	signatureBytes, err := hexutil.Decode(msg.Tx.Tx.Signature)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to decode signature")
	}

	txResp, err := k.erc20.CallEVM(
		ctx,
		contracts.SwapperContract.ABI,
		common.HexToAddress(params.SwapperCallerAddress), // the address which calls swapper contract
		common.HexToAddress(params.SwapperAddress),
		true,
		withdrawSwapAndRouteMethod,
		swaptypes.SwapperWithdrawParams{
			Token:      common.HexToAddress(msg.Tx.Tx.WithdrawalToken),
			Amount:     amountIn,
			TxHash:     txHashToBytes32(msg.Tx.Tx.DepositTxHash),
			TxNonce:    new(big.Int).SetUint64(msg.Tx.Tx.DepositTxIndex),
			IsWrapped:  msg.Tx.Tx.IsWrapped,
			Signatures: [][]byte{signatureBytes},
		},
		swaptypes.SwapperSwapParams{
			AmountIn:                 amountIn,
			MinDestinationAmount:     amountOutMin,
			SwapDeadline:             new(big.Int).SetUint64(msg.Tx.SwapDeadline),
			Path:                     path,
			IsDestinationTokenNative: isZeroAddress(finalDestinationTokenInfo.Address),
		},
		swaptypes.SwapperDepositParams{
			Receiver:   msg.Tx.FinalReceiver,
			Network:    msg.Tx.FinalChainId,
			IsWrapped:  finalDestinationTokenInfo.IsWrapped,
			ReferralId: uint16(msg.Tx.Tx.ReferralId)},
		swaptypes.SwapperDepositParams{
			Receiver:   msg.Tx.Tx.Depositor,
			Network:    msg.Tx.Tx.DepositChainId,
			IsWrapped:  msg.Tx.Tx.IsWrapped,
			ReferralId: uint16(msg.Tx.Tx.ReferralId)},
	)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to execute swapper withdraw swap and route")
	}

	k.Logger(ctx).Info("swap executed successfully", "txHash", txResp.Hash)
	msg.Tx.FinalDepositTxHash = txResp.Hash
	return msg.Tx, nil
}

func (k Keeper) buildSwapPath(ctx sdk.Context, sourceToken string, destinationToken string, destinationChain string) ([]common.Address, error) {
	params := k.GetParams(ctx)
	if !common.IsHexAddress(sourceToken) {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid source token address: %s", sourceToken)
	}
	if !common.IsHexAddress(params.WrappedBridge) {
		return nil, errorsmod.Wrap(swaptypes.ErrInvalidConfig, "wrapped bridge address is not configured")
	}

	token, found := k.bridge.GetDstToken(ctx, destinationToken, destinationChain, getChainId(ctx))
	if !found {
		return nil, errorsmod.Wrapf(bridgetypes.ErrTokenInfoNotFound, "no token info found for destination token %s on chain %s", destinationToken, getChainId(ctx))
	}
	if !common.IsHexAddress(token.Address) {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid bridgeless token address: %s", token.Address)
	}

	// if one of tokens is WrappedBridge, we can skip it in the path and
	// swap directly between the other token and WrappedBridge
	if sourceToken == params.WrappedBridge || isZeroAddress(sourceToken) {
		return []common.Address{
			common.HexToAddress(sourceToken),
			common.HexToAddress(token.Address),
		}, nil
	}
	if token.Address == params.WrappedBridge || isZeroAddress(token.Address) {
		return []common.Address{
			common.HexToAddress(sourceToken),
			common.HexToAddress(token.Address),
		}, nil
	}

	return []common.Address{
		common.HexToAddress(sourceToken),
		common.HexToAddress(params.WrappedBridge),
		common.HexToAddress(token.Address),
	}, nil
}
