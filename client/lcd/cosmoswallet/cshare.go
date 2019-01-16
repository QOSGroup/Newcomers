package cosmoswallet

import "github.com/cosmos/cosmos-sdk/client/lcd/cosmoswallet/sdksource"

//set KeyBase for the key store directory via input
func SetKeyBase(rootDir string) {
	sdksource.SetKeyBase(rootDir)
}

//get the seed(mnemonic) for the account generation
func GetSeed() string {
	output := sdksource.GetSeed()
	return output
}

//create account
func CreateAccount(name, password, seed string) string {
	output := sdksource.CreateAccount(name, password, seed)
	return output
}


//recover key
func RecoverKey(name,password,seed string) string {
	output := sdksource.RecoverKey(name, password, seed)
	return output
}

