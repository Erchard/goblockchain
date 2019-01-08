package mempool

import "../transaction"

var Mempool = make(map[string]transaction.Transaction)
