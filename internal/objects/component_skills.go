package objects

import (
	"RainbowRunner/pkg/byter"
)

type Skills struct {
	Component
	*GCObject
}

func (n *Skills) WriteInit(b *byter.Byter) {
	// Skills::readInit()
	b.WriteUInt32(0xFFFFFFFF)

	// GCObject::readChildData<Skill>
	b.WriteByte(0x04) // Count

	b.WriteByte(0xFF)
	b.WriteCString("skills.generic.Butcher")
	b.WriteUInt32(0x02)
	b.WriteByte(0x03) // Level

	b.WriteByte(0xFF)
	b.WriteCString("skills.generic.Stomp")
	b.WriteUInt32(0x04)
	b.WriteByte(0x05) // Level

	b.WriteByte(0xFF)
	b.WriteCString("skills.generic.FighterClassPassive")
	b.WriteUInt32(0x06)
	b.WriteByte(0x07) // Level

	b.WriteByte(0xFF)
	b.WriteCString("skills.generic.MeleeAttackSpeedModPassive")
	b.WriteUInt32(0x08)
	b.WriteByte(0x09) // Level

	// GCObject::readChildData<SkillProfession>
	b.WriteByte(0x01)
	b.WriteByte(0xFF)
	b.WriteCString("skills.professions.Warrior")
}

func NewSkills(gcType string) *Skills {
	gcObject := NewGCObject("Skills")
	gcObject.GCType = gcType

	return &Skills{
		GCObject: gcObject,
	}
}
