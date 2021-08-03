package token

import (
	"math/big"
)

func (o *Token) Onit_LoadContract(tokenAddr string) error {
	err := o.eth.Onit_LoadContract(tokenAddr)
	return err
}

func (o *Token) Onit_LoadContractInfo() error {
	var err error
	o.tokenName, err = o.eth.Onit_GetName()
	if err != nil {
		return err
	}
	o.tokenSymbol, err = o.eth.Onit_GetSymbol()
	if err != nil {
		return err
	}
	return err
}

func (o *Token) Onit_GetBalanceOf(walletAddr string) (int64, error) {
	bal, err := o.eth.Onit_GetBalanceOf(walletAddr)
	ne := big.NewInt(1000000000000000000)

	baseBal := big.NewInt(0)
	baseBal = baseBal.Div(bal, ne)

	return baseBal.Int64(), err
}
