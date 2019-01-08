package blockchain

import (
	"../block"
	"errors"
)

var blockchain = make(map[string]block.Block)

var lastBlock *block.Block

func AddBlock(bl block.Block) {
	blockchain[bl.BlHash] = bl
	lastBlock = &bl
}

func GetBlock(BlHash string) (*block.Block, error) {

	if bl, ok := blockchain[BlHash]; ok {
		return &bl, nil
	}
	return nil, errors.New("Block not found")
}

func GetLastBlock() *block.Block {
	return lastBlock
}
