package sdksource

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/viper"
	"os"
	"path/filepath"

	"github.com/tendermint/tendermint/libs/log"
	tmliteProxy "github.com/tendermint/tendermint/lite/proxy"
	rpcclient "github.com/tendermint/tendermint/rpc/client"

	"github.com/cosmos/cosmos-sdk/client/context"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
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
		AccountStore:  ctxAccStoreName,
		Verifier:      verifier,

	}
	return CliContext

}

// NewTxBuilderFromCLI returns a new initialized TxBuilder with parameters input
func newTxBuilderFromCLI(chainID string) authtxb.TxBuilder {
	return authtxb.TxBuilder{
		ChainID:       chainID,
		AccountNumber: uint64(viper.GetInt64(client.FlagAccountNumber)),
		Gas:           client.GasFlagVar.Gas,
		GasAdjustment: viper.GetFloat64(client.FlagGasAdjustment),
		Sequence:      uint64(viper.GetInt64(client.FlagSequence)),
		SimulateGas:   client.GasFlagVar.Simulate,
		Fee:           viper.GetString(client.FlagFee),
		Memo:          viper.GetString(client.FlagMemo),
	}
}
