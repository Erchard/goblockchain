package transaction

import (
	"../wallet"
	"encoding/hex"
	"fmt"
	"testing"
	"time"
)

func TestFromJSON(t *testing.T) {

	json := GetNewTransaction()

	tx := FromJSON(json)

	fmt.Print(tx)

	fmt.Println("Sender len: %n", len(tx.Sender))
	fmt.Println("Receiver len: %n", len(tx.Receiver))
	fmt.Println("PublicKey len: %n", len(tx.PublicKey))
	fmt.Println("Signature len: %n", len(tx.Signature))

}

func TestToRaw(t *testing.T) {
	tx := FromJSON(GetNewTransaction())

	raw := ToRaw(tx)

	fmt.Printf("raw: %x\n", raw)
}

func GetNewTransaction() TransactionJSON {

	walletSender := wallet.CreateWallet()
	walletReceiver := wallet.CreateWallet()

	pubKey := hex.EncodeToString(walletSender.PublicKey)

	return TransactionJSON{
		Sender:    walletSender.Address,
		Receiver:  walletReceiver.Address,
		Amount:    1234,
		Fee:       12,
		Timestamp: uint32(time.Now().Unix()),
		PublicKey: pubKey,
		Signature: "2b4a9f00",
	}
}
