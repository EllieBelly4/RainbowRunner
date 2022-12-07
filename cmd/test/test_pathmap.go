package main

import (
	"encoding/json"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

type Node struct {
	Solid bool `json:"solid"`
	Unk0  int  `json:"unk0"`
}

func main() {
	rawData, err := os.ReadFile("./tmp/pathmap.json")

	if err != nil {
		panic(err)
	}

	var pathMap [][]Node

	err = json.Unmarshal(rawData, &pathMap)

	if err != nil {
		panic(err)
	}

	sizeX := 25
	sizeY := 25

	img := image.NewRGBA(image.Rectangle{
		Max: image.Point{
			X: sizeX * 16,
			Y: sizeY * 16,
		},
	})

	high, low := getHeightRange(sizeX, sizeY, pathMap)
	heightRange := high - low

	for cx := 0; cx < sizeX; cx++ {
		for cy := 0; cy < sizeY; cy++ {

			nodes := pathMap[cx+cy*sizeX]

			//if nodes == nil {
			//	continue
			//}

			for x := 0; x < 16; x++ {
				for y := 0; y < 16; y++ {
					chunkX := ((sizeX - 1) - (cx)) * 16
					chunkY := ((sizeY - 1) - (cy)) * 16

					colour := color.RGBA{R: 255, G: 255, B: 255, A: 255}

					if nodes == nil || len(nodes) == 0 {
						img.Set(chunkX-x, chunkY-y, colour)
						continue
					}

					node := nodes[x+y*16]

					heightPercent := float64(node.Unk0-low) / float64(heightRange)
					heightColour := uint8(255 * heightPercent)

					//offsetIndex :=
					//	cx*16 + // Chunk X * Chunk Pixels(16)
					//		x + // Inner X
					//		cy*sizeX*16*16 + // Chunk Y * Chunk Count X * Full Chunk Pixels(16*16)
					//		y*sizeX*16 // Inner Y * Chunk Count X * Chunk Pixels(16)

					if node.Solid {
						colour = color.RGBA{R: heightColour, G: heightColour, B: heightColour, A: 255}
					}

					img.Set(chunkX-y, chunkY-x, colour)
				}
			}
		}
	}

	f, _ := os.OpenFile("tmp/pathmap.png", os.O_CREATE, 0755)
	f.Truncate(0)

	png.Encode(f, img)
}

func getHeightRange(sizeX int, sizeY int, pathMap [][]Node) (highest int, lowest int) {
	highest = math.MinInt32
	lowest = math.MaxInt32

	for cx := 0; cx < sizeX; cx++ {
		for cy := 0; cy < sizeY; cy++ {

			nodes := pathMap[cx+cy*sizeX]

			//if nodes == nil {
			//	continue
			//}

			for x := 0; x < 16; x++ {
				for y := 0; y < 16; y++ {
					if nodes == nil || len(nodes) == 0 {
						continue
					}

					node := nodes[x+y*16]

					if node.Unk0 > highest {
						highest = node.Unk0
					}

					if node.Unk0 < lowest {
						lowest = node.Unk0
					}
				}
			}
		}
	}

	return highest, lowest
}
