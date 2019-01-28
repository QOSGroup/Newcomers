package sdksource

import (
	"os/user"
	"testing"
)

func TestGetSeed(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	output := GetSeed(rootDir)
	t.Log(output)
}

func TestCreateAccount(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	name := "transfercall"
	password := "qstars"
	seed := "call clip before orchard unveil salute rebel pipe armor dignity topic term rural lady noodle emerge power kite hint curve embody manage tribe pulse"
	output := CreateAccount(rootDir,name,password,seed)
	t.Log(output)
}

func TestRecoverKey(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	name := "cosmoslcd8"
	password := "qstars"
	seed := "blue cash manage net peace diary system wine cool picture minimum earth parent gadget useful dose pear cycle legend buyer leopard spy giggle bamboo"
	output := RecoverKey(rootDir,name,password,seed)
	t.Log(output)
}
