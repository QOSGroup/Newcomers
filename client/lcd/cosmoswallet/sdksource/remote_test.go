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
	password := "qstars"
	toStr := "cosmos1fqr280v6x00uylwczh969vyfz4rfwsn5q9pvsg"
	coinStr := "1token"
	feeStr := "1token"
	transout := Transfer(rootDir,node,chainId,fromName,password,toStr,coinStr,feeStr)
	t.Log(transout)
}