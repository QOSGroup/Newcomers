package sdksource

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"os"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/tendermint/tendermint/libs/log"
	tmliteProxy "github.com/tendermint/tendermint/lite/proxy"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
)


const ctxAccStoreName = "acc"

//NewCLIContext is used to init the config context without using Viper, the argues are all from the input of the func
func newCLIContext(rootDir,node,chainID string) context.CLIContext {
	var (
		rpc rpcclient.Client

	)

	//init the rpc instance
	nodeURI := node
	if nodeURI == "" {
		fmt.Printf("The nodeURI can not be nil for the rpc connection!")
	}
	rpc = rpcclient.NewHTTP(nodeURI, "/websocket")

	//create the verifier for the LCD verification
	var trustNode bool
	trustNode = false
	if trustNode {
		fmt.Printf("The default value for the trustNode is false!")
	}
	//chainID := ChainID
	//home := rootDir

	cacheSize := 10 // TODO: determine appropriate cache size
	verifier, err := tmliteProxy.NewVerifier(
		chainID, filepath.Join(rootDir, ".gaiacli", ".gaialite"),
		rpc, log.NewNopLogger(), cacheSize,
	)


	if err != nil {
		fmt.Printf("Create verifier failed: %s\n", err.Error())
		fmt.Printf("Please check network connection and verify the address of the node to connect to\n")
		os.Exit(1)
	}

	CliContext := context.CLIContext{
		Client:        rpc,
		Output:        os.Stdout,
		NodeURI:       nodeURI,
		AccountStore:  auth.StoreKey,
		Verifier:      verifier,

	}
	return CliContext

}

// TxBuilder implements a transaction context created in SDK modules.
//type TxBuilder struct {
//	txEncoder          sdk.TxEncoder
//	keybase            crkeys.Keybase
//	accountNumber      uint64
//	sequence           uint64
//	gas                uint64
//	gasAdjustment      float64
//	simulateAndExecute bool
//	chainID            string
//	memo               string
//	fees               sdk.Coins
//	gasPrices          sdk.DecCoins
//}

// NewTxBuilderFromCLI returns a new initialized TxBuilder with parameters input
//func newTxBuilderFromCLI(ChainID string) authtxb.TxBuilder {
//	txBldr := authtxb.TxBuilder{
//		chainID:            ChainID,
//	}
//	var txBldr authtxb.TxBuilder
//	return txBldr
//}
