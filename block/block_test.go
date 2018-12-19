package block

import (
	"bytes"
	"fmt"
	"log"
	"testing"
)

func TestToRaw(t *testing.T) {
	bl := GetTestBlock()

	log.Printf("BlHash: %s\n", bl.BlHash)
	log.Printf("block: %x/n", bl)
}

func TestFromRaw(t *testing.T) {

	blA := GetTestBlock()
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

func TestCheckSignature(t *testing.T) {

	bl := GetTestBlock()
	raw := ToRaw(bl)

	if !CheckSignature(raw) {
		fmt.Printf("Signature: %x \n", raw.Signature)
		t.Error("signature is not correct")
	}
}
