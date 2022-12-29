package types

import (
	"RainbowRunner/pkg/datatypes"
	log "github.com/sirupsen/logrus"
	"math"
)

type PathMap struct {
	WorldOffsetX float32      `json:"worldOffsetX"`
	WorldOffsetY float32      `json:"worldOffsetY"`
	WorldOffsetZ float32      `json:"worldOffsetZ"`
	CoordLimitX  int          `json:"coordLimitX"`
	CoordLimitY  int          `json:"coordLimitY"`
	TileWidth    float32      `json:"tileWidth"`
	TileHeight   float32      `json:"tileHeight"`
	ChunkWidth   int          `json:"width"`
	ChunkHeight  int          `json:"height"`
	Nodes        [][]PathNode `json:"nodes"`
	Offset       datatypes.Vector3Float32
}

type PathNode struct {
	Solid      bool    `json:"solid"`
	Height     float32 `json:"height"`
	WorldPosX  float32 `json:"worldX"`
	WorldPosY  float32 `json:"worldY"`
	GridCoordX int     `json:"gridX"`
	GridCoordY int     `json:"gridY"`
}

// WorldPosToGridCoords Converts world position to a pathmap grid coord
func (p PathMap) WorldPosToGridCoords(pos datatypes.Vector2Float32) datatypes.Vector2 {
	offsetPos := pos.Add(p.Offset.ToVector2Float32())

	return datatypes.Vector2{
		X: int32((offsetPos.X - 10*p.TileWidth) / 10),
		Y: int32((offsetPos.Y - 10*p.TileHeight) / 10),
	}
}

// GridCoordsToWorldPos Converts a pathmap grid coord to a world position
func (p PathMap) GridCoordsToWorldPos(coords datatypes.Vector2) datatypes.Vector3Float32 {
	return datatypes.Vector3Float32{
		X: float32(coords.X)*10 + 10*p.TileWidth,
		Y: float32(coords.Y)*10 + 10*p.TileHeight,
		Z: p.HeightAtGridCoords(coords),
	}.Sub(p.Offset)
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

func (p PathMap) HeightAt(position datatypes.Vector2Float32) float32 {
	gridCoords := p.WorldPosToGridCoords(position)
	node := p.GetNode(gridCoords)

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

	return node.Height
}

func (p PathMap) GridCoordsToAbsolute(coords datatypes.Vector2) datatypes.Vector2 {
	return datatypes.Vector2{
		X: coords.X + int32(p.ChunkWidth*16)/2,
		Y: coords.Y + int32(p.ChunkHeight*16)/2,
	}
}

func (p PathMap) AbsoluteGridCoordsToRelative(coords datatypes.Vector2) datatypes.Vector2 {
	return datatypes.Vector2{
		X: coords.X - int32(p.ChunkWidth*16)/2,
		Y: coords.Y - int32(p.ChunkHeight*16)/2,
	}
}

func (p PathMap) GetNode(coords datatypes.Vector2) *PathNode {
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

func (p PathMap) HeightAtGridCoords(coords datatypes.Vector2) float32 {
	node := p.GetNode(coords)

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

	return node.Height
}

func (p *PathMap) Init() {
	for _, nodeList := range p.Nodes {
		if nodeList == nil {
			continue
		}

		for _, node := range nodeList {
			if node.GridCoordX == 0 && node.GridCoordY == 0 {
				continue
			}

			nodeRelativeCoords := p.AbsoluteGridCoordsToRelative(datatypes.Vector2{
				X: int32(node.GridCoordX),
				Y: int32(node.GridCoordY),
			})

			calcPos := p.GridCoordsToWorldPos(nodeRelativeCoords)

			nodePos := datatypes.Vector3Float32{
				X: node.WorldPosX,
				Y: node.WorldPosY,
				Z: node.Height,
			}

			p.Offset = calcPos.Sub(nodePos)
			return
		}
	}

	log.Errorf("could not find offset for path map")
}
