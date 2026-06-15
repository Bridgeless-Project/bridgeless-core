package keeper

import (
	"math/big"

	errorsmod "cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
)

// The commission amount decimals is mapped to brideless decimals (18)

func (k Keeper) setCommission(sdkCtx sdk.Context, epochId uint32, commission types.Commission) {
	cStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreCommissionPrefix))
	eStore := prefix.NewStore(cStore, types.KeyEpoch(epochId))

	eStore.Set(types.KeyEpochCommission(epochId, commission.TokenId), k.cdc.MustMarshal(&commission))
}

// SetCommissionNormalized stores a non-negative commission amount with 18 decimals.
func (k Keeper) SetCommissionNormalized(ctx sdk.Context, epochId uint32, tokenId uint64, amount *big.Int) error {
	if amount == nil || amount.Sign() < 0 {
		return errorsmod.Wrap(types.ErrInvalidCommission, "amount must be non-negative")
	}

	k.setCommission(ctx, epochId, types.Commission{
		TokenId: tokenId,
		Amount:  amount.String(),
	})
	return nil
}

func (k Keeper) AddCommissionNormalized(ctx sdk.Context, epochId uint32, tokenId uint64, amount *big.Int) error {
	if amount == nil || amount.Sign() < 0 {
		return errorsmod.Wrap(types.ErrInvalidCommission, "amount must be non-negative")
	}

	total := new(big.Int).Set(amount)
	if commission, found := k.GetCommission(ctx, epochId, tokenId); found {
		current, ok := new(big.Int).SetString(commission.Amount, 10)
		if !ok || current.Sign() < 0 {
			return errorsmod.Wrapf(types.ErrInvalidCommission, "invalid stored amount %q", commission.Amount)
		}
		total.Add(total, current)
	}

	return k.SetCommissionNormalized(ctx, epochId, tokenId, total)
}

func (k Keeper) AddCommissionNative(ctx sdk.Context, epochId uint32, token types.TokenInfo, amount *big.Int) error {
	if amount == nil || amount.Sign() < 0 {
		return errorsmod.Wrap(types.ErrInvalidCommission, "native amount must be non-negative")
	}

	return k.AddCommissionNormalized(
		ctx,
		epochId,
		token.TokenId,
		TransformAmount(amount, token.Decimals, types.DefaultChainDecimals),
	)
}

func (k Keeper) SubtractCommissionNative(ctx sdk.Context, epochId uint32, token types.TokenInfo, amount *big.Int) error {
	if amount == nil || amount.Sign() < 0 {
		return errorsmod.Wrap(types.ErrInvalidCommission, "native amount must be non-negative")
	}

	commission, found := k.GetCommission(ctx, epochId, token.TokenId)
	if !found {
		return errorsmod.Wrapf(types.ErrCommissionNotFound, "token %d", token.TokenId)
	}

	current, ok := new(big.Int).SetString(commission.Amount, 10)
	if !ok || current.Sign() < 0 {
		return errorsmod.Wrapf(types.ErrInvalidCommission, "invalid stored amount %q", commission.Amount)
	}

	deduction := TransformAmount(amount, token.Decimals, types.DefaultChainDecimals)
	if current.Cmp(deduction) < 0 {
		return errorsmod.Wrapf(
			types.ErrInvalidCommission,
			"native deduction %s exceeds stored commission %s",
			amount.String(),
			commission.Amount,
		)
	}

	return k.SetCommissionNormalized(ctx, epochId, token.TokenId, current.Sub(current, deduction))
}

func (k Keeper) SubtractCommissionNormalized(ctx sdk.Context, epochId uint32, tokenId uint64, amount *big.Int) error {
	if amount == nil || amount.Sign() < 0 {
		return errorsmod.Wrap(types.ErrInvalidCommission, "native amount must be non-negative")
	}

	commission, found := k.GetCommission(ctx, epochId, tokenId)
	if !found {
		return errorsmod.Wrapf(types.ErrCommissionNotFound, "token %d", tokenId)
	}

	current, ok := new(big.Int).SetString(commission.Amount, 10)
	if !ok || current.Sign() < 0 {
		return errorsmod.Wrapf(types.ErrInvalidCommission, "invalid stored amount %q", commission.Amount)
	}

	if current.Cmp(amount) < 0 {
		return errorsmod.Wrapf(
			types.ErrInvalidCommission,
			"native deduction %s exceeds stored commission %s",
			amount.String(),
			commission.Amount,
		)
	}

	return k.SetCommissionNormalized(ctx, epochId, tokenId, current.Sub(current, amount))
}

func (k Keeper) GetCommission(sdkCtx sdk.Context, epochId uint32, tokenId uint64) (types.Commission, bool) {
	cStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreCommissionPrefix))
	eStore := prefix.NewStore(cStore, types.KeyEpoch(epochId))

	var commission types.Commission
	bz := eStore.Get(types.KeyEpochCommission(epochId, tokenId))
	if bz == nil {
		return commission, false
	}
	k.cdc.MustUnmarshal(bz, &commission)
	return commission, true
}

func (k Keeper) RemoveCommission(sdkCtx sdk.Context, epochId uint32, tokenId uint64) {
	cStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreCommissionPrefix))
	eStore := prefix.NewStore(cStore, types.KeyEpoch(epochId))

	eStore.Delete(types.KeyEpochCommission(epochId, tokenId))
}

func (k Keeper) GetCommissionsWithPagination(sdkCtx sdk.Context, epochId uint32, pagination *query.PageRequest) ([]types.Commission, *query.PageResponse, error) {
	cStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreCommissionPrefix))
	eStore := prefix.NewStore(cStore, types.KeyEpoch(epochId))

	var commissions []types.Commission
	pageRes, err := query.Paginate(eStore, pagination, func(key []byte, value []byte) error {
		var commission types.Commission
		if err := k.cdc.Unmarshal(value, &commission); err != nil {
			return err
		}
		commissions = append(commissions, commission)
		return nil
	})
	if err != nil {
		return nil, nil, errorsmod.Wrap(err, "failed to get paginated commissions")
	}

	return commissions, pageRes, nil
}

func (k Keeper) GetAllCommissions(sdkCtx sdk.Context, epochId uint32) (commissions []types.Commission) {
	cStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreCommissionPrefix))
	eStore := prefix.NewStore(cStore, types.KeyEpoch(epochId))

	iterator := eStore.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var commission types.Commission
		k.cdc.MustUnmarshal(iterator.Value(), &commission)
		commissions = append(commissions, commission)
	}

	return
}
