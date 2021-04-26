package main

import (
	"crypto/des"
	"encoding/hex"
	"fmt"
)

var blowfishKey = "[;'.]94-31==-%&@!^+]"

func main() {
	var encrypted = make([]byte, 1024)

	// Manually extracted
	encryptedSize, err := hex.Decode(encrypted, []byte("554F8A9C3B24261F394BAE5A2F3117009301CD012CBE78E6"))
	//encryptedSize, err := hex.Decode(encrypted, []byte("dd66b254dfd5994466c151405c506caad09873ea09c7a5d1a1f5e02a4048ec776f1ae7c373a3a89f58380ac8d5e75cc8"))

	decrypted := DecryptDES(err, encryptedSize, encrypted)

	fmt.Printf("%s\n", decrypted[0:encryptedSize])
}

func DecryptDES(err error, encryptedSize int, encrypted []byte) []byte {
	start := 0
	blockSize := 8

	if err != nil {
		panic(err)
	}

	block, err := des.NewCipher([]byte{'T', 'E', 'S', 'T', 0, 0, 0, 0})

	if err != nil {
		panic(err)
	}

	decrypted := make([]byte, 1024)

	for ; start < encryptedSize; start += blockSize {
		end := start + blockSize

		block.Decrypt(decrypted[start:], encrypted[start:end])
	}
	return decrypted
}

