package keeper

import (
	"cosmossdk.io/errors"
	v2 "github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/migrations/v2"
	v3 "github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/migrations/v3"
	v4 "github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/migrations/v4"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Migrator struct {
	keeper Keeper
}

func NewMigrator(keeper Keeper) Migrator {
	return Migrator{
		keeper: keeper,
	}
}

func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	if err := v2.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc); err != nil {
		return errors.Wrap(err, "failed to migrate store from v1 to v2")
	}

	return nil
}

func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	return v3.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc)
}

func (m Migrator) Migrate3to4(ctx sdk.Context) error {
	keyUniswapAddress := []byte(types.ParamUniswapRouterAddress)
	if !m.keeper.paramstore.Has(ctx, keyUniswapAddress) {
		m.keeper.paramstore.Set(ctx, keyUniswapAddress, "")
	}

	keyBridgeAddress := []byte(types.ParamBridgeAddress)
	if !m.keeper.paramstore.Has(ctx, keyBridgeAddress) {
		m.keeper.paramstore.Set(ctx, keyBridgeAddress, "")
	}

	return v4.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc)
}
