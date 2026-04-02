package swap_test

import (
	"testing"

	keepertest "github.com/Bridgeless-Project/bridgeless-core/v12/testutil/keeper"
	"github.com/Bridgeless-Project/bridgeless-core/v12/testutil/nullify"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/swap"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/swap/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:	types.DefaultParams(),
		
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SwapKeeper(t)
	swap.InitGenesis(ctx, *k, genesisState)
	got := swap.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	

	// this line is used by starport scaffolding # genesis/test/assert
}
