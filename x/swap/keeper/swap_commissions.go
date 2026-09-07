package keeper

import (
	"math/big"

	errorsmod "cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/contracts"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/swap/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/Zano-Execution-Layer/go-ethereum/common"
)

// retuns the amount of pair SourceToken -> BridgeToken
func (k Keeper) ComputeSwapPrice(ctx sdk.Context, sourceToken string, amountIn *big.Int) (*big.Int, []common.Address, error) {
	params := k.GetParams(ctx)

	pairPath := make([]common.Address, 0, 2)
	pairPath = append(pairPath, common.HexToAddress(sourceToken), common.HexToAddress(params.WrappedBridge))

	// TODO: add slipage
	resp, err := k.erc20.CallEVM(
		ctx,
		contracts.UniswapV2RouterV2Contract.ABI,
		common.HexToAddress(params.SwapperCallerAddress),
		common.HexToAddress(params.UniswapRouterAddress),
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
