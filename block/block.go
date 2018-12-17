package block

import (
	"../transaction"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"log"
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

	timestamp := make([]byte, 4)
	binary.BigEndian.PutUint32(timestamp, block.Timestamp)

	nonce, err := hex.DecodeString(block.Nonce)
	if err != nil {
		log.Fatal(err)
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
