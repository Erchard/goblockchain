package wallet

import (
	"../block"
	"../transaction"
)

type Wallet struct {
	Address   string
	Amount    uint64
	Timestamp uint32
	Height    uint32
	TxList    []transaction.Transaction
	Fee       []block.Block
}

var Wallets = make(map[string]Wallet)

func TransferMoney(bl block.Block) {

	var miner string
	var fee uint64 = 0

	for _, tx := range bl.TxList {
		if tx.Sender == "0000000000000000000000000000000000000000000000000000000000000000" {
			miner = tx.Receiver
		} else {
			Withdraw(bl, tx)
		}
		Refill(bl, tx)
		MinerFee(bl, miner, fee)
	}

}

func MinerFee(bl block.Block, miner string, fee uint64) {
	if w, ok := Wallets[miner]; ok {
		w.Amount += fee
		w.Timestamp = bl.Timestamp
		w.Height = bl.Height
		w.Fee = append(w.Fee, bl)
		Wallets[w.Address] = w
	} else {
		panic("Miner not found")
	}
}

func Refill(bl block.Block, tx transaction.Transaction) {
	if w, ok := Wallets[tx.Receiver]; ok {
		w.Amount += tx.Amount
		w.Timestamp = tx.Timestamp
		w.Height = bl.Height
		w.TxList = append(w.TxList, tx)
		Wallets[w.Address] = w
	} else {
		txList := []transaction.Transaction{tx}
		blList := []block.Block{}
		w = Wallet{
			Address:   tx.Receiver,
			Amount:    tx.Amount,
			Timestamp: tx.Timestamp,
			Height:    bl.Height,
			TxList:    txList,
			Fee:       blList,
		}
		Wallets[w.Address] = w
	}
}

func Withdraw(bl block.Block, tx transaction.Transaction) {
	if w, ok := Wallets[tx.Sender]; ok {
		if tx.Amount+tx.Fee > w.Amount {
			panic("Not enough money")
		}
		w.Amount -= tx.Amount
		w.Timestamp = tx.Timestamp
		w.Height = bl.Height
		w.TxList = append(w.TxList, tx)
		Wallets[w.Address] = w
	} else {
		panic("Sender wallet not found!")
	}
}
