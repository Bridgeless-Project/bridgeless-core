package keeper

import (
	"math/big"

	bridgetypes "github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	evmtypes "github.com/Bridgeless-Project/bridgeless-core/v12/x/evm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// Wrapper struct
type Hooks struct {
	k Keeper
}

var (
	_ bridgetypes.BridgeHook = Hooks{}
	_ evmtypes.EvmHooks      = Hooks{}
)

// Hooks creates new nft hooks
func (k Keeper) Hooks() Hooks { return Hooks{k} }

func (h Hooks) PostTxProcessing(ctx sdk.Context, msg core.Message, receipt *ethtypes.Receipt) error {
	return h.k.PostTxProcessing(ctx, msg, receipt)
}

func (h Hooks) GetTokenPrice(ctx sdk.Context, tokenAddress string, amountIn *big.Int) (*big.Int, []common.Address, error) {
	return h.k.ComputeSwapPrice(ctx, tokenAddress, amountIn)
}
