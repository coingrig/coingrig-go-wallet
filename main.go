package CGWallet

import (
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
	case "LTC":
		return getLTC(master)
	case "DOGE":
		return getDOGE(master)
	case "ETC":
		return getETC(master)
	case "BCH":
		return getBCH(master)
	case "DASH":
		return getBCH(master)
	case "SOLANA":
		seed, err := hdwallet.NewSeed(mnemonic, "", "english")
		if err != nil {
			return "", err
		}
		return getSOLANA(seed)
	}
	return "", nil
}
