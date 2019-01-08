package wallet

import "../transaction"

type Wallet struct {
	Address   string
	Amount    uint64
	Timestamp uint32
	Height    uint32
	TxList    []transaction.Transaction
}
