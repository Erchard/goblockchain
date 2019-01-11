package transaction

import (
	"../account"
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"log"
	"math/big"
	"time"
)

type Transaction struct {
	TxHash    string `json:"tx_hash"`
	Block     string `json:"block"`
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

func Sign(raw *TransactionRaw, privateKey ecdsa.PrivateKey) {

	hash := sha256.Sum256(raw.TxData)

	if !bytes.Equal(raw.TxHash, hash[:]) {
		log.Fatal("Hash Transaction Error")
		raw.TxHash = hash[:]
	}

	x509EncodedPub := account.PublicToBytes(privateKey.PublicKey)

	if !bytes.Equal(x509EncodedPub, raw.PublicKey) {
		log.Fatal("Public key is not correct")
		raw.PublicKey = x509EncodedPub
	}

	r, s, err := ecdsa.Sign(rand.Reader, &privateKey, hash[:])
	if err != nil {
		log.Fatal(err)
	}

	raw.Signature = append(r.Bytes(), s.Bytes()...)
}

func CheckSignature(raw TransactionRaw) bool {

	pubKey := account.RestorePubKey(raw.PublicKey)

	sig := raw.Signature

	r := new(big.Int).SetBytes(sig[:len(sig)/2])
	s := new(big.Int).SetBytes(sig[len(sig)/2:])

	return ecdsa.Verify(&pubKey, raw.TxHash, r, s)
}

func GetTestTransaction() Transaction {
	accountSender := account.CreateAccount()
	accountReceiver := account.CreateAccount()

	tx := Transaction{
		Sender:    accountSender.Address,
		Receiver:  accountReceiver.Address,
		Amount:    1234,
		Fee:       12,
		Timestamp: uint32(time.Now().Unix()),
		PublicKey: hex.EncodeToString(accountSender.PublicKey),
	}

	tx.TxHash = hex.EncodeToString(ToRaw(tx).TxHash)

	raw := ToRaw(tx)
	privateKey := account.RestorePrivKey(accountSender.PrivateKey)
	Sign(&raw, privateKey)

	tx.Signature = hex.EncodeToString(raw.Signature)

	return tx
}

func CreateMinerTx(miner account.Account) Transaction {
	minerTx := Transaction{
		Sender:    "0000000000000000000000000000000000000000000000000000000000000000",
		Receiver:  miner.Address,
		Amount:    100000000000,
		Fee:       0,
		Timestamp: uint32(time.Now().Unix()),
		PublicKey: hex.EncodeToString(miner.PublicKey),
	}

	minerTx.TxHash = hex.EncodeToString(ToRaw(minerTx).TxHash)

	raw := ToRaw(minerTx)
	privateKey := account.RestorePrivKey(miner.PrivateKey)
	Sign(&raw, privateKey)

	minerTx.Signature = hex.EncodeToString(raw.Signature)
	return minerTx
}
