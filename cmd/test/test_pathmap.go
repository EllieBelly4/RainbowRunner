package main

import (
	"RainbowRunner/internal/types"
	"encoding/json"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	rawData, err := os.ReadFile("./data/pathmaps/tutorial_pathmap.json")

	if err != nil {
		panic(err)
	}

	var pathMap types.PathMap

	err = json.Unmarshal(rawData, &pathMap)

	if err != nil {
		panic(err)
	}

	sizeX := pathMap.ChunkWidth
	sizeY := pathMap.ChunkHeight

	img := image.NewRGBA(image.Rectangle{
		Max: image.Point{
			X: sizeX * 16,
			Y: sizeY * 16,
		},
	})

	high, low := pathMap.GetHeightRange()
	heightRange := high - low

	for cx := 0; cx < sizeX; cx++ {
		for cy := 0; cy < sizeY; cy++ {
			nodes := pathMap.Nodes[cx+cy*sizeX]

			//if nodes == nil {
			//	continue
			//}

			for x := 0; x < 16; x++ {
				for y := 0; y < 16; y++ {
					//chunkX := cx * 16
					chunkX := ((sizeX - 1) - (cx)) * 16
					chunkY := cy * 16
					//chunkY := ((sizeY - 1) - (cy)) * 16

					colour := color.RGBA{R: 255, G: 255, B: 255, A: 255}

					px := chunkY + x
					py := chunkX - y

					if nodes == nil || len(nodes) == 0 {
						img.Set(px, py, colour)
						continue
					}

					node := nodes[x+y*16]

					heightPercent := float64(node.Height-low) / float64(heightRange)
					heightColour := uint8(255 * heightPercent)

					//offsetIndex :=
					//	cx*16 + // Chunk X * Chunk Pixels(16)
					//		x + // Inner X
					//		cy*sizeX*16*16 + // Chunk Y * Chunk Count X * Full Chunk Pixels(16*16)
					//		y*sizeX*16 // Inner Y * Chunk Count X * Chunk Pixels(16)

					if node.Solid {
						colour = color.RGBA{R: heightColour, G: heightColour, B: heightColour, A: 255}
					}

					img.Set(px, py, colour)
				}
			}
		}
	}

	f, _ := os.OpenFile("tmp/pathmap.png", os.O_CREATE, 0755)
	f.Truncate(0)

	png.Encode(f, img)
}
