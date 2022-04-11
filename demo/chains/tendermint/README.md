# Tendermint Application
This is packages used by go implementations of Tendermint CLI which contains not only Tendermint application `simd` but also various functionalities like key generations.

## Directory structure
This directory consist of the following files/directories  
- scripts
  - build-go-embedded ... build go packages
  - entrypoint.sh ... entrypoint for tendermint container
  - tm-chain ... generating keys 
- simapp ... packages for `simd`
- docker-compose.yaml ... tendermint network.
- Dockerfile ... Tendermint application `simd` is deployed as Docker container 
