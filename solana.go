package CGWallet

import (
	"encoding/json"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"github.com/portto/solana-go-sdk/types"
	"github.com/stellar/go/exp/crypto/derivation"
)

func getSOLANA(seed []byte) (address string, err error) {

	kk, err := derivation.DeriveForPath("m/44'/501'/0'/0'", seed)
	if err != nil {
		return "", err
	}
	wallet, err := types.AccountFromSeed(kk.Key)
	// fmt.Println("Private Key:", base58.Encode(wallet.PrivateKey))

	if err != nil {
		return "", err
	}
	newWallet := &Wallet{Address: wallet.PublicKey.ToBase58(), PrivateKey: base58.Encode(wallet.PrivateKey)}
	wf, err := json.Marshal(newWallet)
	if err != nil {
		fmt.Println(err)
		return
	}
	return string(wf), nil
}
