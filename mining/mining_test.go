package mining

import (
	"../account"
	"../block"
	"log"
	"testing"
)

func TestMineBig(t *testing.T) {

	miner := account.CreateAccount()
	bl := block.GetTestBlock()
	privKey := account.RestorePrivKey(miner.PrivateKey)

	log.Print("Start minig...")
	MineBig(&bl, privKey)

	log.Print(bl)
}
