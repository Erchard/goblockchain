package transaction

import (
	"fmt"
	"testing"
)

func TestFromJSON(t *testing.T) {
	json := TransactionJSON{
		Sender:    "23af",
		Receiver:  "5f12",
		Amount:    1234,
		Fee:       12,
		PublicKey: "9083456f",
		Signature: "2b4a9f00",
	}

	tx := FromJSON(json)

	fmt.Print(tx)
}
