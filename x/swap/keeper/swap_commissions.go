package keeper

import (
	"math/big"

	errorsmod "cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/contracts"
	bridgetypes "github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/swap/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

func (k Keeper) computeCommission(ctx sdk.Context, tx *types.SwapTransaction) (*bridgetypes.Commission, error) {
	if !tx.IsFeeDistribution {
		return nil, nil
	}
	depositTokenInfo, found := k.bridge.GetTokenInfo(ctx, tx.Tx.DepositToken, tx.Tx.DepositChainId)
	if !found {
		return nil, errorsmod.Wrapf(bridgetypes.ErrTokenInfoNotFound, "token info not found for %s on chain %s", tx.Tx.WithdrawalToken, tx.Tx.WithdrawalChainId)
	}

	commission, found := k.bridge.GetCommission(ctx, tx.Tx.EpochId, depositTokenInfo.TokenId)
	if !found {
		return nil, errorsmod.Wrapf(bridgetypes.ErrCommissionNotFound, "commission not found for token %s", depositTokenInfo.TokenId)
	}

	commissionAmount, ok := new(big.Int).SetString(commission.Amount, 10)
	if !ok {
		return nil, errorsmod.Wrapf(bridgetypes.ErrInvalidCommission, "invalid commission amount: %s", commission.Amount)
	}

	withdrawalAmount, ok := new(big.Int).SetString(commission.Amount, 10)
	if !ok {
		return nil, errorsmod.Wrapf(bridgetypes.ErrInvalidAmount, "invalid withdrawal amount: %s", tx.Tx.WithdrawalAmount)
	}

	commissionAmount.Sub(commissionAmount, withdrawalAmount)
	if commissionAmount.Sign() < 0 {
		return nil, errorsmod.Wrapf(bridgetypes.ErrInvalidCommission, "withdrawal amount %s exceeds commission amount %s", withdrawalAmount.String(), commission.Amount)
	}

	commission.Amount = commissionAmount.String()
	return &commission, nil
}

// retuns the amount of pair SourceToken -> BridgeToken
func (k Keeper) ComputeSwapPrice(ctx sdk.Context, sourceToken string, amountIn *big.Int) (*big.Int, []common.Address, error) {
	params := k.GetParams(ctx)

	pairPath := make([]common.Address, 0, 2)
	pairPath = append(pairPath, common.HexToAddress(sourceToken), common.HexToAddress(params.WrappedBridge))

	// TODO: add slipage
	resp, err := k.erc20.CallEVM(
		ctx,
		contracts.UniswapV2RouterV2Contract.ABI,
		common.HexToAddress(params.SwapperAddress),
		common.HexToAddress(params.SwapperCallerAddress),
		false,
		getAmountsOutMethod,
		amountIn,
		pairPath,
	)
	if err != nil {
		return nil, nil, errorsmod.Wrap(err, "failed to compute price")
	}

	var amountRes types.AmountsListResponse
	if err = contracts.UniswapV2RouterV2Contract.ABI.UnpackIntoInterface(&amountRes, getAmountsOutMethod, resp.Ret); err != nil {
		return nil, nil, errorsmod.Wrap(err, "failed to extract amount from response")
	}

	// the getAmountsOutMethod returns the list of prices [In, Out].
	// this function MUST return second element of array
	if len(amountRes.Amounts) < 2 {
		return nil, nil, errorsmod.Wrapf(types.ErrInvalidRoute, "unexpected number of amounts in response: expected 2, got %d", len(amountRes.Amounts))
	}

	return amountRes.Amounts[1], pairPath, nil
}
