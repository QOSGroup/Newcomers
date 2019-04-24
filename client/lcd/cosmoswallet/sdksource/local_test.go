package sdksource

import (
	"os/user"
	"testing"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
	name := "cosmos341"
	password := "wm131421"
	seed := ""
	output := CreateAccount(rootDir,name,password,seed)
	t.Log(output)
}

func TestRecoverKey(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	name := "c34banker"
	password := "wm131421"
	seed := "wood render impose elegant gravity adapt buffalo during husband never stem text lesson public boring street interest sphere imitate margin lift rival invest nature"
	output := RecoverKey(rootDir,name,password,seed)
	t.Log(output)
}

func TestUpdateKey(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	name := "c33"
	oldpass := "wm131421"
	newpass := "wm131422"
	output := UpdateKey(rootDir, name, oldpass, newpass)
	t.Log(output)
}

func TestToken2Power(t *testing.T) {
	tokenInt := sdk.NewInt(int64(1000000))
	power := sdk.TokensToTendermintPower(tokenInt)
	t.Log(power)
}