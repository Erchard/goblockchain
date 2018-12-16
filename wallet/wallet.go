package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"log"
)

type Wallet struct {
	Address    string
	PrivateKey []byte
	PublicKey  []byte
}

func CreateWallet() Wallet {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	x509EncodedPriv, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	return RestoreWallet(x509EncodedPriv)
}

func RestoreWallet(x509EncodedPriv []byte) Wallet {
	privateKey, err := x509.ParseECPrivateKey(x509EncodedPriv)
	if err != nil {
		log.Fatal(err)
	}

	x509EncodedPub, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		log.Fatal(err)
	}

	address := sha256.Sum256(x509EncodedPub)

	return Wallet{
		Address:    hex.EncodeToString(address[:]),
		PrivateKey: x509EncodedPriv,
		PublicKey:  x509EncodedPub,
	}
}
