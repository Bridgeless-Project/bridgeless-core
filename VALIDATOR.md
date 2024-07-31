# Validators’ Guide

* Chain id: `bridge_2607-1`
* Bridgeless-core binary (linux/amd64): https://github.com/hyle-team/bridgeless-core/releases/tag/v12.1.7
* Genesis file: https://github.com/hyle-team/bridgeless-core/blob/chains/devnet/config/genesis.json

## Local installation

You can get binary file from github release and also, you can build core from sources by yourself:
use “https://github.com/hyle-team/bridgeless-core/releases/tag/v12.1.7” release information.

The binary is built under Alpine linux. If you are using Ubuntu linux, please install musl-dev
using `sudo apt install musl-dev` command to be able to use Alpine binary on your machine

To generate configs file download the core binary file and rename it to `bridgeless-core` if needed.

Set envs:

    export MONIKER_NAME=YOU_VALIDATOR_NAME
    export BRIDGELESS_HOME=YOU_CORE_HOME_PATH
    export BRIDGELESS_NODE=tcp://167.99.26.8:26657

Then to init node struct and generate configs u should execute this

    bridgeless-core init $MONIKER_NAME main --chain-id bridge_2607-1  --home=BRIDGELESS_HOME --keyring-backend test

Move the genesis file to `$BRIDGELESS_HOME/config`

Create validator private key:

    bridgeless-core keys add <key-name> --keyring-backend test --home=$BRIDGELESS_HOME

Dont forget to save the mnemonic and address. That address will be used for your validator staking.
Send you address (bridge…) to our team.

Please, backup the following files and folders:

    $BRIDGELESS_HOME/config/priv_validator_key.json
    $BRIDGELESS_HOME/config/node_key.json
    $BRIDGELESS_HOME/keyring-test

Get the node id:

    bridgeless-core tendermint show-node-id --home=$BRIDGELESS_HOME

The next step is a setting peers into config.toml. Find a `persistent_peers` field and set here at least your node
info(`node_id@node_ip:26656`) and one or two the others.
In our case, past should past `0b872fe5d809863a94fe0e666a334ff06cd8d2cf@167.99.26.8:26656`. It's good to have at least 3
peers.

To run use env variables:

    - name: DAEMON_NAME
      value: "bridgeless-core"

    - name: DAEMON_HOME
      value: $BRIDGELESS_HOME

    - name: DAEMON_ALLOW_DOWNLOAD_BINARIES
      value: "true"

To the simplest updates, run a chain using a `cosmovisor` tool.

    mkdir -p $DAEMON_HOME/cosmovisor/genesis/bin && cp YOU_STORE_CORE_BIN(name bridgeless-core) $DAEMON_HOME/cosmovisor/genesis/bin && cosmovisor run start --home=$BRIDGELESS_HOME --rpc.laddr tcp://0.0.0.0:26657


Note, that your node has several ports that is required to be omen on your machine:

    - "26656" # P2P port
    - "26657" and "9090" # Cosmos RPC
    - "1317" # REST Swagger API
    - "8545" # EVM HTTP RPC
    - "8546" # EVM WS RPC

Also, please note that `app.toml` file contains the configuration for the gas price accepted by your validator. 
    - `minimum-gas-prices = "1abridge"`

### Finishing with the validator

After receiving confirmation from ourside about token transfer, execute command to stake tokens
and become a validator. You need to stake exactly `100000000000000000000abridge` (10^18) that is equal to 1 BRIDGE.

    bridgeless-core tx staking create-validator --amount 100000000000000000000abridge --commission-max-change-rate "0.01" --commission-max-rate "0.2" --commission-rate "0.1" --min-self-delegation "1" --details "Meet new bridgeless validator" --pubkey $(bridgless-core tendermint show-validator --home=$BRIDGELESS_HOME) --moniker $MONIKER_NAME --chain-id bridge_2607-1 --fees 0bridge --from $LOCAL_VALIDATOR_ADDRESS --home=$BRIDGELESS_HOME --node=$BRIDGELESS_NODE --keyring-backend=test


## Docker compose

First of all init folders struct, it should be structured like this

```
BRIDGELESS_HOME
│   docker-compose.yaml
└───config
   └───validator
        └───config
            │  app.toml     
            │  genesis.json     
            │  client.toml
            │  node_key.json
            │  priv_validator_key.json
            │  config.toml

```

To generate folder struct execute the command:

    docker run --volume $BRIDGELESS_HOME/config/validator:/config/validator  ghcr.io/hyle-team/bridgeless-core:{last_tag}   init $MONIKER_NAME main--chain-id bridge_2607-1  --home=$BRIDGELESS_HOME --keyring-backend test

Modify docker-compose.yml:

    version: "3.7"

    services:
    
    validator:
    image: ghcr.io/hyle-team/bridgeless-core:{last_tag}
    entrypoint: sh -c "bridgeless-core start --home=$BRIDGELESS_HOME --rpc.laddr tcp://0.0.0.0:26657 --p2p.external-address tcp://validator:26656"
    volumes:
      - $BRIDGELESS_HOME/config:/config/validator/config
      - .$BRIDGELESS_HOME/data:/config/validator/data
      ports:
      - "26657:26657" # RPC
      - "1317:1317" # REST
      - "9090:9090" # Cosmos gRPC
      - "8545:8545" # EVM HTTP RPC
      - "8546:8546" # EVM WS RPC

Move the genesis file to `$BRIDGELESS_HOME/config`

Create validator private key:

    bridgeless-core keys add <key-name> --keyring-backend test --home=$BRIDGELESS_HOME

Dont forget to save the mnemonic and address. That address will be used for your validator staking.
Send you address (bridge…) to our team.

Please, backup the following files and folders:

    $BRIDGELESS_HOME/config/priv_validator_key.json
    $BRIDGELESS_HOME/config/node_key.json
    $BRIDGELESS_HOME/keyring-test

Get the node id:

    docker run ghcr.io/hyle-team/bridgeless-core:{last_tag} tendermint show-node-id --home=$BRIDGELESS_HOME

The next step is a setting peers into config.toml. Find a `persistent_peers` field and set here at least your node
info(`node_id@node_ip:26656`) and one or two the others.
In our case, past should past `0b872fe5d809863a94fe0e666a334ff06cd8d2cf@167.99.26.8:26656`. It's good to have at least 3
peers

Run a command `docker compose start` to launch the node.

### Finishing with the validator

After receiving confirmation from ourside about token transfer, execute command to stake tokens
and become a validator. You need to stake exactly `100000000000000000000abridge` (10^18) that is equal to 1 BRIDGE.

Find docker container id

    docker ps

Init validator

    docker exec -it <container-id> tx staking create-validator --amount 100000000000000000000abridge --commission-max-change-rate "0.01" --commission-max-rate "0.2" --commission-rate "0.1" --min-self-delegation "1" --details "Meet new bridgeless validator" --pubkey $(bridgless-core tendermint show-validator --home=$BRIDGELESS_HOME) --moniker $MONIKER_NAME --chain-id bridge_2607-1 --fees 0bridge --from $LOCAL_VALIDATOR_ADDRESS --home=$BRIDGELESS_HOME --node=$BRIDGELESS_NODE --keyring-backend=test
