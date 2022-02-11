package CGWallet

import (
	"encoding/json"
	"fmt"

	"github.com/foxnut/go-hdwallet"
)

func getBTC(master *hdwallet.Key) (address string, err error) {
	wallet, _ := master.GetWallet(hdwallet.CoinType(hdwallet.BTC), hdwallet.Path("m/84'/0'/0'/0/0"))
	addressP2WPKH, _ := wallet.GetKey().AddressP2WPKH()
	kk, err := wallet.GetKey().PrivateWIF(true)
	if err != nil {
		return "", err
	}
	newWallet := &Wallet{Address: addressP2WPKH, PrivateKey: kk}
	wf, err := json.Marshal(newWallet)
	if err != nil {
		fmt.Println(err)
		return
	}
	return string(wf), nil
}
