package mempool

import (
	"../transaction"
	"log"
	"testing"
)

func TestGetTransactions(t *testing.T) {

	for i := 0; i < 25; i++ {
		tx := transaction.GetTestTransaction()
		AddTransaction(tx)
	}

	list := GetTransactions()

	for _, tx := range list {
		log.Printf("tx: %x", tx)
	}
	log.Println("Size list: ", len(list))
}
