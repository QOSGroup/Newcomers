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
	toStr := "cosmos1kklk4eqye6pla97dzmc03pw5lst7x0n4zt8syw"
	coinStr := "1000 stake"
	feeStr := "1stake"
	async := true
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
	delegationCoinStr := "300000stake"
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
	delegatorAddr := "cosmos1e5yhp5lkhjuautf4mhhll7l733za8tgpj329d9"
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
	feeStr := "1"
	async := false
	unbondDel := UnbondingDelegation(rootDir, node, chainId, delegatorName, password, delegatorAddr, validatorAddr, feeStr, async)
	t.Log(unbondDel)
}

func TestGetAllUnbondingDelegations(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "cosmosv33"
	delegatorAddr := "cosmos1e5yhp5lkhjuautf4mhhll7l733za8tgpj329d9"
	//validatorAddr := "cosmosvaloper1a8e4nvxw26c9ug9x687s65vxquszu3j82zezuc"
	getUbns := GetAllUnbondingDelegations(rootDir,node,chainId,delegatorAddr)
	t.Log(getUbns)
}

func TestGetBondValidators(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "cosmosv33"
	delegatorAddr := "cosmos1e5yhp5lkhjuautf4mhhll7l733za8tgpj329d9"
	getBd := GetBondValidators(rootDir,node,chainId,delegatorAddr)
	t.Log(getBd)
}

func TestGetAllValidators(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://47.105.142.9:26657"
	chainId := "cosmoshub-1"
	getVals := GetAllValidators(rootDir,node,chainId)
	t.Log(getVals)
}

func TestGetAllDelegations(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "cosmosv33"
	delegatorAddr := "cosmos1e5yhp5lkhjuautf4mhhll7l733za8tgpj329d9"
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
	delegatorAddr := "cosmos1e5yhp5lkhjuautf4mhhll7l733za8tgpj329d9"
	validatorAddr := "cosmosvaloper1yqlq9kg2txmwc606apvmen0ssag20dsfvnaq0v"
	getWithdraw := GetDelegationRewards(rootDir,node,chainId,delegatorAddr,validatorAddr)
	t.Log(getWithdraw)
}

func TestQueryTx(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://47.105.156.172:36657"
	chainId := "testing"
	txHash := "11E8BF5FF8EDEB33FEF736E81782903A1C129C48789AB4CF0C4D74B044C54375"
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

func TestGetDelegtorRewardsShares(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "cosmosv33"
	delegatorAddr := "cosmos1e5yhp5lkhjuautf4mhhll7l733za8tgpj329d9"
	daa := GetDelegtorRewardsShares(rootDir,node,chainId,delegatorAddr)
	t.Log(daa)
}

func TestTransferB4send(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://47.105.156.172:36657"
	chainId := "testing"
	fromName := "cosmos"
	password := "wm131421"
	toStr := "cosmos1kklk4eqye6pla97dzmc03pw5lst7x0n4zt8syw"
	coinStr := "100stake"
	feeStr := "1stake"
	async := true
	Tx := TransferB4send(rootDir, node, chainId, fromName, password, toStr, coinStr, feeStr,async)
	t.Log(Tx)

}

func TestBroadcastTransferTx(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://47.105.156.172:36657"
	chainId := "testing"
	fromName := "cosmos"
	password := "wm131421"
	toStr := "cosmos1kklk4eqye6pla97dzmc03pw5lst7x0n4zt8syw"
	coinStr := "10000stake"
	feeStr := "1stake"
	async := true
	txString := TransferB4send(rootDir, node, chainId, fromName, password, toStr, coinStr, feeStr,async)
	Bt:= BroadcastTransferTx(rootDir, node, chainId, txString, async)
	t.Log(Bt)
}