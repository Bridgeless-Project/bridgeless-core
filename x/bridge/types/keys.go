package types

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	// ModuleName defines the module name
	ModuleName = "bridge"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_bridge"

	// ----- Param Keys -----
	ParamModuleAdminKey   = "ModuleAdmin"
	ParamModulePartiesKey = "Parties"
	ParamTssThresholdKey  = "TssThreshold"

	// ---- Store Prefixes ------
	StoreTokenPrefix                  = "token"
	StoreTokenInfoPrefix              = "token-info"
	StoreTokenPairsPrefix             = "token-pairs"
	StoreChainPrefix                  = "chain"
	StoreTransactionPrefix            = "transaction"
	StoreTransactionSubmissionsPrefix = "transaction-submissions"

	// Attributes keys for bridge events
	AttributeKeyDepositTxHash     = "deposit_tx_hash"
	AttributeKeyDepositNonce      = "deposit_nonce"
	AttributeKeyDepositChainId    = "deposit_chain_id"
	AttributeKeyDepositBlock      = "deposit_block"
	AttributeKeyDepositAmount     = "deposit_amount"
	AttributeKeyDepositToken      = "deposit_token"
	AttributeKeyWithdrawalAmount  = "withdrawal_amount"
	AttributeKeyDepositor         = "depositor"
	AttributeKeyReceiver          = "receiver"
	AttributeKeyWithdrawalChainID = "withdrawal_chain_id"
	AttributeKeyWithdrawalTxHash  = "withdrawal_tx_hash"
	AttributeKeyWithdrawalToken   = "withdrawal_token"
	AttributeKeySignature         = "signature"
	AttributeKeyIsWrapped         = "is_wrapped"
	AttributeKeyCommissionAmount  = "commission_amount"
)

func Prefix(p string) []byte {
	return []byte(p + "/")
}

func TokenPairPrefix(srcChain, srcAddr string) []byte {
	return []byte(fmt.Sprintf("%s/%s/", srcChain, strings.ToLower(srcAddr)))
}

func KeyToken(id uint64) []byte {
	return []byte(strconv.FormatInt(int64(id), 10))
}

func KeyTokenPair(dstChain string) []byte {
	return []byte(dstChain)
}

func KeyTokenInfo(chain, addr string) []byte {
	return []byte(fmt.Sprintf("%s/%s", chain, strings.ToLower(addr)))
}

func KeyChain(chain string) []byte {
	return []byte(chain)
}

func KeyTransaction(id string) []byte {
	return []byte(id)
}
