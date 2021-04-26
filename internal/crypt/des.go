package crypt

import "crypto/des"

func DecryptDES(encrypted []byte, encryptedSize int) []byte {
	start := 0
	blockSize := 8

	block, err := des.NewCipher([]byte{'T', 'E', 'S', 'T', 0, 0, 0, 0})

	if err != nil {
		panic(err)
	}

	decrypted := make([]byte, encryptedSize)

	for ; start < encryptedSize; start += blockSize {
		end := start + blockSize

		block.Decrypt(decrypted[start:], encrypted[start:end])
	}
	return decrypted
}


