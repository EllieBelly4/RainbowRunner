package objects

import (
	"RainbowRunner/pkg/byter"
)

type ItemObject struct {
	*WorldEntity
	Item DRObject
}

func (n *ItemObject) Type() DRObjectType {
	return DRObjectEntity
}

func (n *ItemObject) WriteInit(b *byter.Byter) {
	n.WorldEntity.WriteInit(b)

	//ItemObject::readInit
	b.WriteByte(0x11)
	b.WriteUInt16(0x2233)
	b.WriteUInt32(0x00000000) // If this is not 0 then it reads a string

	// String here if above is 0

	b.WriteInt32(n.WorldPosition.X)
	b.WriteInt32(n.WorldPosition.Y)
	b.WriteByte(0xBA)

	// At some point here it expects a GCClass of type Item, Manipulator

	n.Item.WriteInit(b)
}

func NewItemObject(gcType string, item DRObject) *ItemObject {
	worldEntity := NewWorldEntity(gcType)

	return &ItemObject{
		WorldEntity: worldEntity,
		Item:        item,
	}
}
