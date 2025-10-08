package main

import (
	"log"

	fabric "github.com/hyperledger-labs/yui-fabric-ibc/relay/module"
	tendermint "github.com/hyperledger-labs/yui-relayer/chains/tendermint/module"
	"github.com/hyperledger-labs/yui-relayer/cmd"
)

func main() {
	if err := cmd.Execute(
		fabric.Module{},
		tendermint.Module{},
	); err != nil {
		log.Fatal(err)
	}
}
