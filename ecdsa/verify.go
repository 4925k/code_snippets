package main

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"math/big"
	"os"
)

//Public key verification
func publicKeyVerify(data, rText, sText []byte, publicKeyPemFileName string) bool {
	//Read public key
	publicFile, err := os.Open(publicKeyPemFileName)
	if err != nil {
		panic(err)
	}
	defer publicFile.Close()
	fileInfo, err := publicFile.Stat()
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, fileInfo.Size())
	publicFile.Read(buffer)
	block, _ := pem.Decode(buffer)
	publicKeyType, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//Get the type of publickey, and make assertion judgment RSA / ECC
	publicKey := publicKeyType.(*ecdsa.PublicKey)
	//Convert [] byte signature data into big data
	var r, s big.Int
	r.UnmarshalJSON(rText)
	s.UnmarshalJSON(sText)
	if !ecdsa.Verify(publicKey, getHash(data), &r, &s) {
		return false
	} else {
		return true
	}
}
