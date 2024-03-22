## Running a node

The following commands assume that the `lorenzod` executable has been
installed. If the repository was only built, then `./build/lorenzod` should be
used in its place.

### Generating the node configuration
The configuration for a single node can be created through the `testnet`
command. While the `testnet` command can create an arbitrary number of nodes that
communicate on a testnet, here we focus on the setup of a single node.
```console
lorenzod testnet init-files \
    --v                     1 \
    --output-dir            ./.testnet \
    --starting-ip-address   192.168.10.2 \
    --keyring-backend       test \
    --chain-id              chain-test \
    --btc-network string    testnet
```

The flags specify the following:
- `--output-dir <testnet-dir>`: Specifies that the testnet files should
  reside under this directory.
- `--v <N>`: Leads to the creation of `N` nodes, each one residing under the
  `<testnet-dir>/node{i}`. In this case `i={0..N-1}`.
- `--starting-ip-address <ip>`: Specifies the IP address for the nodes. For example,
  `192.168.10.2` leads to the first node running on `192.168.10.2:46656`, the
  second one on `192.168.10.3:46656` etc.
- `--keyring-backend {os,file,test}`: Specifies the backend to use for the keyring. Available
  choices include `os`, `file`, and `test`. We use `test` for convenience.
- `--chain-id`: An identifier for the chain. Useful when perrforming operations
  later.

In this case, we generated a single node. If we take a look under `.testnet`:
```console
$ ls .testnet
gentxs node0
```

The `gentxs` directory contains the genesis transactions. It contains
transactions that assign bbn tokens to a single address that is defined for each
node.

The `node0` directory contains the the following,
```console
$ ls .testnet/node0/lorenzod
config        data          key_seed.json keyring-test
```

A brief description of the contents:
- `config`: Contains the configuration files for the node.
- `data`: Contains the database storage for the node.
- `key_seed.json`: Seed to generate the keys maintained by the keyring.
- `keyring-test`: Contains the test keyring. This directory was created because
  we provided the `--keyring-backend test` flag. The `testnet` command, creates
  a validator node named `node{i}` (depends on the node name), and assigns
  bbn tokens to it through a transaction written to `.testnet/gentxs/node{i}.json`.
  The keys for this node can be pointed to by the `node{i}` name.

### Running the node
```console
lorenzod start --home ./.testnet/node0/lorenzod
```

### Logs

The logs for a particular node can be found under
`.testnets/node{id}/lorenzod/lorenzod.log`.

### Performing queries

After building a node and starting it, you can perform queries.
```console
lorenzod --home .testnet/node{i}/lorenzod/ --chain-id <chain-id> \
    query <module-name> <query-name>
```

For example, in order to get the hashes maintained by the `btcligthclient`
module:
```console
$ lorenzod --home .testnet/node0/lorenzod/ --chain-id chain-test query btclightclient hashes

hashes:
- 00000000000000000002bf1c218853bc920f41f74491e6c92c6bc6fdc881ab47
pagination:
  next_key: null
  total: "1"
```

### Submitting transactions

After building a node and running it, one can send transactions as follows:
```console
lorenzod --home .testnet/node{i}/lorenzod --chain-id <chain-id> \
         --keyring-backend {os,file,test} --fees <amount><denom> \
         --from <key-name> --broadcast-mode {sync,async,block} \
         tx <module-name> <tx-name> [data]
```

The `--fees` flag specifies the amount of fees that we are willing to pay and
the denomination and the `--from` flag denotes the name of the key that we want
to use to sign the transaction (i.e. from which account we want this
transaction to happen). The `--broadcast-mode` specifies how long we want to
wait until we receive a response from the CLI: `async` means immediately,
`sync` means after the transaction has been validated through `CheckTx`,
and `block` means after the transaction has been processed by the next block.

For example, in the `btclightclient` module, in order
to submit a header, one should:
```console
lorenzod --home .testnet/node0/lorenzod --chain-id chain-test \
         --keyring-backend test --fees 100bbn \
         --from node0 --broadcast-mode block \
         tx btclightclient insert-header <header-hex>
```

## Running a multi-node testnet

We provide support for running a multi-node testnet using Docker. To build it

```console
cp -R .testnets/* build/
docker compose up
```
