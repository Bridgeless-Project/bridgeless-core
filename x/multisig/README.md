---
layout: default
title: x/multisig
---

# `x/multisig`

## Abstract

The main goal of this module is to provide implementation which allows the creation and management of on-chain
multi-signature accounts and enables message execution based on the multi-signed message.
As an example of implementation ```x/group``` module from Cosmos SDK will be taken.

----

## Concepts

The main differences between implementations:

- Decision policy will be stored in the group entity.
- There will be only one decision policy – threshold decision policy.
- No group member entity - the base account will be used instead.
- No group administrator – the group will be managed by itself. The group member has to send a multi-signed message to
  remove or add a new member. The same flow works for other operations unrelated to the group management.
- The group member can’t leave the group at will.
- If the message is valid – it will be executed automatically by the module, instead of manually execution.

The multi-signature feature works in the following way:

- Every group has its own subaccount that holds the tokens.
- Any of group members creates the proposal to execute some messages.
- Group members votes for the proposal
- After voting period module calculates the results.
- If 'YES' votes count reaches the threshold module executes the messages.
  By executing messages in module we can skip the signature check that default transaction do automatically.
  So the module 'replaces' the default tx signature check with its own rules.

----

## State

### Params

Definition:

```protobuf
message Params {
  option (gogoproto.goproto_stringer) = false;
  uint64 groupSequence = 1;
  uint64 proposalSequence = 2;
  uint64 prunePeriod = 3;
  uint64 votingPeriod = 4;
}
```

Example:

```json
{
  "groupSequence": 0,
  "proposalSequence": 0,
  "prunePeriod": 240,
  "votingPeriod": 120
}
```

### Group

**Group** - stores the aggregation of group member account addresses. It is a module’s sub-account and has a balance.
The group’s account address is derived from the module account address and group ID.
The module uses the group sequence parameter to generate a group ID for producing the group account address.
The group sequence is an integer stored in the module's parameters and increments during new group creation.

Definition:

```protobuf
message Group {
  string account = 1;
  repeated string members = 2;
  uint64 threshold = 3;
}
```

Example:

```json
{
  "account": "rarimo....",
  "members": [
    "rarimo...",
    "rarimo..."
  ],
  "threshold": 1
}
```

### Proposal

**Proposal** - proposal to execute some operation that should be signed by group.

Definition:

```protobuf
enum VoteOption {
  YES = 0;
  NO = 1;
}

// ProposalStatus defines proposal statuses.
enum ProposalStatus {
  // Initial status of a proposal when submitted.
  SUBMITTED = 0;
  // Status of a proposal when it passes the group's decision policy.
  ACCEPTED = 1;
  // Status of a proposal when it is rejected by the group's decision policy.
  REJECTED = 2;
  // Status of a proposal when it is successfully executed by the module.
  EXECUTED = 3;
  // Status of a proposal when execution is failed.
  FAILED = 4;
}

// Proposal defines a group proposal. Any member of a group can submit a proposal
// for a module to decide upon.
// A proposal consists of a set of `sdk.Msg`s that will be executed if the proposal
// passes as well.
message Proposal {
  // Account address of the proposer.
  string proposer = 1;
  // Unique id of the proposal.
  uint64 id = 2;
  // Account address of the group.
  string group = 3;
  // Block height when the proposal was submitted.
  uint64 submitBlock = 5;
  // Status represents the high level position in the life cycle of the proposal. Initial value is Submitted.
  ProposalStatus status = 8;
  // Contains the sums of all votes for this proposal for each vote option.
  // It is empty at submission, and only populated after tallying, at voting end block.
  TallyResult finalTallyResult = 9;
  // Block height before which voting must be done.
  uint64 votingEndBlock = 10;
  // List of `sdk.Msg`s that will be executed if the proposal passes.
  repeated google.protobuf.Any messages = 12;
}

// TallyResult represents the sum of votes for each vote option.
message TallyResult {
  // Sum of yes votes.
  uint64 yesCount = 1;
  // Sum of no votes.
  uint64 noCount = 3;
}

// Vote represents a vote for a proposal.
message Vote {
  // Unique ID of the proposal.
  uint64 proposalId = 1;
  // Voter is the account address of the voter.
  string voter = 2;
  // Option is the voter's choice on the proposal.
  VoteOption option = 3;
  // Block height when the vote was submitted.
  uint64 submitBlock = 5;
}
```

----

## Transactions

## RPC

### SubmitProposal

**SubmitProposal** - creates proposal to execute some set of messages signed by group account.

```protobuf
message MsgSubmitProposal {
  string creator = 1;
  string group = 2;
  repeated google.protobuf.Any messages = 3;
}
```

### Vote

**Vote** - vote for proposal.

```protobuf
message MsgVote {
  string creator = 1;
  uint64 proposalId = 2;
  VoteOption option = 3;
}
```

### CreateGroup

**CreateGroup** - creating of the new group.

```protobuf
message MsgCreateGroup {
  string creator = 1;
  repeated string members = 2;
  uint64 threshold = 3;
}
```

### ChangeGroup

**ChangeGroup** - changing group parameters or set of participants.
The signer of that message should be a group account, so that message can be executed only from that module using
multisig flow.

```protobuf
message MsgChangeGroup {
  string creator = 1;
  string group = 2;
  repeated string members = 3;
  uint64 threshold = 4;
}
```

----

## CLI

## Queries
### Params

The params command allow users to query `multisig` module params.
```
 simd query multisig params
```
Example output:
```
params:
  groupSequence: "0"
  proposalSequence: "0"
  prunePeriod: "240"
  votingPeriod: "120"

```
----
### GroupAll

The list-all command lists all existing groups.

```
simd query multisig groups list-all 
```
Example output:
```
group:
- account: bridge17fcahfu87uy6ejaapgx5xfsq3snrct6lq88km79zw9zckum70r4sdsx6vy
  members:
  - bridge1...
  - bridge1...
  threshold: "1"
- account: bridge1n77flmf2dhxzv7wh2ekxzn4wuz7rna06yqyw6sqp0pqte0fjue3snl5rvc
  members:
  - bridge1...
  - bridge1...
  threshold: "1"
pagination:
  next_key: null
  total: "2"


```

### Group

The group command lists info about certain group.

```
simd query multisig groups group bridge1...
```
Example output:
```
group:
  account: bridge1n77flmf2dhxzv7wh2ekxzn4wuz7rna06yqyw6sqp0pqte0fjue3snl5rvc
  members:
  - bridge1...
  - bridge1...
  threshold: "1"
```

----
### ProposalAll

The list-all lists all existing proposals in multisig module.

```
simd query multisig proposals list-all
```
Example output:

```
pagination:
  next_key: null
  total: "2"
proposal:
- finalTallyResult: null
  group: bridge15wm6yd4ktqsmr5gzccytsu3wdzamzl4gtst6r3d3jv5g53yk8m8qtl6mts
  id: "1"
  messages:
  - '@type': /core.multisig.MsgChangeGroup
    creator: bridge15wm6yd4ktqsmr5gzccytsu3wdzamzl4gtst6r3d3jv5g53yk8m8qtl6mts
    group: bridge15wm6yd4ktqsmr5gzccytsu3wdzamzl4gtst6r3d3jv5g53yk8m8qtl6mts
    members:
    - bridge1...
    - bridge1...
    threshold: "1"
  proposer: bridge1...
  status: SUBMITTED
  submitBlock: "1356"
  votingEndBlock: "1476"
- finalTallyResult: null
  group: bridge15wm6yd4ktqsmr5gzccytsu3wdzamzl4gtst6r3d3jv5g53yk8m8qtl6mts
  id: "2"
  messages:
  - '@type': /core.multisig.MsgChangeGroup
    creator: bridge15wm6yd4ktqsmr5gzccytsu3wdzamzl4gtst6r3d3jv5g53yk8m8qtl6mts
    group: bridge15wm6yd4ktqsmr5gzccytsu3wdzamzl4gtst6r3d3jv5g53yk8m8qtl6mts
    members:
    - bridge1...
    - bridge1...
    threshold: "1"
  proposer: bridge1...
  status: SUBMITTED
  submitBlock: "1357"
  votingEndBlock: "1477"
```

### Proposal
The proposal command lists info about proposal with given id.

```
simd query multisig proposals proposal 1
```
Example output:
```
proposal:
  finalTallyResult: null
  group: bridge15wm6yd4ktqsmr5gzccytsu3wdzamzl4gtst6r3d3jv5g53yk8m8qtl6mts
  id: "1"
  messages:
  - '@type': /core.multisig.MsgChangeGroup
    creator: bridge15wm6yd4ktqsmr5gzccytsu3wdzamzl4gtst6r3d3jv5g53yk8m8qtl6mts
    group: bridge15wm6yd4ktqsmr5gzccytsu3wdzamzl4gtst6r3d3jv5g53yk8m8qtl6mts
    members:
    - bridge1...
    - bridge1...
    threshold: "1"
  proposer: bridge1...
  status: SUBMITTED
  submitBlock: "1356"
  votingEndBlock: "1476"

```
----

### VotesAll
The  list-all command list all votes made by users to proposals in multisig module.

```
simd query multisig votes list-all
```
Example output:
```
pagination:
  next_key: null
  total: "2"
vote:
- option: "YES"
  proposalId: "1"
  submitBlock: "1467"
  voter: bridge1...
- option: "YES"
  proposalId: "3"
  submitBlock: "1511"
  voter: bridge1...

```
### Vote
The vote command list info about votes to given proposal id.

```
simd query multisig votes vote 3  
```
Example output:
```
pagination:
  next_key: null
  total: "1"
vote:
- option: "YES"
  proposalId: "3"
  submitBlock: "1511"
  voter: bridge1...
```
### VoteDetailed
The vote-detailed command gives info about vote with given proposal id and voter address.

```
simd query multisig votes vote-detailed 3 bridge1...
```

Example output:

```
vote:
  option: "YES"
  proposalId: "3"
  submitBlock: "1511"
  voter: bridge1...
```
----
## Transactions

### Create

The create command creates a multisig group with provided members and threshold.

```
simd tx multisig groups create bridge1... bridge1...,bridge1... 2 
```

### Update

The update command allows to update group changing the list of it`s members and threshold. (_NOTE: this command is executed only from the group account. You have to create proposal to update group_)

```
simd tx multisig groups update bridge1... bridge1...,bridge1... 1
```

### SubmitProposal

The submit proposal command allows to create multisig proposals.

```
simd tx multisig proposals submit-proposal /path/to/proposal.json --from mykey
```
Proposal to change group example:

```
{
  "creator": "bridge1...",
  "group": "bridge1...",
  "messages": [{
    "@type": "/core.multisig.MsgChangeGroup",
    "creator": "bridge1...", //creator and group must be equal
    "group": "bridge1...",
    "members": ["bridge1...",
      "bridge1...",
      "bridge1..."],
    "threshold": 3
  }]
}
```

### Vote

The vote command submits the vote to proposal with given id.

```
simd tx multisig votes vote bridge1... 1 0
```

_NOTE: 0 - Yes, 1 - No_