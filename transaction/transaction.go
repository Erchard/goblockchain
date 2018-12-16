package transaction

import (
	"../wallet"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"log"
	"math/big"
)

type Transaction struct {
	TxHash    string `json:"tx_hash,hex"`
	Block     uint32 `json:"block"`
	NoInBlock uint32 `json:"no_in_block"`
	Sender    string `json:"sender"`
	Receiver  string `json:"receiver"`
	Amount    uint64 `json:"amount"`
	Fee       uint64 `json:"fee"`
	Timestamp uint32 `json:"timestamp"`
	PublicKey string `json:"pub_key"`
	Signature string `json:"signature"`
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

	pubKey, err := hex.DecodeString(transaction.PublicKey)
	if err != nil {
		log.Fatal(err)
	}

	signature, err := hex.DecodeString(transaction.Signature)
	if err != nil {
		log.Fatal(err)
	}

	return TransactionRaw{
		TxHash:    txHash[:],
		TxData:    data,
		PublicKey: pubKey,
		Signature: signature,
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
		TxHash:    hex.EncodeToString(raw.TxHash),
		Sender:    sender,
		Receiver:  receiver,
		Amount:    amount,
		Fee:       fee,
		Timestamp: timestamp,
		PublicKey: hex.EncodeToString(raw.PublicKey),
		Signature: hex.EncodeToString(raw.Signature),
	}
}

func ToJson(transaction Transaction) string {
	jsonBytes, err := json.Marshal(transaction)
	if err != nil {
		log.Fatal(err)
	}

	return string(jsonBytes)
}

func FromJSON(jsonString string) Transaction {
	var tx Transaction
	err := json.Unmarshal([]byte(jsonString), &tx)
	if err != nil {
		log.Fatal(err)
	}
	return tx
}

func CheckSignature(raw TransactionRaw) bool {

	pubKey := wallet.RestorePubKey(raw.PublicKey)

	sig := raw.Signature

	r := new(big.Int).SetBytes(sig[:len(sig)/2])
	s := new(big.Int).SetBytes(sig[len(sig)/2:])

	return ecdsa.Verify(&pubKey, raw.TxData, r, s)
}
