package sdksource

import (
	"os/user"
	"testing"
)

func TestCreateSeed(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	output := CreateSeed(rootDir)
	t.Log(output)
}

func TestCreateAccount(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	name := "fp3"
	password := "wm131421"
	seed := ""
	output := CreateAccount(rootDir,name,password,seed)
	t.Log(output)
}

func TestRecoverKey(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	name := "test2213"
	password := "wm131421"
	seed := "math away vicious collect pole few reduce undo meadow dawn mesh nature pet guide valve behave hammer paddle number battle flavor two urge concert"
	output := RecoverKey(rootDir,name,password,seed)
	t.Log(output)
}

func TestUpdateKey(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	name := "test2"
	oldpass := "wm131422"
	newpass := "wm131421"
	output := UpdateKey(rootDir, name, oldpass, newpass)
	t.Log(output)
}