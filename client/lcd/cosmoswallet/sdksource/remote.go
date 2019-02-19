package sdksource

import (
	"bytes"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/cmd/gaia/app"
	"github.com/cosmos/cosmos-sdk/codec"
	cskeys "github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/cosmos/cosmos-sdk/x/bank"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distritypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/cli"
	"os"
)

var cdc = app.MakeCodec()
const (
	storeStake = "staking"
	//storeDistri = "distr"
	)
//get account from /auth/accounts/{address}
func GetAccount(rootDir,node,chainID,addr string) string {
	key, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return err.Error()
	}

	//to be fixed, the trust-node was set true to passby the verifier function, need improvement
	cliCtx := newCLIContext(rootDir,node,chainID).
		WithCodec(cdc).
		WithAccountDecoder(cdc).WithTrustNode(true)
	//cliCtx := context.NewCLIContext().
	//	WithCodec(cdc).WithAccountDecoder(cdc)

	if err = cliCtx.EnsureAccountExistsFromAddr(key); err != nil {
		return err.Error()
	}

	acc, err := cliCtx.GetAccount(key)
	if err != nil {
		return err.Error()
	}

	var output []byte
	if cliCtx.Indent {
		output, err = cdc.MarshalJSONIndent(acc, "", "  ")
	} else {
		output, err = cdc.MarshalJSON(acc)
	}
	if err != nil {
		return err.Error()
	}

	return string(output)

}

//complete the whole process with following sequence {Send coins (build -> sign -> send)}
func Transfer(rootDir, node, chainID, fromName, password, toStr, coinStr, feeStr string) string {
	//get the Keybase
	viper.Set(cli.HomeFlag, rootDir)
	kb, err1 := keys.NewKeyBaseFromHomeFlag()
	if err1 != nil {
		fmt.Println(err1)
	}
	//SetKeyBase(rootDir)
	//fromName generated from keyspace locally
	if fromName == "" {
		fmt.Println("no fromName input!")
	}
	var info cskeys.Info
	var err error
		info, err = kb.Get(fromName)
		if err != nil {
			fmt.Printf("could not find key %s\n", fromName)
			os.Exit(1)
		}

	fromAddr := info.GetAddress()
	cliCtx := newCLIContext(rootDir,node,chainID).
		WithCodec(cdc).
		WithAccountDecoder(cdc).WithTrustNode(true)
	if err := cliCtx.EnsureAccountExistsFromAddr(fromAddr); err != nil {
		return err.Error()
	}

	to, err := sdk.AccAddressFromBech32(toStr)
	if err != nil {
		return err.Error()
	}

	// parse coins trying to be sent
	coins, err := sdk.ParseCoins(coinStr)
	if err != nil {
		return err.Error()
	}

	account, err := cliCtx.GetAccount(fromAddr)
	if err != nil {
		return err.Error()
	}

	// ensure account has enough coins
	if !account.GetCoins().IsAllGTE(coins) {
		return fmt.Sprintf("Address %s doesn't have enough coins to pay for this transaction.", fromAddr)
	}

	// build and sign the transaction, then broadcast to Tendermint
	msg := bank.NewMsgSend(fromAddr, to, coins)

	//init a txBuilder for the transaction with fee
	txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc)).WithFees(feeStr).WithChainID(chainID)

	//accNum added to txBldr
	accNum, err := cliCtx.GetAccountNumber(fromAddr)
	if err != nil {
		return err.Error()
	}
	txBldr = txBldr.WithAccountNumber(accNum)

	//accSequence added
	accSeq, err := cliCtx.GetAccountSequence(fromAddr)
	if err != nil {
		return err.Error()
	}
	txBldr = txBldr.WithSequence(accSeq)


	// build and sign the transaction
	txBytes, err := txBldr.BuildAndSign(fromName, password, []sdk.Msg{msg})
	if err != nil {
		return err.Error()
	}
	// broadcast to a Tendermint node
	res, err := cliCtx.BroadcastTx(txBytes)
	if err != nil {
		return err.Error()
	}
	resbyte, err := cdc.MarshalJSON(res)
	if err != nil {
		return err.Error()
	}
	return string(resbyte)
}

//do Delegate operation
func Delegate(rootDir, node, chainID, delegatorName, password, delegatorAddr, validatorAddr, delegationCoinStr, feeStr string) string  {
	//get the Keybase
	viper.Set(cli.HomeFlag, rootDir)
	kb, err1 := keys.NewKeyBaseFromHomeFlag()
	if err1 != nil {
		fmt.Println(err1)
	}
	//delegatorName generated from keyspace locally
	if delegatorName == "" {
		fmt.Println("no delegatorName input!")
	}
	info, err := kb.Get(delegatorName)
	if err != nil {
		return err.Error()
	}
	//checkout with rule of own deligation
	DelegatorAddr, err := sdk.AccAddressFromBech32(delegatorAddr)
	if err != nil {
		return err.Error()
	}
	if !bytes.Equal(info.GetPubKey().Address(), DelegatorAddr) {
		return fmt.Sprintf("Must use own delegator address")
	}

	//init a context for this delegate tx
	cliCtx := newCLIContext(rootDir,node,chainID).
		WithCodec(cdc).
		WithAccountDecoder(cdc).WithTrustNode(true)
	if err := cliCtx.EnsureAccountExistsFromAddr(DelegatorAddr); err != nil {
		return err.Error()
	}

	//validator to address type []byte
	ValidatorAddr, err := sdk.ValAddressFromBech32(validatorAddr)
	if err != nil {
		return err.Error()
	}

	// parse coin from the delegation
	Delegation, err := sdk.ParseCoin(delegationCoinStr)
	if err != nil {
		return err.Error()
	}

	//check out the account enough money for the delegation
	account, err := cliCtx.GetAccount(DelegatorAddr)
	if err != nil {
		return err.Error()
	}

	DelegationToS := sdk.Coins{Delegation}
	if !account.GetCoins().IsAllGTE(DelegationToS) {
		return fmt.Sprintf("Delegator address %s doesn't have enough coins to perform this transaction.", delegatorAddr)
	}

	//build the stake message
	msg := staking.NewMsgDelegate(DelegatorAddr, ValidatorAddr, Delegation)
	err = msg.ValidateBasic()
	if err != nil {
		return err.Error()
	}

	//sign the stake message
	//init the txbldr
	txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc)).WithFees(feeStr).WithChainID(chainID)

	//accNum added to txBldr
	accNum, err := cliCtx.GetAccountNumber(DelegatorAddr)
	if err != nil {
		return err.Error()
	}
	txBldr = txBldr.WithAccountNumber(accNum)

	//accSequence added
	accSeq, err := cliCtx.GetAccountSequence(DelegatorAddr)
	if err != nil {
		return err.Error()
	}
	txBldr = txBldr.WithSequence(accSeq)

	// build and sign the transaction
	txBytes, err := txBldr.BuildAndSign(delegatorName, password, []sdk.Msg{msg})
	if err != nil {
		return err.Error()
	}
	// broadcast to a Tendermint node
	res, err := cliCtx.BroadcastTx(txBytes)
	if err != nil {
		return err.Error()
	}
	resbyte, err := cdc.MarshalJSON(res)
	if err != nil {
		return err.Error()
	}
	return string(resbyte)

}
//get the delegation share under a specific validator
func GetDelegationShares(rootDir, node, chainID, delegatorAddr, validatorAddr string) string {
	//convert the delegator string address to sdk form
	DelAddr, err := sdk.AccAddressFromBech32(delegatorAddr)
	if err != nil {
		return err.Error()
	}

	//convert the validator string address to sdk form
	ValAddr, err := sdk.ValAddressFromBech32(validatorAddr)
	if err != nil {
		return err.Error()
	}

	//to be fixed, the trust-node was set true to passby the verifier function, need improvement
	cliCtx := newCLIContext(rootDir,node,chainID).
		WithCodec(cdc).
		WithAccountDecoder(cdc).WithTrustNode(true)
	if err := cliCtx.EnsureAccountExistsFromAddr(DelAddr); err != nil {
		return err.Error()
	}

	// make a query to get the existing delegation shares
	key := staking.GetDelegationKey(DelAddr, ValAddr)
	res, err := cliCtx.QueryStore(key, storeStake)
	if err != nil {
		return err.Error()
	}

	// parse out the delegation
	delegation, err := types.UnmarshalDelegation(cdc, res)
	if err != nil {
		return err.Error()
	}

	//json output the result
	output, err := codec.MarshalJSONIndent(cdc, delegation)
	if err != nil {
		return err.Error()
	}

	return string(output)

}


//for unbond delegation shares from specific validator
func UnbondingDelegation(rootDir, node, chainID, delegatorName, password, delegatorAddr, validatorAddr, feeStr string) string {
	//build procedure
	//get the Keybase
	viper.Set(cli.HomeFlag, rootDir)
	kb, err1 := keys.NewKeyBaseFromHomeFlag()
	if err1 != nil {
		fmt.Println(err1)
	}

	//delegatorName generated from keyspace locally
	if delegatorName == "" {
		fmt.Println("no delegatorName input!")
	}
	info, err := kb.Get(delegatorName)
	if err != nil {
		return err.Error()
	}
	//checkout with rule of own deligation
	DelegatorAddr, err := sdk.AccAddressFromBech32(delegatorAddr)
	if err != nil {
		return err.Error()
	}
	if !bytes.Equal(info.GetPubKey().Address(), DelegatorAddr) {
		return fmt.Sprintf("Must use own delegator address")
	}

	////to be fixed, the trust-node was set true to passby the verifier function, need improvement
	cliCtx := newCLIContext(rootDir,node,chainID).
		WithCodec(cdc).
		WithAccountDecoder(cdc).WithTrustNode(true)
	if err := cliCtx.EnsureAccountExistsFromAddr(DelegatorAddr); err != nil {
		return err.Error()
	}

	//validator to address type []byte
	ValidatorAddr, err := sdk.ValAddressFromBech32(validatorAddr)
	if err != nil {
		return err.Error()
	}

	// make a query to get the existing delegation shares
	key := staking.GetDelegationKey(DelegatorAddr, ValidatorAddr)
	res, err := cliCtx.QueryStore(key, storeStake)
	if err != nil {
		return err.Error()
	}

	// parse out the delegation
	delegation, err := types.UnmarshalDelegation(cdc, res)
	if err != nil {
		return err.Error()
	}

	//create the unbond message
	sharesAmount := delegation.Shares
	msg := staking.NewMsgUndelegate(DelegatorAddr, ValidatorAddr, sharesAmount)

	//build-->sign-->broadcast
	//sign the stake message
	//init the txbldr
	txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc)).WithFees(feeStr).WithChainID(chainID)

	//accNum added to txBldr
	accNum, err := cliCtx.GetAccountNumber(DelegatorAddr)
	if err != nil {
		return err.Error()
	}
	txBldr = txBldr.WithAccountNumber(accNum)

	//accSequence added
	accSeq, err := cliCtx.GetAccountSequence(DelegatorAddr)
	if err != nil {
		return err.Error()
	}
	txBldr = txBldr.WithSequence(accSeq)

	// build and sign the transaction
	txBytes, err := txBldr.BuildAndSign(delegatorName, password, []sdk.Msg{msg})
	if err != nil {
		return err.Error()
	}
	// broadcast to a Tendermint node
	resb, err := cliCtx.BroadcastTx(txBytes)
	if err != nil {
		return err.Error()
	}
	resbyte, err := cdc.MarshalJSON(resb)
	if err != nil {
		return err.Error()
	}
	return string(resbyte)

}

//get all unbonding delegations from a specific delegator
func GetAllUnbondingDelegations (rootDir, node, chainID, delegatorAddr string) string {
	//convert the delegator string address to sdk form
	DelAddr, err := sdk.AccAddressFromBech32(delegatorAddr)
	if err != nil {
		return err.Error()
	}


	//to be fixed, the trust-node was set true to passby the verifier function, need improvement
	cliCtx := newCLIContext(rootDir,node,chainID).
		WithCodec(cdc).WithTrustNode(true)

	resKVs, err := cliCtx.QuerySubspace(staking.GetUBDsKey(DelAddr), storeStake)
	if err != nil {
		return err.Error()
	}

	var ubds staking.UnbondingDelegations
	for _, kv := range resKVs {
		ubds = append(ubds, types.MustUnmarshalUBD(cdc, kv.Value))
	}

	//json output the result
	output, err := codec.MarshalJSONIndent(cdc, ubds)
	if err != nil {
		return err.Error()
	}
	return string(output)
}

//Get bonded validators
func GetBondValidators(rootDir, node, chainID, delegatorAddr string) string {
	//convert the delegator string address to sdk form
	DelAddr, err := sdk.AccAddressFromBech32(delegatorAddr)
	if err != nil {
		return err.Error()
	}

	//generate paras for next query
	params := staking.NewQueryDelegatorParams(DelAddr)
	bz, err := cdc.MarshalJSON(params)
	if err != nil {
		return err.Error()
	}

	//to be fixed, the trust-node was set true to passby the verifier function, need improvement
	cliCtx := newCLIContext(rootDir,node,chainID).
		WithCodec(cdc).WithTrustNode(true)

	//query with data
	res, err := cliCtx.QueryWithData("custom/staking/delegatorValidators", bz)
	if err != nil {
		return err.Error()
	}

	return string(res)
}

//get all the validators
func GetAllValidators(rootDir, node, chainID string) string {
	key := staking.ValidatorsKey
	//to be fixed, the trust-node was set true to passby the verifier function, need improvement
	cliCtx := newCLIContext(rootDir,node,chainID).
		WithCodec(cdc).WithTrustNode(true)

	resKVs, err := cliCtx.QuerySubspace(key, storeStake)
	if err != nil {
		return err.Error()
	}

	// parse out the validators
	var validators staking.Validators
	for _, kv := range resKVs {
		validators = append(validators, types.MustUnmarshalValidator(cdc, kv.Value))
	}

	output, err := codec.MarshalJSONIndent(cdc, validators)
	if err != nil {
		return err.Error()
	}
	return string(output)
}

//get all delegations from the delegator
func GetAllDelegations(rootDir, node, chainID, delegatorAddr string) string {
	//convert the delegator string address to sdk form
	DelAddr, err := sdk.AccAddressFromBech32(delegatorAddr)
	if err != nil {
		return err.Error()
	}

	key := staking.GetDelegationsKey(DelAddr)
	//to be fixed, the trust-node was set true to passby the verifier function, need improvement
	cliCtx := newCLIContext(rootDir,node,chainID).
		WithCodec(cdc).WithTrustNode(true)

	resKVs, err := cliCtx.QuerySubspace(key, storeStake)
	if err != nil {
		return err.Error()
	}

	// parse out the delegations
	var delegations staking.Delegations
	for _, kv := range resKVs {
		delegations = append(delegations, types.MustUnmarshalDelegation(cdc, kv.Value))
	}

	output, err := codec.MarshalJSONIndent(cdc, delegations)
	if err != nil {
		return err.Error()
	}

	return string(output)
}

//Withdraw rewards from a specific validator
func WithdrawDelegationReward(rootDir, node, chainID, delegatorName, password, delegatorAddr, validatorAddr, feeStr string) string {
	//build procedure
	//get the Keybase
	viper.Set(cli.HomeFlag, rootDir)
	kb, err1 := keys.NewKeyBaseFromHomeFlag()
	if err1 != nil {
		fmt.Println(err1)
	}

	//delegatorName generated from keyspace locally
	if delegatorName == "" {
		fmt.Println("no delegatorName input!")
	}
	info, err := kb.Get(delegatorName)
	if err != nil {
		return err.Error()
	}
	//checkout with rule of own deligation
	DelegatorAddr, err := sdk.AccAddressFromBech32(delegatorAddr)
	if err != nil {
		return err.Error()
	}
	if !bytes.Equal(info.GetPubKey().Address(), DelegatorAddr) {
		return fmt.Sprintf("Must use own delegator address")
	}

	////to be fixed, the trust-node was set true to passby the verifier function, need improvement
	cliCtx := newCLIContext(rootDir,node,chainID).
		WithCodec(cdc).
		WithAccountDecoder(cdc).WithTrustNode(true)
	if err := cliCtx.EnsureAccountExistsFromAddr(DelegatorAddr); err != nil {
		return err.Error()
	}

	//validator to address type []byte
	ValidatorAddr, err := sdk.ValAddressFromBech32(validatorAddr)
	if err != nil {
		return err.Error()
	}

	//generate messages betweeb delegator and validator
	msgs := []sdk.Msg{distritypes.NewMsgWithdrawDelegatorReward(DelegatorAddr, ValidatorAddr)}

	//build-->sign-->broadcast
	//sign the stake message
	//init the txbldr
	txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc)).WithFees(feeStr).WithChainID(chainID)

	//accNum added to txBldr
	accNum, err := cliCtx.GetAccountNumber(DelegatorAddr)
	if err != nil {
		return err.Error()
	}
	txBldr = txBldr.WithAccountNumber(accNum)

	//accSequence added
	accSeq, err := cliCtx.GetAccountSequence(DelegatorAddr)
	if err != nil {
		return err.Error()
	}
	txBldr = txBldr.WithSequence(accSeq)

	// build and sign the transaction
	txBytes, err := txBldr.BuildAndSign(delegatorName, password, msgs)
	if err != nil {
		return err.Error()
	}
	// broadcast to a Tendermint node
	resb, err := cliCtx.BroadcastTx(txBytes)
	if err != nil {
		return err.Error()
	}
	resbyte, err := cdc.MarshalJSON(resb)
	if err != nil {
		return err.Error()
	}
	return string(resbyte)

}

//get a delegation reward between delegator and validator
func GetDelegationRewards(rootDir, node, chainID, delegatorAddr, validatorAddr string) string {
	//convert the delegator string address to sdk form
	DelAddr, err := sdk.AccAddressFromBech32(delegatorAddr)
	if err != nil {
		return err.Error()
	}

	//convert the validator string address to sdk form
	ValAddr, err := sdk.ValAddressFromBech32(validatorAddr)
	if err != nil {
		return err.Error()
	}

	//to be fixed, the trust-node was set true to passby the verifier function, need improvement
	cliCtx := newCLIContext(rootDir,node,chainID).
		WithCodec(cdc).
		WithAccountDecoder(cdc).WithTrustNode(true)
	if err := cliCtx.EnsureAccountExistsFromAddr(DelAddr); err != nil {
		return err.Error()
	}

	//query the delegation rewards
	resp, err := cliCtx.QueryWithData("custom/distr/delegation_rewards", cdc.MustMarshalJSON(distr.NewQueryDelegationRewardsParams(DelAddr, ValAddr)))
	if err != nil {
		return err.Error()
	}

	var result sdk.DecCoins
	cdc.MustUnmarshalJSON(resp, &result)

	resbyte, err := cdc.MarshalJSON(result)
	if err != nil {
		return err.Error()
	}
	return string(resbyte)
}