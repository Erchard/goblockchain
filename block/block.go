package block

import (
	"../account"
	"../transaction"
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"log"
	"math/big"
)

type Block struct {
	BlHash    string                    `json:"bl_hash"`
	Height    uint32                    `json:"height"`
	Previous  string                    `json:"prev"`
	Timestamp uint32                    `json:"timestamp"`
	Nonce     string                    `json:"nonce"`
	TxList    []transaction.Transaction `json:"tx_list"`
	PublicKey string                    `json:"pub_key"`
	Signature string                    `json:"signature"`
}

type BlockRaw struct {
	BlHash    []byte
	BlData    []byte
	PublicKey []byte
	Signature []byte
}

func ToRaw(block Block) BlockRaw {

	height := make([]byte, 4)
	binary.BigEndian.PutUint32(height, block.Height)

	previous, err := hex.DecodeString(block.Previous)
	if err != nil {
		log.Fatal(err)
	}
	if len(previous) == 0 {
		previous = make([]byte, 32)
	}

	timestamp := make([]byte, 4)
	binary.BigEndian.PutUint32(timestamp, block.Timestamp)

	nonce, err := hex.DecodeString(block.Nonce)
	if err != nil {
		log.Fatal(err)
	}
	if len(nonce) < 32 {
		pad := make([]byte, 32-len(nonce))
		nonce = append(pad, nonce...)
	}

	data := append(height, previous...)
	data = append(data, timestamp...)
	data = append(data, nonce...)

	for _, tx := range block.TxList {
		txhash, err := hex.DecodeString(tx.TxHash)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, txhash...)
	}

	blHash := sha256.Sum256(data)

	pubKey, err := hex.DecodeString(block.PublicKey)
	if err != nil {
		log.Fatal(err)
	}

	signature, err := hex.DecodeString(block.Signature)
	if err != nil {
		log.Fatal(err)
	}

	raw := BlockRaw{
		BlHash:    blHash[:],
		BlData:    data,
		PublicKey: pubKey,
		Signature: signature,
	}

	return raw
}

func FromRaw(raw BlockRaw) Block {
	pointer := 0
	height := binary.BigEndian.Uint32(raw.BlData[pointer : pointer+4])
	pointer += 4

	previous := hex.EncodeToString(raw.BlData[pointer : pointer+32])
	pointer += 32

	timestamp := binary.BigEndian.Uint32(raw.BlData[pointer : pointer+4])
	pointer += 4

	nonce := hex.EncodeToString(raw.BlData[pointer : pointer+32])
	pointer += 32

	var txList []transaction.Transaction

	for pointer < len(raw.BlData) {
		tx := hex.EncodeToString(raw.BlData[pointer : pointer+32])
		txList = append(txList, transaction.Transaction{
			TxHash: tx,
		})
		pointer += 32
	}

	bl := Block{
		BlHash:    hex.EncodeToString(raw.BlHash),
		Height:    height,
		Previous:  previous,
		Timestamp: timestamp,
		Nonce:     nonce,
		TxList:    txList,
		PublicKey: hex.EncodeToString(raw.PublicKey),
		Signature: hex.EncodeToString(raw.Signature),
	}

	return bl
}

func Sign(raw *BlockRaw, privateKey ecdsa.PrivateKey) {

	hash := sha256.Sum256(raw.BlData)

	if !bytes.Equal(raw.BlHash, hash[:]) {
		log.Fatal("Hash Transaction Error")
		raw.BlHash = hash[:]
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

func CheckSignature(raw BlockRaw) bool {

	pubKey := account.RestorePubKey(raw.PublicKey)

	sig := raw.Signature

	r := new(big.Int).SetBytes(sig[:len(sig)/2])
	s := new(big.Int).SetBytes(sig[len(sig)/2:])

	return ecdsa.Verify(&pubKey, raw.BlHash, r, s)
}

func AddTransaction(block *Block, tx *transaction.Transaction) {
	tx.Block = block.BlHash
	tx.NoInBlock = uint32(len(block.TxList))
	block.TxList = append(block.TxList, *tx)
}
