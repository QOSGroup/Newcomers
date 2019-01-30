package sdksource

import (
	"os/user"
	"testing"
)

func TestGetAccount(t *testing.T) {
	addr := "cosmos1x0pssqzp4tqwf5vktts838em6el694hmmkm4nt"
	node := "tcp://localhost:26657"
	chainId := "test4matt"
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	acout := GetAccount(rootDir,node,chainId,addr)
	t.Log(acout)
}

func TestTransfer(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://localhost:26657"
	chainId := "test4matt"
	fromName := "cosmoslcd8"
	password := "lcdtest"
	toStr := "cosmos1mrf49r22adtd8juv6kvg8dxly32qlj7rg47644"
	coinStr := "10token"
	feeStr := "1token"
	transout := Transfer(rootDir,node,chainId,fromName,password,toStr,coinStr,feeStr)
	t.Log(transout)
}

func TestDelegate(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://localhost:26657"
	chainId := "test4matt"
	delegatorName := "cosmoslcd8"
	password := "lcdtest"
	delegatorAddr := "cosmos1x0pssqzp4tqwf5vktts838em6el694hmmkm4nt"
	validatorAddr := "cosmosvaloper1a8e4nvxw26c9ug9x687s65vxquszu3j82zezuc"
	delegationCoinStr := "20stake"
	feeStr := "1token"
	delout := Delegate(rootDir, node, chainId, delegatorName, password, delegatorAddr, validatorAddr, delegationCoinStr, feeStr)
	t.Log(delout)
}