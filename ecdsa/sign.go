package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"os"
)

//Private key signature
func privateKeySignature(data []byte, privateKeyPemFileName string) (rText, sText []byte) {
	//Read private key
	privateFile, err := os.Open(privateKeyPemFileName)
	if err != nil {
		panic(err)
	}
	defer privateFile.Close()
	//Read file source information
	fileInfo, err := privateFile.Stat()
	if err != nil {
		panic(err)
	}
	buffer := make([]byte, fileInfo.Size())
	privateFile.Read(buffer)
	//PEM decoding
	block, _ := pem.Decode(buffer)
	//X509 der decoding
	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//Sign
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, getHash(data))
	if err != nil {
		panic(err)
	}
	//Convert big data to [] byte
	rText, _ = r.MarshalText()
	sText, _ = s.MarshalText()
	return
}
