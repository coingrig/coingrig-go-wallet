package CGWallet

import (
	"encoding/json"
	"fmt"

	"github.com/foxnut/go-hdwallet"
)

func getDASH(master *hdwallet.Key) (address string, err error) {
	wallet, _ := master.GetWallet(hdwallet.CoinType(hdwallet.DASH))
	addr, _ := wallet.GetAddress()
	kk, err := wallet.GetKey().PrivateWIF(true)
	if err != nil {
		return "", err
	}
	newWallet := &Wallet{Address: addr, PrivateKey: kk}
	wf, err := json.Marshal(newWallet)
	if err != nil {
		fmt.Println(err)
		return
	}
	return string(wf), nil
}
