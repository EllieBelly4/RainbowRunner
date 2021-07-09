package main

import (
	byter "RainbowRunner/pkg/byter"
	"fmt"
)

func main() {
	b := byter.NewLEByter([]byte{00, 01, 01, 00})

	f := b.Fixed32()

	fmt.Printf("%f", f)
}
