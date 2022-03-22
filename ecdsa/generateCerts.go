package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"os"
)

const (
	PRIVATEKEYFILE = "ECDSAPrivateKey.pem"
	PRIVATETYPE    = "ECC privatekey"

	PUBLICKEYFILE = "ECDSAPublicKey.pem"
	PUBLICTYPE    = "Ecc PublicKey"
)

// generateECDSAkey uses ecdsa to generate a public and private key
func generateECDSAkey() {
	//Generate the key of ECC algorithm
	privateKey, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	if err != nil {
		panic(err)
	}
	//Localize the private key and serialize it using x509
	privateDerBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		panic(err)
	}
	//Then convert it into PEM format coding
	privatePemBlock := pem.Block{
		Type:  PRIVATETYPE, // introduction to using algorithm type RSA / ECC
		Bytes: privateDerBytes,
	}
	//Create PEM files locally
	privateFile, err := os.Create(PRIVATEKEYFILE)
	if err != nil {
		panic(err)
	}
	defer privateFile.Close()
	//PEM coding
	err = pem.Encode(privateFile, &privatePemBlock)
	if err != nil {
		panic(err)
	}
	//Public key isomorphism
	publicKey := privateKey.PublicKey
	publicDerBytes, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}
	publicPemBlock := pem.Block{
		Type:  "",
		Bytes: publicDerBytes,
	}
	publicFile, err := os.Create(PUBLICKEYFILE)
	if err != nil {
		panic(err)
	}
	defer publicFile.Close()
	pem.Encode(publicFile, &publicPemBlock)
}
