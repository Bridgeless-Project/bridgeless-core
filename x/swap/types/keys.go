package types

import (
	"fmt"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ethereum/go-ethereum/common"
)

const (
	// ModuleName defines the module name
	ModuleName = "swap"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_swap"

	ParamModuleAdminKey          = "ModuleAdmin"
	ParamUniswapRouterAddressKey = "UniswapRouterAddress"
	ParamWrappedBridgeKey        = "WrappedBridge"
	ParamSwapDeadlineSecondsKey  = "SwapDeadlineSeconds"

	StorePoolPrefix           = "pool"
	StoreSwapPrefix           = "swap"
	StoreSwapSubmissionPrefix = "swap_submission"
)

// ModuleAddress is the native module address for EVM
var ModuleAddress common.Address

func init() {
	ModuleAddress = common.BytesToAddress(authtypes.NewModuleAddress(ModuleName).Bytes())
}

func Prefix(p string) []byte {
	return []byte(p + "/")
}

func KeyPool(tokenID string) []byte {
	return []byte(tokenID)
}

func KeySwap(txHash string, txNonce uint64, chainID string) []byte {
	return []byte(KeySwapString(txHash, txNonce, chainID))
}

func KeySwapString(txHash string, txNonce uint64, chainID string) string {
	return fmt.Sprintf("%s/%d/%s", txHash, txNonce, chainID)
}
