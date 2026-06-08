package keeper

import (
	"math/big"

	bridgetypes "github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

// Wrapper struct
type Hooks struct {
	k Keeper
}

var _ bridgetypes.BridgeHook = Hooks{}

// Hooks creates new nft hooks
func (k Keeper) Hooks() Hooks { return Hooks{k} }

func (h Hooks) GetTokenPrice(ctx sdk.Context, tokenAddress string, amountIn *big.Int) (*big.Int, []common.Address, error) {
	return h.k.ComputeSwapPrice(ctx, tokenAddress, amountIn)
}
