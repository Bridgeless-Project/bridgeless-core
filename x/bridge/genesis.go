package bridge

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/hyle-team/bridgeless-core/v12/x/bridge/keeper"
	"github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
	"github.com/pkg/errors"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
	for _, chain := range genState.Chains {
		k.SetChain(ctx, chain)
	}
	for _, token := range genState.Tokens {
		k.SetToken(ctx, token)
		for _, info := range token.Info {
			k.SetTokenInfo(ctx, info)
			k.SetTokenPairs(ctx, info, token.Info...)
		}
	}
	for _, tx := range genState.Transactions {
		k.SetTransaction(ctx, tx)
	}

	for _, txSubmissions := range genState.TransactionsSubmissions {
		k.SetTransactionSubmissions(ctx, &txSubmissions)
	}

	if err := genState.Validate(); err != nil {
		panic(errors.Wrap(err, "invalid genesis state"))
	}
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	txs, _, err := k.GetPaginatedTransactions(ctx, &query.PageRequest{Limit: query.MaxLimit})
	if err != nil {
		panic(errorsmod.Wrap(err, "failed to export genesis transactions"))
	}

	tokens, _, err := k.GetTokensWithPagination(ctx, &query.PageRequest{Limit: query.MaxLimit})
	if err != nil {
		panic(errorsmod.Wrap(err, "failed to export genesis tokens"))
	}

	chains, _, err := k.GetChainsWithPagination(ctx, &query.PageRequest{Limit: query.MaxLimit})
	if err != nil {
		panic(errorsmod.Wrap(err, "failed to export genesis chains"))
	}

	txsWithSubmissions, _, err := k.GetPaginatedTransactionsSubmissions(ctx, &query.PageRequest{Limit: query.MaxLimit})

	return &types.GenesisState{
		Params:                  k.GetParams(ctx),
		Chains:                  chains,
		Tokens:                  tokens,
		Transactions:            txs,
		TransactionsSubmissions: txsWithSubmissions,
	}
}
