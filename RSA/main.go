package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"hash"
)

const (
	generationBits = 2048
)

var (
	message   = []byte("hello from Adamnite")
	hashToUse = sha256.New()
	hashCode  = crypto.SHA256
)

func foo() {
	//the technical details regarding how to generate and encrypt keys need to be discussed and finalized

	// Key Generation
	privKey, pubKey, err := certGeneration()
	checkErr(err)

	// ENCRYPTION AND DECRYPTION

	// Encryption
	cipher, err := encrypt(hashToUse, pubKey, message)
	checkErr(err)
	fmt.Printf("cipher: %q\n", cipher)

	// Decryption
	plainText, err := decrypt(hashCode, privKey, cipher)
	checkErr(err)
	fmt.Printf("plainText: %s\n", plainText)

	// SIGNING AND VERIFICATION

	// Signing
	sum, signature, err := sign(hashCode, privKey, message)
	checkErr(err)

	// Verification
	status, err := verify(hashCode, pubKey, sum, signature)
	checkErr(err)

	fmt.Println(status)

}

// sign
// returns the sum, signature and error
func sign(hashCode crypto.Hash, privateKey *rsa.PrivateKey, message []byte) ([]byte, []byte, error) {
	// We hash our message first
	_, err := hashToUse.Write(message)
	if err != nil {
		return nil, nil, err
	}
	contentHashSum := hashToUse.Sum(nil)

	// signing
	signature, err := rsa.SignPSS(rand.Reader, privateKey, hashCode, contentHashSum, nil)
	if err != nil {
		return nil, nil, err
	}

	return contentHashSum, signature, nil
}

// verify
// return bool only or error too?
func verify(hashCode crypto.Hash, publicKey rsa.PublicKey, sum, signature []byte) (bool, error) {
	err := rsa.VerifyPSS(&publicKey, hashCode, sum, signature, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}

// generate a pair of private and public key
func certGeneration() (*rsa.PrivateKey, rsa.PublicKey, error) {
	// Key Generation
	privateKey, err := rsa.GenerateKey(rand.Reader, generationBits)
	if err != nil {
		return nil, rsa.PublicKey{}, err
	}

	publicKey := privateKey.PublicKey

	return privateKey, publicKey, nil
}

func encrypt(hash hash.Hash, publicKey rsa.PublicKey, content []byte) ([]byte, error) {
	encryptedBytes, err := rsa.EncryptOAEP(
		hash,
		rand.Reader,
		&publicKey,
		[]byte(content),
		nil,
	)
	if err != nil {
		return nil, err
	}
	return encryptedBytes, nil
}

func decrypt(hashCode crypto.Hash, privateKey *rsa.PrivateKey, cipher []byte) ([]byte, error) {
	plainText, err := privateKey.Decrypt(nil, cipher, &rsa.OAEPOptions{Hash: hashCode})
	if err != nil {
		return nil, err
	}

	return plainText, nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
