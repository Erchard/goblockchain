package account

import (
	"fmt"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	account := CreateAccount()

	fmt.Printf("Account: %x \n", account)
}

func TestRestoreAccount(t *testing.T) {
	fmt.Println("Test Restore Account")
	accountA := CreateAccount()

	x509EncodedPriv := accountA.PrivateKey

	accountB := RestoreAccount(x509EncodedPriv)

	if accountA.Address != accountB.Address {
		t.Error("Accounts not equals")
	}
}
