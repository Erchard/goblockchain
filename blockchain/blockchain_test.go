package blockchain

import (
	"../block"
	"log"
	"testing"
)

func TestGetBlock(t *testing.T) {
	blA := block.GetTestBlock()
	AddBlock(blA)

	blB, err := GetBlock(blA.BlHash)
	if err != nil {
		log.Fatal(err)
	}
	if blA.Timestamp != blB.Timestamp {
		log.Fatal("Timestamp blocks don't equals")
	}

	log.Printf("BlB: %x \n", blB)

	blC, err := GetBlock("00000")
	if err != nil {
		log.Print(err)
	}
	log.Print(blC)
}
