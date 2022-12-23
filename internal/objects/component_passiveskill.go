package objects

//go:generate go run ../../scripts/generatelua -type=PassiveSkill -extends=Skill
type PassiveSkill struct {
	*Skill
}

func NewPassiveSkill(gcType string) *PassiveSkill {
	skill := NewSkill(gcType)
	skill.GCNativeType = "PassiveSkill"

	return &PassiveSkill{
		Skill: skill,
	}
}
