package keeper

import (
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetEpochesWithPagination(ctx sdk.Context, pagination *query.PageRequest) ([]types.Epoch, *query.PageResponse, error) {
	cStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.Prefix(types.StoreEpochesPrefix))
	var epoches []types.Epoch

	pageRes, err := query.Paginate(cStore, pagination, func(key []byte, value []byte) error {
		var epoch types.Epoch

		k.cdc.MustUnmarshal(value, &epoch)

		epoches = append(epoches, epoch)
		return nil
	})

	if err != nil {
		return nil, pageRes, status.Error(codes.Internal, err.Error())
	}

	return epoches, pageRes, nil
}

func (k Keeper) SetEpoch(sdkCtx sdk.Context, epoch types.Epoch) {
	cStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreEpochesPrefix))
	cStore.Set(types.KeyEpoch(epoch.Id), k.cdc.MustMarshal(&epoch))
}

func (k Keeper) GetEpoch(sdkCtx sdk.Context, id uint64) (epoch types.Epoch, found bool) {
	cStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreEpochesPrefix))
	bz := cStore.Get(types.KeyEpoch(id))
	if bz == nil {
		return
	}

	k.cdc.MustUnmarshal(bz, &epoch)
	found = true

	return
}

func (k Keeper) GetEpochId(ctx sdk.Context) uint64 {
	params := k.GetParams(ctx)
	params.EpochSequence += 1

	for {
		_, ok := k.GetEpoch(ctx, params.EpochSequence)
		if ok {
			params.EpochSequence += 1
			continue
		}

		break
	}
	k.SetParams(ctx, params)

	return params.EpochSequence
}
