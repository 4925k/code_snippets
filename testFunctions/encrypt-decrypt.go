package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"strings"
)

// // encrypt using the given public key and SHA512 to encrypt the given []bytes
// // returns encrypted []byte and error if any
// func Encrypt(publicKey ecdsa.PublicKey, content []byte) ([]byte, error) {
// 	encryptedBytes, err := rsa.EncryptOAEP(
// 		sha512.New(),
// 		rand.Reader,
// 		&publicKey,
// 		[]byte(content),
// 		nil,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return encryptedBytes, nil
// }

// func Decrypt(privateKey *ecdsa.PrivateKey, cipher []byte) ([]byte, error) {
// 	plainText, err := privateKey.Decrypt(nil, cipher, &rsa.OAEPOptions{Hash: crypto.SHA512})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return plainText, nil
// }

// generate a pair of private and public key
func certGeneration() (*ecdsa.PrivateKey, ecdsa.PublicKey, error) {
	// Key Generation
	privateKey, err := ecdsa.GenerateKey(elliptic.P384(), strings.NewReader("password"))
	if err != nil {
		return nil, ecdsa.PublicKey{}, err
	}

	publicKey := privateKey.PublicKey

	return privateKey, publicKey, nil
}

// func main() {
// 	priv, pub, err := certGeneration()
// 	checkErr(err)

// 	privBlock := &pem.Block{
// 		Type:  "RSA PRIVATE KEY",
// 		Bytes: x509.MarshalPKCS1PrivateKey(priv),
// 	}

// 	os.WriteFile("private.pem", pem.EncodeToMemory(privBlock), 0666)

// 	publicKey, err := x509.MarshalPKIXPublicKey(pub)
// 	checkErr(err)
// 	publicBlock := &pem.Block{
// 		Type:  "PUBLIC KEY",
// 		Bytes: publicKey,
// 	}

// 	os.WriteFile("public.pem", pem.EncodeToMemory(publicBlock), 0666)
// 	// message := []byte("hello World")
// 	// fmt.Printf("Message: %s\n", message)

// 	// cipher, err := Encrypt(pub, message)
// 	// checkErr(err)
// 	// fmt.Printf("Cipher: %v\n")

// 	// plaintext, err := Decrypt(priv, cipher)
// 	// checkErr(err)
// 	// fmt.Printf("Plaintext: %s\n", plaintext)
// }
