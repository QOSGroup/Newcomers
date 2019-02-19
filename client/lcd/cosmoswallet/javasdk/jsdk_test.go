package javasdk

import (
	"os/user"
	"testing"
)

func TestTransferAsync(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://localhost:26657"
	chainId := "test4matt"
	fromName := "local"
	password := "wm131421"
	toStr := "cosmos1mrf49r22adtd8juv6kvg8dxly32qlj7rg47644"
	coinStr := "1stake"
	feeStr := "1token"
	transout := TransferAsync(rootDir,node,chainId,fromName,password,toStr,coinStr,feeStr)
	t.Log(transout)
}

func TestQueryTx(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://localhost:26657"
	chainId := "test4matt"
	Txhash := "FE92F68247DB4435E4E6D1B66C7289F8C0654D297DAD8F1F6306556780C77EED"
	qout := QueryTx(rootDir, node, chainId, Txhash)
	t.Log(qout)
}