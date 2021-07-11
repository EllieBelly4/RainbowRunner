package objects

import "RainbowRunner/pkg/byter"

type NPC struct {
	*GCObject
}

func (n *NPC) WriteInit(b *byter.Byter) {
}

func NewNPC(gcType string) *NPC {
	gcObject := NewGCObject("unk")
	gcObject.GCType = gcType

	return &NPC{
		GCObject: gcObject,
	}
}
