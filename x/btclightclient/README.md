# BTC light client

The BTC light client module is essentially a BTC light client that maintains
the canonical header chain of Bitcoin.

Babylon chain needs to know about different events happening on Bitcoin chain.
To make it possible in a secure way, Babylon needs to know the current
state of the Bitcoin chain i.e., what is the canonical chain of the Bitcoin network.

## Table of contents

- [Table of contents](#table-of-contents)
- [Concepts](#concepts)
  - [Problem statement](#problem-statement)
  - [Babylon's BTC light client design](#babylons-btc-light-client-design)
- [States](#states)
  - [Parameters](#parameters)
  - [Headers storage](#headers-storage)
  - [HashToHeight storage](#hashtoheight-storage)
- [Messages](#messages)
  - [MsgInsertHeaders](#msginsertheaders)
  - [MsgUpdateParams](#msgupdateparams)
- [Hooks](#hooks)
  - [Hooks exposed by BTC light client](#hooks-exposed-by-btc-light-client)
- [Events](#events)

## Concepts

### Problem statement

The Babylon chain needs to learn and validate a number of events that had
happened on the Bitcoin chain. Those events are:

1. `New checkpoint event` - Bitcoin Timestamping protocol requires checkpoints
on Bitcoin to be reported back to Babylon. To do it securely, each checkpoint
must be reported back along with the inclusion proof of transactions
which carry this checkpoint.
2. `New BTC delegation event` - Bitcoin Staking protocol requires staking
transactions to be deep enough in the Bitcoin chain. Thus, the staking
transactions also must be accompanied by the inclusion proof.

To properly validate those inclusion proofs, the Babylon chain needs to know the
current state of BTC chain i.e., what is current canonical chain recognized by BTC.

### Babylon's BTC light client design

Babylon maintains a BTC light client so that it can verify the inclusion
of various Bitcoin transactions.

In a high-level overview, the Babylon BTC light client starts from some base BTC
header existing on the BTC network and allows extending this header by applying the
same rules as a normal BTC node.

The base BTC header must:
- Be deep enough in the BTC chain so that it will never be reverted by the BTC network.
- Be at a height of BTC difficulty adjustment [boundary](https://en.bitcoin.it/wiki/Difficulty#How_often_does_the_network_difficulty_change.3F).
This is required, to properly validate all future difficulty adjustments.

The base BTC header is defined in the [genesis](../../proto/babylon/btclightclient/v1/genesis.proto) module.

The Babylon BTC light client module stores only BTC headers from the canonical
chain, and does not store the headers on the forks.
The BTC canonical chain can only be extended by processing
valid [MsgInsertHeaders](#msginsertheaders) messages.

If a better fork is encountered:
1. The current chain is rolled back to the parent of the received fork.
2. The chain is extend with new headers from the fork.

A better fork is defined as the fork with higher total difficulty, summing the
difficulties for each block in the fork.

## States

The BTC light client module maintains the following KV stores.

### Parameters

The [parameter storage](./keeper/params.go) maintains the BTC light client module's
parameters. The BTC light client module's parameters are represented as a `Params`
[object](../../proto/babylon/btclightclient/v1/params.proto) defined as follows:

```protobuf
// Params defines the parameters for the module.
message Params {
  option (gogoproto.equal) = true;

  // List of addresses which are allowed to insert headers to btc light client
  // if the list is empty, any address can insert headers
  repeated string insert_headers_allow_list = 1;
}
```

In a nutshell, `insert_headers_allow_list` makes it possible to set up
restrictions about who is able to update the BTC light client module state.
If `insert_headers_allow_list` is not empty, only addresses in the list can send
`MsgInsertHeaders` messages.

### Headers storage

The [Headers storage](./keeper/state.go) maintains all headers on the canonical
chain of Bitcoin.
The key is the header height in the BTC chain, and the value is a `BTCHeaderInfo`
[object](../../proto/babylon/btclightclient/v1/btclightclient.proto)
which contains the BTC header along with some metadata.

```protobuf
// BTCHeaderInfo is a structure that contains all relevant information about a
// BTC header
//  - Full header bytes
//  - Header hash for easy retrieval
//  - Height of the header in the BTC chain
//  - Total work spent on the header. This is the sum of the work corresponding
//  to the header Bits field and the total work of the header.
message BTCHeaderInfo {
  bytes header = 1
      [ (gogoproto.customtype) =
            "github.com/Lorenzo-Protocol/lorenzo/types.BTCHeaderBytes" ];
  bytes hash = 2
      [ (gogoproto.customtype) =
            "github.com/Lorenzo-Protocol/lorenzo/types.BTCHeaderHashBytes" ];
  uint64 height = 3;
  bytes work = 4
      [ (gogoproto.customtype) = "cosmossdk.io/math.Uint" ];
}
```

### HashToHeight storage

The [HashToHeight storage](./keeper/state.go) maintains an index in which the key is
the BTC header hash, and the value is the BTC header's height.

This index enables efficient lookup of BTC headers by their hash. This is useful
in many situations, notably when receiving a potential chain extension which
does not point to the current BTC chain tip.

## Messages

### MsgInsertHeaders

`MsgInsertHeaders` is the main message processed by the BTC light client module.
Its purpose is to update the state of the BTC chain as viewed by Babylon chain.

The handler of this message is defined
at [x/btclightclient/keeper/msg_server.go](./keeper/msg_server.go).

The message contains a list of BTC headers encoded in Bitcoin format.

```proto
message MsgInsertHeaders {
  option (cosmos.msg.v1.signer) = "signer";

  string signer = 1;
  repeated bytes headers = 2
      [ (gogoproto.customtype) =
            "github.com/Lorenzo-Protocol/lorenzo/types.BTCHeaderBytes" ];
}
```

Upon receiving a `MsgInsertHeaders` message, a Babylon node applies the following
verification rules. This is a subset of the BTC
[protocol](https://en.bitcoin.it/wiki/Protocol_rules#.22block.22_messages) rules:

- The `headers` list must not be empty.
- The headers in the list must be connected by parent-child relationships i.e.
the header at position `i + 1`, must have its `PrevBlock` field set to header's `i` hash.
- The first header of the list must point to a header already maintained by the BTC
light client module.
- Each header must be correctly encoded.
- Each header in the list must have valid proof of work and difficulty.
- Each header in the list must have a `Timestamp` that is greater than the median
of last 11 ancestors.
- If the first header of the list does not point to the current tip of the
chain maintained by the BTC light client, it means that the message contains a fork. For
the fork to be valid, the forked chain must be better than the current chain maintained by
the BTC light client. The fork is better when its total work is greater than the work
of current the [chain](https://en.bitcoin.it/wiki/Protocol_rules#Blocks).

All those rules are the same rules which are applied by BTC nodes when receiving
headers from the BTC network.

Processing of the message is atomic, so, if just one header in the list is
invalid, the state of the BTC light client module won't be updated.

In case of receiving a valid chain extension, the chain maintained by
the BTC light client module will be extended, and the received headers will be saved
in the storage.

In case of receiving a valid and better fork, i.e., the first header of the `headers` list
does not point to the current BTC chain tip, and the fork's total work is larger than
the current BTC chain's total work, the chain maintained by the BTC light client will
be rolled back to the header that is the fork's header, and it will then be
extended with the headers received in `headers`.

### MsgUpdateParams

The `MsgUpdateParams` message is used for updating the module parameters for the
BTC light client module. It can only be executed via a governance proposal.

```protobuf
// MsgUpdateParams defines a message for updating btc light client module parameters.
message MsgUpdateParams {
  option (cosmos.msg.v1.signer) = "authority";

  // authority is the address of the governance account.
  // just FYI: cosmos.AddressString marks that this field should use type alias
  // for AddressString instead of string, but the functionality is not yet implemented
  // in cosmos-proto
  string authority = 1 [(cosmos_proto.scalar) = "cosmos.AddressString"];

  // params defines the btc light client parameters to update.
  //
  // NOTE: All parameters must be supplied.
  Params params = 2 [(gogoproto.nullable) = false];
}
```

## Hooks

The BTC light client module exposes a set of hooks to inform other modules
about the updates to the maintained Bitcoin light client:

### Hooks exposed by BTC light client

```go
type BTCLightClientHooks interface {
	AfterBTCRollBack(ctx context.Context, headerInfo *BTCHeaderInfo)       // Must be called after the chain is rolled back
	AfterBTCRollForward(ctx context.Context, headerInfo *BTCHeaderInfo)    // Must be called after the chain is rolled forward
	AfterBTCHeaderInserted(ctx context.Context, headerInfo *BTCHeaderInfo) // Must be called after a header is inserted
}

```

## Events

The BTC light client module exposes a set of events about the updates to the
maintained Bitcoin best chain:

```protobuf

// The header included in the event is the block in the history
// of the current mainchain to which we are rolling back to.
// In other words, there is one rollback event emitted per re-org, to the
// greatest common ancestor of the old and the new fork.
message EventBTCRollBack { BTCHeaderInfo header = 1; }

// EventBTCRollForward is emitted on Msg/InsertHeader
// The header included in the event is the one the main chain is extended with.
// In the event of a reorg, each block on the new fork that comes after
// the greatest common ancestor will have a corresponding roll forward event.
message EventBTCRollForward { BTCHeaderInfo header = 1; }

// EventBTCHeaderInserted is emitted on Msg/InsertHeader
// The header included in the event is the one that was added to the
// on chain BTC storage.
message EventBTCHeaderInserted { BTCHeaderInfo header = 1; }

```

