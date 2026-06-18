package swap

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/swap/keeper"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/swap/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)

	for _, tx := range genState.Transactions {
		k.SetSwap(ctx, tx)
	}
}

// ExportGenesis returns the module's exported genesisy
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	txs, _, err := k.GetSwapsWithPagination(ctx, &query.PageRequest{Limit: query.MaxLimit})
	if err != nil {
		panic(errorsmod.Wrap(err, "failed to export genesis transactions"))
	}

	// this line is used by starport scaffolding # genesis/module/export
	genesis.Transactions = txs
	return genesis
}
