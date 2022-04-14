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

## How to add ERC20 modules into simapp
### [simapp/app.go](https://github.com/datachainlab/fabric-tendermint-cross-demo/blob/main/demo/chains/tendermint/simapp/app.go)
- Create `BasicManager` with additional `AppModuleBasic{}`. [code](https://github.com/datachainlab/fabric-tendermint-cross-demo/blob/main/demo/chains/tendermint/simapp/app.go#L157-L159)
```go
ModuleBasics = module.NewBasicManager(
	...
	erc20.AppModuleBasic{},
	erc20mgr.AppModuleBasic{},
	erc20contract.AppModuleBasic{},
	```
)
```

- Add `Keeper`s into `SimApp` struct. [code](https://github.com/datachainlab/fabric-tendermint-cross-demo/blob/main/demo/chains/tendermint/simapp/app.go#L219-L221)
```go
type SimApp struct {
	...
	ERC20Keeper         erc20keeper.Keeper
	ERC20mgrKeeper      erc20mgrkeeper.Keeper
	ERC20contractKeeper erc20contractkeeper.Keeper
	...
}
```

- Add `StoreKeys` when creating `KVStoreKeys`. [code](https://github.com/datachainlab/fabric-tendermint-cross-demo/blob/main/demo/chains/tendermint/simapp/app.go#L272)
```go
keys := sdk.NewKVStoreKeys(
	...
	erc20types.StoreKey, erc20mgrtypes.StoreKey, erc20contracttypes.StoreKey,
)
```

- Create modules in `NewSimApp()`. [code](https://github.com/datachainlab/fabric-tendermint-cross-demo/blob/main/demo/chains/tendermint/simapp/app.go#L361-L377)
  - Note. [cross-cdt](https://github.com/datachainlab/cross-cdt) data type is used for `Store` in this demo. So some code depend on CDT.
```go
// Create a additional modules
// Create CDT Store
schema := cdttypes.NewSchema()
schemaERC20Prefix := []byte(erc20mgrtypes.ModuleName + "/")
schema.Set(schemaERC20Prefix, cdttypes.CDT_TYPE_INT64)
cdtStore := cdtkeeper.NewStore(appCodec, keys[crosstypes.StoreKey], schema)

// Create ERC20mgr module
erc20Int64Store := cdtStore.GetInt64Store(schemaERC20Prefix)
app.ERC20Keeper = erc20keeper.NewKeeper(erc20Int64Store)
erc20Module := erc20.NewAppModule(app.ERC20Keeper)
app.ERC20mgrKeeper = erc20mgrkeeper.NewKeeper(appCodec, app.ERC20Keeper, app.GetSubspace(erc20mgrtypes.ModuleName))
erc20mgrModule := erc20mgr.NewAppModule(app.ERC20mgrKeeper)

// Create erc20contract module
app.ERC20contractKeeper = erc20contractkeeper.NewKeeper(appCodec, keys[erc20contracttypes.StoreKey], app.ERC20mgrKeeper)
erc20contractModule := erc20contract.NewAppModule(app.ERC20contractKeeper)
```

- Setup a cross module using CDT. [code](https://github.com/datachainlab/fabric-tendermint-cross-demo/blob/main/demo/chains/tendermint/simapp/app.go#L379-L387)
```go
// Setup a cross module
app.XCCResolver = xcctypes.NewChannelInfoResolver(app.IBCKeeper.ChannelKeeper)
cmgr := contractkeeper.NewContractManager(
	appCodec,
	crosstypes.NewPrefixStoreKey(keys[crosstypes.StoreKey], crosstypes.ContractManagerPrefix),
	erc20contractModule,
	cdtStore,
	erc20mgrtypes.CDTContractHandleDecorators(),
)
```

- Create `Manager` with additional modules in `NewSimApp()`. [code](https://github.com/datachainlab/fabric-tendermint-cross-demo/blob/main/demo/chains/tendermint/simapp/app.go#L460-L462)
```go
	app.mm = module.NewManager(
		...
		erc20Module,
		erc20mgrModule,
		erc20contractModule,
	)
```

- Add `ModuleName` when calling `SetOrderInitGenesis` in `NewSimApp()`. [code](https://github.com/datachainlab/fabric-tendermint-cross-demo/blob/main/demo/chains/tendermint/simapp/app.go#L484)
```go
app.mm.SetOrderInitGenesis(
	...
	erc20types.ModuleName, erc20mgrtypes.ModuleName,
)
```

- Add `genesisState` in InitChainer(). [code](https://github.com/datachainlab/fabric-tendermint-cross-demo/blob/main/demo/chains/tendermint/simapp/app.go#L600-L610)
```
// erc20mgr module
erc20mgrGenesisState := erc20mgrtypes.DefaultGenesis()
erc20mgrGenesisState.Params = erc20mgrtypes.NewParams(
	ERC20Admin,
	false,
)
genesisState[erc20mgr.AppModuleBasic{}.Name()] = app.AppCodec().MustMarshalJSON(erc20mgrGenesisState)
// erc20 module in cross-cdt
genesisState[erc20.AppModuleBasic{}.Name()] = app.AppCodec().MustMarshalJSON(erc20types.DefaultGenesis())
// erc20cotract module
genesisState[erc20contract.AppModuleBasic{}.Name()] = app.AppCodec().MustMarshalJSON(erc20contracttypes.DefaultGenesis())
```

- Add `Subspace` if used. [code](https://github.com/datachainlab/fabric-tendermint-cross-demo/blob/main/demo/chains/tendermint/simapp/app.go#L762)
```go
paramsKeeper.Subspace(erc20mgrtypes.ModuleName)
```
