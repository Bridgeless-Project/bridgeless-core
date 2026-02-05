# `x/bridge`

## Abstract
Bridge module is developed to store bridge data on-chain. It stores chains and tokens which bridge operates with.
BRidge transactions also stored in module to prevent spam and double-spending's.

---

## State
### Params

Bridge module params contains next fields:
- **Module admin** address - module admin is responsible for making changes in parties lists and thresholds, also only him
  can add new info about chains, tokens and transfers
- **Parties list** - parties list is the list of active validators which are parts of TSS protocol.
- **Tss threshold** - minimum required number of signers.

Definition:

```protobuf
// Params defines the parameters for the module.
message Params {
  string module_admin = 1;
  repeated Party parties = 2;
  uint32 tss_threshold = 3;
}
```

Example:

```json
{
  "params": {
    "module_admin": "bridge1...",
    "parties": []
  }
}
```

### Chain
**Chain** defines necessary chain properties used during bridging process.

Definition:

```protobuf
message Chain {
  string id = 1;
  ChainType type =2;
  // bridge_address is the address of the bridge contract on the chain
  string bridge_address = 3;
  // operator is the address of the operator of the bridge contract
  string operator = 4;
}
```

Example:
```json
{
      "id": "0",
      "type": "BITCOIN",
      "bridge_address": "addr...",
      "operator": "addr...."
}
```

### Referral
**Referral** defines referral properties that include the referral address to process withdrawal and the commission rate.

Definition:
```protobuf
message Referral {
      // stores a 16-bit unsigned integer
      uint32 id = 1; 
      string withdrawal_address = 2;
      string commission_rate = 3;
}
```

Example:
```json
{
  "id": "0",
  "withdrawal_address": "bridge1qqqqqnqaea03090u62sd7e42jfv2lhllsckms0",
  "commission_rate": "0.1"
}
```


### Referral Rewards
**ReferralRewards** defines referral rewards properties that include the referral id, token id, amount to claim and total collected amount.

Definition 
```protobuf
message ReferralRewards {
  uint32 referral_id = 1;
  uint64 token_id = 2;
  string to_claim = 3;
  string total_claimed_amount = 4;
}
```

Example
```json
{
  "referral_id": 1,
  "token_id": 1,
  "to_claim": "10",
  "total_claimed_amount": "100"
}
```

### Token
**Token** defines necessary token properties like id or decimals used during bridging process.

Definition:
```protobuf
message TokenInfo {
  string address = 1;
  uint64 decimals = 2;
  string chain_id = 3;;
  uint64 token_id = 4;
  bool is_wrapped = 5;
  string min_withdrawal_amount = 6;
}

message TokenMetadata {
  string name = 1;
  string symbol = 2;
  string uri = 3;
  string dex_name = 4;
}

message Token {
  uint64 id = 1;
  TokenMetadata metadata = 2 [(gogoproto.nullable) = false];
  // info is the token information on different chains
  repeated TokenInfo info = 3 [(gogoproto.nullable) = false];
  string commission_rate = 4;
}
```

Example:
```json
{
  "id": "0",
  "commission_rate": "0.1",
  "metadata": {
    "name": "TESTNET",
    "symbol": "TEST",
    "uri": "",
    "dex_name": ""
  },
  "info": [
    {
      "address": "0x0000000000000000000000000000000000000000",
      "decimals": "18",
      "chain_id": "00000",
      "token_id": "1",
      "is_wrapped": true,
      "min_withdrawal_amount": "0"
    },
    {
      "address": "0x0000000000000000000000000000000000000000",
      "decimals": "18",
      "chain_id": "00000",
      "token_id": "1",
      "is_wrapped": false,
      "min_withdrawal_amount": "0"
    }
  ]
}
```

### Transaction
**Transaction** defines bridge transaction details.

Definition:

```protobuf
message Transaction {
  string deposit_chain_id = 1;
  string deposit_tx_hash = 2;
  uint64 deposit_tx_index = 3;
  uint64 deposit_block = 4;
  string deposit_token = 5;
  string deposit_amount = 6;
  string depositor = 7;
  string receiver = 8;
  string withdrawal_chain_id = 9;
  string withdrawal_tx_hash = 10;
  string withdrawal_token = 11;
  string signature = 12;
  bool is_wrapped = 13;
  string withdrawal_amount = 14;
  string commission_amount = 15;
  string tx_data = 16;
  int32 referral_id = 17;
}
```

Example:
```json
{
  "deposit_chain_id": "00000",
  "deposit_tx_hash": "0x0000000000000000000000000000000000000000",
  "deposit_tx_index": "0",
  "deposit_block": "0",
  "deposit_token": "0x0000000000000000000000000000000000000000",
  "deposit_amount": "00000",
  "depositor": "0x0000000000000000000000000000000000000000",
  "receiver": "0x0000000000000000000000000000000000000000",
  "withdrawal_chain_id": "00000",
  "withdrawal_tx_hash": "",
  "withdrawal_token": "0x0000000000000000000000000000000000000000",
  "signature": "0x0000000000000000000000000000000000000000",
  "is_wrapped": true,
  "withdrawal_amount": "0",
  "commission_amount": "0",
  "tx_data": "",
  "referral_id": 0
}
```

### TransactionSubmission

```protobuf
message TransactionSubmissions{
  string tx_hash = 1;
  repeated string submitters = 2;
}
```
Example:

```json
    {
      "tx_hash": "0x0000000000000000000000000000000000000000",
      "submitters": [
        "bridge1..."
      ]
    }
```
___

### Epoch
**Epoch** defines the epoch state for TSS key rotation and signature management.

Definition:

```protobuf
enum EpochStatus {
  INITIATED = 0;
  FINALIZING = 1;
  COMPLETE = 2;
}

message Epoch {
  uint32 id = 1;
  EpochStatus status = 2;
  uint64 start_block = 3;
  uint64 end_block = 4;
  repeated Party parties = 5;
  uint32 tss_threshold = 6;
  repeated TSSInfo tss_info = 7;
  string pubkey = 8;
}
```

Example:
```json
{
  "id": 1,
  "status": "INITIATED",
  "start_block": "100",
  "end_block": "0",
  "parties": [
    {"address": "bridge1..."}
  ],
  "tss_threshold": 2,
  "tss_info": [
    {
      "certificate": "cert_pem_content",
      "domen": "tss-node1.example.com",
      "address": "0x1234567890abcdef",
      "active": true
    }
  ],
  "pubkey": "02abc..."
}
```
___

### TSSInfo
**TSSInfo** defines the TSS node information for an epoch.

Definition:

```protobuf
message TSSInfo {
  string certificate = 1;
  string domen = 2;
  string address = 3;
  bool active = 4;
}
```

Example:
```json
{
  "certificate": "-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----",
  "domen": "tss-node1.example.com",
  "address": "0x1234567890abcdef",
  "active": true
}
```
___

### EpochChainSignatures
**EpochChainSignatures** defines epoch signatures for a specific chain type.

Definition:

```protobuf
message EpochChainSignatures {
   uint32 epoch_id = 1;
   ChainType chain_type = 2;
   EpochSignature added_signature = 3;
   EpochSignature removed_signature = 4;
   string address = 5;
   uint32 submittions = 6;
}

message EpochSignature {
  EpochSignatureMod mod = 1;
  uint32 epoch_id = 2;
  string signature = 3;
  EpochSignatureData data = 4;
}

message EpochSignatureData {
  string new_signer = 1;
  uint64 start_time = 2;
  uint64 end_time = 3;
  string nonce = 4;
}
```

Example:
```json
{
  "epoch_id": 1,
  "chain_type": 1,
  "added_signature": {
    "mod": 1,
    "epoch_id": 1,
    "signature": "0xsignature...",
    "data": {
      "new_signer": "0xnewsigner...",
      "start_time": 1234567890,
      "end_time": 1234567899,
      "nonce": "abc123"
    }
  },
  "removed_signature": null,
  "address": "0xcontractaddress...",
  "submittions": 3
}
```
___

## Transactions
## RPC
___
### Chains
___
### InsertChain

**InsertChain** - adds new chain data to core.

Message example:
```protobuf
message MsgInsertChain {
  string creator = 1;
  Chain chain  = 2 [(gogoproto.nullable) = false];
}
```
___
### DeleteChain
**DeleteChain**-deletes chain data from core.

Message example:
```protobuf
message MsgDeleteChain {
  string creator = 1;
  string chain_id = 2;
}
```
___

### Tokens

___

### InsertToken

**InsertToken** - inserts token data to core.

Message example:

```protobuf
message MsgInsertToken {
  string creator = 1;
  Token token  = 2 [(gogoproto.nullable) = false];
}

```
___
### UpdateToken
**UpdateToken** - updates queried token metadata on core.

Message example:

```protobuf
message MsgUpdateToken {
  string creator = 1;
  uint64 token_id = 2;
  TokenMetadata metadata = 3 [(gogoproto.nullable) = false];
}
```
___

### DeleteToken
**DeleteToken** - deletes token data from core.

Message example:

```protobuf
message MsgDeleteToken {
  string creator = 1;
  uint64 token_id = 2;
}
```

___
### Transactions
___

### SubmitTransactions

**SubmitTransactions** - stores bridge transactions data to core.

Message definition:

```protobuf
message MsgSubmitTransactions {
  option (cosmos.msg.v1.signer) = "submitter";

  string submitter = 1 [(cosmos_proto.scalar) = "cosmos.AddressString"];
  repeated core.bridge.Transaction transactions = 2 [(gogoproto.nullable) = false];
}
```
___
### RemoveTransaction
**RemoveTransaction** - removes bridge transaction data from core.

Message definition:

```protobuf
message MsgRemoveTransaction{
  string creator = 1;

  string deposit_chain_id = 2;
  string deposit_tx_hash = 3;
  uint64 deposit_tx_index = 4;
}
```

___
### TokenInfo
___

### AddTokenInfo
**AddTokenInfo** - adds new token info to existing one on core.

Message example:

```protobuf
message MsgAddTokenInfo {
  string creator = 1;
  TokenInfo info = 3 [(gogoproto.nullable) = false];
}
```
___

### RemoveTokenInfo

Message example:

```protobuf
message MsgRemoveTokenInfo {
  string creator = 1;
  uint64 token_id = 2;
  string chain_id = 3;
}
```
___

### Referrals
___

### SetReferral

Message definition:
```protobuf
message MsgSetReferral {
  string creator = 1;
  Referral referral = 2 [(gogoproto.nullable) = false];
}
```
___

### RemoveReferral
Message definition:
```protobuf
message MsgRemoveReferral {
  string creator = 1;
  uint32 id = 2;
}
```
___
### SetReferralRewards
Message definition:
```protobuf
message MsgSetReferralRewards {
  string creator = 1;
  ReferralRewards rewards = 4 [(gogoproto.nullable) = false];
}
```
___
### RemoveReferralRewards
Message definition 
```protobuf
message MsgRemoveReferralRewards {
  string creator = 1;
  uint32 referrer_id = 2;
  uint64 token_id = 3;
}
```
___
### Parties
___

### SetParties

Message example:

```protobuf
message MsgSetPartiesList {
  string creator =1;
  repeated Party parties = 2;
}
```
___
### Threshold
___

### SetTssThreshold

Message example:
```protobuf
message MsgSetTssThreshold {
  string creator =1;
  uint32  amount = 2;
}
```
___

### Epochs
___

### StartEpoch

**StartEpoch** - initiates a new epoch with TSS information. Only the module admin can start a new epoch.

Message definition:
```protobuf
message MsgStartEpoch {
  option (cosmos.msg.v1.signer) = "creator";

  string creator = 1 [(cosmos_proto.scalar) = "cosmos.AddressString"];
  uint32 epoch_id = 2;
  repeated TSSInfo info = 3 [(gogoproto.nullable) = false];
  uint32 tss_threshold = 4;
}
```
___

### SetEpochSignature

**SetEpochSignature** - submits epoch chain signatures. Only authorized parties can submit signatures. When the threshold is reached, the signatures are stored and the epoch status is updated.

Message definition:
```protobuf
message MsgSetEpochSignature {
  option (cosmos.msg.v1.signer) = "creator";

  string creator = 1 [(cosmos_proto.scalar) = "cosmos.AddressString"];
  repeated EpochChainSignatures epoch_chain_signatures = 2 [(gogoproto.nullable) = false];
}
```
___

### SetEpochPubkey

**SetEpochPubkey** - submits a public key for an epoch. Only authorized parties can submit pubkeys. When the threshold is reached, the pubkey is stored for the epoch.

Message definition:
```protobuf
message MsgSetEpochPubkey {
  option (cosmos.msg.v1.signer) = "creator";

  string creator = 1 [(cosmos_proto.scalar) = "cosmos.AddressString"];
  string pubkey = 2;
  uint32 epoch_id = 3;
}
```
___

## CLI
___

## Transactions
___
### Chains
___
### InsertChain

**InsertChain** - adds new chain data to core.

```
bridgeless-cored tx bridge chains insert bridge1... chain.json
```

Example of `chain.json`:
```json
{
  "id": "0",
  "type": 0,
  "bridge_address": "0x0000000000000000000000000000000000000000",
  "operator": "0x0000000000000000000000000000000000000000"
}
```
___
### DeleteChain
**DeleteChain** - deletes chain data by chain id from core.

```
 bridgeless-cored tx bridge chains remove bridge1... 1
```
___

### Referrals
___

### SetReferral
**SetReferral** - adds new referral data to core.

```
bridgeless-cored tx bridge referral set bridge1... referral-id=0 referral-withdrawal-address=bridge1... referral-commission-rate=5
```

___
### RemoveReferral
**RemoveReferral** - removes referral data by referral id from core.

```
 bridgeless-cored tx bridge referral remove bridge1... 0
```

___
### SetReferralRewards
**SetReferralRewards** - adds new referral rewards data to core.

```
 bridgeless-cored tx bridge referral-rewards set bridge1... referral-id=1 token-id=1 to-claim=10abridge total-collected-amount=100abridge
```
___
### RemoveReferralRewards
**RemoveReferralRewards** - removes referral rewards data by referral id and token id from core.

```
 bridgeless-cored tx bridge referral-rewards remove bridge1... 1 1
```
___

### Tokens

___

### InsertToken

**InsertToken** - inserts token data to core.


```
 bridgeless-cored tx bridge tokens insert bridge1... token.json
```

Example of `token.json`: 
```json
{
  "id": 1,
  "commission_rate": "0.1",
  "metadata": {
    "name": "TESTNET",
    "symbol": "TEST",
    "uri": "https://example.com"
  },
  "info": [
    {
      "address": "0x0000000000000000000000000000000000000000",
      "decimals": 18,
      "chain_id": "00000",
      "token_id": 1,
      "is_wrapped": true
    },
    {
      "address": "0x0000000000000000000000000000000000000000",
      "decimals": 18,
      "chain_id": "00000",
      "token_id": 1,
      "is_wrapped": false
    }
  ]
}

```
___
### UpdateToken
**UpdateToken** - updates queried token metadata on core.

```
bridgeless-cored tx bridge tokens update bridge1... token.json
```
___

### DeleteToken
**DeleteToken** - deletes token data by token id from core.

```
 bridgeless-cored tx bridge tokens remove bridge1... 1
```

___
### Transactions
___

### SubmitTransactions

**SubmitTransactions** - stores bridge transactions data to core.

```
bridgeless-cored tx bridge transactions submit bridge1... tx.json
```
___
### TokenInfo
___

### AddTokenInfo
**AddTokenInfo** - adds new token info to existing one on core.

```
bridgeless-cored tx bridge tokens add-info bridge1... info.json
```
___

### RemoveTokenInfo

```
bridgeless-cored tx bridge tokens remove-info bridge1... 1 [token-id] 1 [chain-id]
```
___
### Parties
___

### SetParties

```
 bridgeless-cored tx bridge parties set default bridge1... bridge1...,bridge1...
```
___

### SetTssThreshold

```
 bridgeless-cored tx bridge threshold set-tss-threshold bridge1... 5
```
___

### Transactions Stop List

___

### AddTxToStopList

```
bridgeless-cored tx bridge stop-list add-tx bridge1... tx.json
```

Example of tx.json: 

```json
{
  "deposit_chain_id": "0",
  "deposit_tx_hash": "0x0000000000000000000000000000000000000000",
  "deposit_tx_index": 0,
  "deposit_block": 0,
  "deposit_token": "0x0000000000000000000000000000000000000000",
  "deposit_amount": "0",
  "depositor": "0x0000000000000000000000000000000000000000",
  "receiver": "0x0000000000000000000000000000000000000000",
  "withdrawal_chain_id": "0",
  "withdrawal_tx_hash": "",
  "withdrawal_token": "0x0000000000000000000000000000000000000000",
  "signature": "0x0000000000000000000000000000000000000000",
  "is_wrapped": true,
  "withdrawal_amount": "0",
  "commission_amount": "0",
  "tx_data": "",
  "referral_id": 0
}
```
___

### RemoveTxFromStopList

```
 bridgeless-cored tx bridge stop-list remove-tx bridge1... 0 0x0000000000000000000000000000000000000000 0
```
___

### Epochs
___

### StartEpoch

**StartEpoch** - initiates a new epoch with TSS information from a JSON file.

```
bridgeless-cored tx bridge epochs start bridge1... epoch.json
```

Example of `epoch.json`:
```json
{
  "epoch_id": 1,
  "tss_info": [
    {
      "certificate": "-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----",
      "domen": "tss-node1.example.com",
      "address": "0x1234567890abcdef",
      "active": true
    },
    {
      "certificate": "-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----",
      "domen": "tss-node2.example.com",
      "address": "0xabcdef1234567890",
      "active": true
    }
  ],
  "tss_threshold": 2
}
```
___

### SetEpochSignature

**SetEpochSignature** - submits epoch chain signatures from a JSON file.

```
bridgeless-cored tx bridge epochs set-signature bridge1... signatures.json
```

Example of `signatures.json`:
```json
[
  {
    "epoch_id": 1,
    "chain_type": 1,
    "added_signature": {
      "mod": 1,
      "epoch_id": 1,
      "signature": "0xsignature...",
      "data": {
        "new_signer": "0xnewsigner...",
        "start_time": 1234567890,
        "end_time": 1234567899,
        "nonce": "abc123"
      }
    },
    "removed_signature": null,
    "address": "0xcontractaddress...",
    "submittions": 0
  }
]
```
___

### SetEpochPubkey

**SetEpochPubkey** - sets the public key for an epoch.

```
bridgeless-cored tx bridge epochs set-pubkey bridge1... 02abc123... 1
```

Arguments:
- `bridge1...` - the sender address (must be an authorized party)
- `02abc123...` - the public key
- `1` - the epoch ID
___
___
## Query
___

### Params

```
bridgeless-cored query bridge params
```

Response example:

```
params:
  module_admin: bridge1...
  parties:
  - address: bridge1...
  - address: bridge1...
  tss_threshold: 0

```
___
### Chains
___

### QueryChains

```
bridgeless-cored query bridge chains
```

Response example:
```
chains:
- bridge_address: m4...
  id: "1"
  operator: m4...
  type: BITCOIN
- bridge_address: Zx...
  id: "2"
  operator: Zx...
  type: ZANO
- bridge_address: 0x...
  id: "3"
  operator: 0x...
  type: EVM
- bridge_address: 0x...
  id: "4"
  operator: 0x...
  type: EVM
```
___

### QueryChainById

```
bridgeless-cored query bridge chain 1
```

Response example:
```
chain:
  bridge_address: m4...
  id: "1"
  operator: m4...
  type: BITCOIN
```

___
### Referrals
___
### GetReferralById
```
 bridgeless-cored query bridge referrals 1
```

Response example:
```
referral:
  commission_rate: "0.1"
  id: "1"
  withdrawal_address: bridge1qqqqqnqaea03090u62sd7e42jfv2lhllsckms0
```
___
### GetQueryGetReferrals
```
bridgeless-cored query bridge referrals
```
Response example:
```
referrals: 
  - commission_rate: "0.1"
    id: "1"
    withdrawal_address: bridge1qqqqqnqaea03090u62sd7e42jfv2lhllsckms0
```
___
### GetReferralRewardsByToken
```
 bridgeless-cored query bridge referral-rewards 1 1
```
Response example:
```
referral_rewards:
  referral_id: "1"
  to_claim: "10"
  token_id: "1"
  total_claimed_amount: "100"
```
___
### GetQueryGetReferralsRewardsById
```
bridgeless-cored query bridge referrals-rewards
```

Response example:
```
referrals_rewards:
- referral_id: "1"
  to_claim: "10"
  token_id: "1"
  total_claimed_amount: "100"
```

___
### Tokens
___
### QueryAllTokens

```
bridgeless-cored query bridge tokens
```

Response example:

```
tokens:
- commission_rate: "0.5"
  id: "1"
  info:
  - address: 0x0000000000000000000000000000000000000000
    chain_id: "00000"
    decimals: "18"
    is_wrapped: true
    token_id: "1"
  - address: "0x0000000000000000000000000000000000000000"
    chain_id: "00000"
    decimals: "18"
    is_wrapped: false
    token_id: "1"
  metadata:
    name: TESTNET
    symbol: TEST
    uri: https://example.com
```

___
### QueryTokenById

```
 bridgeless-cored query bridge token 1
```

Response example:
```
tokens:
- commission_rate: "0.5"
  id: "1"
  info:
  - address: 0x0000000000000000000000000000000000000000
    chain_id: "00000"
    decimals: "18"
    is_wrapped: true
    token_id: "1"
  - address: "0x0000000000000000000000000000000000000000"
    chain_id: "00000"
    decimals: "18"
    is_wrapped: false
    token_id: "1"
  metadata:
    name: TESTNET
    symbol: TEST
    uri: https://example.com
```
___

### QueryTokenPairs

```
bridgeless-cored query bridge token-pairs 1 0x0000000000000000000000000000000000000000 2
```

Response example:

```
    address: 0x0000000000000000000000000000000000000000
    chain_id: "2"
    decimals: "18"
    is_wrapped: true
    token_id: "1"
```
___

### QueryTokenInfo

```
bridgeless-cored query bridge token-info 1 0x...
```

Response example:
```
info:
  address: 0x0000000000000000000000000000000000000000
  chain_id: "00000"
  decimals: "18"
  is_wrapped: true
  token_id: "1"
```
___
### Transactions
___

### QueryAllTransactions

```
bridgeless-cored query bridge transactions
```

Response example:

```
transactions:
- commission_amount: "0"
  deposit_amount: "00000"
  deposit_block: "0"
  deposit_chain_id: "0000"
  deposit_token: "0x0000000000000000000000000000000000000000"
  deposit_tx_hash: 0x0000000000000000000000000000000000000000
  deposit_tx_index: "0"
  depositor: 0x0000000000000000000000000000000000000000
  is_wrapped: true
  receiver: 0x0000000000000000000000000000000000000000
  signature: 0x0000000000000000000000000000000000000000
  tx_data: ""
  withdrawal_amount: "0"
  withdrawal_chain_id: "0"
  withdrawal_token: 0x0000000000000000000000000000000000000000
  withdrawal_tx_hash: ""
```
___

### QueryTransactionById
```
bridgeless-cored query bridge transaction 0x/2/0x
```

Response example:

```
transactions:
- commission_amount: "0"
  deposit_amount: "00000"
  deposit_block: "0"
  deposit_chain_id: "0000"
  deposit_token: "0x0000000000000000000000000000000000000000"
  deposit_tx_hash: 0x0000000000000000000000000000000000000000
  deposit_tx_index: "0"
  depositor: 0x0000000000000000000000000000000000000000
  is_wrapped: true
  receiver: 0x0000000000000000000000000000000000000000
  signature: 0x0000000000000000000000000000000000000000
  tx_data: ""
  withdrawal_amount: "0"
  withdrawal_chain_id: "0"
  withdrawal_token: 0x0000000000000000000000000000000000000000
  withdrawal_tx_hash: ""
```
___

### QueryAllTransactionsSubmissions

```
bridgeless-cored query bridge transactions-submissions
```

Response example:
```
txs_submissions:
- submitters:
  - bridge1....
  tx_hash: 0x0000000000000000000000000000000000000000

```
___

### QueryAllTransactionsSubmissions

```
bridgeless-cored query bridge transaction-submissions 0x0000000000000000000000000000000000000000
```

Response example:
```
txs_submissions:
- submitters:
  - bridge1....
  tx_hash: 0x0000000000000000000000000000000000000000

```
___
### Transactions Stop List
___

### QueryAllStopListTxs

```
 bridgeless-cored query bridge stop-list
```

Response example:

```
pagination:
  next_key: null
  total: "1"
transactions:
- commission_amount: "0"
  deposit_amount: "0"
  deposit_block: "0"
  deposit_chain_id: "0"
  deposit_token: "0x0000000000000000000000000000000000000000"
  deposit_tx_hash: 0x0000000000000000000000000000000000000000
  deposit_tx_index: "0"
  depositor: 0x0000000000000000000000000000000000000000
  is_wrapped: true
  receiver: 0x0000000000000000000000000000000000000000
  referral_id: 0
  signature: 0x0000000000000000000000000000000000000000
  tx_data: ""
  withdrawal_amount: "0"
  withdrawal_chain_id: "0"
  withdrawal_token: 0x0000000000000000000000000000000000000000
  withdrawal_tx_hash: ""
```

___

### QueryStopListTxById

```
bridgeless-cored query bridge stop-list-tx 0 0x0000000000000000000000000000000000000000 0
```

Response example:

```
transaction:
  commission_amount: "0"
  deposit_amount: "0"
  deposit_block: "0"
  deposit_chain_id: "0"
  deposit_token: "0x0000000000000000000000000000000000000000"
  deposit_tx_hash: 0x0000000000000000000000000000000000000000
  deposit_tx_index: "0"
  depositor: 0x0000000000000000000000000000000000000000
  is_wrapped: true
  receiver: 0x0000000000000000000000000000000000000000
  referral_id: 0
  signature: 0x0000000000000000000000000000000000000000
  tx_data: ""
  withdrawal_amount: "0"
  withdrawal_chain_id: "0"
  withdrawal_token: 0x0000000000000000000000000000000000000000
  withdrawal_tx_hash: ""
```
___

### Epochs
___

### QueryEpochById

```
bridgeless-cored query bridge epoch 1
```

Response example:

```
epoch:
  id: 1
  status: INITIATED
  start_block: "100"
  end_block: "0"
  parties:
  - address: bridge1...
  - address: bridge1...
  tss_threshold: 2
  tss_info:
  - certificate: "-----BEGIN CERTIFICATE-----..."
    domen: "tss-node1.example.com"
    address: "0x1234567890abcdef"
    active: true
  pubkey: "02abc..."
```
___

### QueryEpochTransactions

```
bridgeless-cored query bridge epoch-transactions 1
```

Response example:

```
transactions:
- deposit_tx_hash: "0x..."
  deposit_tx_index: "0"
  deposit_chain_id: "1"
pagination:
  next_key: null
  total: "1"
```
___

### QueryChainsByType

```
bridgeless-cored query bridge chains-by-type 1
```

Response example:

```
chains:
- bridge_address: 0x...
  id: "3"
  operator: 0x...
  type: EVM
- bridge_address: 0x...
  id: "4"
  operator: 0x...
  type: EVM
pagination:
  next_key: null
  total: "2"
```
___


