package cosmoswallet

import "github.com/cosmos/cosmos-sdk/client/lcd/cosmoswallet/sdksource"


//get the seed(mnemonic) for the account generation
func GetSeed(rootDir string) string {
	output := sdksource.GetSeed(rootDir)
	return output
}

//create account
func CreateAccount(rootDir, name, password, seed string) string {
	output := sdksource.CreateAccount(rootDir,name, password, seed)
	return output
}


//recover key
func RecoverKey(rootDir, name,password,seed string) string {
	output := sdksource.RecoverKey(rootDir, name, password, seed)
	return output
}

//get account info
func GetAccount(rootDir,node,chainID,addr string) string {
	output := sdksource.GetAccount(rootDir,node,chainID,addr)
	return output
}


//transfer
func Transfer(rootDir,node,chainId,fromName,password,toStr,coinStr,feeStr string) string  {
	output := sdksource.Transfer(rootDir,node,chainId,fromName,password,toStr,coinStr,feeStr)
	return output
}