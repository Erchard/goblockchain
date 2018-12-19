package block

import (
	"crypto/ecdsa"
	"encoding/hex"
	"log"
	"math/big"
)

var difficult big.Int = *new(big.Int)

func MineBig(block *Block, key ecdsa.PrivateKey) {

	dif, _ := hex.DecodeString("00000fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")
	difficult.SetBytes(dif)

	if len(block.Nonce) == 0 {
		block.Nonce = "0000000000000000000000000000000000000000000000000000000000000000"
	}

	nonce := *new(big.Int)
	nonceBytes, err := hex.DecodeString(block.Nonce)
	if err != nil {
		log.Fatal(err)
	}
	nonce.SetBytes(nonceBytes)

	for {
		hashBytes, err := hex.DecodeString(block.BlHash)
		if err != nil {
			log.Fatal(err)
		}
		blHash := new(big.Int)
		blHash.SetBytes(hashBytes)

		if difficult.Cmp(blHash) > 0 {

			log.Printf("Difficult: %x\n", difficult.Bytes())
			log.Printf("BlHash:    %x\n", blHash.Bytes())
			raw := ToRaw(*block)
			Sign(&raw, key)
			fromRaw := FromRaw(raw)
			block.Nonce = fromRaw.Nonce
			block.Signature = fromRaw.Signature
			break
		}

		nonce.Add(&nonce, big.NewInt(1))

		block.Nonce = hex.EncodeToString(nonce.Bytes())
		//log.Printf("Nonce: %s\n",block.Nonce)

		raw := ToRaw(*block)
		block.BlHash = hex.EncodeToString(raw.BlHash)
	}

}
