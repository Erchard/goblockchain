package block

import (
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

func GetNewBlock() Block {

	bl := Block{
		Height:    height,
		Timestamp: uint32(time.Now().Unix()),
	}

	raw := ToRaw(bl)

	bl.BlHash = hex.EncodeToString(raw.BlHash)

	return bl
}
