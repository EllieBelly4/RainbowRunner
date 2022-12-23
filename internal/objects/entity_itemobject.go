package objects

import (
	"RainbowRunner/internal/types/drobjecttypes"
	"RainbowRunner/pkg/byter"
)

//go:generate go run ../../scripts/generatelua -type=ItemObject -extends=WorldEntity
type ItemObject struct {
	*WorldEntity
	Item drobjectypes.DRObject
}

func (n *ItemObject) Type() drobjectypes.DRObjectType {
	return drobjectypes.DRObjectEntity
}

func (n *ItemObject) WriteInit(b *byter.Byter) {
	n.WorldEntity.WriteInit(b)

	//ItemObject::readInit
	b.WriteByte(0x11)
	b.WriteUInt16(0x2233)
	b.WriteUInt32(0x00000000) // If this is not 0 then it reads a string

	// String here if above is 0

	b.WriteInt32(int32(n.WorldPosition.X * 256))
	b.WriteInt32(int32(n.WorldPosition.Y * 256))
	b.WriteByte(0xBA)

	// At some point here it expects a GCClass of type Item, Manipulator

	n.Item.WriteInit(b)
}

func NewItemObject(gcType string, item drobjectypes.DRObject) *ItemObject {
	worldEntity := NewWorldEntity(gcType)

	return &ItemObject{
		WorldEntity: worldEntity,
		Item:        item,
	}
}
