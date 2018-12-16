package transaction

import (
	"../wallet"
	"encoding/hex"
	"fmt"
	"testing"
	"time"
)

func TestFromRaw(t *testing.T) {

	tx1 := GetNewTransaction()
	raw := ToRaw(tx1)
	tx2 := FromRaw(raw)

	if tx1 != tx2 {
		fmt.Printf("Tx1: %x \n", tx1)
		fmt.Printf("Tx2: %x \n", tx2)
		t.Error("tx1 != tx2")
	}

}

func TestFromJSON(t *testing.T) {

	jsonStringA := ToJson(GetNewTransaction())
	tx := FromJSON(jsonStringA)
	jsonStringB := ToJson(tx)

	if jsonStringA != jsonStringB {
		fmt.Println("A: ", jsonStringA)
		fmt.Println("B: ", jsonStringB)
		t.Error("A != B")
	}

}

func GetNewTransaction() Transaction {

	walletSender := wallet.CreateWallet()
	walletReceiver := wallet.CreateWallet()

	tx := Transaction{
		Sender:    walletSender.Address,
		Receiver:  walletReceiver.Address,
		Amount:    1234,
		Fee:       12,
		Timestamp: uint32(time.Now().Unix()),
		PublicKey: hex.EncodeToString(walletSender.PublicKey),
		Signature: "0f23ab2356",
	}

	tx.TxHash = hex.EncodeToString(ToRaw(tx).TxHash)

	return tx
}
