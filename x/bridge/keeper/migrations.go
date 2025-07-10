package keeper

import (
	"cosmossdk.io/errors"
	v2 "github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/migrations/v2"
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
