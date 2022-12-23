package objects

//go:generate go run ../../scripts/generatelua -type=SkillProfession -extends=Component
type SkillProfession struct {
	*Component

	/**
	PropertySkillProfessionDescSlotID
	PropertySkillProfessionDescRequiredLevel
	PropertySkillProfessionDescCost
	PropertySkillProfessionDescCost
	*/
}

func NewSkillProfession(gcType string) *SkillProfession {
	component := NewComponent(gcType, "SkillProfession")

	return &SkillProfession{
		Component: component,
	}
}
