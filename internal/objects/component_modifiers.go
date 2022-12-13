package objects

import (
	"RainbowRunner/pkg/byter"
)

//go:generate go run ../../scripts/generatelua -type=Modifiers -extends=Component
type Modifiers struct {
	*Component
}

func (n *Modifiers) WriteInit(b *byter.Byter) {
	// Modifiers::readInit
	b.WriteUInt32(0x00) //
	b.WriteUInt32(0x00) //

	// GCObject::readChildData<Modifier>
	b.WriteByte(0x00)
}

func NewModifiers(gcType string) *Modifiers {
	component := NewComponent(gcType, "Modifiers")

	return &Modifiers{
		Component: component,
	}
}
