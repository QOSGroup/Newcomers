package sdksource

import (
	"bytes"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/cmd/gaia/app"
	"github.com/cosmos/cosmos-sdk/codec"
	cskeys "github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankClient "github.com/cosmos/cosmos-sdk/x/bank/client"
	"github.com/cosmos/cosmos-sdk/x/stake"
	"github.com/cosmos/cosmos-sdk/x/stake/types"
	"os"
)

var cdc = app.MakeCodec()
var storeStake = "stake"
//get account from /auth/accounts/{address}
func GetAccount(rootDir,node,chainID,addr string) string {
	key, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return err.Error()
	}

	cliCtx := newCLIContext(rootDir,node,chainID).
		WithCodec(cdc).
		WithAccountDecoder(cdc)

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
	//build procedure
	SetKeyBase(rootDir)
	//fromName generated from keyspace locally
	if fromName == "" {
		fmt.Println("no fromName input!")
	}
	var info cskeys.Info
	var err error
		info, err = keybase.Get(fromName)
		if err != nil {
			fmt.Printf("could not find key %s\n", fromName)
			os.Exit(1)
		}

	fromAddr := info.GetAddress()
	cliCtx := newCLIContext(rootDir,node,chainID).
		WithCodec(cdc).
		WithAccountDecoder(cdc)
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
	msg := bankClient.CreateMsg(fromAddr, to, coins)

	//init a txBuilder for the transaction with fee
	txBldr := newTxBuilderFromCLI(chainID).WithTxEncoder(utils.GetTxEncoder(cdc)).WithFee(feeStr)

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
	//build procedure
	SetKeyBase(rootDir)
	//delegatorName generated from keyspace locally
	if delegatorName == "" {
		fmt.Println("no delegatorName input!")
	}
	info, err := keybase.Get(delegatorName)
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
		WithAccountDecoder(cdc)
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
	msg := stake.NewMsgDelegate(DelegatorAddr, ValidatorAddr, Delegation)
	err = msg.ValidateBasic()
	if err != nil {
		return err.Error()
	}

	//sign the stake message
	//init the txbldr
	txBldr := newTxBuilderFromCLI(chainID).WithTxEncoder(utils.GetTxEncoder(cdc)).WithFee(feeStr)

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
//get the delegation share under the specific validator
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

	//init a context for the rpc connection
	cliCtx := newCLIContext(rootDir,node,chainID).
		WithCodec(cdc).
		WithAccountDecoder(cdc)
	if err := cliCtx.EnsureAccountExistsFromAddr(DelAddr); err != nil {
		return err.Error()
	}

	// make a query to get the existing delegation shares
	key := stake.GetDelegationKey(DelAddr, ValAddr)
	res, err := cliCtx.QueryStore(key, storeStake)
	if err != nil {
		return err.Error()
	}

	// parse out the delegation
	delegation, err := types.UnmarshalDelegation(cdc, key, res)
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
	SetKeyBase(rootDir)
	//delegatorName generated from keyspace locally
	if delegatorName == "" {
		fmt.Println("no delegatorName input!")
	}
	info, err := keybase.Get(delegatorName)
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
		WithAccountDecoder(cdc)
	if err := cliCtx.EnsureAccountExistsFromAddr(DelegatorAddr); err != nil {
		return err.Error()
	}

	//validator to address type []byte
	ValidatorAddr, err := sdk.ValAddressFromBech32(validatorAddr)
	if err != nil {
		return err.Error()
	}

	// make a query to get the existing delegation shares
	key := stake.GetDelegationKey(DelegatorAddr, ValidatorAddr)
	res, err := cliCtx.QueryStore(key, storeStake)
	if err != nil {
		return err.Error()
	}

	// parse out the delegation
	delegation, err := types.UnmarshalDelegation(cdc, key, res)
	if err != nil {
		return err.Error()
	}

	//create the unbond message
	sharesAmount := delegation.Shares
	msg := stake.NewMsgBeginUnbonding(DelegatorAddr, ValidatorAddr, sharesAmount)

	//build-->sign-->broadcast
	//sign the stake message
	//init the txbldr
	txBldr := newTxBuilderFromCLI(chainID).WithTxEncoder(utils.GetTxEncoder(cdc)).WithFee(feeStr)

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
