package sdksource

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/keys"
	crkeys "github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/cli"
)
// keybase is used to make GetKeyBase a singleton
var keybase crkeys.Keybase

type KeyOutput struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Address string `json:"address"`
	PubKey  string `json:"pub_key"`
	Seed    string `json:"seed,omitempty"`
}

// SetKeyBase initialized the LCD keybase. It also requires rootDir as input for the directory for key storing.
func SetKeyBase(rootDir string) crkeys.Keybase {
	viper.Set(cli.HomeFlag, rootDir)

	var err error
	keybase = nil
	keybase, err = keys.GetKeyBaseWithWritePerm()
	if err != nil {
		fmt.Println(err)
	}
	return keybase
}


func GetSeed() string {
	// algo type defaults to secp256k1
	algo := crkeys.SigningAlgo("secp256k1")
	pass := "throwing-this-key-away"
	name := "inmemorykey"
	_, seed, _ := keybase.CreateMnemonic(name, crkeys.English, pass, algo)
	return seed

}

//errors on account creation
func errKeyNameConflict(name string) error {
	return fmt.Errorf("acount with name %s already exists", name)
}

func errMissingName() error {
	return fmt.Errorf("you have to specify a name for the locally stored account")
}

func errMissingPassword() error {
	return fmt.Errorf("you have to specify a password for the locally stored account")
}

func errMissingSeed() error {
	return fmt.Errorf("you have to specify seed for key recover")
}


func CreateAccount(name, password, seed string) string {
	var (
		err  error
		info crkeys.Info
	)
	//check out the input
	if name == "" {
		err = errMissingName()
		return err.Error()
	}
	if password == "" {
		err = errMissingPassword()
		return err.Error()
	}
	// check if already exists
	infos, err := keybase.List()
	for _, info := range infos {
		if info.GetName() == name {
			err = errKeyNameConflict(name)
			return err.Error()
		}
	}

	//create account
	if seed == "" {
		seed = GetSeed()
	}


	info, err1 := keybase.CreateKey(name, seed, password)
	if err1 != nil {
		return err1.Error()
	}

	keyOutput, err2 := keys.Bech32KeyOutput(info)
	if err2 != nil {
		return err2.Error()
	}

	keyOutput.Seed = seed

	respbyte, _ := json.Marshal(keyOutput)
	return string(respbyte)
}

func RecoverKey(name,password,seed string) string {
	var (
		err  error
		info crkeys.Info
	)
	if name == "" {
		err = errMissingName()
		return err.Error()
	}
	if password == "" {
		err = errMissingPassword()
		return err.Error()
	}
	if seed == "" {
		err = errMissingSeed()
		return err.Error()
	}
	if err != nil {
		return err.Error()
	}
	info, err1 := keybase.CreateKey(name, seed, password)
	if err1 != nil {
		return err1.Error()
	}

	keyOutput, err2 := keys.Bech32KeyOutput(info)
	if err2 != nil {
		return err2.Error()
	}

	keyOutput.Seed = seed

	respbyte, _ := json.Marshal(keyOutput)
	return string(respbyte)
}