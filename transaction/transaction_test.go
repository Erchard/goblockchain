package transaction

import (
	"../wallet"
	"fmt"
	"testing"
	"time"
)

func TestFromRaw(t *testing.T) {

	tx1 := GetNewTransaction()
	raw := ToRaw(tx1)
	tx2 := FromRaw(raw)

	fmt.Printf("Tx1: %x \n", tx1)
	fmt.Printf("Tx2: %x \n", tx2)
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
		PublicKey: walletSender.PublicKey,
		Signature: []byte{12, 43, 76, 87},
	}

	tx.TxHash = ToRaw(tx).TxHash

	return tx
}
