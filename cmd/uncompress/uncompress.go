package main

import (
	"compress/zlib"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fileName := "resources/Dumps/010/game_pkg_gcdictionary.dict_compressed_zlib_deflate_best_compression"
	//fileName := "resources/Dumps/010/game_pkg_gcdictionary.dict_header_compressed"

	data, err := os.Open(fileName)
	offset := int64(0x00)

	if err != nil {
		panic(err)
	}

	data.Seek(offset, 0)

	reader, err := zlib.NewReader(data)

	if err != nil {
		panic(err)
	}

	buffer := make([]byte, 1024*10000)

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

	ioutil.WriteFile("resources/Dumps/010/game_pkg_gcdictionary.dict_uncompressed_body", buffer[0:totalBytes], 755)
}
