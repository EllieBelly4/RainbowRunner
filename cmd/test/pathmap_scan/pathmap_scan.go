package main

import (
	"RainbowRunner/internal/pathfinding"
	"RainbowRunner/pkg/datatypes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	//highlight := datatypes.Vector2{X: 357, Y: -180}
	//highlight := datatypes.Vector2{X: 1, Y: -203}
	//highlight := datatypes.Vector2{X: 190, Y: 325}
	//highlight := datatypes.Vector2{X: 26, Y: -112}
	highlight := datatypes.Vector2{X: 760, Y: 49}

	pathMap := pathfinding.LoadPathMap("town")

	img := image.NewRGBA(image.Rectangle{
		Max: image.Point{
			X: pathMap.ChunkWidth * 16,
			Y: pathMap.ChunkHeight * 16,
		},
	})

	highlightGrid := pathMap.WorldPosToGridCoords(highlight.ToVector3Float32())

	high, low := pathMap.GetHeightRange()
	heightRange := high - low

	for x := -1000; x < 1500; x++ {
		for y := -1500; y < 1000; y++ {
			worldPos := datatypes.Vector3Float32{
				X: float32(x),
				Y: float32(y),
			}

			height := pathMap.HeightAt(worldPos)

			gridCoord := pathMap.WorldPosToGridCoords(worldPos)
			absoluteCoords := pathMap.GridCoordsToAbsolute(gridCoord)

			dist := gridCoord.Distance(highlightGrid)

			heightPercent := float64(height-low) / float64(heightRange)
			bwVal := uint8(255 * heightPercent)
			colour := color.RGBA{R: bwVal, G: bwVal, B: bwVal, A: 255}

			if height == 0 {
				colour = color.RGBA{A: 255}
			}

			if dist < 0.5 {
				colour = color.RGBA{R: 255, A: 255}
			}

			img.Set(int(absoluteCoords.X), (pathMap.ChunkHeight*16-1)-int(absoluteCoords.Y), colour)
		}
	}

	fmt.Println(pathMap)
	f, _ := os.OpenFile("tmp/pathmap_scan.png", os.O_CREATE, 0755)
	f.Truncate(0)

	png.Encode(f, img)
}
