package block

import (
	"../account"
	"encoding/hex"
	"log"
	"testing"
)

func TestMineBig(t *testing.T) {

	miner := account.CreateAccount()
	bl := GetTestBlock()
	bl.PublicKey = hex.EncodeToString(miner.PublicKey)
	privKey := account.RestorePrivKey(miner.PrivateKey)

	log.Print("Start minig...")
	MineBig(&bl, privKey)

	log.Print(bl)
}
