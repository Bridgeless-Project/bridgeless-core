package types

import (
	"fmt"
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
	ParamWrappedBridgeKey        = "WrappedBridge"
	ParamSwapperAddressKey       = "SwapperAddress"
	ParamSwapperCallerAddressKey = "SwapperCallerAddress"

	StoreSwapPrefix           = "swap"
	StoreSwapSubmissionPrefix = "swap_submission"
)

func Prefix(p string) []byte {
	return []byte(p + "/")
}

func KeySwap(txHash string, txNonce uint64, chainID string) []byte {
	return []byte(KeySwapString(txHash, txNonce, chainID))
}

func KeySwapString(txHash string, txNonce uint64, chainID string) string {
	return fmt.Sprintf("%s/%d/%s", txHash, txNonce, chainID)
}
