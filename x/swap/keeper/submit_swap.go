package keeper

import (
	"math/big"
	"time"

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

type swapperWithdrawParams struct {
	Token      common.Address
	Amount     *big.Int
	TxHash     common.Hash
	TxNonce    *big.Int
	IsWrapped  bool
	Signatures [][]byte
}

type swapperSwapParams struct {
	AmountIn                 *big.Int
	MinDestinationAmount     *big.Int
	SwapDeadline             *big.Int
	Path                     []common.Address
	IsDestinationTokenNative bool
}

type swapperDepositParams struct {
	Receiver   string
	Network    string
	IsWrapped  bool
	ReferralId uint16
}

func (k Keeper) executeSwap(ctx sdk.Context, msg *swaptypes.MsgSubmitSwapTx) (*swaptypes.SwapTransaction, error) {
	params := k.GetParams(ctx)
	if !common.IsHexAddress(params.SwapperAddress) {
		return nil, errorsmod.Wrap(swaptypes.ErrInvalidConfig, "swapper address is not configured")
	}
	if !common.IsHexAddress(params.WrappedBridge) {
		return nil, errorsmod.Wrap(swaptypes.ErrInvalidConfig, "wrapped bridge address is not configured")
	}

	if _, found := k.bridge.GetChain(ctx, msg.Tx.Tx.WithdrawalChainId); !found {
		return nil, errorsmod.Wrapf(bridgetypes.ErrChainNotFound, "withdrawal chain not found: %s", msg.Tx.Tx.WithdrawalChainId)
	}

	destinationInfo, found := k.bridge.GetTokenInfo(ctx, msg.Tx.Tx.WithdrawalChainId, msg.Tx.Tx.WithdrawalToken)
	if !found {
		return nil, errorsmod.Wrapf(bridgetypes.ErrTokenInfoNotFound, "token info not found for %s on chain %s", msg.Tx.Tx.WithdrawalToken, msg.Tx.Tx.WithdrawalChainId)
	}

	path, err := k.buildSwapPath(ctx, msg.Tx.Tx.DepositToken, msg.Tx.Tx.WithdrawalToken, msg.Tx.Tx.WithdrawalChainId)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to build swap path")
	}

	amountIn, err := parseUintString(msg.Tx.Tx.DepositAmount, "deposit amount")
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to parse deposit amount")
	}

	amountOutMin, err := parseUintString(msg.Tx.AmountOutMin, "amount_out_min")
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to parse amount_out_min")
	}

	signatureBytes, err := hexutil.Decode(msg.Tx.Tx.Signature)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to decode signature")
	}

	deadlineSeconds := params.SwapDeadlineSeconds
	if deadlineSeconds == 0 {
		deadlineSeconds = swaptypes.DefaultSwapDeadlineSeconds
	}
	deadline := big.NewInt(ctx.BlockTime().Add(time.Duration(deadlineSeconds) * time.Second).Unix())

	txResp, err := k.erc20.CallEVM(
		ctx,
		contracts.SwapperContract.ABI,
		swaptypes.ModuleAddress,
		common.HexToAddress(params.SwapperAddress),
		true,
		withdrawSwapAndRouteMethod,
		swapperWithdrawParams{
			Token:      common.HexToAddress(msg.Tx.Tx.DepositToken),
			Amount:     amountIn,
			TxHash:     common.HexToHash(msg.Tx.Tx.DepositTxHash),
			TxNonce:    new(big.Int).SetUint64(msg.Tx.Tx.DepositTxIndex),
			IsWrapped:  msg.Tx.Tx.IsWrapped,
			Signatures: [][]byte{signatureBytes},
		},
		swapperSwapParams{
			AmountIn:                 amountIn,
			MinDestinationAmount:     amountOutMin,
			SwapDeadline:             deadline,
			Path:                     path,
			IsDestinationTokenNative: isZeroAddress(destinationInfo.Address),
		},
		swapperDepositParams{
			Receiver:   msg.Tx.FinalReceiver,
			Network:    msg.Tx.Tx.WithdrawalChainId,
			IsWrapped:  destinationInfo.IsWrapped,
			ReferralId: uint16(msg.Tx.Tx.ReferralId),
		},
		swapperDepositParams{
			Receiver:   msg.Tx.Tx.TxData,
			Network:    msg.Tx.Tx.DepositChainId,
			IsWrapped:  msg.Tx.Tx.IsWrapped,
			ReferralId: uint16(msg.Tx.Tx.ReferralId),
		},
	)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to execute swapper withdraw swap and route")
	}

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

	token, found := k.bridge.GetDstToken(ctx, destinationToken, destinationChain, ctx.ChainID())
	if !found {
		return nil, errorsmod.Wrapf(bridgetypes.ErrTokenInfoNotFound, "no token info found for destination token %s on chain %s", destinationToken, destinationChain)
	}
	if !common.IsHexAddress(token.Address) {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid bridgeless token address: %s", token.Address)
	}

	return []common.Address{
		common.HexToAddress(sourceToken),
		common.HexToAddress(params.WrappedBridge),
		common.HexToAddress(token.Address),
	}, nil
}

func parseUintString(value, field string) (*big.Int, error) {
	parsed, ok := new(big.Int).SetString(value, 10)
	if !ok {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid %s: %s", field, value)
	}
	if parsed.Sign() < 0 {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "%s cannot be negative: %s", field, value)
	}

	return parsed, nil
}

func isZeroAddress(address string) bool {
	return common.HexToAddress(address) == (common.Address{})
}
