package transaction

import (
	"../wallet"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestFromJSON(t *testing.T) {

	walletSender := wallet.CreateWallet()
	walletReceiver := wallet.CreateWallet()

	pubKey := hex.EncodeToString(walletSender.PublicKey)

	json := TransactionJSON{
		Sender:    walletSender.Address,
		Receiver:  walletReceiver.Address,
		Amount:    1234,
		Fee:       12,
		PublicKey: pubKey,
		Signature: "2b4a9f00",
	}

	tx := FromJSON(json)

	fmt.Print(tx)
}
