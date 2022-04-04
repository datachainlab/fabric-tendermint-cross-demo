# ERC20 Contract

Each of modules is in `modules` directory.  
These contract modules are developed based of [Cosmos SDK](https://docs.cosmos.network/v0.44/). 
Note that these modules depend on version `v0.43.0-beta1` for now. See [go.mod](https://github.com/datachainlab/fabric-tendermint-cross-demo/blob/main/contracts/erc20/go.mod#L7).

Each of modules needs to meet specific [Interfaces](https://github.com/datachainlab/fabric-tendermint-cross-demo/blob/main/contracts/erc20/modules/erc20contract/module.go#L26-L28).
```go
_ module.AppModule             = AppModule{}
_ module.AppModuleBasic        = AppModuleBasic{}
_ contracttypes.ContractModule = AppModule{}
```
- [module.AppModule](https://github.com/cosmos/cosmos-sdk/blob/v0.43.0-beta1/types/module/module.go#L156-L183)
- [module.AppModuleBasic](https://github.com/cosmos/cosmos-sdk/blob/v0.43.0-beta1/types/module/module.go#L47-L60)
- [contracttypes.ContractModule](https://github.com/datachainlab/cross/blob/v0.2.2/x/core/contract/types/types.go#L13-L15)

See section [Introduction to SDK Modules](https://docs.cosmos.network/v0.44/building-modules/intro.html) to know Cosmos SDK Modules in detail.

## erc20mgr module
This module includes `mint`, `burn`, `transfer`, `approve`, `transferFrom` functionalities according to [EIP-20: Token Standard](https://eips.ethereum.org/EIPS/eip-20).  
See [HandleContractCall function](https://github.com/datachainlab/fabric-tendermint-cross-demo/blob/main/contracts/erc20/modules/erc20mgr/keeper/keeper.go#L32-L57).

1. Mint
  - mint token for given account by signer account
2. Burn
  - burn token for given account by signer account
3. Transfer
  - transfer token from signer account to recipient account
4. Approve
  - allows spender account to withdraw from signer account up to the value amount
5. TransferFrom
  - transfer token from owner account to recipient account by spender account 

## erc20contract module
This module includes `transfer` functionality called from cross framework for cross-chain. 
See [HandleContractCall function](https://github.com/datachainlab/fabric-tendermint-cross-demo/blob/main/contracts/erc20/modules/erc20contract/keeper/keeper.go#L29-L54).  

1. Transfer
  - transfer token from signer account to recipient account
