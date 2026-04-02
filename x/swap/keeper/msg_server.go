package keeper

import (
	"context"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/swap/types"
)

type msgServer struct {
	Keeper
}

func (m msgServer) UpdatePool(ctx context.Context, pool *types.MsgUpdatePool) (*types.MsgUpdatePoolResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m msgServer) SubmitSwapTx(ctx context.Context, tx *types.MsgSubmitSwapTx) (*types.MsgSubmitSwapTxResponse, error) {
	//TODO implement me
	panic("implement me")
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}
