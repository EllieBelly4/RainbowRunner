package crypt

import (
	"bytes"
	"encoding/binary"
	"golang.org/x/crypto/blowfish"
)

var blowfishKey = "[;\x27.]94-31==-%&@!^+]\x00"

const blockSize = 8

func EncryptBlowfish(plainText []byte, length int) []byte {
	//TODO cleanup
	cipher, err := blowfish.NewCipher([]byte(blowfishKey))

	if err != nil {
		panic(err)
	}

	start := 0
	var encrypted = make([]byte, length)

	for ; start < length; start += blockSize {
		end := start + blockSize

		endian, err := convertEndian(plainText[start:end])

		if err != nil {
			panic(err)
		}

		tmp := make([]byte, blockSize)

		cipher.Encrypt(tmp, endian)

		tmp, err = convertEndian(tmp)

		if err != nil {
			panic(err)
		}

		copy(encrypted[start:], tmp)
	}

	return encrypted
}

func DecryptBlowfish(encrypted []byte, encryptedSize int) []byte {
	//TODO cleanup
	cipher, err := blowfish.NewCipher([]byte(blowfishKey))

	if err != nil {
		panic(err)
	}

	start := 0
	var decrypted = make([]byte, encryptedSize)

	for ; start < encryptedSize; start += blockSize {
		end := start + blockSize

		endian, err := convertEndian(encrypted[start:end])

		if err != nil {
			panic(err)
		}

		decryptedTmp := make([]byte, blockSize)

		cipher.Decrypt(decryptedTmp, endian)

		decryptedTmp, err = convertEndian(decryptedTmp)

		if err != nil {
			panic(err)
		}

		copy(decrypted[start:], decryptedTmp)
	}

	return decrypted
}

func convertEndian(in []byte) ([]byte, error) {
	//Read byte array as uint32 (little-endian)
	var v1, v2 uint32
	buf := bytes.NewReader(in)
	if err := binary.Read(buf, binary.LittleEndian, &v1); err != nil {
		return nil, err
	}
	if err := binary.Read(buf, binary.LittleEndian, &v2); err != nil {
		return nil, err
	}

	//convert uint32 to byte array
	out := make([]byte, 8)
	binary.BigEndian.PutUint32(out, v1)
	binary.BigEndian.PutUint32(out[4:], v2)

	return out, nil
}
