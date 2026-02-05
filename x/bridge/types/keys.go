package types

import (
	"encoding/binary"
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
	ParamRelayerAccounts  = "RelayerAccounts"
	ParamEpochId          = "EpochId"
	ParamSupportingTime   = "SupportingTime"

	// ---- Store Prefixes ------
	StoreTokenPrefix                         = "token"
	StoreTokenInfoPrefix                     = "token-info"
	StoreTokenPairsPrefix                    = "token-pairs"
	StoreChainPrefix                         = "chain"
	StoreChainTypePrefix                     = "chain"
	StoreTransactionPrefix                   = "transaction"
	StoreTransactionSubmissionsPrefix        = "transaction-submissions"
	StoreReferralPrefix                      = "referral"
	StoreReferralRewardsPrefix               = "referral_rewards"
	StoreStopListTransactionsPrefix          = "stop_list_transactions"
	StoreEpochPrefix                         = "epoch"
	StoreEpochChainSignaturePrefix           = "epoch_chain_signature"
	StoreEpochChainSignatureSubmissionPrefix = "epoch_chain_signature_submission"
	StoreEpochTransactionPrefix              = "epoch_transaction"
	StoreEpochPubkeyPrefix                   = "epoch_pubkey"
	StoreEpochPubkeySubmissionPrefix         = "epoch_pubkey_submission"

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
	AttributeKeyMerkleProof       = "merkle_proof"
	AttributeEpochId              = "epoch_id"
	AttributeTssInfo              = "tss_info"
	AttributeEpochSignature       = "epoch_signature"
	AttributeEpochSignatureData   = "epoch_signature_data"
	AttributeEpochSigner          = "epoch_signer"
	AttributeEpochNonce           = "epoch_nonce"
	AttributeEpochStartTime       = "epoch_start_time"
	AttributeEpochEndTime         = "epoch_end_time"
	AttributeEpochSignatureMode   = "epoch_signature_mode"

	AttributeEpochChainType        = "epoch_chain_type"
	AttributeChainId               = "chain_id"
	AttributeEpochSignatureAddress = "epoch_signature_address"
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

func KeyChainType(chainType ChainType) []byte {
	return []byte(chainType.String())
}

func KeyReferralRewards(referraId uint32, tokenId uint64) []byte {
	return []byte(fmt.Sprintf("%d/%d", referraId, tokenId))
}

func KeyReferral(referralId uint32) []byte {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, referralId)
	return bytes
}

func KeyTransaction(id string) []byte {
	return []byte(id)
}

func KeyTransactionSubmissions(txHash string) []byte {
	return []byte(txHash)
}

func KeyEpoch(epochId uint32) []byte {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, epochId)
	return bytes
}

func KeyEpochChainSignature(chainType ChainType, epochId uint32) []byte {
	return []byte(fmt.Sprintf("%d/%d", chainType, epochId))
}

func KeyEpochChainSignatureSubmission(chainType ChainType, epochId uint32, hash string) []byte {
	return []byte(fmt.Sprintf("%d/%d/%s", chainType, epochId, hash))
}

func KeyEpochTransaction(epochId uint32, txNonce uint64, txHash string) []byte {
	return []byte(fmt.Sprintf("%s/%s/%s/%d", txHash, txNonce, epochId))
}

func KeyEpochPubkey(epochId uint32) []byte {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, epochId)
	return bytes
}

func KeyEpochPubkeySubmission(epochId uint32, pubkeyHash string) []byte {
	return []byte(fmt.Sprintf("%d/%s", epochId, pubkeyHash))
}
