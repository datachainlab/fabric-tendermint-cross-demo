# fabric tendermint cross demo

This is example application for atomic swap through cross chain transaction using [Cross Framework](https://github.com/datachainlab/cross) between [Hyperledger Fabric](https://github.com/hyperledger/fabric) and [Tendermint](https://github.com/tendermint/tendermint). 

As Atomic commit protocol, [Simple commit protocol](https://datachainlab.github.io/cross-docs/architecture/atomic-commit-protocol/#simple-commit-protocol) is used.

## Requirements
- [Go](https://go.dev/) 1.16+
- [Docker](https://www.docker.com/products/docker-desktop)
- [jq](https://stedolan.github.io/jq/)

## Install required tools
The below command installs Relayer, Fabric tools, Fabric CLI, Tendermint CLI
```Makefile
# install fabric tools and build Fabric/Tendermint CLI, Relayer
make -C demo build
```

## Directory structure
- cmds
  - alpha ... Tendermint CLI
  - beta ... Fabric CLI
- contracts
  - erc20 ... contract modules for Tendermint/Fabric
- demo ... demo environment including 
  - chains
    - fabric ... fabric related files including docker-compose.yaml for network
      - chaincode
        - fabibc ... fabric chaincode application
    - tendermint ... tendermint simapp working as CLI and tendermint application including docker-compose.yaml for network
  - configs ... config files for alpha cli, beta cli, fabric app, relayer
  - scripts ... shell scripts for alpha cli, beta cli, relayer, snenario
- relayer ... Relayer CLI

### How tools work 
- Relayer 
  - relaying packets between different chains
- Fabric tools 
  - genesis block generation
  - configuration of a channel
  - cryptographic generation
  - deployment of chaincode
- Fabric CLI
  - creating tx, submitting tx, querying tx on Fabric node
- Tendermint CLI
  - creating tx, submitting tx, querying tx on Tendermint node

### ERC20 Contracts
See [README.md in contracts/erc20](https://github.com/datachainlab/fabric-tendermint-cross-demo/tree/main/contracts/erc20)

## Startup fabric and tendermint network
```Makefile
# prepare fabric network and tendermint network using docker containers
make -C demo network
```

## Testing the demo
```Makefile
# initialize relayer, fabric CLI, tendermint CLI, run handshake for IBC between fabric and tendermint by creating transactions.
make -C demo run-init
# run ./scripts/scenario/sample-scenario. See `About sample-scenario` section for more detail.
make -C demo run
```

## End network
```Makefile
# down related containers, remove volumes
make -C demo network-down
# remove any generated data
make -C demo clean
```

## Restart network if needed
```Makefile
make -C demo network-down
make -C demo clean
make -C demo network
```

## About sample-scenario
There are 3 actors, Alice, Bob, TokenOwner.  
Alice transfers 10 token to Bob on Tendermint. Then Bob transfers 10 token to Alice on Fabric.

See [scripts/scenario/sample-scenario](https://github.com/datachainlab/fabric-tendermint-cross-demo/blob/main/demo/scripts/scenario/sample-scenario)
