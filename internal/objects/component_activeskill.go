package objects

import "RainbowRunner/pkg/byter"

//go:generate go run ../../scripts/generatelua -type=ActiveSkill -extends=Skill
type ActiveSkill struct {
	*Skill

	/*
		PropertyActiveSkillCoolDownTimer
		PropertyActiveSkillDescSpellType
		PropertyActiveSkillDescSpellUse
		PropertyActiveSkillDescSpellUseRange
		PropertyActiveSkillDescSpellUseMinimumRange
		PropertyActiveSkillDescRange
		PropertyActiveSkillDescTargetType
		PropertyActiveSkillDescWeaponType
		PropertyActiveSkillDescRepeatAnimation
		PropertyActiveSkillDescRepeatCount
		PropertyActiveSkillDescAddModifierWhileClosing
		PropertyActiveSkillDescAnimationID
		PropertyActiveSkillDescAnimationLength
		PropertyActiveSkillDescTriggerTime
		PropertyActiveSkillDescAnimationOffset
		PropertyActiveSkillDescIsPrimaryAttack
		PropertyActiveSkillDescSelfHealthPct
		PropertyActiveSkillDescTargetHealthPct
	*/
}

func (s *ActiveSkill) WriteInit(b *byter.Byter) {
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

func NewActiveSkill(gcType string) *ActiveSkill {
	skill := NewSkill(gcType)
	skill.GCNativeType = "ActiveSkill"

	return &ActiveSkill{
		Skill: skill,
	}
}
