package block

import (
	"../account"
	"../transaction"
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
	"testing"
	"time"
)

var height uint32 = 0

func TestToRaw(t *testing.T) {
	bl := GetNewBlock()

	log.Printf("BlHash: %s\n", bl.BlHash)
	log.Printf("block: %x/n", bl)
}

func TestFromRaw(t *testing.T) {

	blA := GetNewBlock()
	rawA := ToRaw(blA)
	blB := FromRaw(rawA)
	rawB := ToRaw(blB)

	if !bytes.Equal(rawA.BlData, rawB.BlData) {
		log.Printf("Block A: %x\n", blA)
		log.Printf("Block B: %x\n", blB)

		log.Printf("A: %x\n", rawA.BlData)
		log.Printf("B: %x\n", rawB.BlData)

		log.Printf("A tx: %x\n", blA.TxList)
		log.Printf("B tx: %x\n", blB.TxList)
		t.Fatal("Block A != Block B")
	}

	log.Printf("blA: %x\n", rawA)
	log.Printf("blB: %x\n", rawB)
}

func GetNewBlock() Block {

	miner := account.CreateAccount()
	privKey := account.RestorePrivKey(miner.PrivateKey)

	bl := Block{
		Height:    height,
		Timestamp: uint32(time.Now().Unix()),
		PublicKey: hex.EncodeToString(miner.PublicKey),
	}

	FillTransaction(&bl)

	raw := ToRaw(bl)

	bl.BlHash = hex.EncodeToString(raw.BlHash)

	Sign(&raw, privKey)
	bl.Signature = hex.EncodeToString(raw.Signature)

	return bl
}

func FillTransaction(block *Block) {

	for i := 0; i < 3; i++ {
		tx := GetNewTransaction()
		AddTransaction(block, &tx)
	}
}

func GetNewTransaction() transaction.Transaction {
	accountSender := account.CreateAccount()
	accountReceiver := account.CreateAccount()

	tx := transaction.Transaction{
		Sender:    accountSender.Address,
		Receiver:  accountReceiver.Address,
		Amount:    1234,
		Fee:       12,
		Timestamp: uint32(time.Now().Unix()),
		PublicKey: hex.EncodeToString(accountSender.PublicKey),
	}

	tx.TxHash = hex.EncodeToString(transaction.ToRaw(tx).TxHash)

	raw := transaction.ToRaw(tx)
	privateKey := account.RestorePrivKey(accountSender.PrivateKey)
	transaction.Sign(&raw, privateKey)

	tx.Signature = hex.EncodeToString(raw.Signature)

	return tx
}

func TestCheckSignature(t *testing.T) {

	bl := GetNewBlock()
	raw := ToRaw(bl)

	if !CheckSignature(raw) {
		fmt.Printf("Signature: %x \n", raw.Signature)
		t.Error("signature is not correct")
	}
}
