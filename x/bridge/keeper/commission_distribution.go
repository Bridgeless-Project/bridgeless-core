package keeper

import (
	"cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/utils"
	bridgetypes "github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetCommissionPrices(ctx sdk.Context, epochId uint32) ([]bridgetypes.CommissionDistributionInfo, error) {
	commissionsInfo := make([]bridgetypes.CommissionDistributionInfo, 0)
	for _, commission := range k.GetAllCommissions(ctx, epochId) {
		amount, found := sdk.NewIntFromString(commission.Amount)
		if !found {
			continue
		}

		info, err := k.tokenOnBridgeless(ctx, commission.TokenId)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get token info for commission token")
		}

		amountOut, path, err := k.hooks.GetTokenPrice(ctx, info.Address, amount.BigInt())
		if err != nil {
			return nil, errors.Wrap(err, "failed to get token price")
		}

		pathStr := make([]string, 0)
		for _, address := range path {
			pathStr = append(pathStr, address.String())
		}

		commissionsInfo = append(commissionsInfo, bridgetypes.CommissionDistributionInfo{
			AmountIn:  amount.String(),
			AmountOut: amountOut.String(),
			Path:      pathStr,
		})
	}

	return commissionsInfo, nil
}

func (k Keeper) tokenOnBridgeless(ctx sdk.Context, tokenId uint64) (*bridgetypes.TokenInfo, error) {
	token, found := k.GetToken(ctx, tokenId)
	if !found {
		return nil, errors.Wrap(bridgetypes.ErrTokenInfoNotFound, "token info not found for commission token")
	}

	for _, tokenInfo := range token.Info {
		if tokenInfo.ChainId == utils.GetChainId(ctx) {
			return &tokenInfo, nil
		}
	}

	return nil, errors.Wrap(bridgetypes.ErrTokenInfoNotFound, "token not found")
}
