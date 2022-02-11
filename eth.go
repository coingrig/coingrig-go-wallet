package CGWallet

import (
	"encoding/json"

	"github.com/foxnut/go-hdwallet"
)

func getETH(master *hdwallet.Key) (address string, err error) {
	wallet, _ := master.GetWallet(hdwallet.CoinType(hdwallet.ETH))
	addr, _ := wallet.GetAddress()
	newWallet := &Wallet{Address: addr, PrivateKey: "0x" + wallet.GetKey().PrivateHex()}
	wf, err := json.Marshal(newWallet)
	if err != nil {
		return "", err
	}
	return string(wf), nil
}
