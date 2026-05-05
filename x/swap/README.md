# `x/swap`

## Abstract

The swap module stores submitted swap requests and coordinates the Core-side swap flow after the required bridge party threshold is reached.

Swap execution is delegated to the EVM `Swapper` contract. The module builds the `withdrawSwapAndRoute` request from the submitted bridge transaction, resolves the destination token address on the Cosmos chain through bridge token mappings, calls the contract, and stores the processed swap record.

The module no longer stores swap pools. Route construction uses bridge token metadata plus the configured wrapped bridge token.

---

## State

### Params

```protobuf
message Params {
  string module_admin = 1;
  string wrapped_bridge = 2;
  string swapper_address = 3;
  uint64 swaper_caller_address = 4;
}
```

- `module_admin` is retained for module configuration ownership.
- `wrapped_bridge` is the middle token used in the swap path.
- `swapper_address` is the EVM contract called by the module.
- `swaper_caller_address` is the EVM address used by the module to call the Swapper contract.

### SwapTransaction

```protobuf
message SwapTransaction {
  core.bridge.Transaction tx = 1 [(gogoproto.nullable) = false]; // Base transaction from bridge module
  string final_receiver = 2; // in case of native swap the final receiver is the end user
  string final_token = 3; // token user wanna get after swap (can be on Bridgeless chain or another one)
  string final_chain_id = 4; // chain where user wanna get the final token (can be bridgeless chain or another one)
  uint64 swap_deadline = 5; // timestamp until when the swap is valid, provided by
  string swap_out_amount = 6; // minimum acceptable output amount, provided by backend
  string final_deposit_tx_hash = 7; // do not used on the submit endpoint
}
```


## Messages

### MsgSubmitSwapTx

```protobuf
message MsgSubmitSwapTx {
  string creator = 1;
  SwapTransaction tx = 2;
}
```

Only authorized bridge parties can submit swap transactions. The module hashes the request payload, tracks party submissions, and executes the Swapper call once the TSS threshold is reached.

The module calls:

```text
Swapper.withdrawSwapAndRoute(
  withdrawParams,
  swapParams,
  destinationDepositParams,
  fallbackDepositParams
)
```

Argument sources:

- `withdrawParams`: deposit token, deposit amount, deposit transaction hash/index, original wrapped flag, and decoded signatures.
- `swapParams`: deposit amount, `amount_out_min`, deadline, path, and destination-native flag.
- `destinationDepositParams`: `final_receiver`, withdrawal chain id, destination token wrapped flag, and referral id.
- `fallbackDepositParams`: `tx.tx_data`, deposit chain id, wrapped flag, and referral id.

The swap path is:

```text
[deposit_token, wrapped_bridge, destination_token_on_cosmos_chain]
```

The last address is resolved with bridge token mappings from the withdrawal token and withdrawal chain into the current Cosmos chain.

---

## Queries

### Params

```protobuf
rpc Params(QueryParamsRequest) returns (QueryParamsResponse);
```

### AllSwaps

```protobuf
rpc AllSwaps(QueryAllSwaps) returns (QueryAllSwapsResponse);
```

### GetSwapById

```protobuf
rpc GetSwapById(QueryGetSwapById) returns (QueryGetSwapByIdResponse);
```

---

## Dependencies

### Bridge keeper

The bridge keeper is used to authorize party submissions, read bridge threshold params, validate destination chains, and resolve token mappings for the Cosmos-chain route token.

### EVM keeper

The EVM keeper is used to call the Swapper contract ABI from the swap module address.

---

## Processing Flow

During each swap, the module:

- validates the submitter is an authorized bridge party,
- rejects duplicate submitters and already processed swaps,
- waits until the configured bridge TSS threshold is reached,
- resolves the destination token address on the Cosmos chain,
- builds Swapper withdraw, swap, destination deposit, and fallback deposit params,
- calls `withdrawSwapAndRoute`,
- stores the `SwapTransaction` with the Swapper EVM response hash.

Fallback routing is handled by the Swapper contract using `tx.tx_data` as the fallback receiver or recovery payload.
