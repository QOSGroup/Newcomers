package sdksource

import (
	"os/user"
	"testing"
)

func TestGetAccount(t *testing.T) {
	addr := "cosmos1fqnc6mzrz7slm74cmh2mcxngllqujvm7y2lzfj"
	node := "tcp://192.168.1.184:26657"
	chainId := "cosmosv33"
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	acout := GetAccount(rootDir,node,chainId,addr)
	t.Log(acout)
}

func TestTransfer(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "cosmosv33"
	fromName := "c33banker"
	password := "wm131421"
	toStr := "cosmos1e5yhp5lkhjuautf4mhhll7l733za8tgpj329d9"
	coinStr := "10000000stake"
	feeStr := "1stake"
	async := false
	transout := Transfer(rootDir,node,chainId,fromName,password,toStr,coinStr,feeStr, async)
	t.Log(transout)
}

func TestDelegate(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "cosmosv33"
	delegatorName := "c33test"
	password := "wm131421"
	delegatorAddr := "cosmos1fqnc6mzrz7slm74cmh2mcxngllqujvm7y2lzfj"
	validatorAddr := "cosmosvaloper1yqlq9kg2txmwc606apvmen0ssag20dsfvnaq0v"
	delegationCoinStr := "3000000stake"
	feeStr := "1stake"
	async := true
	delout := Delegate(rootDir, node, chainId, delegatorName, password, delegatorAddr, validatorAddr, delegationCoinStr, feeStr, async)
	t.Log(delout)
}

func TestGetDelegationShares(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "cosmosv33"
	delegatorAddr := "cosmos1fqnc6mzrz7slm74cmh2mcxngllqujvm7y2lzfj"
	validatorAddr := "cosmosvaloper1yqlq9kg2txmwc606apvmen0ssag20dsfvnaq0v"
	getDelout := GetDelegationShares(rootDir,node,chainId,delegatorAddr,validatorAddr)
	t.Log(getDelout)
}

func TestUnbondingDelegation(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "cosmosv33"
	delegatorName := "c33test"
	password := "wm131421"
	delegatorAddr := "cosmos1fqnc6mzrz7slm74cmh2mcxngllqujvm7y2lzfj"
	validatorAddr := "cosmosvaloper1yqlq9kg2txmwc606apvmen0ssag20dsfvnaq0v"
	feeStr := "1stake"
	async := false
	unbondDel := UnbondingDelegation(rootDir, node, chainId, delegatorName, password, delegatorAddr, validatorAddr, feeStr, async)
	t.Log(unbondDel)
}

func TestGetAllUnbondingDelegations(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "cosmosv33"
	delegatorAddr := "cosmos1fqnc6mzrz7slm74cmh2mcxngllqujvm7y2lzfj"
	//validatorAddr := "cosmosvaloper1a8e4nvxw26c9ug9x687s65vxquszu3j82zezuc"
	getUbns := GetAllUnbondingDelegations(rootDir,node,chainId,delegatorAddr)
	t.Log(getUbns)
}

func TestGetBondValidators(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "cosmosv33"
	delegatorAddr := "cosmos1fqnc6mzrz7slm74cmh2mcxngllqujvm7y2lzfj"
	getBd := GetBondValidators(rootDir,node,chainId,delegatorAddr)
	t.Log(getBd)
}

func TestGetAllValidators(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "cosmosv33"
	getVals := GetAllValidators(rootDir,node,chainId)
	t.Log(getVals)
}

func TestGetAllDelegations(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "cosmosv33"
	delegatorAddr := "cosmos1fqnc6mzrz7slm74cmh2mcxngllqujvm7y2lzfj"
	getDels := GetAllDelegations(rootDir,node,chainId,delegatorAddr)
	t.Log(getDels)
}

func TestWithdrawDelegationReward(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "cosmosv33"
	delegatorName := "c33test"
	password := "wm131421"
	delegatorAddr := "cosmos1fqnc6mzrz7slm74cmh2mcxngllqujvm7y2lzfj"
	validatorAddr := "cosmosvaloper1yqlq9kg2txmwc606apvmen0ssag20dsfvnaq0v"
	feeStr := "1stake"
	async := false
	withdrawRew := WithdrawDelegationReward(rootDir, node, chainId, delegatorName, password, delegatorAddr, validatorAddr, feeStr, async)
	t.Log(withdrawRew)
}

func TestGetDelegationRewards(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "cosmosv33"
	delegatorAddr := "cosmos1fqnc6mzrz7slm74cmh2mcxngllqujvm7y2lzfj"
	validatorAddr := "cosmosvaloper1yqlq9kg2txmwc606apvmen0ssag20dsfvnaq0v"
	getWithdraw := GetDelegationRewards(rootDir,node,chainId,delegatorAddr,validatorAddr)
	t.Log(getWithdraw)
}

func TestQueryTx(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "cosmosv33"
	txHash := "F1CB0364201B9414829353366D1855EF9F6E94E61BAEC74AF775EA8172C5F39F"
	qTx := QueryTx(rootDir,node,chainId,txHash)
	t.Log(qTx)
}

func TestGetValSelfBondShares(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "cosmosv33"
	validatorAddr := "cosmosvaloper1yqlq9kg2txmwc606apvmen0ssag20dsfvnaq0v"
	vsb := GetValSelfBondShares(rootDir, node, chainId, validatorAddr)
	t.Log(vsb)
}