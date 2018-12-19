package transaction

import (
	"../account"
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

	accountSender := account.CreateAccount()
	accountReceiver := account.CreateAccount()

	tx := Transaction{
		Sender:    accountSender.Address,
		Receiver:  accountReceiver.Address,
		Amount:    1234,
		Fee:       12,
		Timestamp: uint32(time.Now().Unix()),
		PublicKey: hex.EncodeToString(accountSender.PublicKey),
	}

	tx.TxHash = hex.EncodeToString(ToRaw(tx).TxHash)

	raw := ToRaw(tx)
	privateKey := account.RestorePrivKey(accountSender.PrivateKey)
	Sign(&raw, privateKey)

	tx.Signature = hex.EncodeToString(raw.Signature)

	return tx
}

func TestCheckSignature(t *testing.T) {

	tx := GetNewTransaction()
	raw := ToRaw(tx)

	if !CheckSignature(raw) {
		fmt.Printf("Signature: %x \n", raw.Signature)
		t.Error("signature is not correct")
	}
}
