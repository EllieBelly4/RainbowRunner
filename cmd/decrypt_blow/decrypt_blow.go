package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/blowfish"
)

//var blowfishKey = "[;'.]94-31==-%&@!^+]"
var blowfishKey = "[;\x27.]94-31==-%&@!^+]\x00"

const blockSize = 8

//var blowfishKey = "["

func main() {
	blowfishKeyBytes := []byte(blowfishKey)

	fmt.Printf("Length %d, Key: %s\n", len(blowfishKeyBytes), blowfishKeyBytes)

	expected := make([]byte, 1024)

	expectedLength, err := hex.Decode(expected, []byte("00554F8A9C3B24261F394BAE5A2F3117009301CD012CBE78E6000000000000ADBA080000000100ADBAABABABABABABABAB00"))

	if err != nil {
		panic(err)
	}

	var encrypted = make([]byte, 1024)

	encryptedSize, err := hex.Decode(encrypted, []byte("3200dd66b254dfd5994466c151405c506caad09873ea09c7a5d1a1f5e02a4048ec776f1ae7c373a3a89f58380ac8d5e75cc8"))
	encryptedSize -= 2

	payloadLength := binary.LittleEndian.Uint16(encrypted[0:2])

	encrypted = encrypted[2:]

	fmt.Printf("Encrypted:\n%s\n", hex.Dump(encrypted[0:encryptedSize]))

	// All 0s
	//encryptedSize, err := hex.Decode(encrypted, []byte("624A79FE01ED0600"))
	//encryptedSize, err := hex.Decode(encrypted, []byte("DD66B254DFD59944"))

	if err != nil {
		panic(err)
	}

	block, err := blowfish.NewCipher(blowfishKeyBytes)

	if err != nil {
		panic(err)
	}

	if uint16(encryptedSize+2) != payloadLength {
		panic(fmt.Sprintf("Expected data length %d, got %d", payloadLength, encryptedSize))
	}

	fmt.Println()

	decrypted := decryptBlowfish(encrypted, encryptedSize, block)

	//hexString := hex.EncodeToString(decrypted[0:encryptedSize])

	fmt.Printf("Expected result:\n%s\n", hex.Dump(expected[0:expectedLength]))
	fmt.Printf("Actual result  :\n%s\n", hex.Dump(decrypted))
}

func decryptBlowfish(encrypted []byte, encryptedSize int, cipher *blowfish.Cipher) []byte {
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
