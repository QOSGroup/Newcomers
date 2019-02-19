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
	name := "test21"
	password := "wm131421"
	seed := ""
	output := CreateAccount(rootDir,name,password,seed)
	t.Log(output)
}

func TestRecoverKey(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	name := "test2"
	password := "wm131421"
	seed := "fine vintage pottery fortune brick inherit tiny play child alter unfold region skin few duty false heavy bounce danger corn relax tomato describe audit"
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