# Local Testnet Scrips

Under this folder many scripts are used to help with the local testnet interaction.

You should only use this scripts when local testnet is running.

## Transaction

Here are two snippets for the transaction. Copy and modify them as required to save effort.

Execute on docker:

```bash
# Note by default user keyring folder is not mounted on docker.
docker exec -it lorenzonode0 /lorenzod/lorenzod tx \
      <command> \
      --from node0 \
      --chain-id lorenzo_83291-1 \
      --gas 400000 \
      --gas-prices 10alrz \
      --keyring-backend test \
      --home "/data/node0/lorenzod" \
      --node "tcp://localhost:26657" \
      --yes
```

Execute on host machine:

```bash
lorenzod tx \
      <command> \
      --from <user> \
      --chain-id lorenzo_83291-1 \
      --gas 400000 \
      --gas-prices 10alrz \
      --keyring-backend test \
      --node "tcp://localhost:26657" \
      --yes
```