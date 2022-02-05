package CGWallet

import (
	"encoding/json"
	"fmt"

	"github.com/foxnut/go-hdwallet"
)

type Wallet struct {
	Address    string `json:"address"`
	PrivateKey string `json:"privateKey"`
}

func GenerateMnemonic(size int) string {
	mnemonic, err := hdwallet.NewMnemonic(size, "english")
	if err != nil {
		panic(err)
	}
	return mnemonic
}

func GenerateWallet(mnemonic string, chain string) (address string, err error) {
	master, err := hdwallet.NewKey(
		hdwallet.Mnemonic(mnemonic),
	)
	if err != nil {
		return "", err
	}

	switch chain {
	case "BTC":
		return getBTC(master)
	case "ETH":
		return getETH(master)
	}
	return "", nil
}

func getBTC(master *hdwallet.Key) (address string, err error) {
	wallet, _ := master.GetWallet(hdwallet.CoinType(hdwallet.BTCTestnet), hdwallet.Path("m/84'/0'/0'/0/0"))
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

func getETH(master *hdwallet.Key) (address string, err error) {
	wallet, _ := master.GetWallet(hdwallet.CoinType(hdwallet.ETH))
	addr, _ := wallet.GetAddress()
	kk, err := wallet.GetKey().PrivateWIF(true)
	if err != nil {
		return "", err
	}
	newWallet := &Wallet{Address: addr, PrivateKey: kk}
	wf, err := json.Marshal(newWallet)
	if err != nil {
		return "", err
	}
	return string(wf), nil
}
