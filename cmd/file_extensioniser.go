package main

import (
	byter "RainbowRunner/pkg/byter"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"os"
	"path/filepath"
)

func main() {
	root := "D:\\Work\\dungeon-runners\\666 dumps"

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if len(filepath.Ext(path)) > 0 {
			fmt.Printf("%s already has extension\n", path)
			return nil
		}

		ext, err := GetExtensionForFile(path)

		if err != nil {
			if err.Error() == "cannot determine type" {
				fmt.Printf("cannot determine type for %s\n", path)
			}
			return nil
		}

		newPath := fmt.Sprintf("%s.%s", path, ext)

		fmt.Printf("%s\n", newPath)

		err = os.Rename(path, newPath)

		if err != nil {
			panic(err)
		}

		return nil
	})
}

func GetExtensionForFile(path string) (ext string, err error) {
	file, err := os.Open(path)

	if err != nil {
		return "", err
	}

	defer file.Close()

	buf := make([]byte, 1024)

	read, err := file.Read(buf)

	if err != nil {
		return "", err
	}

	b := byter.NewByter(buf)

	if b.CompareString("DDS ") {
		return "dds", nil
	}

	dumpLen := int(math.Min(float64(read), 128))

	fmt.Printf("%s\n", hex.Dump(buf[:dumpLen]))

	return "", errors.New("cannot determine type")
}
