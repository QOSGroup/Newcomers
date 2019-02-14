package sdksource

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/keys"
	crkeys "github.com/cosmos/cosmos-sdk/crypto/keys"

)
// keybase is used to make GetKeyBase a singleton
var keybase crkeys.Keybase
const DenomName = "ATOM"

type KeyOutput struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Address string `json:"address"`
	PubKey  string `json:"pub_key"`
	Seed    string `json:"seed,omitempty"`
	Denom  string `json:"denom"`
}

// SetKeyBase initialized the LCD keybase. It also requires rootDir as input for the directory for key storing.
func SetKeyBase(rootDir string) crkeys.Keybase {
	var err error
	keybase, err = keys.NewKeyBaseFromDir(rootDir)
	if err != nil {
		fmt.Println(err)
	}
	return keybase
}


func GetSeed(rootDir string) string {
	//initialize keybase
	SetKeyBase(rootDir)
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


func CreateAccount(rootDir, name, password, seed string) string {
	var (
		err  error
		info crkeys.Info
	)
	//initialize keybase
	SetKeyBase(rootDir)

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
		seed = GetSeed(rootDir)
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
	//add new field denom for the coin name
	var Ko KeyOutput
	Ko = KeyOutput{keyOutput.Name, keyOutput.Type, keyOutput.Address,keyOutput.PubKey,keyOutput.Seed,DenomName}
	respbyte, _ := json.Marshal(Ko)
	return string(respbyte)
}

//for recover key with name, password and seed input
func RecoverKey(rootDir,name,password,seed string) string {
	var (
		err  error
		info crkeys.Info
	)
	//initialize keybase
	SetKeyBase(rootDir)
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
	//add new field denom for the coin name
	var Ko KeyOutput
	Ko = KeyOutput{keyOutput.Name, keyOutput.Type, keyOutput.Address,keyOutput.PubKey,keyOutput.Seed,DenomName}
	respbyte, _ := json.Marshal(Ko)
	return string(respbyte)
}

//for update the password of the name key stored in level db
func UpdateKey(rootDir, name, oldpass, newpass string) string {
	SetKeyBase(rootDir)

	getNewpass := func() (string, error) {
		return newpass, nil
	}

	err2 := keybase.Update(name, oldpass, getNewpass)
	if err2 != nil {
		return err2.Error()
	}
	return fmt.Sprintf("Password is successfully updated!")

}