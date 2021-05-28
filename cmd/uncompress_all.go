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
	fileName := "E:\\Games\\DungeonRunners v666\\game.pkg"

	data, err := os.Open(fileName)
	offset := int64(0x00)

	if err != nil {
		panic(err)
	}

	for {
		_, err = data.Seek(offset, 0)
		offset++

		if err != nil {
			panic(err)
		}

		reader, err := zlib.NewReader(data)

		if err != nil && err.Error() == "unexpected EOF" {
			break
		}

		if err != nil {
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

			if bytesRead == 0 {
				break
			}

			totalBytes += bytesRead
		}

		if failed {
			continue
		}

		fmt.Printf("Read file at 0x%x, size: %d\n", offset, totalBytes)
		outFilename := fmt.Sprintf("D:\\Work\\dungeon-runners\\666 dumps\\gamepkg_%x", offset)

		ioutil.WriteFile(outFilename, buffer[0:totalBytes], 755)

		offset += int64(totalBytes)
	}
	fmt.Printf("Done\n")
}
