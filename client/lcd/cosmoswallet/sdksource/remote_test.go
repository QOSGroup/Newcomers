package sdksource

import (
	"fmt"
	"os/user"
	"testing"
)

func TestGetAccount(t *testing.T) {
	addr := "cosmos1a8e4nvxw26c9ug9x687s65vxquszu3j80kdhst"
	node := "tcp://localhost:26657"
	chainId := "test4matt"
	usr, _ := user.Current()
	rootDir := usr.HomeDir

	SetParas(rootDir,node,chainId)
	fmt.Println(Node)
	acout := GetAccount(addr)
	t.Log(acout)
}

//func TestInitConfig(t *testing.T) {
//	usr, _ := user.Current()
//	rootDir := usr.HomeDir
//	InitConfig(rootDir)
//}