package mining

import (
	"../account"
	"../block"
	"../blockchain"
	"../mempool"
	"../transaction"
	"crypto/ecdsa"
	"encoding/hex"
	"log"
	"math/big"
	"time"
)

var difficult big.Int = *new(big.Int)

var miner account.Account

func MineBig(bl *block.Block, key ecdsa.PrivateKey) {

	dif, _ := hex.DecodeString("00000fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")
	difficult.SetBytes(dif)

	if len(bl.Nonce) == 0 {
		bl.Nonce = "0000000000000000000000000000000000000000000000000000000000000000"
	}

	if len(bl.BlHash) == 0 {
		raw := block.ToRaw(*bl)
		bl.BlHash = hex.EncodeToString(raw.BlHash)
	}

	nonce := *new(big.Int)
	nonceBytes, err := hex.DecodeString(bl.Nonce)
	if err != nil {
		log.Fatal(err)
	}
	nonce.SetBytes(nonceBytes)

	for {
		hashBytes, err := hex.DecodeString(bl.BlHash)
		if err != nil {
			log.Fatal(err)
		}
		blHash := new(big.Int)
		blHash.SetBytes(hashBytes)

		if difficult.Cmp(blHash) > 0 {
			break
		}

		nonce.Add(&nonce, big.NewInt(1))

		bl.Nonce = hex.EncodeToString(nonce.Bytes())
		//log.Printf("Nonce: %s\n",bl.Nonce)

		raw := block.ToRaw(*bl)
		bl.BlHash = hex.EncodeToString(raw.BlHash)
	}

}

func MineLoop(ending chan bool) {

	miner = account.CreateAccount()

	log.Print("Miner: ", miner.Address)

	privKey := account.RestorePrivKey(miner.PrivateKey)

	stop := false

	for !stop {
		select {
		case stop = <-ending:
			log.Println("Stop: ", stop)
			log.Println("Press Enter")
			break
		default:
			lastBlock := blockchain.GetLastBlock()
			bl := block.Block{
				Height:    0,
				Previous:  "0000000000000000000000000000000000000000000000000000000000000000",
				Timestamp: uint32(time.Now().Unix()),
			}

			if lastBlock != nil {
				bl.Height = lastBlock.Height + 1
				bl.Previous = lastBlock.BlHash
			}

			minerTx := transaction.CreateMinerTx(miner)

			bl.TxList = append(mempool.GetTransactions(), minerTx)

			MineBig(&bl, privKey)
			log.Printf("%d %s %s \n", bl.Height, bl.Previous, bl.BlHash)
			blockchain.AddBlock(bl)
		}

	}

}
