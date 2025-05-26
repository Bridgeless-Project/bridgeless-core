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
}


message TokenMetadata {
  string name = 1;
  string symbol = 2;
  string uri = 3;
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
  "commission_rate": "0.5",
  "metadata": {
    "name": "TESTNET",
    "symbol": "TEST",
    "uri": ""
  },
  "info": [
    {
      "address": "0x0000000000000000000000000000000000000000",
      "decimals": "18",
      "chain_id": "00000",
      "token_id": "1",
      "is_wrapped": true
    },
    {
      "address": "0x0000000000000000000000000000000000000000",
      "decimals": "18",
      "chain_id": "00000",
      "token_id": "1",
      "is_wrapped": false
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
  "tx_data": ""
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

Message example:

```protobuf
message MsgSubmitTransactions {
  option (cosmos.msg.v1.signer) = "submitter";

  string submitter = 1 [(cosmos_proto.scalar) = "cosmos.AddressString"];
  repeated core.bridge.Transaction transactions = 2 [(gogoproto.nullable) = false];
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
**DeleteChain**-deletes chain data by chain id from core.

```
 bridgeless-cored tx bridge chains remove bridge1... 1
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
  "commission_rate": "0.5",
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

