package objects

import (
	"RainbowRunner/internal/types/drobjecttypes"
	"RainbowRunner/pkg/byter"
	"RainbowRunner/pkg/datatypes"
)

//go:generate go run ../../scripts/generateLua/ -type=Unit -extends=WorldEntity
type Unit struct {
	*WorldEntity
	HP        int
	MP        int
	UnitFlags byte
	Level     byte

	Unk10Case         byte
	Unk20CaseEntityID uint16
	Unk40Case0        uint16
	Unk40Case1        uint16
	Unk40Case2        uint16
	Unk40Case3        byte
	Unk80Case         byte
	UnitUnkUint16_0   uint16
	UnitUnkUint16_1   uint16
}

func (u *Unit) AddChild(child drobjecttypes.DRObject) {
	u.WorldEntity.AddChild(child)
	child.SetParent(u)
}

func (u *Unit) WriteInit(b *byter.Byter) {
	u.WorldEntity.WriteInit(b)

	// Unit::readInit()
	// Next 4 values always used
	// Same flag as above? + has extras
	// 0x01 - has parent/player owner?
	// 0x02 - add HP
	// 0x04 -
	//b.WriteByte(0x07) // HasParent + Unk
	//u.UnitFlags := 0x01 | 0x02 | 0x04 | 0x10 | 0x20 | 0x40 | 0x80
	b.WriteByte(u.UnitFlags)
	b.WriteByte(u.Level) // Level
	b.WriteUInt16(u.UnitUnkUint16_0)
	b.WriteUInt16(u.UnitUnkUint16_1)

	if u.UnitFlags&0x01 > 0 {
		if u.RREntityProperties().OwnerID != 0 {
			b.WriteUInt16(uint16(Players.Players[int(u.RREntityProperties().OwnerID)].CurrentCharacter.RREntityProperties().ID)) // Parent ID!!!!!
		} else {
			b.WriteUInt16(0x00) // Parent ID!!!!!
		}
	}

	if u.UnitFlags&0x02 > 0 {
		//u.HP = 1150
		// 0x02 case
		// Multiply HP by 256
		b.WriteUInt32(uint32(u.HP) * 256) // Current HP
	}

	if u.UnitFlags&0x04 > 0 {
		// 0x04 case
		// Multiply MP by 256
		b.WriteUInt32(uint32(u.MP) * 256) // MP
	}

	if u.UnitFlags&0x010 > 0 {
		// 0x10 case
		b.WriteByte(u.Unk10Case) // Unk
	}

	if u.UnitFlags&0x020 > 0 {
		// 0x20 case
		b.WriteUInt16(u.Unk20CaseEntityID) // Entity ID, Includes a call to IsKindOf<EncounterObject,Entity>(Entity *)
	}

	if u.UnitFlags&0x040 > 0 {
		// 0x40 case
		b.WriteUInt16(u.Unk40Case0) // Unk
		b.WriteUInt16(u.Unk40Case1) // Unk
		b.WriteUInt16(u.Unk40Case2) // Unk
		b.WriteByte(u.Unk40Case3)
	}

	if u.UnitFlags&0x080 > 0 {
		//0x80 case
		b.WriteByte(u.Unk80Case)
	}
}

func (n *Unit) Warp(pos datatypes.Vector3Float32) {
	unitBehavior := n.GetChildByGCNativeType("UnitBehavior").(*UnitBehavior)
	unitBehavior.Warp(pos.X, pos.Y, pos.Z)
}

func NewUnit(gcType string) *Unit {
	worldEntity := NewWorldEntity(gcType)
	worldEntity.GCType = gcType

	//worldEntity.WorldEntityFlags = 0x07
	//worldEntity.WorldEntityInitFlags

	return &Unit{
		Level:       1,
		HP:          100,
		MP:          100,
		WorldEntity: worldEntity,
		UnitFlags:   0x01 | 0x02 | 0x04 | 0x20,
	}
}
