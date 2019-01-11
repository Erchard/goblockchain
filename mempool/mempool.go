package mempool

import "../transaction"

var Mempool = make(map[string]transaction.Transaction)

func GetTransactions() []transaction.Transaction {

	var result = make([]transaction.Transaction, 10)
	i := 0
	for _, v := range Mempool {
		result = append(result, v)
		i++
		if i >= 10 {
			break
		}
	}
	return result
}
