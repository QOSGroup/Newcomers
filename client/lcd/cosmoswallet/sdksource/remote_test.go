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
	toStr := "cosmos1uzsu64daw0javd5ntg7ee64pj7frfx46xs7684"
	coinStr := "1000000stake"
	feeStr := "1token"
	async := false
	transout := Transfer(rootDir,node,chainId,fromName,password,toStr,coinStr,feeStr, async)
	t.Log(transout)
}

func TestDelegate(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://localhost:26657"
	chainId := "test4matt"
	delegatorName := "testunbond"
	password := "wm131421"
	delegatorAddr := "cosmos1ecnlyjrd4wr724wczxsjvgrkcq4f2sudq08cun"
	validatorAddr := "cosmosvaloper1a8e4nvxw26c9ug9x687s65vxquszu3j82zezuc"
	delegationCoinStr := "100stake"
	feeStr := "2token"
	async := false
	delout := Delegate(rootDir, node, chainId, delegatorName, password, delegatorAddr, validatorAddr, delegationCoinStr, feeStr, async)
	t.Log(delout)
}

func TestGetDelegationShares(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://localhost:26657"
	chainId := "test4matt"
	delegatorAddr := "cosmos1a8e4nvxw26c9ug9x687s65vxquszu3j80kdhst"
	validatorAddr := "cosmosvaloper1a8e4nvxw26c9ug9x687s65vxquszu3j82zezuc"
	getDelout := GetDelegationShares(rootDir,node,chainId,delegatorAddr,validatorAddr)
	t.Log(getDelout)
}

func TestUnbondingDelegation(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://localhost:26657"
	chainId := "test4matt"
	delegatorName := "testunbond"
	password := "wm131421"
	delegatorAddr := "cosmos1ecnlyjrd4wr724wczxsjvgrkcq4f2sudq08cun"
	validatorAddr := "cosmosvaloper1a8e4nvxw26c9ug9x687s65vxquszu3j82zezuc"
	feeStr := "1token"
	async := false
	unbondDel := UnbondingDelegation(rootDir, node, chainId, delegatorName, password, delegatorAddr, validatorAddr, feeStr, async)
	t.Log(unbondDel)
}

func TestGetAllUnbondingDelegations(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://localhost:26657"
	chainId := "test4matt"
	delegatorAddr := "cosmos1u4mxu4u52qgdxn56kuaew4jr6ewr2g2yxe2vcx"
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
	delegatorAddr := "cosmos1u4mxu4u52qgdxn56kuaew4jr6ewr2g2yxe2vcx"
	getDels := GetAllDelegations(rootDir,node,chainId,delegatorAddr)
	t.Log(getDels)
}

func TestWithdrawDelegationReward(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://localhost:26657"
	chainId := "test4matt"
	delegatorName := "testunbond"
	password := "wm131421"
	delegatorAddr := "cosmos1ecnlyjrd4wr724wczxsjvgrkcq4f2sudq08cun"
	validatorAddr := "cosmosvaloper1a8e4nvxw26c9ug9x687s65vxquszu3j82zezuc"
	feeStr := "1token"
	async := false
	withdrawRew := WithdrawDelegationReward(rootDir, node, chainId, delegatorName, password, delegatorAddr, validatorAddr, feeStr, async)
	t.Log(withdrawRew)
}

func TestGetDelegationRewards(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://localhost:26657"
	chainId := "test4matt"
	delegatorAddr := "cosmos1ecnlyjrd4wr724wczxsjvgrkcq4f2sudq08cun"
	validatorAddr := "cosmosvaloper1a8e4nvxw26c9ug9x687s65vxquszu3j82zezuc"
	getWithdraw := GetDelegationRewards(rootDir,node,chainId,delegatorAddr,validatorAddr)
	t.Log(getWithdraw)
}

func TestQueryTx(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://localhost:26657"
	chainId := "test4matt"
	txHash := "0BA029449967228DB14E7ECCFF9B97C5963807DB07D32CF180CBD545BBE59CFC"
	qTx := QueryTx(rootDir,node,chainId,txHash)
	t.Log(qTx)
}

func TestGetValSelfBondShares(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://localhost:26657"
	chainId := "test4matt"
	validatorAddr := "cosmosvaloper1a8e4nvxw26c9ug9x687s65vxquszu3j82zezuc"
	vsb := GetValSelfBondShares(rootDir, node, chainId, validatorAddr)
	t.Log(vsb)
}