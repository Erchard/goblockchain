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
	Sender    []byte
	Receiver  []byte
	Amount    uint64
	Fee       uint64
	Timestamp uint32
	PublicKey []byte
	Signature []byte
}

type TransactionJSON struct {
	TxHash    string
	Block     uint32
	NoInBlock uint32
	Sender    string
	Receiver  string
	Amount    uint64
	Fee       uint64
	Timestamp uint32
	PublicKey string
	Signature string
}

type TransactionRaw struct {
	TxHash    []byte
	TxData    []byte
	PublicKey []byte
	Signature []byte
}

func FromJSON(json TransactionJSON) Transaction {

	Sender, err := hex.DecodeString(json.Sender)
	if err != nil {
		log.Fatal(err)
	}

	Receiver, err := hex.DecodeString(json.Receiver)
	if err != nil {
		log.Fatal(err)
	}

	PublicKey, err := hex.DecodeString(json.PublicKey)
	if err != nil {
		log.Fatal(err)
	}

	Signature, err := hex.DecodeString(json.Signature)
	if err != nil {
		log.Fatal(err)
	}

	tx := Transaction{
		Sender:    Sender,
		Receiver:  Receiver,
		Amount:    json.Amount,
		Fee:       json.Fee,
		Timestamp: json.Timestamp,
		PublicKey: PublicKey,
		Signature: Signature,
	}

	tx.TxHash = ToRaw(tx).TxHash

	return tx
}

func ToRaw(transaction Transaction) TransactionRaw {

	amount := make([]byte, 8)
	binary.BigEndian.PutUint64(amount, transaction.Amount)

	fee := make([]byte, 8)
	binary.BigEndian.PutUint64(fee, transaction.Fee)

	timestamp := make([]byte, 4)
	binary.BigEndian.PutUint32(timestamp, transaction.Timestamp)

	data := append(transaction.Sender, transaction.Receiver...)
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
	pointer := 32 + 32

	amount := binary.BigEndian.Uint64(raw.TxData[pointer : pointer+8])
	pointer += 8

	fee := binary.BigEndian.Uint64(raw.TxData[pointer : pointer+8])
	pointer += 8

	timestamp := binary.BigEndian.Uint32(raw.TxData[pointer : pointer+4])

	return Transaction{
		TxHash:    raw.TxHash,
		Sender:    raw.TxData[0:32],
		Receiver:  raw.TxData[32 : 32+32],
		Amount:    amount,
		Fee:       fee,
		Timestamp: timestamp,
		PublicKey: raw.PublicKey,
		Signature: raw.Signature,
	}
}
