package transaction

import (
	"fmt"
	"testing"
)

func TestFromRaw(t *testing.T) {

	tx1 := GetTestTransaction()
	raw := ToRaw(tx1)
	tx2 := FromRaw(raw)

	if tx1 != tx2 {
		fmt.Printf("Tx1: %x \n", tx1)
		fmt.Printf("Tx2: %x \n", tx2)
		t.Error("tx1 != tx2")
	}

}

func TestFromJSON(t *testing.T) {

	jsonStringA := ToJson(GetTestTransaction())
	tx := FromJSON(jsonStringA)
	jsonStringB := ToJson(tx)

	if jsonStringA != jsonStringB {
		fmt.Println("A: ", jsonStringA)
		fmt.Println("B: ", jsonStringB)
		t.Error("A != B")
	}

}

func TestCheckSignature(t *testing.T) {

	tx := GetTestTransaction()
	raw := ToRaw(tx)

	if !CheckSignature(raw) {
		fmt.Printf("Signature: %x \n", raw.Signature)
		t.Error("signature is not correct")
	}
}
