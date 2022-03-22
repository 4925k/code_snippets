package main

import "fmt"

func main() {
	generateECDSAkey()
	message := []byte("test")

	r, s := privateKeySignature(message, "ECDSAPrivateKey.pem")
	fmt.Println(publicKeyVerify(message, r, s, "ECDSAPublicKey.pem"))
}
