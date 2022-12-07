package types

import (
	"RainbowRunner/pkg/datatypes"
	"math"
)

type PathMap struct {
	CoordLimitX int          `json:"coordLimitX"`
	CoordLimitY int          `json:"coordLimitY"`
	TileWidth   float32      `json:"tileWidth"`
	TileHeight  float32      `json:"tileHeight"`
	ChunkWidth  int          `json:"width"`
	ChunkHeight int          `json:"height"`
	Nodes       [][]PathNode `json:"nodes"`
}

type PathNode struct {
	Solid  bool    `json:"solid"`
	Height float32 `json:"height"`
}

// WorldPosToGridCoords Converts world position to a pathmap grid coord
// WARNING THIS CURRENTLY CONTAINS FUNKY OFFSETS TO THE RESULT TO MAKE TOWNSTON WORK
// IF YOU TRY TO USE THIS FOR SOMETHING ELSE BEFORE FIXING THINGS IT WILL PROBABLY BE WRONG
func (p PathMap) WorldPosToGridCoords(pos datatypes.Vector3Float32) datatypes.Vector2 {
	return datatypes.Vector2{
		// TODO figure these offsets out
		// The two following offsets to X and Y were found to be reasonably accurate in Townston
		// They probably do not work in other zones, but I do not know where these come from or where I'm
		// calculating the position incorrectly
		X: int32((pos.X-10*p.TileWidth)/10) - 72,
		Y: int32((pos.Y-10*p.TileHeight)/10) - 8,
	}
}

func (p PathMap) GetHeightRange() (highest float32, lowest float32) {
	highest = math.SmallestNonzeroFloat32
	lowest = math.MaxFloat32

	nodes := p.Nodes

	for cx := 0; cx < p.ChunkWidth; cx++ {
		for cy := 0; cy < p.ChunkHeight; cy++ {

			subNodes := nodes[cx+cy*p.ChunkWidth]

			//if nodes == nil {
			//	continue
			//}

			for x := 0; x < 16; x++ {
				for y := 0; y < 16; y++ {
					if subNodes == nil || len(subNodes) == 0 {
						continue
					}

					node := subNodes[x+y*16]

					if node.Height > highest {
						highest = node.Height
					}

					if node.Height < lowest {
						lowest = node.Height
					}
				}
			}
		}
	}

	return highest, lowest
}

func (p PathMap) HeightAt(position datatypes.Vector3Float32) float32 {
	//position.X -= 737
	//position.Y -= 95
	gridCoords := p.WorldPosToGridCoords(position)
	node := p.getPathNode(gridCoords)

	if node == nil {
		//log.Warningf("position is not on path map: %f, %f, %f (%d,%d)",
		//	position.X,
		//	position.Y,
		//	position.Z,
		//	gridCoords.X,
		//	gridCoords.Y,
		//)
		return 0
	}

	//log.Infof("%d,%d height: %f", gridCoords.X, gridCoords.Y, node.Height)

	return node.Height
}

func (p PathMap) GridCoordsToAbsolute(coords datatypes.Vector2) datatypes.Vector2 {
	return datatypes.Vector2{
		X: coords.X + int32(p.ChunkWidth*16)/2,
		Y: coords.Y + int32(p.ChunkHeight*16)/2,
	}
}

func (p PathMap) getPathNode(coords datatypes.Vector2) *PathNode {
	absCoords := p.GridCoordsToAbsolute(coords)

	//ccX := ()

	//chunkX := (p.ChunkWidth - 1) - int(absCoords.X/16)
	chunkX := int(absCoords.X / 16)
	chunkY := int(absCoords.Y / 16)

	chunkIndex := chunkY + chunkX*p.ChunkHeight

	if chunkIndex >= len(p.Nodes) {
		return nil
	}

	nodes := p.Nodes[chunkIndex]

	if len(nodes) == 0 {
		return nil
	}

	remainderX := absCoords.X % 16
	remainderY := absCoords.Y % 16
	innerIndex := int(remainderX + remainderY*16)

	if innerIndex > len(nodes) {
		return nil
	}

	node := nodes[innerIndex]

	return &node
}
