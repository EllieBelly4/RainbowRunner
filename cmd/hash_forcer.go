package main

import (
	"RainbowRunner/internal/objects"
	"fmt"
	"os"
)

//const chars =/**/ "1234567890-abcdefghijklmnopqrstuvwxyz"

const chars = "abcdefghijklmnopqrstuvwxyz"

//const chars = "01"

var maxLen = 12
var currentLength = 1
var crackHash = uint32(0x120ab043)
var beCrackHash = uint32(0)
var done = make(chan bool)
var alive = 0

//var crackHash = uint32(0xb88c9fb)

func main() {
	beCrackHash = ((crackHash & 0xFF000000) >> 24) | ((crackHash & 0x00FF0000) >> 8) | ((crackHash & 0x0000FF00) << 8) | ((crackHash & 0x000000FF) << 24)

	fmt.Printf("LE %x\nBE %x\n", crackHash, beCrackHash)

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
			hash := objects.GetTypeHash(string(str))
			if hash == crackHash || hash == beCrackHash {
				fmt.Printf("%x (%s)\n", hash, string(str))
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
