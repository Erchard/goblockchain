package blockchain

import (
	"../block"
	"../wallet"
	"errors"
)

var blockchain = make(map[string]block.Block)

var lastBlock *block.Block

func AddBlock(bl block.Block) {
	if block.BlockIsValid(bl) {
		blockchain[bl.BlHash] = bl
		lastBlock = &bl
		wallet.TransferMoney(bl)
	}
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
