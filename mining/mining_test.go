package mining

import (
	"../account"
	"../block"
	"encoding/hex"
	"log"
	"testing"
)

func TestMineBig(t *testing.T) {

	miner := account.CreateAccount()
	bl := block.GetTestBlock()
	bl.PublicKey = hex.EncodeToString(miner.PublicKey)
	privKey := account.RestorePrivKey(miner.PrivateKey)

	log.Print("Start minig...")
	MineBig(&bl, privKey)

	log.Print(bl)
}
