package block

import (
	"../transaction"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"log"
	"time"
)

type Block struct {
	BlHash    string                    `json:"bl_hash"`
	Height    uint32                    `json:"height"`
	Previous  string                    `json:"prev"`
	Timestamp uint32                    `json:"timestamp"`
	Nonce     string                    `json:"nonce"`
	TxList    []transaction.Transaction `json:"tx_list"`
}

type BlockRaw struct {
	BlHash []byte
	BlData []byte
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

	raw := BlockRaw{
		BlHash: blHash[:],
		BlData: data,
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
	}

	return bl
}

func AddTransaction(block *Block, tx *transaction.Transaction) {
	tx.Block = block.BlHash
	tx.NoInBlock = uint32(len(block.TxList))
	block.TxList = append(block.TxList, *tx)
}

var testheight uint32 = 0

func GetTestBlock() Block {

	bl := Block{
		Height:    testheight,
		Timestamp: uint32(time.Now().Unix()),
	}

	FillTestTransaction(&bl)

	raw := ToRaw(bl)

	bl.BlHash = hex.EncodeToString(raw.BlHash)

	return bl
}

func FillTestTransaction(bl *Block) {

	for i := 0; i < 3; i++ {
		tx := transaction.GetTestTransaction()
		AddTransaction(bl, &tx)
	}
}

func BlockIsValid(bl Block) bool {
	return true
}
