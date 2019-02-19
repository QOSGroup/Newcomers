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
//var keybase crkeys.Keybase
const (
	DenomName = "ATOM"
	defaultBIP39pass = "12345678"
	)

type KeyOutput struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Address string `json:"address"`
	PubKey  string `json:"pub_key"`
	Seed    string `json:"seed,omitempty"`
	Denom  string `json:"denom"`
}

type SeedOutput struct {
	Seed string `json:"seed"`
}

// To be depreacted! SetKeyBase initialized the LCD keybase. It also requires rootDir as input for the directory for key storing.
//func SetKeyBase(rootDir string) crkeys.Keybase {
//	var err error
//	keybase = nil
//	keybase, err = keys.NewKeyBaseFromDir(rootDir)
//	if err != nil {
//		fmt.Println(err)
//	}
//	return keybase
//}


func CreateSeed(rootDir string) string {
	//get the Keybase
	viper.Set(cli.HomeFlag, rootDir)
	kb, err1 := keys.NewKeyBaseFromHomeFlag()
	if err1 != nil {
		fmt.Println(err1)
	}
	// algo type defaults to secp256k1
	algo := crkeys.SigningAlgo("secp256k1")
	pass := defaultBIP39pass
	name := "inmemorykey"
	_, seed, _ := kb.CreateMnemonic(name, crkeys.English, pass, algo)
	//json output the seed
	var So SeedOutput
	So.Seed = seed
	respbyte, _ := json.Marshal(So)
	return string(respbyte)

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
	//SetKeyBase(rootDir)
	viper.Set(cli.HomeFlag, rootDir)
	kb, errz := keys.NewKeyBaseFromHomeFlag()
	if errz != nil {
		fmt.Println(errz)
	}

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
	infos, err := kb.List()
	for _, info := range infos {
		if info.GetName() == name {
			err = errKeyNameConflict(name)
			return err.Error()
		}
	}

	//create account
	if seed == "" {
		algo := crkeys.SigningAlgo("secp256k1")
		pass := defaultBIP39pass
		name := "inmemorykey"
		_, Seed, _ := kb.CreateMnemonic(name, crkeys.English, pass, algo)
		seed = Seed
	}


	info, err1 := kb.CreateAccount(name, seed, defaultBIP39pass, password, 0,0)
	if err1 != nil {
		return err1.Error()
	}

	keyOutput, err2 := keys.Bech32KeyOutput(info)
	if err2 != nil {
		return err2.Error()
	}

	keyOutput.Mnemonic = seed
	//add new field denom for the coin name
	var Ko KeyOutput
	Ko = KeyOutput{keyOutput.Name, keyOutput.Type, keyOutput.Address,keyOutput.PubKey,keyOutput.Mnemonic,DenomName}
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
	//SetKeyBase(rootDir)
	viper.Set(cli.HomeFlag, rootDir)
	kb, errz := keys.NewKeyBaseFromHomeFlag()
	if errz != nil {
		fmt.Println(errz)
	}

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
	info, err1 := kb.CreateAccount(name, seed, defaultBIP39pass, password, 0,0)
	if err1 != nil {
		return err1.Error()
	}

	keyOutput, err2 := keys.Bech32KeyOutput(info)
	if err2 != nil {
		return err2.Error()
	}

	keyOutput.Mnemonic = seed
	//add new field denom for the coin name
	var Ko KeyOutput
	Ko = KeyOutput{keyOutput.Name, keyOutput.Type, keyOutput.Address,keyOutput.PubKey,keyOutput.Mnemonic,DenomName}
	respbyte, _ := json.Marshal(Ko)
	return string(respbyte)
}

type UpdateKeyOutput struct {
	PasswordUpdate string `json:"pass_update"`

}

//for update the password of the name key stored in level db
func UpdateKey(rootDir, name, oldpass, newpass string) string {
	//SetKeyBase(rootDir)
	viper.Set(cli.HomeFlag, rootDir)
	kb, errz := keys.NewKeyBaseFromHomeFlag()
	if errz != nil {
		fmt.Println(errz)
	}
	getNewpass := func() (string, error) {
		return newpass, nil
	}

	err2 := kb.Update(name, oldpass, getNewpass)
	if err2 != nil {
		return err2.Error()
	}
	res := fmt.Sprintf("Password is successfully updated!")

	//json output the result
	var Po UpdateKeyOutput
	Po.PasswordUpdate = res
	respbyte, _ := json.Marshal(Po)
	return string(respbyte)

}