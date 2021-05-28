package main

import (
	"RainbowRunner/internal/byter"
	"bytes"
	"compress/zlib"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fileName := "E:\\Games\\DungeonRunners v666\\game.pki"

	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	outFile, err := os.OpenFile("D:\\Work\\dungeon-runners\\game_decompressed.pki", os.O_CREATE|os.O_TRUNC, 644)

	if err != nil {
		panic(err)
	}

	inFileByter := byter.NewLEByter(data)

	// Copying header
	// Unk uint32
	// Unk uint32
	// GUID uint128
	outFile.Write(inFileByter.Bytes(24))

	//i := inFileByter.I
	inFileReader := bytes.NewReader(inFileByter.RemainingBytes())

	reader, err := zlib.NewReader(inFileReader)

	if err != nil {
		panic(err)
	}

	buffer := make([]byte, 1024*100000)

	var bytesRead = 0
	var totalBytes = 0

	for {
		bytesRead, err = reader.Read(buffer[totalBytes:])

		if err != nil && !errors.Is(err, io.EOF) {
			panic(err)
		}

		fmt.Printf("Read %d uncompressed bytes\n", bytesRead)

		if bytesRead == 0 {
			break
		}

		totalBytes += bytesRead
	}

	outFile.Write(buffer[:totalBytes])

	ioutil.WriteFile("resources/Dumps/010/game_pkg_gcdictionary.dict_uncompressed_body", buffer[0:totalBytes], 755)
}
