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

	pathMap := pathfinding.LoadPathMap("dungeon00_level01")
	//highlight := datatypes.Vector2Float32{X: -200, Y: 0} // dungeon00_level01 (deepfriedextracrispy)
	highlight := datatypes.Vector2Float32{X: -360, Y: 50} // dungeon00_level01 in hook (deepfriedextracrispy)

	//pathMap := pathfinding.LoadPathMap("town")
	//highlight := datatypes.Vector2Float32{X: 31, Y: -108} // town bridge
	//highlight := datatypes.Vector2Float32{X: 415, Y: -180} // town
	//highlight := datatypes.Vector2Float32{X: -140, Y: -20} // town behind mechanic

	img := image.NewRGBA(image.Rectangle{
		Max: image.Point{
			X: pathMap.ChunkWidth * 16,
			Y: pathMap.ChunkHeight * 16,
		},
	})

	/*.Sub(datatypes.Vector2{
		X: 11,
		Y: 40,
	})*/

	highlightGrid := pathMap.WorldPosToGridCoords(highlight)

	// This works in town
	//highlightGrid.X -= int32(pathMap.WorldOffsetX/10) + 1
	//highlightGrid.Y -= int32(pathMap.WorldOffsetZ * 2 / 10)

	// town
	//highlightGrid.X -= 73
	//highlightGrid.Y -= 9

	// dungeon00_level01
	//highlightGrid.X -= 10
	//highlightGrid.Y -= 40

	//highlightGrid.X -= int32(pathMap.WorldOffsetX/256) * 10
	//highlightGrid.Y -= int32(pathMap.WorldOffsetY/256) * 10

	high, low := pathMap.GetHeightRange()
	heightRange := high - low

	for x := -1000; x < 1500; x += 10 {
		for y := -1500; y < 1000; y += 10 {
			worldPos := datatypes.Vector3Float32{
				X: float32(x),
				Y: float32(y),
			}

			gridCoords := pathMap.WorldPosToGridCoords(worldPos.ToVector2Float32())

			node := pathMap.GetNode(gridCoords)

			if node != nil && (node.GridCoordX != 0 || node.GridCoordY != 0) {
				nodeRelativeCoords := pathMap.AbsoluteGridCoordsToRelative(datatypes.Vector2{
					X: int32(node.GridCoordX),
					Y: int32(node.GridCoordY),
				})
				calcPos := pathMap.GridCoordsToWorldPos(nodeRelativeCoords)

				nodePos := datatypes.Vector3Float32{
					X: node.WorldPosX,
					Y: node.WorldPosY,
					Z: node.Height,
				}

				fmt.Printf("pos %s, nPos %s -- diff %s\n", calcPos, nodePos, calcPos.Sub(nodePos))
				//fmt.Printf("Xl %d Xg %f,Yl %d Yg %f\n", node.GridCoordX, node.WorldPosX, node.GridCoordY, node.WorldPosY)
			}

			height := pathMap.HeightAtGridCoords(gridCoords)
			absoluteCoords := pathMap.GridCoordsToAbsolute(gridCoords)

			dist := gridCoords.Distance(highlightGrid)

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

	f, _ := os.OpenFile("tmp/pathmap_scan.png", os.O_CREATE, 0755)
	f.Truncate(0)

	png.Encode(f, img)
}
