# Lorenzo

[![Website](https://badgen.net/badge/icon/website?label=)](http://www.lorenzo-protocol.xyz/)

## Build and install

The lorenzod application based on the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk) is the main application of the Lorenzo network. 
This repository is used to build the Lorenzo core application to join the Lorenzo network.

### Requirements

To build and install, you need to have Go 1.19 available.
Follow the instructions on the [Golang page](https://go.dev/doc/install) to do that.

To build the binary:

```console
make build
```

The binary will then be available at `./build/lorenzod` .

To install:

```console
make install
```

## Contributing

The [docs](./docs) directory contains the necessary information on how to get started using the lorenzod executable for development purposes.


