package keeper

import (
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
)

type queryServer struct {
	Keeper
}

var _ types.QueryServer = queryServer{}

// NewQueryServerImpl returns an implementation of the QueryServer interface
// for the provided Keeper.
func NewQueryServerImpl(keeper Keeper) types.QueryServer {
	return &queryServer{Keeper: keeper}
}
