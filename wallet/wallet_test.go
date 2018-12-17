package wallet

import (
	"fmt"
	"testing"
)

func TestCreateWallet(t *testing.T) {
	wallet := CreateWallet()

	fmt.Printf("Wallet: %x \n", wallet)
}

func TestRestoreWallet(t *testing.T) {
	fmt.Println("Test Restore Wallet")
	wallet1 := CreateWallet()

	x509EncodedPriv := wallet1.PrivateKey

	wallet2 := RestoreWallet(x509EncodedPriv)

	if wallet1.Address != wallet2.Address {
		t.Error("Wallets not equals")
	}
}
