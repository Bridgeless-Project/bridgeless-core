<!--
Guiding Principles:

Changelogs are for humans, not machines.
There should be an entry for every single version.
The same types of changes should be grouped.
Versions and sections should be linkable.
The latest version comes first.
The release date of each version is displayed.
Mention whether you follow Semantic Versioning.

Usage:

Change log entries are to be added to the Unreleased section under the
appropriate stanza (see below). Each entry should ideally include a tag and
the Github issue reference in the following format:

* (<tag>) \#<issue-number> message

The issue numbers will later be link-ified during the release process so you do
not have to worry about including a link manually, but you can if you wish.

Types of changes (Stanzas):

"Features" for new features.
"Improvements" for changes in existing functionality.
"Deprecated" for soon-to-be removed features.
"Bug Fixes" for any bug fixes.
"Client Breaking" for breaking CLI commands and REST routes used by end-users.
"API Breaking" for breaking exported APIs used by developers building on SDK.
"State Machine Breaking" for any changes that result in a different AppState given same genesisState and txList.

Ref: https://keepachangelog.com/en/1.0.0/
-->

# Changelog

## [v12.1.19] - 2025-05-20

### Improvement 
- (deps) [#87](https://github.com/Bridgeless-Project/bridgeless-core/pull/87) Added detailed docs for `x/bridge` module. Updated ibc-go import in go.mod to v6.1.9

### Client Breaking
- (deps) [#87](https://github.com/Bridgeless-Project/bridgeless-core/pull/87) Simplified bridge module CLI transaction queries

### Bug Fixes
- (deps) [#92](https://github.com/Bridgeless-Project/bridgeless-core/pull/92) Remove the ability to send feeless transactions over the EVM module.


## [v12.1.18] 

### Improvement
- (deps) [#86](https://github.com/Bridgeless-Project/bridgeless-core/pull/86) Updated cosmos-sdk version import to v0.46.33
- Added app.toml and genesis.json configs

## [v12.1.17-rc3] 

### Bug Fixes
  - (deps) [#83](https://github.com/Bridgeless-Project/bridgeless-core/pull/83) Removed unnecessary `commission_amount` field from `bridge.TransactionSubmissions`

## [v12.1.17-rc2] - 2025-05-05

### Bug Fixes
- (deps) [#81](https://github.com/Bridgeless-Project/bridgeless-core/pull/81) Removed checking for tx existence on core while submission

## [v12.1.17-rc1] - 2025-05-01

### Improvement
- (deps) [#74](https://github.com/Bridgeless-Project/bridgeless-core/pull/74) Added new param `commission_rate` to `x/bridge` token. Added new param `commission_amount` to `x/bridge` transaction, new param `tss_threshold` to `x/bridge` module params.  Added storage of transaction submissions to `x/bridge` module. Updated `SubmitTx` logic with added `tss_threshold` param

### Bug Fixes
- (deps) [#74](https://github.com/Bridgeless-Project/bridgeless-core/pull/74) Fixed query params in `GetTokenById` message

## [v12.1.16-rc1] - 2025-02-03

### Improvement

- (deps) [#64](https://github.com/Bridgeless-Project/bridgeless-core/pull/64) Added new param `parties` to `x/bridge` module params. Improved message handling in `x/bridge` module - added `SubmitEvent` event on transaction submission to core.
- (deps) [#65](https://github.com/Bridgeless-Project/bridgeless-core/pull/65) Changed Transaction structure in `x/bridge` module

## [v12.1.15-rc1]

### Improvement
- (deps) [#57](https://github.com/Bridgeless-Project/bridgeless-core/pull/57) Added ZANO chain type support to `x/bridge` module
- (deps) [#60](https://github.com/Bridgeless-Project/bridgeless-core/pull/60) Removed tokens burning while slashing validator
- (deps) [#61](https://github.com/Bridgeless-Project/bridgeless-core/pull/61) Added custom ZANO chain validation

## [v12.1.14-rc1]

### Bug Fixes
- (deps) [#53](https://github.com/Bridgeless-Project/bridgeless-core/pull/53) Fix `wrong fee denomination; got: ; required: abridge: insufficient fee: insufficient fee` error

### Improvement
- (deps) [#52](https://github.com/Bridgeless-Project/bridgeless-core/pull/52) Added `x/multisig` module


## [v12.1.13-rc1]

### Improvement
- (deps) [#49](https://github.com/Bridgeless-Project/bridgeless-core/pull/49) Bump cosmos-sdk version. Fixed module path, NFT module logic, adding hook to bank module

## [v12.1.12-rc1]

### Improvement
- (deps) [#47](https://github.com/Bridgeless-Project/bridgeless-core/pull/47) Upgrading bridge module with new logic for TokenE bridging contracts (adding signature, adding is_wrapped fields)

## [v12.1.11-rc1]

### Improvement
- (deps) [#45](https://github.com/Bridgeless-Project/bridgeless-core/pull/45) Upgraded ibc-go version to v6.1.6
- (deps) [#46](https://github.com/Bridgeless-Project/bridgeless-core/pull/46) Upgraded comsos-sdk version

## [v12.1.10-rc1]

### Improvement
- (deps) [#43](https://github.com/Bridgeless-Project/bridgeless-core/pull/43) Added `token_info` fields to `x/bridge` Token

## [v12.1.9-rc4]

### Bug Fixes
- Disabled repeated call of InitGenesis for new bridge module

## [v12.1.9-rc3]

### Improvement
- Removed unused field `evm_admin` from `x/bridge` module

## [v12.1.9-rc2]
 
### Improvement
- Added migrations for v12.1.9-rc2

## [v12.1.9-rc1]

### Improvement
- (deps) [#40](https://github.com/Bridgeless-Project/bridgeless-core/pull/40) Added store for `x/bridge` module
- (deps) [#42](https://github.com/Bridgeless-Project/bridgeless-core/pull/42) Added `x/bridge` module

## [v12.1.8-rc2]
### Improvement
- Updated tutorial docs at VALIDATOR.md. Updated cosmos-sdk version to v0.46.24

## [v12.1.7-dev.2]
### Improvement
- (deps) [#29](https://github.com/Bridgeless-Project/bridgeless-core/pull/29) Added `x/nft` module
- (deps) [#30](https://github.com/Bridgeless-Project/bridgeless-core/pull/30) Moved `x/nft` module to cosmos-sdk
- (deps) [#31](https://github.com/Bridgeless-Project/bridgeless-core/pull/31) Added CI
- (deps) [#32](https://github.com/Bridgeless-Project/bridgeless-core/pull/32) Updated cosmos-sdk version to v0.46.16
- (deps) [#33](https://github.com/Bridgeless-Project/bridgeless-core/pull/33) Updated cosmos-sdk version to v0.46.17
- (deps) [#34](https://github.com/Bridgeless-Project/bridgeless-core/pull/34) Removed unused modules and fix staking hooks initialization
- (deps) [#35](https://github.com/Bridgeless-Project/bridgeless-core/pull/35) Added genesis to Dockerfile

## [v12.1.6] - 2023-07-04

### Improvement

- (deps) [#1635](https://github.com/evmos/evmos/pull/1635) Update cometbft `v0.34.29` with several minor bug fixes and low-severity security-fixes

## [v12.1.5] - 2023-06-08

### Bug Fixes

- (vesting) [GHSA-2q3r-p2m3-898g](https://github.com/evmos/evmos/commit/39b750cdaf1d69158ab93da85bd43ae4a7da1456) Apply ClawbackVestingAccount Barberry patch & Bump SDK to v0.46.13


## [v12.1.4] - 2023-05-26

### Improvements

- (deps) [#1571](https://github.com/evmos/evmos/pull/1571) Bump IBC-go version to [`v6.1.1`](https://github.com/cosmos/ibc-go/releases/tag/v6.1.1)

### Bug Fixes

- (ci) [#1546](https://github.com/evmos/evmos/pull/1546) Fix docker image push on release action
- (ci) [#1475](https://github.com/evmos/evmos/pull/1475) Fix version of GitHub action to push to the [Buf Schema Registry](https://buf.build/evmos/evmos) upon releases

## [v12.1.3] - 2023-05-24

### Improvements

- (cli) [#1556](https://github.com/evmos/evmos/pull/1556) Add CLI subcommand to debug legacy EIP712 transaction data

### Bug Fixes

- (deps) [#1566](https://github.com/evmos/evmos/pull/1566) Bump cosmos-sdk version to `v0.46.10-ledger.3`.
  Fix memory leak in `cosmos/iavl` package.

## [v12.1.2] - 2023-04-14

### Bug Fixes

- (rpc) [#1431](https://github.com/evmos/evmos/pull/1431) Fix websocket connection id parsing

## [v12.1.1] - 2023-04-14

### Improvements

- (config) [#1513](https://github.com/evmos/evmos/pull/1513) Set default `timeout_commit` to `3s`

## [v12.1.0] - 2023-03-24

### State Machine Breaking

- (evm)[#1308](https://github.com/evmos/evmos/pull/1308) Migrate `evm` and `feemarket` types
- (contracts) [#1306](https://github.com/evmos/evmos/pull/1306) Migrate `contracts` directory to evmos repository
- (proto) [#1305](https://github.com/evmos/evmos/pull/1305) Migrate Ethermint proto files
- (ante) [#1266](https://github.com/evmos/evmos/pull/1266) Use `DynamicFeeChecker` for Cosmos txs.
- (ante) [#1403](https://github.com/evmos/evmos/pull/1403) Update `AnteHandler` decorator for `x/authz` messages to run in deliverTx mode
- (eip712) [#1390](https://github.com/evmos/evmos/pull/1390) Refactor EIP-712 message handling to support multiple message schemas
- (ante) [#1405](https://github.com/evmos/evmos/pull/1405) Enable fees to be deducted from unclaimed staking rewards

### API Breaking

- [#1426](https://github.com/evmos/evmos/pull/1426) Move `revenue` module files into `v1` directory.
- [#1355](https://github.com/evmos/evmos/pull/1355) Remove `vm` package from EVM.

### Improvements

- (tests) [#1434](https://github.com/evmos/evmos/pull/1434) Set default staking denom to `aevmos` in `evm` and `feemarket` tests
- (test) [#1402](https://github.com/evmos/evmos/pull/1402) Refactor NewTx function arguments
- (test) [#1415](https://github.com/evmos/evmos/pull/1415) Refactor InvalidTx type and NextFn used in AnteHandler tests
- (vesting) [#1400](https://github.com/evmos/evmos/pull/1400) Add convert vesting account message
- (test) [#1393](https://github.com/evmos/evmos/pull/1393) Move utilities from `tests` folder to `testutil` package
- (test) [\#1391](https://github.com/evmos/evmos/pull/1391) Refactor test files
- (claims) [#1378](https://github.com/evmos/evmos/pull/1378) Validate authorized channels when updating claims params
- (test) [#1348](https://github.com/evmos/evmos/pull/1348) Add query executions to e2e upgrade test suite
- (deps) [#1370](https://github.com/evmos/evmos/pull/1370) Bump Cosmos SDK version to [`v0.46.9-ledger`](https://github.com/evmos/cosmos-sdk/releases/tag/v0.46.9-ledger)
- (deps) [#1370](https://github.com/evmos/evmos/pull/1370) Bump Tendermint version to [`v0.34.26`](https://github.com/informalsystems/tendermint/releases/tag/v0.34.26)
- (evm) [#1354](https://github.com/evmos/evmos/pull/1354) Expose `Context` from the `StateDB` instance.
- (proto)[#1311](https://github.com/evmos/evmos/pull/1311) Also generate common types with `make proto-gen`
- (revenue)[#1153](https://github.com/evmos/evmos/pull/1153) Migrate revenue module event emitting to `TypedEvent`
- (erc20) [#1152](https://github.com/evmos/evmos/pull/1152) Migrate event emitting to `TypedEvent`
- (claims) [#1126](https://github.com/evmos/evmos/pull/1126) Remove old x/params migration logic
- (vesting) [#1155](https://github.com/evmos/evmos/pull/1155) Migrate deprecated event emitting to new `TypedEvent`
- (docs) [#1361](https://github.com/evmos/evmos/pull/1361) Update `vesting` module docs with new behavior for `ClawbackVestingAccounts`
- (evm) [#1349](https://github.com/evmos/evmos/pull/1349) Restrict the Evmos codebase from working with chain IDs other than `9000` and `9001`
- (test) [#1352](https://github.com/evmos/evmos/pull/1352) Deprecate usage of `aphoton` as denomination on tests
- (test) [#1369](https://github.com/evmos/evmos/pull/1369) Refactor code to use `BaseDenom` for simplification
- (cli) [#1371](https://github.com/evmos/evmos/pull/1371) Improve cli error messages
- (ante) [#1380](https://github.com/evmos/evmos/pull/1380) Split vesting decorators between `evm` and `cosmos` packages
- (cli) [#1386](https://github.com/evmos/evmos/pull/1386) Use required fees (i.e `--fees=auto`) as default if fees are not specified
- (test) [#1408](https://github.com/evmos/evmos/pull/1408) Refactor `DeployContract` and `DeployContractWithFactory` functions used for tests
- (test) [#1417](https://github.com/evmos/evmos/pull/1417) Refactor EIP-712 transactions helper functions used on tests
- (ante) [#1468](https://github.com/evmos/evmos/pull/1468) Add TxFeeChecker requirement
- (deps) [#1473](https://github.com/evmos/evmos/pull/1473) Bump Cosmos SDK version to [v0.46.10-alpha.ledger.2](https://github.com/evmos/cosmos-sdk/releases/tag/v0.46.10-alpha.ledger.2)
- (ante) [#1470](https://github.com/evmos/evmos/pull/1470) Improve error message on `DynamicFeeChecker` ante handler
- (test) [#1484](https://github.com/evmos/evmos/pull/1484) Update e2e test: refactor Makefile command and use latest changes for the tests
- (deps) [#1498](https://github.com/evmos/evmos/pull/1498) Bump Cosmos SDK version to [v0.46.10-ledger.1](https://github.com/evmos/cosmos-sdk/releases/tag/v0.46.10-ledger.1)
- (lint) [#1487](https://github.com/evmos/evmos/pull/1487) Fix lint issues created by new `golangci-lint` version

### Bug Fixes

- (ante) [#1433](https://github.com/evmos/evmos/pull/1433) Add max priority fee check on `FeeChecker`.
- (ci) [#1383](https://github.com/evmos/evmos/pull/1383) Fix go-releaser error when building macOS binaries
- (ante) [#1435](https://github.com/evmos/evmos/pull/1435) Add block gas limit check for cosmos transactions
- (evm) [#1452](https://github.com/evmos/evmos/pull/1452) Consider refund amount on `gasUsed` calculation
- (evm) [#1466](https://github.com/evmos/evmos/pull/1466) Add `gasUsed` field in Ethereum transaction receipt
- (cli) [#1467](https://github.com/evmos/evmos/pull/1467) Rollback fees `auto` flag logic
- (ci) [#1476](https://github.com/evmos/evmos/pull/1476) Fix go-releaser configuration to be consistent with previous version binaries naming
- (upgrade) [#1493](https://github.com/evmos/evmos/pull/1493) Add decay bug affected accounts

## [v11.0.2] - 2023-02-10

### Improvements

- (deps) [#1370](https://github.com/evmos/evmos/pull/1370) Bump Cosmos SDK version to [`v0.46.9-ledger`](https://github.com/evmos/cosmos-sdk/releases/tag/v0.46.9-ledger)
- (deps) [#1370](https://github.com/evmos/evmos/pull/1370) Bump Tendermint version to [`v0.34.26`](https://github.com/informalsystems/tendermint/releases/tag/v0.34.26)
- (deps) [#1374](https://github.com/evmos/evmos/pull/1374) Bump Gin version to [`v1.7.7`](https://github.com/gin-gonic/gin/releases/tag/v1.7.7)
- (ante) [#1382](https://github.com/evmos/evmos/pull/1382) Add `AnteHandler` decorator for `x/authz` messages

## [v11.0.1] - 2023-02-04

### Improvements

- (deps) [#1248](https://github.com/evmos/evmos/pull/1248) Use the Informal Systems Tendermint Core fork

### Bug Fixes

- (deps) [#1342](https://github.com/evmos/evmos/pull/1342) Bump `tendermint` to [`v0.34.25`](https://github.com/informalsystems/tendermint/releases/tag/v0.34.25)

## [v11.0.0] - 2023-01-27

### State Machine Breaking

- (deps) [#1288](https://github.com/evmos/evmos/pull/1288) Bump `ethermint` to [`v0.21.0`](https://github.com/evmos/ethermint/releases/v0.21.0)
- (ica) [#1101](https://github.com/evmos/evmos/pull/1101) Add ICA host submodule
- (inflation) [#1210](https://github.com/evmos/evmos/pull/1210) Delete `EpochMintProvision` from `KVStore` in a migration
- (deps) [\#1196](https://github.com/evmos/evmos/pull/1196) Bump `ibc-go` to [`v6.1.0`](https://github.com/cosmos/ibc-go/releases/tag/v6.1.0)
- (inflation) [#1193](https://github.com/evmos/evmos/pull/1193) Remove `EpochMintProvision` setters and getters to calculate on the fly
- (erc20) [#1100](https://github.com/evmos/evmos/pull/1100) Deprecate usage of `x/params` in `x/erc20`
- (inflation) [#1107](https://github.com/evmos/evmos/pull/1107) Deprecate usage of `x/params` in `x/inflation`
- (incentives) [#1130](https://github.com/evmos/evmos/pull/1130) Deprecate usage of `x/params` in `x/incentives`
- (claims) [#1125](https://github.com/evmos/evmos/pull/1125) Deprecate usage of `x/params` in `x/claims`
- (revenue) [#1129](https://github.com/evmos/evmos/pull/1129) Deprecate usage of `x/params` in `x/revenue`
- (vesting) [#1268](https://github.com/evmos/evmos/pull/1268) Allow usage of vested and unlocked tokens in EVM interactions

### Features

- (upgrade) [#1209](https://github.com/evmos/evmos/pull/1209) Incentivized testnet reward distribution logic.

### Improvements

- (tests) [#1283](https://github.com/evmos/evmos/pull/1283) Enable multiple upgrades for automated upgrade tests
- (deps) [#1279](https://github.com/evmos/evmos/pull/1279) Bump Cosmos SDK version to [`v0.46.8-ledger`](https://github.com/evmos/cosmos-sdk/releases/tag/v0.46.8-ledger)
- (inflation) [#1258](https://github.com/evmos/evmos/pull/1258) Remove unnecessary `Coin` validation and store calls for `Params`

### Bug Fixes

- (app) [#1276](https://github.com/evmos/evmos/pull/1276) Fix store uploader for `x/recovery` module.
- (inflation) [#1259](https://github.com/evmos/evmos/pull/1259) Re-add missing key to not disrupt order in store
- (upgrade) [#1257](https://github.com/evmos/evmos/pull/1257) Add `recovery` module store to `StoreUpgrades`
- (upgrade) [#1252](https://github.com/evmos/evmos/pull/1252) Add account number and sequence to migrated IBC transfer escrow accounts.
- (upgrade) [#1242](https://github.com/evmos/evmos/pull/1242) Fix Ethermint params upgrade
- (ibc) [#1156](https://github.com/evmos/evmos/pull/1156) Migrate IBC transfer escrow accounts to `ModuleAccount` type.
- (upgrade) [#1252](https://github.com/evmos/evmos/pull/1252) Add account number and sequence to migrated IBC transfer escrow accounts.

## [v10.0.1] - 2023-01-03

### Improvements

- (deps) [#1201](https://github.com/evmos/evmos/pull/1201) Bump `ics23/go` to v0.9.0

## [v10.0.0] - 2022-12-28

### State Machine Breaking

- (deps) [#1184](https://github.com/evmos/evmos/pull/1184) Bump Ethermint version to [`v0.20.0-rc5`](https://github.com/evmos/ethermint/releases/tag/v0.20.0-rc5)
- (deps) [\#1176](https://github.com/evmos/evmos/pull/1176) Bump `ibc-go` to [`v5.2.0`](https://github.com/cosmos/ibc-go/releases/tag/v5.2.0)
- (vesting) [\#1070](https://github.com/evmos/evmos/pull/1070) Add Amino encoding support to the vesting module for EIP-712 signing.
- (ante) [#1054](https://github.com/evmos/evmos/pull/1054) Remove validator commission `AnteHandler` decorator and replace it with the new `MinCommissionRate` staking parameter.
- (deps) [\#1041](https://github.com/evmos/evmos/pull/1041) Add ICS-23 dragon-berry replace in `go.mod` as mentioned in the [Cosmos SDK release](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.46.4)

### API Breaking

- (erc20) [\#914](https://github.com/evmos/evmos/pull/914) Support registering multiple assets on `RegisterCoinProposal` and `RegisterERC20Proposal`

### Features

- (app) [\#1114](https://github.com/evmos/evmos/pull/1114) Add default File store listener for application from [ADR38](https://docs.cosmos.network/v0.47/architecture/adr-038-state-listening)
- (transfer, erc20) [\#1085](https://github.com/evmos/evmos/pull/1085) Added wrapper for ICS-20 `transfer` to automatically convert ERC-20 tokens to native Cosmos coins.

### Improvements

- (tests) [\1194](https://github.com/evmos/evmos/pull/1194) Lint tests so they are consistent with non-test code.
- (deps) [\#1176](https://github.com/evmos/evmos/pull/1176) Bump Cosmos SDK to [`v0.46.7`](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.46.7)
- (ci) [#1138](https://github.com/evmos/evmos/pull/1138) Add Golang dependency vulnerability checker.
- (docs) [\#1090](https://github.com/evmos/evmos/pull/1090) Add audits page to documentation.
- (vesting) [\#1087](https://github.com/evmos/evmos/pull/1087) Add new `MsgUpdateVestingFunder` that updates the `Funder` field of a given clawback vesting account
- (ibc) [\#1081](https://github.com/evmos/evmos/pull/1081) Added utils functions for ibc denoms.
- (erc20) [\#1059](https://github.com/evmos/evmos/pull/1059) Add util functions (iterator and params) for ERC20 module.
- (go) [\#1039](https://github.com/evmos/evmos/pull/1039) Bump go v1.19
- (test) [#1028](https://github.com/evmos/evmos/pull/1028) Add node upgrade end-to-end test suite.
- (cmd) [\#1027](https://github.com/evmos/evmos/pull/1027) Apply Google CLI Syntax for required and optional args.
- (ante) [\#993](https://github.com/evmos/evmos/pull/993) Re-order AnteHandlers for better performance
- (docs) [\#985](https://github.com/evmos/evmos/pull/985) Specify repo branch name on markdown-link-check configuration.
- (docs) [\#883](https://github.com/evmos/evmos/pull/883) Add Ethereum tx indexer documentation.
- (docs) [\#980](https://github.com/evmos/evmos/pull/980) Fix documentation links to cosmos-sdk docs.
- (cmd) [\#974](https://github.com/evmos/evmos/pull/974) Add `prune` command.
- (cli) [#816](https://github.com/evmos/evmos/pull/816) Add Ledger CLI support.

### Bug Fixes

- (app) [#1165](https://github.com/evmos/evmos/pull/1165) Update Ledger supported algorithms to only consist of `EthSecp256k1`
- (cmd) [#1172](https://github.com/evmos/evmos/pull/1172) Update default node snapshot interval to `5000`
- (cmd) [\#1121](https://github.com/evmos/evmos/pull/1121) Fix `evmosd version` to show either tag or last commit.
- (cmd) [\#1120](https://github.com/evmos/evmos/pull/1120) Fix snapshot configuration
- (app) [\#1118](https://github.com/evmos/evmos/pull/1118) Setup gRPC node service with the application.
- (analytics) [\#1094](https://github.com/evmos/evmos/pull/1094) Fix unbound metrics and remove labels that keep increasing.

## [v9.1.0] - 2022-10-25

### Improvements

- (deps) [\#1011](https://github.com/evmos/evmos/pull/1011) Bump Cosmos SDK to [`v0.45.10`](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.45.10)

## [v9.0.0] - 2022-10-21

### Bug Fixes

- (claims) [#951](https://github.com/evmos/evmos/pull/951) Fix ClawbackEmptyAccounts logic

## [v8.2.3] - 2022-10-15

### Improvements

- (cmd) [#988](https://github.com/evmos/evmos/pull/988) Set IAVL config
- (cli) [#971](https://github.com/evmos/evmos/pull/971) Add `prune` command.

## [v8.2.2] - 2022-10-14

### Improvements

- (deps)[#965](https://github.com/evmos/evmos/pull/965) Bump SDK to v0.45.9 and Ethermint to v0.19.3

## [v8.2.0] - 2022-09-23

### State Machine Breaking

- (app) [\#918](https://github.com/evmos/evmos/pull/918) Fix unregistered `revenue` module for `v8.1.0` store upgrade

### Bug Fixes

- (app,docs) [\#933](https://github.com/evmos/evmos/pull/933) Replace invalid linux `x86_64` [architecture](https://go.dev/doc/install/source#environment) to `amd64`.

## [v8.1.1] - 2022-09-23

### Bug Fixes

- (app) [\#922](https://github.com/evmos/evmos/pull/922) Add hard fork logic for `v8.2.0`

## [v8.1.0] - 2022-08-30

### State Machine Breaking

- (revenue) [\#859](https://github.com/evmos/evmos/pull/859) Add amino codecs to `x/revenue` module to support EIP-712 signatures.
- (deps) Bump Ethermint version to [`v0.19.2`](https://github.com/evmos/ethermint/releases/tag/v0.19.2)

## [v8.0.0] - 2022-08-16

### State Machine Breaking

- (deps) [\#845](https://github.com/evmos/evmos/pull/845) Bump Ethermint version to [`v0.19.0`](https://github.com/evmos/ethermint/releases/tag/v0.19.0)
- (revenue) Add `x/revenue` module

### Improvements

- (deps) [\#839](https://github.com/evmos/evmos/pull/839) Bump ibc-go to [`v3.2.0`](https://github.com/cosmos/ibc-go/releases/tag/v3.2.0) and Cosmos SDK to [`v0.45.7`](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.45.7).
- (build) [\#725](https://github.com/evmos/evmos/pull/725) Migrate Protobuf code generator to [Protobuf Workspaces](https://docs.buf.build/reference/workspaces)

### Bug Fixes

- (build) [\#856](https://github.com/evmos/evmos/pull/856) Update docker base image to use golang:1.18.5-bullseye and expose other relevant ports

## [v7.0.0] - 2022-08-04

### State Machine Breaking

- (deps) Bump Ethermint version to [`v0.18.0`](https://github.com/evmos/ethermint/releases/tag/v0.18.0)

### Bug Fixes

- (app) [\#760](https://github.com/evmos/evmos/pull/760) Migrate inaccessible balance of testnet faucet account to new address
- (inflation) [\#748](https://github.com/evmos/evmos/pull/748) Remove overcounted epochs from `skippedEpochs` value in store

## [v6.0.3] - 2022-07-26

### Bug Fixes

- (deps) [\#803](https://github.com/evmos/evmos/pull/803) Bump Ethermint version to [`v0.17.2`](https://github.com/evmos/ethermint/releases/tag/v0.17.2)

## [v6.0.2] - 2022-07-13

### Bug Fixes

- (deps) [\#769](https://github.com/evmos/evmos/pull/769) Bump Ethermint version to [`v0.17.1`](https://github.com/evmos/ethermint/releases/tag/v0.17.1)

## [v6.0.1] - 2022-06-28

### Improvements

- (ci) [\#729](https://github.com/evmos/evmos/pull/729) Remove unshallow action in goreleaser.

## [v6.0.0] - 2022-06-28

### State Machine Breaking

- (deps) [\#719](https://github.com/evmos/evmos/pull/719) Bump Ethermint version to [`v0.17.0`](https://github.com/evmos/ethermint/releases/tag/v0.17.0)

### API Breaking

- (all) [\#701](https://github.com/evmos/evmos/pull/703) Rename Go module to `evmos/evmos`

### Improvements

- (deps) [\#714](https://github.com/evmos/evmos/pull/714) Bump Go version to `1.18`.
- (cmd) [\#696](https://github.com/evmos/evmos/pull/696) Set a custom tendermint node configuration on initialization.
- (revenue) [\#685](https://github.com/evmos/evmos/pull/685) Internal Specification audit.
- (revenue) [\#691](https://github.com/evmos/evmos/pull/691) Internal API audit.
- (revenue) [\#715](https://github.com/evmos/evmos/pull/715) Internal state machine audit.

## [v5.0.0] - 2022-06-14

### State Machine Breaking

- (deps) [\#684](https://github.com/evmos/evmos/pull/684) Bump ibc-go version to [`v3.1.0`](https://github.com/cosmos/ibc-go/releases/tag/v3.1.0)
- (vesting) [\#666](https://github.com/evmos/evmos/pull/666) Remove support of Cosmos SDK `VestingAccount` types.
- (deps) [\#663](https://github.com/evmos/evmos/pull/663) Bump Ethermint version to [`v0.16.1`](https://github.com/evmos/ethermint/releases/tag/v0.16.1)
- (claims) [\#605](https://github.com/evmos/evmos/pull/605) Remove duplicated `SetClaimsRecord`.
- (erc20) [\#602](https://github.com/evmos/evmos/pull/602) Modified `RegisterERC20` proposals.
  Fix erc20 name sanitization to allow spaces on token name.

### API Breaking

- (claims) [\#605](https://github.com/evmos/evmos/pull/605) Remove `claims-` prefix in CLI query commands.
- (erc20) [\#592](https://github.com/evmos/evmos/pull/592) Finish module completeness audit.
- (analytics) [\#637](https://github.com/evmos/evmos/pull/637) Add telemetry to Evmos modules.
- (vesting) [\#643](https://github.com/evmos/evmos/pull/643) Remove the `create-vesting-account` CLI command from Cosmos SDK in favor of the clawback vesting accounts.

### Improvements

- (erc20) [\#677](https://github.com/evmos/evmos/pull/677) Add Amino registration to `ConvertCoin` and `ConvertERC20` msgs for ERC712 compatibility.
- (deps) [\#668](https://github.com/evmos/evmos/pull/668) Bump Cosmos SDK to [`v0.45.5`](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.45.5)
- (erc20) [\#642](https://github.com/evmos/evmos/pull/642) Remove enforcing ibc and channel names during `RegisterCoin`

### Bug Fixes

- (app) [\#682](https://github.com/evmos/evmos/pull/682) Fix Tendermint consensus params (Evidence `MaxAgeNumBlocks` and `MaxAgeDuration`)
- (incentives) [\#656](https://github.com/evmos/evmos/pull/656) Fix incentives that were previously only allocated to `EthAccount`s.
- (revenue) [\#612](https://github.com/evmos/evmos/pull/612) Fix fees registration cli command and description
- (inflation) [\#554](https://github.com/evmos/evmos/pull/554) Changing erroneous epoch skips to `daily` instead of `weekly`
- (claims) [\#626](https://github.com/evmos/evmos/pull/626) fix durations denominated in `nanoseconds`
- (epochs) [\#629](https://github.com/evmos/evmos/pull/629) fix epochs durations denominated in `nanoseconds`

## [v4.0.1] - 2022-05-10

### Bug Fixes

(erc20) [\#588](https://github.com/evmos/evmos/pull/588) Revert PR [\#556](https://github.com/evmos/evmos/pull/556).

## [v4.0.0] - 2022-05-09

### State Machine Breaking

- (app) [\#537](https://github.com/evmos/evmos/pull/537) Fix router key for IBC client proposals.
- (erc20) [\#530](https://github.com/evmos/evmos/pull/530) Use the highest denom unit when deploying an ERC20 contract.

### API Breaking

- (upgrade) [\#557](https://github.com/evmos/evmos/pull/557) Update Evmos go.mod version `v3` -> `v4`
- (erc20) [\#544](https://github.com/evmos/evmos/pull/544) Remove `updateTokenPairERC20Proposal` functionality rename `relay` to `conversion`
- (inflation) [\#536](https://github.com/evmos/evmos/pull/536) Rename inflation endpoint `/evmos/inflation/v1/total_supply` -> `/evmos/inflation/v1/circulating_supply`

### Improvements

- (deps) [\#580](https://github.com/evmos/evmos/pull/580) Bump Ethermint to [`v0.15.0`](https://github.com/evmos/ethermint/releases/tag/v0.15.0)
- (gitpod) [\#564](https://github.com/evmos/evmos/pull/564) Add one-click development environment
- (erc20) [\#556](https://github.com/evmos/evmos/pull/556) Remove deprecated migrations.
- (incentives) [\#551](https://github.com/evmos/evmos/pull/551) Add additional check to only distribute incentives to EOAs.
- (cmd) [\#543](https://github.com/evmos/evmos/pull/543) Update mainnet default `min-gas-price` to `0.0025aevmos`.
- (epochs) [\#539](https://github.com/evmos/evmos/pull/539) Use constants for epoch identifiers.

### Bug Fixes

- (erc20) [\#530](https://github.com/evmos/evmos/pull/530) Fix `Metadata` equal check for denom units.
- (app) [\#523](https://github.com/evmos/evmos/pull/523) Fix testnet upgrade store loader.

## [v3.0.1] - 2022-05-09

### Improvements

- (app) [\#555](https://github.com/evmos/evmos/pull/555) `v4.0.0` upgrade logic.
- (p2p) [\#541](https://github.com/evmos/evmos/pull/541) Increase default inbound connections and use 8:1 ratio of inbound:outbound.
  Add default seeds to reduce the need for configuration.
- (deps) [\#528](https://github.com/evmos/evmos/pull/528) Bump Cosmos SDK to [`v0.45.4`](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.45.4)

## [v3.0.0] - 2022-04-22

### State Machine Breaking

- [\#342](https://github.com/evmos/evmos/pull/342) Implement IBC middleware to recover stuck funds

### API Breaking

- [\#415](https://github.com/evmos/evmos/pull/415) Bump Evmos go version to v3

### Bug Fixes

- (vesting) [\#502](https://github.com/evmos/evmos/pull/502) Fix gas exhaustion bug by removing `SpendableCoins` during vesting account clawback.
- (vesting) [\#483](https://github.com/evmos/evmos/pull/483) Fix balance clawback when vesting start time is in the future
- (claims) [\#381](https://github.com/evmos/evmos/pull/381) Fix claim migration and deletion for EVM chains via IBC attestation.
- (claims) [\#374](https://github.com/evmos/evmos/pull/374) Fix balance invariant in Claims `InitGenesis`
- (erc20) [\#366](https://github.com/evmos/evmos/issues/366) Delete ERC20 denom map when deleting pair.
- (claims) [\#505](https://github.com/evmos/evmos/pull/505) Fix IBC attestation ordering

### Improvements

- (vesting) [\#486](https://github.com/evmos/evmos/pull/486) Refactor `x/vesting` types and tests.
- (erc20) [\#484](https://github.com/evmos/evmos/pull/484) Avoid unnecessary commits to the StateDB and don't estimate gas when performing a query
- (deps) [\#478](https://github.com/evmos/evmos/pull/478) Bump Cosmos SDK to [`v0.45.3`](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.45.3)
- (deps) [\#478](https://github.com/evmos/evmos/pull/478) Bump Ethermint to [`v0.14.0`](https://github.com/evmos/ethermint/releases/tag/v0.14.0)
- (vesting) [\#468](https://github.com/evmos/evmos/pull/468) Use coins `Min` function from Cosmos SDK.
- (cmd) [\#446](https://github.com/evmos/evmos/pull/446) Update `migrate` command to migrate Evmos, Ethermint and Cosmos SDK modules.
- (app) [\#446](https://github.com/evmos/evmos/pull/446) Refactor upgrade code.
- (ibc) [\#412](https://github.com/evmos/evmos/pull/412) Introduce boilerplate struct for IBC applications.
- (deps) [\#402](https://github.com/evmos/evmos/pull/402) Bump IBC go to [`v3.0.0`](https://github.com/cosmos/ibc-go/releases/tag/v3.0.0)
- (ibctesting) [\#388](https://github.com/evmos/evmos/pull/388) Support Cosmos and EVM chains in IBC testing `Coordinator`.
- (claims) [\#385](https://github.com/evmos/evmos/pull/385) Add claims invariant.
- (inflation) [\#383](https://github.com/evmos/evmos/pull/383) Add gRPC endpoints for inflation rate and total supply
- (inflation) [\#369](https://github.com/evmos/evmos/pull/369) Add `enableInflation` parameter.
- (claims) [\#432](https://github.com/evmos/evmos/pull/432) Add IBC trigger amount to claims merge/migrate IBC callbacks.
- (claims) [\#442](https://github.com/evmos/evmos/pull/443) Remove claims merge/migrate cases where sender already completed an action as they are never reached
- (claims) [\#507](https://github.com/evmos/evmos/pull/507) Always return claimable amount on grpc queries regarding of claims status.
- (claims) [\#516](https://github.com/evmos/evmos/pull/516) Retain claims records when all actions have been completed.

## [v2.0.1] - 2022-03-06

### Bug Fixes

- (upgrade) [#\363](https://github.com/evmos/evmos/pull/363) Don't use `GetParams` for upgrades.

## [v2.0.0] - 2022-03-06

### State Machine Breaking

- (claims) Restrict claiming to a list of authorized IBC channels.

### Improvements

- (deps) [\#360](https://github.com/evmos/evmos/pull/360) Bump Ethermint to [`v0.11.0`](https://github.com/evmos/ethermint/releases/tag/v0.11.0)
- (deps) [\#282](https://github.com/evmos/evmos/pull/282) Bump IBC go to [`v3.0.0-rc1`](https://github.com/cosmos/ibc-go/releases/tag/v3.0.0-rc1)

### Bug Fixes

- (erc20) [\#337](https://github.com/evmos/evmos/pull/337) Ignore errors in ERC20 module's EVM hook.
- (erc20) [\#336](https://github.com/evmos/evmos/pull/336) Return `nil` for disabled ERC20 module or ERC20 EVM hook.

## [v1.1.2] - 2022-03-06

### Bug Fixes

- (app) [\#354](https://github.com/evmos/evmos/pull/354) Add v2 version upgrade logic

## [v1.1.1] - 2022-03-04

### Improvements

- (deps) [\#345](https://github.com/evmos/evmos/pull/345) Bump Ethermint to [`v0.10.2`](https://github.com/evmos/ethermint/releases/tag/v0.10.2)

### Bug Fixes

- (app) [\#341](https://github.com/evmos/evmos/pull/341) Return error when `--ledger` flag is passed in CLI

## [v1.1.0] - 2022-03-02

### Bug Fixes

- (ante) [\#318](https://github.com/evmos/evmos/pull/318) Add authz check in vesting and min commission `AnteHandler` decorators.
- (vesting) [\#317](https://github.com/evmos/evmos/pull/317) Fix clawback for vested coins.

## [v1.0.0] - 2022-02-28

### State Machine Breaking

- (ante) [\#302](https://github.com/evmos/evmos/pull/302) Add AnteHandler decorator to enforce global min validator commission rate.
- (app) [\#224](https://github.com/evmos/evmos/pull/224) Fix power reduction my setting the correct value on app initialization.
- (keys) [\#189](https://github.com/evmos/evmos/pull/189) Remove support for Tendermint's `secp256k1` keys.
- [\#173](https://github.com/evmos/evmos/pull/173) Rename `intrarelayer` module to `erc20`
- [\#190](https://github.com/evmos/evmos/pull/190) Remove governance hook from `erc20` module

### Features

- [\#286](https://github.com/evmos/evmos/pull/286) Add `x/vesting` module.
- [\#184](https://github.com/evmos/evmos/pull/184) Add claims module for claiming the airdrop tokens.
- [\#183](https://github.com/evmos/evmos/pull/183) Add epoch module for incentives.
- [\#202](https://github.com/evmos/evmos/pull/202) Add custom configuration for state sync snapshots and tendermint p2p peers.
  This introduces a custom `InitCmd` function.
- [\#176](https://github.com/evmos/evmos/pull/176) Add `x/incentives` module.

### Improvements

- (deps) Bumped Ethermint to [`v0.10.0`](https://github.com/evmos/ethermint/releases/tag/v0.10.0)
- (deps) Bumped IBC-go to `v3.0.0-rc0`
- (deps) Bumped Cosmos SDK to [`v0.45.1`](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.45.1)
- (deps) bumped Tendermint to `v0.34.15` and tm-db to `v0.6.6`.
- [\#180](https://github.com/evmos/evmos/pull/180) Delete `TokenPair` if ERC20 contract has been selfdestructed.

### Bug Fixes

- (erc20) [\#169](https://github.com/evmos/evmos/pull/169) Fixes several testnet bugs:
  - Check if supply exists for a token before when submitting a `RegisterCoinProposal`, allowing users to create an ERC20 representation of an invalid Cosmos Coin.
  - Sanitize the ERC20 token name when creating coin metadata on ER `RegisterERC20Proposal`.
  - Fix coin metadata validation error when registering an ERC20 with 0 denom units.
- (erc20) [\#191](https://github.com/evmos/evmos/pull/191) Add direct balance protection (IF-ETHERMINT-06).
- (erc20) [\#192](https://github.com/evmos/evmos/pull/192) Add delayed malicious effect protection (IF-ETHERMINT-06).
- (erc20) [\#200](https://github.com/evmos/evmos/pull/200) Match coin and token decimals for erc20 deployment during registerCoin
- (erc20) [\#201](https://github.com/evmos/evmos/pull/201) bug(erc-20): Compile built-in contracts in the build process (IF-ETHERMINT-02).

## [v0.4.2] - 2021-12-11

### Bug Fixes

- (app) [\#166](https://github.com/evmos/evmos/pull/166) Fix `UpgradeHandler`.

## [v0.4.1] - 2021-12-07

### Improvements

- (build) [\#143](https://github.com/evmos/evmos/pull/143) Added `build-reproducible` rule in `Makefile` to build docker containers

### Bug Fixes

- (build) [\#151](https://github.com/evmos/evmos/pull/151) Fixes `version` command by picking the latest tag in the current branch instead of across all branches as the current version

## [v0.4.0] - 2021-12-02

### State Machine Breaking

- (erc20) [\#119](https://github.com/evmos/evmos/issues/119) Register `x/erc20` proposal types on governance module.

### Improvements

- (app) [\#128](https://github.com/evmos/evmos/pull/128) Add ibc-go `TestingApp` interface.
- (ci) [\#117](https://github.com/evmos/evmos/pull/117) Enable automatic backport of PRs.
- (deps) [\#135](https://github.com/evmos/evmos/pull/135) Bump Ethermint version to [`v0.9.0`](https://github.com/evmos/ethermint/releases/tag/v0.9.0)
- (ci) [\#136](https://github.com/evmos/evmos/pull/136) Deploy `evmos` docker container to [docker hub](https://hub.docker.com/u/tharsishq) for every versioned releases

### Bug Fixes

- (build) [\#116](https://github.com/evmos/evmos/pull/116) Fix `build-docker` command

## [v0.3.0] - 2021-11-24

### API Breaking

- (erc20) [\#99](https://github.com/evmos/evmos/pull/99) Rename `enable_e_v_m_hook` json parameter to `enable_evm_hook`.

### Improvements

- (deps) [\#110](https://github.com/evmos/evmos/pull/110) Bump Ethermint version to [`v0.8.1`](https://github.com/evmos/ethermint/releases/tag/v0.8.1)
- (erc20) [\#107](https://github.com/evmos/evmos/pull/107) Add IBC validation
- (cmd) [\#105](https://github.com/evmos/evmos/pull/105) Improve testnet command to include JSON-RPC client.

## Bug Fixes

- (erc20) [\#109](https://github.com/evmos/evmos/pull/109) Fix hardcoded erc20 nonce and `UpdateTokenPairERC20` proposal to support ERC20s with 0 decimals.
- (erc20) [\#102](https://github.com/evmos/evmos/pull/102) Add `convert-erc20` cmd

## [v0.2.0] - 2021-11-17

### Features

- (erc20) [\#82](https://github.com/evmos/evmos/pull/82) ERC20 module
- (cmd) [\#32](https://github.com/evmos/evmos/pull/32) Create `testnet` command that spins up a new local testnet with N nodes.

### Improvements

- (deps) [\#94](https://github.com/evmos/evmos/pull/94) Bump Ethermint version to [`v0.8.0`](https://github.com/evmos/ethermint/releases/tag/v0.8.0)
- (deps) [\#80](https://github.com/evmos/evmos/pull/80) Bump ibc-go to [`v2.0.0`](https://github.com/cosmos/ibc-go/releases/tag/v2.0.0)

## [v0.1.3] - 2021-10-24

### Improvements

- (deps) [\#64](https://github.com/evmos/evmos/pull/64) Bump Ethermint version to `v0.7.2`

### Bug Fixes

- (cmd) [\#41](https://github.com/evmos/evmos/pull/41) Fix `debug` command.

## [v0.1.2] - 2021-10-08

### Improvements

- (deps) [\#34](https://github.com/evmos/evmos/pull/34) Bump Ethermint version to `v0.7.1`

## [v0.1.1] - 2021-10-07

### Bug Fixes

- (build) [\#30](https://github.com/evmos/evmos/pull/30) Fix `version` command.

## [v0.1.0] - 2021-10-07

### Improvements

- (cmd) [\#26](https://github.com/evmos/evmos/pull/26) Use config on genesis accounts.
- (deps) [\#28](https://github.com/evmos/evmos/pull/28) Bump Ethermint version to `v0.7.0`
