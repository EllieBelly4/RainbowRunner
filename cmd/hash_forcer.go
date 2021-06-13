package main

import (
	"fmt"
	"os"
)

//const chars =/**/ "1234567890-abcdefghijklmnopqrstuvwxyz"

const chars = "abcdefghijklmnopqrstuvwxyz"

//const chars = "01"

var maxLen = 12
var currentLength = 1
var crackHash = uint32(0xdc888300)
var done = make(chan bool)
var alive = 0

//var crackHash = uint32(0xb88c9fb)

func main() {
	str := make([]byte, maxLen)

	for ; currentLength < maxLen-1; currentLength++ {
		fmt.Printf("Starting Length %d\n", currentLength+1)
		for _, c := range chars {
			cStr := make([]byte, maxLen)
			copy(cStr, str)
			cStr[0] = byte(c)
			alive++
			go brute(cStr, 1)
		}
		<-done
	}

	os.Exit(1)
}

func brute(str []byte, i int) {
	for _, c := range chars {
		str[i] = byte(c)

		if i < currentLength {
			brute(str, i+1)
		}

		if currentLength > maxLen {
			return
		}

		if i == currentLength {
			hash := GetTypeHash(string(str))
			if hash == crackHash {
				fmt.Printf("%x (%s)\n", hash, string(str))
				return
			}
		}
	}

	if i == 1 {
		alive--
		if alive == 0 {
			done <- true
		}
	}
}

func GetTypeHash(name string) uint32 {
	result := uint32(5381) // eax

	a1 := len(name)

	if a1 > 0 {
		for _, v4 := range name {
			if v4 < 0x41 || v4 > 0x5A {
			} else {
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
