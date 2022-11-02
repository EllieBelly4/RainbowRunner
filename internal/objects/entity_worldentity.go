package objects

import (
	"RainbowRunner/pkg/byter"
	"RainbowRunner/pkg/datatypes"
)

type IWorldEntity interface {
	GetWorldEntity() *WorldEntity
}

type WorldEntity struct {
	*GCObject
	WorldPosition        datatypes.Vector3Float32
	Rotation             float32
	WorldEntityFlags     uint32
	WorldEntityInitFlags byte
	Label                string

	Unk1Case uint16
	Unk2Case byte
	Unk4Case uint32
	Unk8Case uint32
}

func (n *WorldEntity) GetWorldEntity() *WorldEntity {
	return n
}

func (n *WorldEntity) SetPosition(position datatypes.Vector3Float32) {
	n.WorldPosition = position
}

func (n *WorldEntity) SetRotation(degrees float32) {
	n.Rotation = degrees
}

func (n *WorldEntity) Type() DRObjectType {
	return DRObjectEntity
}

func (n *WorldEntity) WriteInit(b *byter.Byter) {
	//WorldEntity::readInit
	// Flags
	// 0x01 Static object?
	// 0x02 Unk
	// 0x04 Makes character appear
	// 0x08 Unk
	// 0x10 Unk
	// 0x20 Unk
	// 0x40 Unk
	// 0x80 Unk
	// 0x100 Unk
	// 0x200 Unk
	// 0x400 Unk
	// 0x800 Breaks everything
	// 0x1000 Makes the character invisible
	// 0x2000 Makes movement very jumpy
	// 0x4000 Unk
	// 0x8000 Unk
	// 0x10000 Unk
	// One of these flags stops the below positions from working
	// With only 0x04 the character can be moved and is the least broken
	// 0x07 is the least required to get NPCs working
	b.WriteUInt32(
		n.WorldEntityFlags, // With this one alone it was working
	)
	// These positions stopped working at some point
	b.WriteInt32(int32(n.WorldPosition.X * 256)) // Pos X
	b.WriteInt32(int32(n.WorldPosition.Y * 256)) // Pos Y
	b.WriteInt32(int32(n.WorldPosition.Z * 256)) // Pos Z
	b.WriteInt32(int32(n.Rotation * 256))

	// Flags
	// Each flag adds one more section of data to read sequentially
	// 0x01 Has Parent?
	// 0x02 Unk
	// 0x04 Makes movement smoother, interpolated position?
	// 0x08 Unk
	//b.WriteByte(1 | 2 | 4 | 8)
	// When this is set to 0 the character is slightly less broken
	// With 1 | 2 | 4 | 8 it was causing the character to have no animations and
	// eventually collapse into itself
	//n.WorldEntityInitFlags := 0x04 | 0x08
	b.WriteByte(byte(n.WorldEntityInitFlags))

	if n.WorldEntityInitFlags&0x01 > 0 {
		// 0x01
		b.WriteUInt16(n.Unk1Case)
	}

	if n.WorldEntityInitFlags&0x02 > 0 {
		// Ox02
		b.WriteByte(n.Unk2Case)
	}

	if n.WorldEntityInitFlags&0x04 > 0 {
		// 0x04
		b.WriteUInt32(n.Unk4Case)
	}

	if n.WorldEntityInitFlags&0x08 > 0 {
		// 0x08
		b.WriteUInt32(n.Unk8Case)
	}
}

func NewWorldEntity(gcType string) *WorldEntity {
	gcObject := NewGCObject(gcType)
	gcObject.GCType = gcType

	return &WorldEntity{
		GCObject:             gcObject,
		WorldEntityFlags:     0x04,
		WorldEntityInitFlags: 0xFF,
	}
}
