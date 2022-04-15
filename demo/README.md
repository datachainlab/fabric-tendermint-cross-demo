# Demo
This directory is used for playground to run demo and several temporary directries are generated after demo executed by scripts.

## Directory structure
This directory consist of the following files/directories  
- chains
  - fabric ... fabric related files including docker-compose.yaml for network
    - chaincode
      - fabibc ... fabric chaincode application
  - tendermint ... tendermint simapp working as CLI and tendermint application including docker-compose.yaml for network
- configs ... config files for alpha cli, beta cli, fabric app, relayer
- scripts ... shell scripts for alpha cli, beta cli, relayer, snenario

## Setup Network
```
# install fabric tools and build Fabric/Tendermint CLI, Relayer
make build -j5

# down related containers, remove volumes
make network-down

# remove any generated data
make clean

# prepare fabric network and tendermint network using docker containers
make network
```

## Run sample scenario
```
# initialize relayer, fabric CLI, tendermint CLI, run handshake for IBC between fabric and tendermint by creating transactions.
make run-init

# run ./scripts/scenario/sample-scenario. See `About sample-scenario` section for more detail.
make run
```
