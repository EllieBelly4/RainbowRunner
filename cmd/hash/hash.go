package main

import (
	"RainbowRunner/internal/objects"
	"fmt"
)

func main() {
	str := "town"
	hash := objects.GetTypeHash(str)
	fmt.Printf("%x\n", hash)
}
