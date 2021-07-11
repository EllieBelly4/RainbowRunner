package objects

import (
	"RainbowRunner/pkg/byter"
)

type Manipulators struct {
	*GCObject
}

func (n *Manipulators) WriteInit(b *byter.Byter) {
	// Manipulators::readInit
	b.WriteByte(0x00) // Some count
}

func NewManipulators(gcType string) *Manipulators {
	gcObject := NewGCObject("Manipulators")
	gcObject.GCType = gcType

	return &Manipulators{
		GCObject: gcObject,
	}
}
