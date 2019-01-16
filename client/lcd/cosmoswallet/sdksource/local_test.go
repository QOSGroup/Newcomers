package sdksource

import (
	"fmt"
	"os/user"
	"testing"
)

func TestGetSeed(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	fmt.Println(rootDir)
	SetKeyBase(rootDir)
	output := GetSeed()
	t.Log(output)
}

func TestCreateAccount(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	fmt.Println(rootDir)
	SetKeyBase(rootDir)
	name := "cosmoslcd"
	password := "qstars"
	seed := "blue cash manage net peace diary system wine cool picture minimum earth parent gadget useful dose pear cycle legend buyer leopard spy giggle bamboo"
	output := CreateAccount(name,password,seed)
	t.Log(output)
}

func TestRecoverKey(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	fmt.Println(rootDir)
	SetKeyBase(rootDir)
	name := "cosmoslcd"
	password := "qstars"
	seed := "blue cash manage net peace diary system wine cool picture minimum earth parent gadget useful dose pear cycle legend buyer leopard spy giggle bamboo"
	output := RecoverKey(name,password,seed)
	t.Log(output)
}
