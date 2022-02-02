package main

import (
	"RainbowRunner/internal/objects"
	"fmt"
)

func main() {
	// town 7c9e936d
	// Townston 7a58db1
	// terrain.city.roads.city_Floor_40_1 5ba75e03
	str := "terrain.city.roads.city_Floor_40_1"
	hash := objects.GetTypeHash(str)
	fmt.Printf("%x\n", hash)
}
