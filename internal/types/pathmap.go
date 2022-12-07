package types

import "math"

type PathMap struct {
	Width  int      `json:"width"`
	Height int      `json:"height"`
	Nodes  [][]Node `json:"nodes"`
}

type Node struct {
	Solid  bool    `json:"solid"`
	Height float32 `json:"height"`
}

func (p PathMap) GetHeightRange() (highest float32, lowest float32) {
	highest = math.MaxFloat32
	lowest = math.SmallestNonzeroFloat32

	nodes := p.Nodes

	for cx := 0; cx < p.Width; cx++ {
		for cy := 0; cy < p.Height; cy++ {

			subNodes := nodes[cx+cy*p.Width]

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
