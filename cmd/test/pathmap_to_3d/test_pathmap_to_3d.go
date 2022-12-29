package main

import (
	"RainbowRunner/cmd/rrcli/modelextractor"
	"RainbowRunner/internal/gosucks"
	"RainbowRunner/internal/types"
	"RainbowRunner/pkg/datatypes"
	"github.com/goccy/go-json"
	"os"
)

func main() {
	rawData, err := os.ReadFile("./data/pathmaps/town_pathmap.json")

	if err != nil {
		panic(err)
	}

	var pathMap types.PathMap

	err = json.Unmarshal(rawData, &pathMap)

	if err != nil {
		panic(err)
	}

	objBuilder := modelextractor.NewOBJBuilder()

	sizeX := pathMap.ChunkWidth
	sizeY := pathMap.ChunkHeight

	high, low := pathMap.GetHeightRange()
	heightRange := high - low

	objBuilder.WriteObject("Townston pathmap")

	for cx := 0; cx < sizeX; cx++ {
		for cy := 0; cy < sizeY; cy++ {

			nodes := pathMap.Nodes[cx+cy*sizeX]

			for x := 0; x < 16; x++ {
				for y := 0; y < 16; y++ {
					chunkX := ((sizeX - 1) - (cx)) * 16
					chunkY := cy * 16

					if nodes == nil || len(nodes) == 0 {
						continue
					}

					node := nodes[x+y*16]

					heightPercent := float64(node.Height-low) / float64(heightRange)

					objBuilder.WriteVert(datatypes.Vector3Float32{
						X: float32(chunkY + x),
						Y: node.Height / 10,
						Z: float32(chunkX - y),
					})

					gosucks.VAR(chunkY, chunkX, heightPercent)

					//img.Set(chunkX-y, chunkY-x, colour)
				}
			}
		}
	}

	obj := objBuilder.String()

	f, _ := os.OpenFile("tmp/pathmap.obj", os.O_CREATE, 0755)
	f.Truncate(0)

	f.WriteString(obj)
	f.Close()
}
