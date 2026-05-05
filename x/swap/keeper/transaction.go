package keeper

import (
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/swap/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SetSwap(ctx sdk.Context, swap types.SwapTransaction) {
	sStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.Prefix(types.StoreSwapPrefix))
	sStore.Set(types.KeySwap(swap.Tx.DepositTxHash, swap.Tx.DepositTxIndex, swap.Tx.DepositChainId), k.cdc.MustMarshal(&swap))
}

func (k Keeper) GetSwap(ctx sdk.Context, txHash string, txNonce uint64, chainID string) (swap types.SwapTransaction, found bool) {
	sStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.Prefix(types.StoreSwapPrefix))
	bz := sStore.Get(types.KeySwap(txHash, txNonce, chainID))
	if bz == nil {
		return
	}

	k.cdc.MustUnmarshal(bz, &swap)
	found = true

	return
}

func (k Keeper) GetSwapsWithPagination(ctx sdk.Context, pagination *query.PageRequest) ([]types.SwapTransaction, *query.PageResponse, error) {
	sStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.Prefix(types.StoreSwapPrefix))
	var swaps []types.SwapTransaction

	pageRes, err := query.Paginate(sStore, pagination, func(key []byte, value []byte) error {
		var swap types.SwapTransaction

		k.cdc.MustUnmarshal(value, &swap)
		swaps = append(swaps, swap)
		return nil
	})
	if err != nil {
		return nil, pageRes, status.Error(codes.Internal, err.Error())
	}

	return swaps, pageRes, nil
}

func (k Keeper) GetAllSwaps(ctx sdk.Context) (swaps []types.SwapTransaction) {
	sStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.Prefix(types.StoreSwapPrefix))
	iterator := sStore.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var swap types.SwapTransaction
		k.cdc.MustUnmarshal(iterator.Value(), &swap)
		swaps = append(swaps, swap)
	}

	return
}
