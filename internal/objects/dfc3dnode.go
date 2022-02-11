package objects

import (
	"RainbowRunner/internal/types"
	"RainbowRunner/pkg/byter"
	"math"
)

type DFC3DNode struct {
	*GCObject
	Matrix   types.Matrix324x4
	UnkFlags uint32
}

func (d *DFC3DNode) ReadData(b *byter.Byter) {
	d.UnkFlags = b.UInt32()
	d.Matrix = types.Matrix324x4{
		Values: [16]float32{
			b.Float32(),
			b.Float32(),
			b.Float32(),
			b.Float32(),

			b.Float32(),
			b.Float32(),
			b.Float32(),
			b.Float32(),

			b.Float32(),
			b.Float32(),
			b.Float32(),
			b.Float32(),

			b.Float32(),
			b.Float32(),
			b.Float32(),
			b.Float32(),
		},
	}

	for i, value := range d.Matrix.Values {
		if math.IsNaN(float64(value)) {
			d.Matrix.Values[i] = 0
		}
	}

	b.Float32() // Unk
}

func NewDFC3DNode() *DFC3DNode {
	return &DFC3DNode{
		GCObject: NewGCObject("DFC3DNode"),
	}
}
