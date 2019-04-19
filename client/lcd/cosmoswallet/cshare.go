package cosmoswallet

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/lcd/cosmoswallet/sdksource"
	"github.com/cosmos/cosmos-sdk/client/lcd/cosmoswallet/slim"
	"strings"
)


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
func Transfer(rootDir,node,chainId,fromName,password,toStr,coinStr,feeStr string, async bool) string  {
	output := sdksource.Transfer(rootDir,node,chainId,fromName,password,toStr,coinStr,feeStr,async)
	return output
}

//delegate
func Delegate(rootDir, node, chainID, delegatorName, password, delegatorAddr, validatorAddr, delegationCoinStr, feeStr string, async bool) string {
	output := sdksource.Delegate(rootDir, node, chainID, delegatorName, password, delegatorAddr, validatorAddr, delegationCoinStr, feeStr, async)
	return output
}

//get a specific delegation shares
func GetDelegationShares(rootDir, node, chainID, delegatorAddr, validatorAddr string) string {
	output := sdksource.GetDelegationShares(rootDir, node, chainID, delegatorAddr, validatorAddr)
	return output
}

//for unbond delegation shares from specific validator
func UnbondingDelegation(rootDir, node, chainID, delegatorName, password, delegatorAddr, validatorAddr, feeStr string, async bool) string {
	output := sdksource.UnbondingDelegation(rootDir, node, chainID, delegatorName, password, delegatorAddr, validatorAddr, feeStr, async)
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
func WithdrawDelegationReward(rootDir, node, chainID, delegatorName, password, delegatorAddr, validatorAddr, feeStr string, async bool) string {
	output := sdksource.WithdrawDelegationReward(rootDir, node, chainID, delegatorName, password, delegatorAddr, validatorAddr, feeStr, async)
	return output
}

//get a delegation reward between delegator and validator
func GetDelegationRewards(rootDir, node, chainID, delegatorAddr, validatorAddr string) string {
	output := sdksource.GetDelegationRewards(rootDir, node, chainID, delegatorAddr, validatorAddr)
	return output
}

//query the tx result by txHash generated via async broadcast
func QueryTx(rootDir,node,chainId,txHash string) string {
	output := sdksource.QueryTx(rootDir,node,chainId,txHash)
	return output
}

func GetValSelfBondShares (rootDir, node, chainID, validatorAddr string) string {
	output := sdksource.GetValSelfBondShares(rootDir, node, chainID, validatorAddr)
	return output
}

func GetDelegtorRewardsShares(rootDir,node,chainId,delegatorAddr string) string {
	output := sdksource.GetDelegtorRewardsShares(rootDir,node,chainId,delegatorAddr)
	return output
}

//QOS wallet part begin from here
func QOSAccountCreate(password string) string {
	output := slim.AccountCreateStr(password)
	return output
}

func QOSAccountCreateFromSeed(mncode string) string {
	output := slim.AccountCreateFromSeed(mncode)
	return output
}

//for QSCKVStoreset
func QSCKVStoreSet(k, v, privkey, chain string) string {
	output := slim.QSCKVStoreSetPost(k, v, privkey, chain)
	return output
}

//for QSCKVStoreGet
func QSCKVStoreGet(k string) string {
	output := slim.QSCKVStoreGetQuery(k)
	return output
}

//for QSCQueryAccount
func QSCQueryAccount(addr string) string {
	output := slim.QSCQueryAccountGet(addr)
	return output
}

//for QOSQueryAccount
func QOSQueryAccount(addr string) string {
	output := slim.QOSQueryAccountGet(addr)
	return output
}

//for AccountRecovery
func QOSAccountRecover(mncode, password string) string {
	output := slim.AccountRecoverStr(mncode, password)
	return output
}

//for IP input
func QOSSetBlockchainEntrance(sh, mh string) {
	slim.SetBlockchainEntrance(sh, mh)
}

//for PubAddrRetrieval
func QOSPubAddrRetrieval(priv string) string {
	//	fmt.Println("Please input host including IP and port for initialization on Qstar deamon:")
	output := slim.PubAddrRetrievalStr(priv)
	return output
}

//for QSCtransferSend
func QSCtransferSend(addrto, coinstr, privkey, chainid string) string {
	output := slim.QSCtransferSendStr(addrto, coinstr, privkey, chainid)
	return output
}

//for QOSCommitResultCheck
func QOSCommitResultCheck(txhash, height string) string {
	output := slim.QOSCommitResultCheck(txhash, height)
	return output
}

func QOSJQInvestAd(QOSchainId, QSCchainId, articleHash, coins, privatekey string) string {
	output := slim.JQInvestAd(QOSchainId, QSCchainId, articleHash, coins, privatekey)
	return output
}

func QOSAesEncrypt(key, plainText string) string {
	output := slim.AesEncrypt(key, plainText)
	return output
}

func QOSAesDecrypt(key, cipherText string) string {
	output := slim.AesDecrypt(key, cipherText)
	return output
}

func QOSTransferRecordsQuery(chainid, addr, cointype, offset, limit string) string {
	output := slim.TransferRecordsQuery(chainid, addr, cointype, offset, limit)
	return output
}

func TransferB4send(rootDir, node, chainID, fromName, password, toStr, coinStr, feeStr string, async bool) string {
	output := sdksource.TransferB4send(rootDir, node, chainID, fromName, password, toStr, coinStr, feeStr, async)
	return output
}

func BroadcastTransferTx(rootDir, node, chainID, txString string, async bool) string {
	output := sdksource.BroadcastTransferTx(rootDir, node, chainID, txString, async)
	return output
}

//for AdvertisersTrue
func AdvertisersTrue( privatekey,  coinsType, coinAmount,qscchainid string) string {
	output := slim.AdvertisersTrue( privatekey,  coinsType, coinAmount,qscchainid )
	return output
}

//for AdvertisersFalse
func AdvertisersFalse( privatekey,  coinsType, coinAmount,qscchainid string) string {
	output := slim.AdvertisersFalse( privatekey,  coinsType, coinAmount,qscchainid )
	return output
}

//for GetTx
func GetTx(tx string)string{
	output := slim.GetTx( tx )
	return output
}


func GetBlance(addrs string)string{
	path := fmt.Sprintf("/store/%s/%s", "aoeaccount", "key")
	output,_ := slim.Query(path,[]byte(addrs))
	return string (output)
}


func GetBlanceByCointype(addrs ,cointype string)string{
	result:=GetBlance(addrs)
    var qsc slim.QSCs
	json.Unmarshal([]byte(result),&qsc)

	for _,v:=range qsc{
		if strings.ToUpper(v.Name)==strings.ToUpper(cointype){
			return v.Amount.String()
		}
	}
	return "0"
}



