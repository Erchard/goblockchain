package block

import (
	"bytes"
	"encoding/hex"
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
		t.Fatal("Block A != Block B")
	}

	log.Printf("blA: %x\n", rawA)
	log.Printf("blB: %x\n", rawB)
}

func GetNewBlock() Block {

	bl := Block{
		Height:    height,
		Timestamp: uint32(time.Now().Unix()),
	}

	raw := ToRaw(bl)

	bl.BlHash = hex.EncodeToString(raw.BlHash)

	return bl
}
