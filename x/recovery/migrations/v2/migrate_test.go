package v2_test

import (
	"testing"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/Bridgeless-Project/bridgeless-core/v12/app"
	"github.com/Bridgeless-Project/bridgeless-core/v12/encoding"
	v2 "github.com/Bridgeless-Project/bridgeless-core/v12/x/recovery/migrations/v2"
	v2types "github.com/Bridgeless-Project/bridgeless-core/v12/x/recovery/migrations/v2/types"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/recovery/types"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

type mockSubspace struct {
	ps           v2types.V2Params
	storeKey     storetypes.StoreKey
	transientKey storetypes.StoreKey
}

func newMockSubspace(ps v2types.V2Params, storeKey, transientKey storetypes.StoreKey) mockSubspace {
	return mockSubspace{ps: ps, storeKey: storeKey, transientKey: transientKey}
}

func (ms mockSubspace) GetParamSetIfExists(_ sdk.Context, ps types.LegacyParams) {
	*ps.(*v2types.V2Params) = ms.ps
}

func (ms mockSubspace) WithKeyTable(keyTable paramtypes.KeyTable) paramtypes.Subspace {
	encCfg := encoding.MakeConfig(app.ModuleBasics)
	cdc := encCfg.Codec
	return paramtypes.NewSubspace(cdc, encCfg.Amino, ms.storeKey, ms.transientKey, "test").WithKeyTable(keyTable)
}

func TestMigrate(t *testing.T) {
	encCfg := encoding.MakeConfig(app.ModuleBasics)
	cdc := encCfg.Codec
	storeKey := sdk.NewKVStoreKey(types.ModuleName)
	tKey := sdk.NewTransientStoreKey("transient_test")
	ctx := testutil.DefaultContext(storeKey, tKey)
	store := ctx.KVStore(storeKey)

	var outputParams v2types.V2Params
	inputParams := v2types.DefaultParams()
	legacySubspace := newMockSubspace(v2types.DefaultParams(), storeKey, tKey).WithKeyTable(v2types.ParamKeyTable())
	legacySubspace.SetParamSet(ctx, &inputParams)
	legacySubspace.GetParamSetIfExists(ctx, &outputParams)

	mockSubspace := newMockSubspace(v2types.DefaultParams(), storeKey, tKey)
	require.NoError(t, v2.MigrateStore(ctx, storeKey, mockSubspace, cdc))

	paramsBz := store.Get(v2types.ParamsKey)
	var params v2types.V2Params
	cdc.MustUnmarshal(paramsBz, &params)

	require.Equal(t, params, outputParams)
}
