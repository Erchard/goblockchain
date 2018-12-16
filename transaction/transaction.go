package transaction

import (
	"encoding/hex"
	"log"
)

type Transaction struct {
	TxHash    []byte
	Block     int32
	Sender    []byte
	Receiver  []byte
	Amount    int64
	Fee       int64
	Timestamp int32
	PublicKey []byte
	Signature []byte
}

type TransactionJSON struct {
	TxHash    string
	Block     int32
	Sender    string
	Receiver  string
	Amount    int64
	Fee       int64
	Timestamp int32
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

	return Transaction{
		Sender:    Sender,
		Receiver:  Receiver,
		Amount:    json.Amount,
		Fee:       json.Fee,
		Timestamp: json.Timestamp,
		PublicKey: PublicKey,
		Signature: Signature,
	}
}
