package objects

import (
	"RainbowRunner/pkg/byter"
)

type Modifiers struct {
	Component
	*GCObject
}

func (n *Modifiers) WriteInit(b *byter.Byter) {
	// Modifiers::readInit
	b.WriteUInt32(0x00) //
	b.WriteUInt32(0x00) //

	// GCObject::readChildData<Modifier>
	b.WriteByte(0x00)
}

func NewModifiers(gcType string) *Modifiers {
	gcObject := NewGCObject("Modifiers")
	gcObject.GCType = gcType

	return &Modifiers{
		GCObject: gcObject,
	}
}
