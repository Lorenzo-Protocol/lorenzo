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

a real example:

```sh
# btc testnet
./build/lorenzod testnet init-files --base-btc-header '{"header": "0000202052d6336f639d03ec6f27638fd34e93d5ba4a971463a8142d0a00000000000000d1cbfdb66dd4131b77114812896e9e9f579e97afeea215fcf7563d77cc89e1103928f865434e2c193514ebdc","hash": "0000000000000023878a8e2ea4ab9d93a5cf7fb07d417dc8a899b9acb33e045f","height": "2582496","work": "96937883"}' --keyring-backend file
# btc mainnet
./build/lorenzod testnet init-files --base-btc-header '{"header": "000018228d263ff4070e2fdb31b654704e99f850a4f31762155003000000000000000000f79ccd11440a9d383b1438ee60692fc7b78a4b8800faf6ec61ec5fc37a078387f998f265595a03175799484b","hash": "000000000000000000032366f4bd696122c3e11096dfdacaf76b428b2a3f2318","height": "834624","work": "83947913181361"}' --keyring-backend file --btc-lightclient-params '{"insert_headers_allow_list": ["lrz1h2eglam9jjqlrjz3dask54mr8xfk47cxx0kh5c"]}' --btc-network mainnet
```

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

We provide support for running a multi-node testnet using Docker. To build and run it

```console
cp -R .testnets/* build/
docker build -t lorenzo/node .
docker compose up
```

## useful commands

init testnet config and genesis

```sh
./build/lorenzod testnet init-files --base-btc-header '{"header": "00000020fe9c2c8d8d1fd467a6e55f7ddaaf9b269d3f591d9539d637170000000000000061469aca436b8f615cb752379638479acfccd44761ceb4a025bca30a8b4e4d74625e086650e22619aa4406c7","hash": "00000000000000019816a254831e576d01c648dccfb333bab4792a49633fde0a","height": "2584512","work": "110454122"}' --btc-lightclient-params '{"insert_headers_allow_list": ["lrz1h2eglam9jjqlrjz3dask54mr8xfk47cxx0kh5c"]}' --keyring-backend file --chain-id lorenzo_26666-1 --btc-staking-params '{"btc_receiving_addr": "tb1q9zkpszu0m06n42jcfw0jhtad4yndsedn4uuye", "btc_confirmations_depth": 3, "burn_fee_factor": "1300000000"}'
```

send some token to a address

```sh
./build/lorenzod tx bank send --home build/node0/lorenzod --keyring-backend file --keyring-dir build/node0/lorenzod --chain-id lorenzo_26666-1 lrz1754tssy23f33rzscnda8xh2wzh738e4e56fywy lrz1h2eglam9jjqlrjz3dask54mr8xfk47cxx0kh5c 10000000alrz
```

add a key from mnemonic

```sh
./build/lorenzod keys add --keyring-backend file --keyring-dir build/node0/lorenzod reporter --recover
```

update btc fee rate

```sh
./build/lorenzod tx btclightclient update-fee-rate 1600 --keyring-backend file --keyring-dir build/node0/lorenzod/ --from lrz1h2eglam9jjqlrjz3dask54mr8xfk47cxx0kh5c --chain-id lorenzo_26666-1
```

query a btc transaction proof

```sh
bitcoin-cli -rpcconnect=<ip> -rpcport=8332 -rpcuser=<...> -rpcpassword=<...> gettxoutproof '["3ccfa1834c5794dc9db1a17a5416a4d1d05afcaa93d4a1621a6f7f1c607eb8e7"]'
```


btcstaking mint

```sh
./build/lorenzod tx btcstaking create-btcstaking-with-btc-proof --keyring-backend file --keyring-dir build/node0/lorenzod --chain-id lorenzo_26666-1 --from lrz1h2eglam9jjqlrjz3dask54mr8xfk47cxx0kh5c 020000000001013412acd95077452c02ac040f320fda6957bd119632ce7f7ca1725fb1f2d44ad20200000000fdffffff030000000000000000166a14bab28ff7659481f1c8516f616a576339936afb06e80300000000000016001428ac180b8fdbf53aaa584b9f2bafada926d865b30774040000000000225120099f46cbdfeeeff497b18f16757da503025da16d6bb3eb0108d48810c310bcf401402f86f1a4c9bc15ee7ed5a5a8c9612650ce233065374deecbe638f1e6a6237351ae3f26ee2674dbbce8240b65566e6a4a042bcc6105d937c995e25061829a22c400000000 0020002007b5a5af1f13f73fa5a7c0cef7c3eb83b6c4260cf0f20e362c3e5b00000000005b7a2c84f6191273ea8571a66c8708cf2714a8294ea60bc60824949e042b1c2c57301666ffff001d0dce81140d0300000b6ae9c2fe4da92665327a95bc635a4f6e38b8c47eb240f9631f42caac0233bfe46df84034cdd39a5099c3a232e4c3c5f4961735cdc2646949a1d6c5dc4f0c29f00b1229247983e492e9263fe785ee363eac04007f100eeedea02ae3b088ea5cade7b87e601c7f6f1a62a1d493aafc5ad0d1a416547aa1b19ddc94574c83a1cf3cae1402dcb2b6fd31bfa9b74f048ef952ac0d07fc23d8d9ded55f485e930e0a83f0628bf9fcc88a177465e01ab81ee0c6e9bdc2c9771290b61c02ecb7e008b09af2b211040f6c6945a98646516b53c7df7789a2672e200f279a26966739661444add035369786ddbc68c6ea03d0f25628c306a31df4d13dced107072e6ccab59392056b77b9f79daf30919e230caa8eb79c277a8d78054e649da7c748891700c5d3a8c99e678c17aa39f27922aafe45bd0cfa4290324d7e1e725aa8f226c21d5bc904c393025b812b97869859153c9cf54ee93cad32eb8aca8d437ec758d2b417035f2f00
```

