package cosmoswallet

import "github.com/cosmos/cosmos-sdk/client/lcd/cosmoswallet/sdksource"


//create the seed(mnemonic) for the account generation
func CreateSeed(rootDir string) string {
	output := sdksource.CreateSeed(rootDir)
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

//update password
func UpdateKey(rootDir, name, oldpass, newpass string) string {
	output := sdksource.UpdateKey(rootDir, name, oldpass, newpass)
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

//delegate
func Delegate(rootDir, node, chainID, delegatorName, password, delegatorAddr, validatorAddr, delegationCoinStr, feeStr string) string {
	output := sdksource.Delegate(rootDir, node, chainID, delegatorName, password, delegatorAddr, validatorAddr, delegationCoinStr, feeStr)
	return output
}

//get a specific delegation shares
func GetDelegationShares(rootDir, node, chainID, delegatorAddr, validatorAddr string) string {
	output := sdksource.GetDelegationShares(rootDir, node, chainID, delegatorAddr, validatorAddr)
	return output
}

//for unbond delegation shares from specific validator
func UnbondingDelegation(rootDir, node, chainID, delegatorName, password, delegatorAddr, validatorAddr, feeStr string) string {
	output := sdksource.UnbondingDelegation(rootDir, node, chainID, delegatorName, password, delegatorAddr, validatorAddr, feeStr)
	return output
}

//get all unbonding delegations from a specific delegator
func GetAllUnbondingDelegations (rootDir, node, chainID, delegatorAddr string) string {
	output := sdksource.GetAllUnbondingDelegations(rootDir, node, chainID, delegatorAddr)
	return output
}

//Get bonded validators
func GetBondValidators(rootDir, node, chainID, delegatorAddr string) string {
	output := sdksource.GetBondValidators(rootDir, node, chainID, delegatorAddr)
	return output
}

//get all the validators
func GetAllValidators(rootDir, node, chainID string) string {
	output := sdksource.GetAllValidators(rootDir, node, chainID)
	return output
}

//get all delegations from the delegator
func GetAllDelegations(rootDir, node, chainID, delegatorAddr string) string {
	output := sdksource.GetAllDelegations(rootDir, node, chainID, delegatorAddr)
	return output
}

//Withdraw rewards from a specific validator
func WithdrawDelegationReward(rootDir, node, chainID, delegatorName, password, delegatorAddr, validatorAddr, feeStr string) string {
	output := sdksource.WithdrawDelegationReward(rootDir, node, chainID, delegatorName, password, delegatorAddr, validatorAddr, feeStr)
	return output
}

//get a delegation reward between delegator and validator
func GetDelegationRewards(rootDir, node, chainID, delegatorAddr, validatorAddr string) string {
	output := sdksource.GetDelegationRewards(rootDir, node, chainID, delegatorAddr, validatorAddr)
	return output
}