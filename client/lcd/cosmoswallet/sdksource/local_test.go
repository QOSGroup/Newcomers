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
	name := "hh"
	password := "wm131421"
	seed := ""
	output := CreateAccount(rootDir,name,password,seed)
	t.Log(output)
}

func TestRecoverKey(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	name := "test2212"
	password := "wm131421"
	seed := "reveal else motion oil dinner disease unveil taxi side volume shiver wheel chat similar flash that three series sign street hill motion silk chat"
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