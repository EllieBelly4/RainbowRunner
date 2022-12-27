package objects

import "RainbowRunner/pkg/byter"

//go:generate go run ../../scripts/generatelua -type=Skill -extends=Manipulator
type Skill struct {
	*Manipulator
	Level        byte
	OriginalSlot int
	SkillFlags   byte

	/**
	PropertySkillLevel
	PropertySkillDescRequiredLevelInc
	PropertySkillDescProfessionType
	PropertySkillDescElementType
	PropertySkillDescElementType
	PropertySkillDescRequiresTrainer
	PropertySkillDescAdjustCooldownByWeapon
	PropertySkillDescRequiredLevel
	PropertySkillDescRequiredLevel
	PropertySkillDescRequiredLevelInc
	PropertySkillDescMaxSkillLevel
	PropertySkillDescMaxSkillLevel
	PropertySkillDescSlotType
	PropertySkillDescSlotType
	PropertySkillDescGoldValueMod
	PropertySkillDescGoldValueMod
	PropertySkillDescManaCostMod
	PropertySkillDescManaCostMod
	PropertySkillDescCoolDown
	*/
}

func (s *Skill) WriteInit(b *byter.Byter) {
	b.WriteByte(s.SkillFlags) // Unk Flags

	if s.SkillFlags&0x01 != 0 {
		b.WriteUInt16(0x00) // Unk
	}

	if s.SkillFlags&0x02 != 0 {
		b.WriteUInt16(0x00) // Unk
	}

	if s.SkillFlags&0x04 != 0 {
		b.WriteUInt32(0x00) // Unk
	}

	if s.SkillFlags&0x08 != 0 {
		b.WriteUInt16(0x00) // Unk
	}

	if s.SkillFlags&0x10 != 0 {
		b.WriteUInt16(0x00) // Unk
	}

	// .text:00539C8B if something is set then write this
	//b.WriteUInt16(0x00) // Unk
}

func (s *Skill) WriteData(b *byter.Byter) {
	b.WriteUInt32(uint32(s.Slot))
	b.WriteByte(s.Level)
}

func NewSkill(gcType string) *Skill {
	manipulator := NewManipulator(gcType, "Skill")

	return &Skill{
		Manipulator: manipulator,
	}
}
