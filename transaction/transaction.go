package transaction

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"log"
)

type Transaction struct {
	TxHash    []byte
	Block     uint32
	NoInBlock uint32
	Sender    string
	Receiver  string
	Amount    uint64
	Fee       uint64
	Timestamp uint32
	PublicKey []byte
	Signature []byte
}

type TransactionRaw struct {
	TxHash    []byte
	TxData    []byte
	PublicKey []byte
	Signature []byte
}

func ToRaw(transaction Transaction) TransactionRaw {
	sender, err := hex.DecodeString(transaction.Sender)
	if err != nil {
		log.Fatal(err)
	}

	receiver, err := hex.DecodeString(transaction.Receiver)
	if err != nil {
		log.Fatal(err)
	}

	amount := make([]byte, 8)
	binary.BigEndian.PutUint64(amount, transaction.Amount)

	fee := make([]byte, 8)
	binary.BigEndian.PutUint64(fee, transaction.Fee)

	timestamp := make([]byte, 4)
	binary.BigEndian.PutUint32(timestamp, transaction.Timestamp)

	data := append(sender, receiver...)
	data = append(data, amount...)
	data = append(data, fee...)
	data = append(data, timestamp...)

	txHash := sha256.Sum256(data)

	return TransactionRaw{
		TxHash:    txHash[:],
		TxData:    data,
		PublicKey: transaction.PublicKey,
		Signature: transaction.Signature,
	}
}

func FromRaw(raw TransactionRaw) Transaction {

	pointer := 0
	sender := hex.EncodeToString(raw.TxData[pointer : pointer+32])
	pointer += 32

	receiver := hex.EncodeToString(raw.TxData[pointer : pointer+32])
	pointer += 32

	amount := binary.BigEndian.Uint64(raw.TxData[pointer : pointer+8])
	pointer += 8

	fee := binary.BigEndian.Uint64(raw.TxData[pointer : pointer+8])
	pointer += 8

	timestamp := binary.BigEndian.Uint32(raw.TxData[pointer : pointer+4])

	return Transaction{
		TxHash:    raw.TxHash,
		Sender:    sender,
		Receiver:  receiver,
		Amount:    amount,
		Fee:       fee,
		Timestamp: timestamp,
		PublicKey: raw.PublicKey,
		Signature: raw.Signature,
	}
}
