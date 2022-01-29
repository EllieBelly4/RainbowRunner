package main

import (
	"RainbowRunner/cmd/files"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	root := "D:\\Work\\dungeon-runners\\666 dumps new"

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if len(filepath.Ext(path)) > 0 {
			//fmt.Printf("%s already has extension\n", path)
			return nil
		}

		//if path != "D:\\Work\\dungeon-runners\\666 dumps new\\world_dungeon11_quest_Q02_a1" {
		//	return nil
		//}

		file, err := os.Open(path)

		if err != nil {
			panic(err)
		}

		data, err := ioutil.ReadAll(file)

		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		file.Close()

		_, ext := files.GetExtensionForFile(data, uint32(len(data)))

		if ext == "" {
			fmt.Printf("cannot determine type for %s\n", path)
			return nil
		}

		newPath := fmt.Sprintf("%s%s", path, ext)

		//fmt.Printf("%s -> %s\n", path, ext)

		err = os.Rename(path, newPath)

		if err != nil {
			panic(err)
		}

		return nil
	})
}

//
//func GetExtensionForFile(path string) (ext string, err error) {
//	file, err := os.Open(path)
//
//	if err != nil {
//		return "", err
//	}
//
//	defer file.Close()
//
//	buf := make([]byte, 1024)
//
//	read, err := file.Read(buf)
//
//	if err != nil {
//		return "", err
//	}
//
//	b := byter.NewByter(buf)
//
//	if b.CompareString("DDS ") {
//		return "dds", nil
//	}
//
//	dumpLen := int(math.Min(float64(read), 128))
//
//	fmt.Printf("%s\n", hex.Dump(buf[:dumpLen]))
//
//	return "", errors.New("cannot determine type")
//}
