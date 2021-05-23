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
	//fileName := "resources/Dumps/010/game_pkg_gcdictionary.dict_compressed_zlib_deflate_best_compression"
	fileName := "resources/Dumps/010/game_pkg_gcdictionary.dict_header_compressed"

	data, err := os.Open(fileName)
	offset := int64(0x00)

	if err != nil {
		panic(err)
	}

	for {
		_, err = data.Seek(offset, 0)

		if err != nil {
			panic(err)
		}

		fmt.Printf("%x\n", offset)

		reader, err := zlib.NewReader(data)

		if err != nil && err.Error() == "unexpected EOF" {
			break
		}

		if err != nil {
			offset++
			continue
		}

		buffer := make([]byte, 1024*10000)

		var bytesRead = 0
		var totalBytes = 0

		failed := false

		for {
			bytesRead, err = reader.Read(buffer[totalBytes:])

			if err != nil && !errors.Is(err, io.EOF) {
				failed = true
				break
			}

			fmt.Printf("Read %d uncompressed bytes\n", bytesRead)

			if bytesRead == 0 {
				break
			}

			totalBytes += bytesRead
		}

		if failed {
			offset++
			continue
		}

		outFilename := fmt.Sprintf("resources/Dumps/010/game_pkg_gcdictionary.dict_header_uncompressed%x", offset)

		ioutil.WriteFile(outFilename, buffer[0:totalBytes], 755)
		break
	}
}
