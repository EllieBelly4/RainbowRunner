package objects

//go:generate go run ../../scripts/generatelua -type=ActivePassiveSkill -extends=Skill
type ActivePassiveSkill struct {
	*Skill

	/*
		PropertyActivePassiveSkillActive
		PropertyActivePassiveSkillModifierID
		PropertyActivePassiveSkillCoolDownTimer
	*/
}

func NewActivePassiveSkill(gcType string) *ActivePassiveSkill {
	skill := NewSkill(gcType)
	skill.GCNativeType = "ActivePassiveSkill"

	return &ActivePassiveSkill{
		Skill: skill,
	}
}
