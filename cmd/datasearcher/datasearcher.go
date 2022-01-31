package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dir := "D:\\Work\\dungeon-runners\\666 dumps new"

	toSearch, err := hex.DecodeString("9ee171b4")

	if err != nil {
		panic(err)
	}

	toSearch = []byte("Townston_tier_1")

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		file, err := os.Open(path)

		if err != nil {
			panic(err)
		}

		stat, err := file.Stat()

		if err != nil {
			panic(err)
		}

		if stat.IsDir() {
			file.Close()
			return nil
		}

		data, err := ioutil.ReadAll(file)

		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		if strings.Contains(string(data), string(toSearch)) {
			fmt.Printf("found in %s\n", path)
		}

		file.Close()

		return nil
	})
}
