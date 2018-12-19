package account

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"log"
)

type Account struct {
	Address    string
	PrivateKey []byte
	PublicKey  []byte
}

func CreateAccount() Account {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	x509EncodedPriv := PrivateToBytes(*privateKey)

	return RestoreAccount(x509EncodedPriv)
}

func RestoreAccount(x509EncodedPriv []byte) Account {

	privateKey := RestorePrivKey(x509EncodedPriv)

	x509EncodedPub := PublicToBytes(privateKey.PublicKey)

	address := sha256.Sum256(x509EncodedPub)

	return Account{
		Address:    hex.EncodeToString(address[:]),
		PrivateKey: x509EncodedPriv,
		PublicKey:  x509EncodedPub,
	}
}

func RestorePrivKey(x509EncodedPriv []byte) ecdsa.PrivateKey {
	privateKey, err := x509.ParseECPrivateKey(x509EncodedPriv)
	if err != nil {
		log.Fatal(err)
	}
	return *privateKey
}

func RestorePubKey(x509EncodedPub []byte) ecdsa.PublicKey {
	genericPublicKey, err := x509.ParsePKIXPublicKey(x509EncodedPub)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := genericPublicKey.(*ecdsa.PublicKey)
	return *publicKey
}

func PublicToBytes(publicKey ecdsa.PublicKey) []byte {
	x509EncodedPub, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		log.Fatal(err)
	}
	return x509EncodedPub
}

func PrivateToBytes(privateKey ecdsa.PrivateKey) []byte {
	x509EncodedPriv, err := x509.MarshalECPrivateKey(&privateKey)
	if err != nil {
		log.Fatal(err)
	}
	return x509EncodedPriv
}
