package objects

import (
	"RainbowRunner/pkg/byter"
	"RainbowRunner/pkg/datatypes"
)

type IUnit interface {
	GetUnit() *Unit
}

type Unit struct {
	*WorldEntity
	CurrentHP int
	UnitFlags byte
}

func (n *Unit) GetUnit() *Unit {
	return n
}

func (n *Unit) WriteInit(b *byter.Byter) {
	n.WorldEntity.WriteInit(b)

	// Unit::readInit()
	// Next 4 values always used
	// Same flag as above? + has extras
	// 0x01 - has parent/player owner?
	// 0x02 - add HP
	// 0x04 -
	//b.WriteByte(0x07) // HasParent + Unk
	//n.UnitFlags := 0x01 | 0x02 | 0x04 | 0x10 | 0x20 | 0x40 | 0x80
	b.WriteByte(n.UnitFlags)
	b.WriteByte(50) // Level
	b.WriteUInt16(0x01)
	b.WriteUInt16(0x02)

	if n.UnitFlags&0x01 > 0 {
		if n.RREntityProperties().OwnerID != 0 {
			b.WriteUInt16(uint16(Players.Players[n.RREntityProperties().OwnerID].CurrentCharacter.RREntityProperties().ID)) // Parent ID!!!!!
		} else {
			b.WriteUInt16(0x00) // Parent ID!!!!!
		}
	}

	if n.UnitFlags&0x02 > 0 {
		n.CurrentHP = 1150 * 256
		// 0x02 case
		// Multiply HP by 256
		b.WriteUInt32(uint32(n.CurrentHP)) // Current HP
	}

	if n.UnitFlags&0x04 > 0 {
		// 0x04 case
		// Multiply MP by 256
		b.WriteUInt32(505 * 256) // MP
	}

	if n.UnitFlags&0x010 > 0 {
		// 0x10 case
		b.WriteByte(0x04) // Unk
	}

	if n.UnitFlags&0x020 > 0 {
		// 0x20 case
		b.WriteUInt16(0x01) // Entity ID, Includes a call to IsKindOf<EncounterObject,Entity>(Entity *)
	}

	if n.UnitFlags&0x040 > 0 {
		// 0x40 case
		b.WriteUInt16(0x02) // Unk
		b.WriteUInt16(0x03) // Unk
		b.WriteUInt16(0x04) // Unk
		b.WriteByte(0x02)
	}

	if n.UnitFlags&0x080 > 0 {
		//0x80 case
		b.WriteByte(0x05)
	}
}

func (n *Unit) Warp(pos datatypes.Vector3Float32) {
	unitBehavior := n.GetChildByGCNativeType("UnitBehavior").(*UnitBehavior)
	unitBehavior.Warp(pos.X, pos.Y, pos.Z)
}

func NewUnit(gcType string) *Unit {
	worldEntity := NewWorldEntity(gcType)
	worldEntity.GCType = gcType

	return &Unit{
		WorldEntity: worldEntity,
		UnitFlags:   0x01 | 0x02 | 0x04,
	}
}
