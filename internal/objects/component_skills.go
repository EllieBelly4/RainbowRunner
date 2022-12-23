package objects

import (
	"RainbowRunner/pkg/byter"
)

//go:generate go run ../../scripts/generatelua -type=Skills -extends=Component
type Skills struct {
	*Component

	/*
		PropertySkillsSkillPoints
	*/
}

func (n *Skills) WriteInit(b *byter.Byter) {
	// Skills::readInit()
	b.WriteUInt32(0xFFFFFFFF)

	skills := n.GetChildrenByGCNativeType("ActiveSkill")
	skills = append(skills, n.GetChildrenByGCNativeType("PassiveSkill")...)
	skills = append(skills, n.GetChildrenByGCNativeType("ActivePassiveSkill")...)

	// GCObject::readChildData<Skill>
	b.WriteByte(byte(len(skills))) // Count

	for _, skillDRObject := range skills {
		skill := skillDRObject.(ISkill).GetSkill()

		b.WriteByte(0xFF)
		b.WriteCString(skill.GCType)
		b.WriteUInt32(skill.Unk0)
		b.WriteByte(skill.Level) // Level
	}

	//b.WriteByte(0xFF)
	//b.WriteCString("skills.generic.Butcher")
	//b.WriteUInt32(0x02)
	//b.WriteByte(0x03) // Level
	//
	//b.WriteByte(0xFF)
	//b.WriteCString("skills.generic.Stomp")
	//b.WriteUInt32(0x04)
	//b.WriteByte(0x05) // Level
	//
	//b.WriteByte(0xFF)
	//b.WriteCString("skills.generic.FighterClassPassive")
	//b.WriteUInt32(0x06)
	//b.WriteByte(0x07) // Level
	//
	//b.WriteByte(0xFF)
	//b.WriteCString("skills.generic.MeleeAttackSpeedModPassive")
	//b.WriteUInt32(0x08)
	//b.WriteByte(0x09) // Level

	// GCObject::readChildData<SkillProfession>
	b.WriteByte(0x01)
	b.WriteByte(0xFF)
	b.WriteCString("skills.professions.Warrior")
}

func NewSkills(gcType string) *Skills {
	component := NewComponent(gcType, "Skills")

	return &Skills{
		Component: component,
	}
}
