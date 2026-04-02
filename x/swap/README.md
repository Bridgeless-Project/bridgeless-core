# `x/swap`

## Abstract
Swap module is developed to store swap data on-chain and process swap requests on the Core side.

The module is responsible for connecting the bridge input flow, the internal swap execution, and the final deposit creation into one logical process. In other words, this module coordinates the full **bridge -> swap -> deposit** lifecycle.

Swap module stores swap transactions, supported Uniswap V2 pools, and module parameters required to execute swaps and recover funds in case of errors.

---

## State

### Params

Swap module params contains next fields:
- **Module admin** address - module admin is responsible for updating supported pools and changing swap-related configuration.

Definition:

```protobuf
// Params defines the parameters for the module.
message Params {
  string module_admin = 1;
}
```

Example:

```json
{
  "params": {
    "module_admin": "bridge1...",
  }
}
```

### SwapPool
**SwapPool** defines supported Uniswap V2 pool properties used during swap execution.

Definition:

```protobuf
message SwapPool {
  string address = 1;
  string token_id = 2;
}
```

Example:

```json
{
  "address": "0x0000000000000000000000000000000000000000",
  "token_id": "1"
}
```

### SwapTransaction
**SwapTransaction** defines swap transaction details.

This entity partially duplicates the Bridge transaction data and extends it with swap-specific fields required to link the original user request with the final internal deposit created after swap execution.

Definition:

```protobuf
message SwapTransaction {
  core.bridge.Transaction tx = 1;
  string final_receiver = 20;
  string final_amount = 21;
  string final_deposit_tx_hash = 23;
}
```

Example:

```json
{
  "tx": {
    "deposit_chain_id": "0",
    "deposit_tx_hash": "0x0000000000000000000000000000000000000000",
    "deposit_tx_index": "0",
    "deposit_block": "0",
    "deposit_token": "0x0000000000000000000000000000000000000000",
    "deposit_amount": "1000000000000000000",
    "depositor": "0x0000000000000000000000000000000000000000",
    "receiver": "0x0000000000000000000000000000000000000000",
    "withdrawal_chain_id": "1",
    "withdrawal_tx_hash": "",
    "withdrawal_token": "0x0000000000000000000000000000000000000000",
    "signature": "",
    "is_wrapped": true,
    "withdrawal_amount": "0",
    "commission_amount": "0",
    "tx_data": "",
    "referral_id": 0
  },
  "final_receiver": "0x0000000000000000000000000000000000000000",
  "final_amount": "990000000000000000",
  "final_deposit_tx_hash": "0x1111111111111111111111111111111111111111"
}
```

---

## Messages

### MsgSubmitSwapTx
**MsgSubmitSwapTx** initiates the swap flow on the Core side.

This message starts the main module logic after the required number of TSS submissions is reached. After threshold validation, the module can mint or unlock wrapped tokens, execute the internal swap sequence, and create the internal deposit used for the final bridging step.

The message must support both:
- **Bridge transaction flow** - when swap is triggered from an external bridge deposit.
- **Native swap flow** - when swap is initiated directly inside the protocol.

Definition:

```protobuf
message MsgSubmitSwapTx {
  string creator = 1;
  SwapTransaction tx = 2;
  bool is_bridge_tx = 3;
}
```

Example:

```json
{
  "creator": "bridge1...",
  "tx": {
    "tx": {
      "deposit_chain_id": "0",
      "deposit_tx_hash": "0xabc",
      "deposit_tx_index": "1",
      "deposit_block": "100",
      "deposit_token": "0xeth",
      "deposit_amount": "1000000000000000000",
      "depositor": "0xuser",
      "receiver": "0xuser",
      "withdrawal_chain_id": "1",
      "withdrawal_tx_hash": "",
      "withdrawal_token": "0xbtc",
      "signature": "",
      "is_wrapped": true,
      "withdrawal_amount": "0",
      "commission_amount": "0",
      "tx_data": "",
      "referral_id": 0
    },
    "final_receiver": "bc1...",
    "final_amount": "0",
    "final_deposit_tx_hash": ""
  },
  "is_bridge_tx": true
}
```

### MsgUpdatePool
**MsgUpdatePool** registers or updates a supported Uniswap V2 pool.

This message can only be called by the module admin.

Definition:

```protobuf
message MsgUpdatePool {
  string creator = 1;
  SwapPool pool = 2;
}
```

Example:

```json
{
  "creator": "bridge1...",
  "pool": {
    "address": "0x0000000000000000000000000000000000000000",
    "token_id": "1"
  }
}
```

---

## Queries

All stored entities must be accessible through GET messages.

### Params
Returns module parameters.

```protobuf
rpc Params(QueryParamsRequest) returns (QueryParamsResponse);
```

### AllPool
Returns the full list of supported pools.

```protobuf
rpc AllPool(QueryAllPools) returns (QueryAllPoolsResponse);
```

### GetPoolByTokenId
Returns the pool associated with the given token id.

```protobuf
rpc GetPoolByTokenId(QueryGetPoolByTokenId) returns (QueryGetPoolByTokenIdResponse);
```

### AllSwaps
Returns the full list of stored swap transactions.

```protobuf
rpc AllSwaps(QueryAllSwaps) returns (QueryAllSwapsResponse);
```

### GetSwapById
Returns a single swap transaction by bridge transaction identifier.

```protobuf
rpc GetSwapById(QueryGetSwapById) returns (QueryGetSwapByIdResponse);
```

---

## Dependencies

Swap module MUST import the **Bridge** and **EVM** keepers.

### Bridge keeper
Bridge keeper is required to:
- create the final internal deposit after swap execution,
- store self-deposits in the Bridge store,
- link the original deposit to the final withdrawal,
- allow deposit creation without TSS governance checks when the caller is the Swap module address.

This exception is required because the second deposit is created internally by the protocol as part of the swap flow, rather than directly from an external user action.

### EVM keeper
EVM keeper is required to:
- call Uniswap contracts,
- call Bridge contracts,
- execute token swaps,
- process contract-based bridging and deposit operations.

To allow the protocol to call these contracts, their ABIs must be stored on the Core side.

---

## Processing Flow

During each swap, the Core MUST:
- validate that the number of TSS submissions has reached the required threshold,
- load the original bridge transaction,
- determine the supported swap route from the stored pools,
- compute transaction parameters, including the final withdrawal amount,
- execute the full **bridge -> swap -> deposit** flow either through a dedicated contract or through native protocol logic,
- create the final internal deposit in the Bridge module,
- store the resulting `SwapTransaction`.

When a user wants to swap **ETH -> BTC**, the Core MUST process two swaps:
1. Core swaps **ETH -> NST**
2. Core swaps **NST -> BTC**

The same slippage value MUST be used for both swaps.

During each swap, the Core MUST update the token-pair relation used by the protocol.

---

## Recovery Flow

If the user-provided slippage is too small, the wrapped tokens MUST be sent to the recovery address. Alternatively, the protocol address MAY hold the tokens, and the user can later request their withdrawal.

In the recovery flow, the user MUST receive the tokens from the **last successfully completed stage** before the error:
- if the user swaps **BTC -> ETH** and the error occurs during the **wBTC -> NST** swap, the user MUST receive **wBTC**
- if the error occurs during the **NST -> wETH** swap, the user MUST receive **NST** tokens at their recovery address
