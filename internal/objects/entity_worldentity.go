package objects

import (
	"RainbowRunner/pkg"
	"RainbowRunner/pkg/byter"
)

type WorldEntity struct {
	*GCObject
	WorldPosition        pkg.Vector3
	Rotation             int
	WorldEntityFlags     uint32
	WorldEntityInitFlags byte
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
	b.WriteUInt32(
		n.WorldEntityFlags, // With this one alone it was working
	)
	// These positions stopped working at some point
	b.WriteInt32(n.WorldPosition.X) // Pos X
	b.WriteInt32(n.WorldPosition.Y) // Pos Y
	b.WriteInt32(n.WorldPosition.Z) // Pos Z
	b.WriteInt32(int32(n.Rotation))

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
		b.WriteUInt16(0x00)
	}

	if n.WorldEntityInitFlags&0x02 > 0 {
		// Ox02
		b.WriteByte(0xFF)
	}

	if n.WorldEntityInitFlags&0x04 > 0 {
		// 0x04
		b.WriteUInt32(0xFFFFFFFF)
	}

	if n.WorldEntityInitFlags&0x08 > 0 {
		// 0x08
		b.WriteUInt32(0xFFFFFFFF)
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
