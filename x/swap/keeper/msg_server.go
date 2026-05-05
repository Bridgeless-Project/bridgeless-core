package keeper

import (
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/swap/types"
)

type msgServer struct {
	Keeper
}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}
