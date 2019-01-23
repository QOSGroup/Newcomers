package sdksource

import (
	"github.com/cosmos/cosmos-sdk/cmd/gaia/app"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var cdc = app.MakeCodec()

var (
	RootDir string
	Node	string
	ChainID string
)

func SetParas(rootDir,node,chainID string) {
	RootDir = rootDir
	Node = node
	ChainID	= chainID
}

func init() {
	var (
		r string
		n string
		c string
	)
	SetParas(r,n,c)
}


//get account from /auth/accounts/{address}
func GetAccount(addr string) string {
	key, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return err.Error()
	}

	cliCtx := NewCLIContext().
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

//
