package main

import (
	"crypto/sha512"
)

//Get hash
func getHash(data []byte) []byte {
	hash := sha512.Sum512(data)
	return hash[:]
}
