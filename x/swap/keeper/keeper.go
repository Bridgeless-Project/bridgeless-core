package keeper

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/swap/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace
	}
)

func (k Keeper) AllPool(ctx context.Context, pools *types.QueryAllPools) (*types.QueryAllPoolsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k Keeper) GetPoolByTokenId(ctx context.Context, id *types.QueryGetPoolByTokenId) (*types.QueryGetPoolByTokenIdResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k Keeper) AllSwaps(ctx context.Context, swaps *types.QueryAllSwaps) (*types.QueryAllSwapsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k Keeper) GetSwapById(ctx context.Context, id *types.QueryGetSwapById) (*types.QueryGetSwapByIdResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
