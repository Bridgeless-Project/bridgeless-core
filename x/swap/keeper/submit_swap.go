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

const (
	withdrawMethodERC20            = "withdrawERC20"
	depositMethodERC20             = "depositERC20"
	swapMethodExactTokensForTokens = "swapExactTokensForTokens"
)

func (k Keeper) executeSwap(ctx sdk.Context, msg *swaptypes.MsgSubmitSwapTx) (*swaptypes.SwapTransaction, error) {
	params := k.GetParams(ctx)
	if !common.IsHexAddress(params.UniswapRouterAddress) {
		return nil, errorsmod.Wrap(swaptypes.ErrInvalidConfig, "uniswap router address is not configured")
	}

	bridgeChain, found := k.bridge.GetChain(ctx, msg.Tx.Tx.WithdrawalChainId)
	if !found {
		return nil, errorsmod.Wrapf(bridgetypes.ErrChainNotFound, "withdrawal chain not found: %s", msg.Tx.Tx.WithdrawalChainId)
	}

	path, err := k.buildSwapPath(
		ctx,
		msg.Tx.Tx.WithdrawalChainId,
		msg.Tx.Tx.DepositToken,
		msg.Tx.Tx.WithdrawalToken,
	)
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

	if err = k.callBridgeWithdrawal(ctx, bridgeChain, msg.Tx); err != nil {
		return nil, errorsmod.Wrap(err, "failed to execute bridge withdrawal")
	}

	deadline := big.NewInt(ctx.BlockTime().Add(time.Duration(params.SwapDeadlineSeconds) * time.Second).Unix())
	routerResp, err := k.erc20.CallEVM(
		ctx,
		contracts.UniswapV2RouterV2Contract.ABI,
		swaptypes.ModuleAddress,
		common.HexToAddress(params.UniswapRouterAddress),
		true,
		swapMethodExactTokensForTokens,
		amountIn,
		amountOutMin,
		path,
		swaptypes.ModuleAddress,
		deadline,
	)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to execute uniswap swap")
	}

	finalAmount, err := unpackFinalAmount(routerResp.Ret)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to execute uniswap swap")
	}

	msg.Tx.FinalAmount = finalAmount

	if msg.IsBridgeTx {
		depositHash, err := k.callBridgeDeposit(ctx, bridgeChain, msg.Tx)
		if err != nil {
			return nil, errorsmod.Wrap(err, "failed to execute bridge deposit")
		}
		msg.Tx.FinalDepositTxHash = depositHash
	}

	return msg.Tx, nil
}

func (k Keeper) callBridgeWithdrawal(ctx sdk.Context, chain bridgetypes.Chain, swap *swaptypes.SwapTransaction) error {
	amount, err := parseUintString(swap.Tx.DepositAmount, "deposit amount")
	if err != nil {
		return errorsmod.Wrap(err, "failed to parse deposit amount")
	}

	bridgeAddr := common.HexToAddress(chain.BridgeAddress)
	txHash := common.HexToHash(swap.Tx.DepositTxHash)
	txNonce := new(big.Int).SetUint64(swap.Tx.DepositTxIndex)
	signatureBytes, err := hexutil.Decode(swap.Tx.Signature)
	if err != nil {
		return errorsmod.Wrap(err, "failed to decode signature")
	}
	signatures := [][]byte{signatureBytes}

	_, err = k.erc20.CallEVM(
		ctx,
		contracts.BridgeContract.ABI,
		swaptypes.ModuleAddress,
		bridgeAddr,
		true,
		withdrawMethodERC20,
		common.HexToAddress(swap.Tx.DepositToken),
		amount,
		swap.Tx.Receiver, // the module address
		txHash,
		txNonce,
		swap.Tx.IsWrapped,
		signatures,
	)
	if err != nil {
		return errorsmod.Wrap(err, "failed to execute bridge withdrawal")
	}

	return nil
}

func (k Keeper) callBridgeDeposit(ctx sdk.Context, chain bridgetypes.Chain, swap *swaptypes.SwapTransaction) (string, error) {
	amount, err := parseUintString(swap.FinalAmount, "final amount")
	if err != nil {
		return "", err
	}

	tokenInfo, found := k.bridge.GetTokenInfo(ctx, swap.Tx.WithdrawalChainId, swap.Tx.WithdrawalToken)
	if !found {
		return "", errorsmod.Wrapf(bridgetypes.ErrTokenInfoNotFound, "token info not found for %s on chain %s", swap.Tx.WithdrawalToken, swap.Tx.WithdrawalChainId)
	}

	txResp, err := k.erc20.CallEVM(
		ctx,
		contracts.BridgeContract.ABI,
		swaptypes.ModuleAddress,
		common.HexToAddress(chain.BridgeAddress),
		true,
		depositMethodERC20,
		common.HexToAddress(swap.Tx.WithdrawalToken),
		amount,
		swap.FinalReceiver,
		swap.Tx.WithdrawalChainId,
		tokenInfo.IsWrapped,
		uint16(swap.Tx.ReferralId),
	)
	if err != nil {
		return "", errorsmod.Wrap(err, "failed to execute bridge deposit")
	}

	return txResp.Hash, nil
}

func (k Keeper) buildSwapPath(ctx sdk.Context, destinationChain string, sourceToken string, destinationToken string) ([]common.Address, error) {
	params := k.GetParams(ctx)
	if !common.IsHexAddress(sourceToken) {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid source token address: %s", sourceToken)
	}
	if !common.IsHexAddress(destinationToken) {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid destination token address: %s", destinationToken)
	}
	if !common.IsHexAddress(params.WrappedBridge) {
		return nil, errorsmod.Wrap(swaptypes.ErrInvalidConfig, "wrapped bridge address is not configured")
	}

	sourceAddr := common.HexToAddress(sourceToken)
	wrappedBridgeAddr := common.HexToAddress(params.WrappedBridge)
	if params.WrappedBridge == destinationToken && destinationChain == ctx.ChainID() {
		return []common.Address{sourceAddr, wrappedBridgeAddr}, nil
	}

	// get the address of token on the Bridgeless
	token, found := k.bridge.GetDstToken(ctx, destinationToken, destinationChain, ctx.ChainID())
	if !found {
		return nil, errorsmod.Wrapf(bridgetypes.ErrTokenInfoNotFound, "no token info found for destination token %s on chain %s", destinationToken, destinationChain)
	}
	if !common.IsHexAddress(token.Address) {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid bridgeless token address: %s", token.Address)
	}

	return []common.Address{sourceAddr, wrappedBridgeAddr, common.HexToAddress(token.Address)}, nil
}

func unpackFinalAmount(ret []byte) (string, error) {
	unpacked, err := contracts.UniswapV2RouterV2Contract.ABI.Unpack(swapMethodExactTokensForTokens, ret)
	if err != nil {
		return "", errorsmod.Wrap(err, "failed to unpack uniswap swap output")
	}
	if len(unpacked) == 0 {
		return "", errorsmod.Wrap(swaptypes.ErrInvalidRoute, "uniswap swap returned no output amounts")
	}

	switch amounts := unpacked[0].(type) {
	case []*big.Int:
		if len(amounts) == 0 {
			return "", errorsmod.Wrap(swaptypes.ErrInvalidRoute, "uniswap swap returned an empty amounts array")
		}
		return amounts[len(amounts)-1].String(), nil
	case []interface{}:
		if len(amounts) == 0 {
			return "", errorsmod.Wrap(swaptypes.ErrInvalidRoute, "uniswap swap returned an empty amounts array")
		}

		last, ok := amounts[len(amounts)-1].(*big.Int)
		if !ok {
			return "", errorsmod.Wrapf(swaptypes.ErrInvalidRoute, "unexpected swap amount type %T", amounts[len(amounts)-1])
		}
		return last.String(), nil
	default:
		return "", errorsmod.Wrapf(swaptypes.ErrInvalidRoute, "unexpected uniswap output type %T", unpacked[0])
	}
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
