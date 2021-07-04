package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	fh, err := os.Open("D:\\Work\\dungeon-runners\\666 dumps\\GCDictionary.txt")

	if err != nil {
		panic(err)
	}

	ofh, err := os.OpenFile("resources/Dumps/GCObjectHashes.txt", os.O_TRUNC|os.O_CREATE, 0644)

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(fh)

loop:
	for {
		str, err := reader.ReadString('\n')

		if err != nil {
			if errors.Is(err, io.EOF) {
				break loop
			}

			panic(err)
		}

		split := strings.Split(str, " ")

		fmt.Printf("%s %s", split[0], split[1])

		index, err := strconv.Atoi(split[0])

		if err != nil {
			panic(err)
		}

		cleanTypeName := split[1][:len(split[1])-2]
		hash := GetTypeHash(cleanTypeName)
		hashStr := fmt.Sprintf("%08X", hash)
		bigHash := reverseEndianness(hashStr)

		_, err = ofh.WriteString(
			fmt.Sprintf("%d %s\nLittleEndian: 0x%s\nBigEndian: 0x%s\n\n",
				index,
				cleanTypeName,
				hashStr,
				bigHash,
			))

		if err != nil {
			panic(err)
		}
	}
}

func reverseEndianness(str string) string {
	newStr := ""

	for i := 3; i >= 0; i-- {
		newStr += str[i*2 : i*2+2]
	}

	return newStr
}

func GetTypeHash(name string) uint32 {
	result := uint32(5381) // eax

	a1 := len(name)

	if a1 > 0 {
		for _, v4 := range name {
			if v4 >= 0x41 && v4 <= 0x5A {
				v4 = v4 + 32
			}

			result += uint32(v4) + 32*result
		}

		if result == 0 {
			result = 1
		}
	}

	return result
}
