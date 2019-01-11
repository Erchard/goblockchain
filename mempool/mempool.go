package mempool

import (
	"../transaction"
)

var Mempool = make(map[string]transaction.Transaction)

func GetTransactions() []transaction.Transaction {

	var result []transaction.Transaction
	i := 0

	for _, v := range Mempool {
		result = append(result, v)
		i++
		delete(Mempool, v.TxHash)
		if i >= 10 {
			break
		}

	}
	return result
}

func AddTransaction(tx transaction.Transaction) {
	Mempool[tx.TxHash] = tx
}
