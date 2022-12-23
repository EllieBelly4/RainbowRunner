package objects

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

func NewActiveSkill(gcType string) *ActiveSkill {
	skill := NewSkill(gcType)
	skill.GCNativeType = "ActiveSkill"

	return &ActiveSkill{
		Skill: skill,
	}
}
