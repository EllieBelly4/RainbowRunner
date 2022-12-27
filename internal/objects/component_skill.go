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
