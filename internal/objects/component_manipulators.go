package objects

import (
	"RainbowRunner/pkg/byter"
)

type Manipulators struct {
	*Component
}

func (n *Manipulators) WriteInit(b *byter.Byter) {
	// Manipulators::readInit
	b.WriteByte(0x00) // Some count
}

func NewManipulators(gcType string) *Manipulators {
	component := NewComponent(gcType, "Manipulators")

	return &Manipulators{
		Component: component,
	}
}
