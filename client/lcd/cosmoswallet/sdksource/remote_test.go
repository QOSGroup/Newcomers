package sdksource

import (
	"os/user"
	"testing"
)

func TestGetAccount(t *testing.T) {
	addr := "cosmos1eet7mg4v8u3lew8vwrtmwpptstn25ysj43q6a6"
	node := "tcp://localhost:26657"
	chainId := "test4matt"
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	acout := GetAccount(rootDir,node,chainId,addr)
	t.Log(acout)
}

func TestTransfer(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://localhost:26657"
	chainId := "test4matt"
	fromName := "local"
	password := "wm131421"
	toStr := "cosmos1nrds9u3kwlsltvk0scayjzq5s6025f25r8l3sf"
	coinStr := "10stake"
	feeStr := "1token"
	transout := Transfer(rootDir,node,chainId,fromName,password,toStr,coinStr,feeStr)
	t.Log(transout)
}

func TestDelegate(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://localhost:26657"
	chainId := "test4matt"
	delegatorName := "local"
	password := "wm131421"
	delegatorAddr := "cosmos1eet7mg4v8u3lew8vwrtmwpptstn25ysj43q6a6"
	validatorAddr := "cosmosvaloper1a8e4nvxw26c9ug9x687s65vxquszu3j82zezuc"
	delegationCoinStr := "20000stake"
	feeStr := "1token"
	delout := Delegate(rootDir, node, chainId, delegatorName, password, delegatorAddr, validatorAddr, delegationCoinStr, feeStr)
	t.Log(delout)
}

func TestGetDelegationShares(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://localhost:26657"
	chainId := "test4matt"
	delegatorAddr := "cosmos1eet7mg4v8u3lew8vwrtmwpptstn25ysj43q6a6"
	validatorAddr := "cosmosvaloper1a8e4nvxw26c9ug9x687s65vxquszu3j82zezuc"
	getDelout := GetDelegationShares(rootDir,node,chainId,delegatorAddr,validatorAddr)
	t.Log(getDelout)
}

func TestUnbondingDelegation(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://localhost:26657"
	chainId := "test4matt"
	delegatorName := "local"
	password := "wm131421"
	delegatorAddr := "cosmos1eet7mg4v8u3lew8vwrtmwpptstn25ysj43q6a6"
	validatorAddr := "cosmosvaloper1a8e4nvxw26c9ug9x687s65vxquszu3j82zezuc"
	feeStr := "1token"
	unbondDel := UnbondingDelegation(rootDir, node, chainId, delegatorName, password, delegatorAddr, validatorAddr, feeStr)
	t.Log(unbondDel)
}

func TestGetAllUnbondingDelegations(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://localhost:26657"
	chainId := "test4matt"
	delegatorAddr := "cosmos1eet7mg4v8u3lew8vwrtmwpptstn25ysj43q6a6"
	//validatorAddr := "cosmosvaloper1a8e4nvxw26c9ug9x687s65vxquszu3j82zezuc"
	getUbns := GetAllUnbondingDelegations(rootDir,node,chainId,delegatorAddr)
	t.Log(getUbns)
}

func TestGetBondValidators(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://localhost:26657"
	chainId := "test4matt"
	delegatorAddr := "cosmos1eet7mg4v8u3lew8vwrtmwpptstn25ysj43q6a6"
	getBd := GetBondValidators(rootDir,node,chainId,delegatorAddr)
	t.Log(getBd)
}

func TestGetAllValidators(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://localhost:26657"
	chainId := "test4matt"
	getVals := GetAllValidators(rootDir,node,chainId)
	t.Log(getVals)
}

func TestGetAllDelegations(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://localhost:26657"
	chainId := "test4matt"
	delegatorAddr := "cosmos1eet7mg4v8u3lew8vwrtmwpptstn25ysj43q6a6"
	getDels := GetAllDelegations(rootDir,node,chainId,delegatorAddr)
	t.Log(getDels)
}

func TestWithdrawDelegationReward(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://localhost:26657"
	chainId := "test4matt"
	delegatorName := "local"
	password := "wm131421"
	delegatorAddr := "cosmos1eet7mg4v8u3lew8vwrtmwpptstn25ysj43q6a6"
	validatorAddr := "cosmosvaloper1a8e4nvxw26c9ug9x687s65vxquszu3j82zezuc"
	feeStr := "1token"
	withdrawRew := WithdrawDelegationReward(rootDir, node, chainId, delegatorName, password, delegatorAddr, validatorAddr, feeStr)
	t.Log(withdrawRew)
}

func TestGetDelegationRewards(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://localhost:26657"
	chainId := "test4matt"
	delegatorAddr := "cosmos1eet7mg4v8u3lew8vwrtmwpptstn25ysj43q6a6"
	validatorAddr := "cosmosvaloper1a8e4nvxw26c9ug9x687s65vxquszu3j82zezuc"
	getWithdraw := GetDelegationRewards(rootDir,node,chainId,delegatorAddr,validatorAddr)
	t.Log(getWithdraw)
}