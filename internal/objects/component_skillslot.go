package objects

//go:generate go run ../../scripts/generatelua -type=SkillSlot -extends=Component
type SkillSlot struct {
	*Component

	/*
		PropertySkillSlotSlotID
		PropertySkillSlotSlotType
	*/
}

func NewSkillSlot(gcType string) *SkillSlot {
	component := NewComponent(gcType, "SkillSlot")

	return &SkillSlot{
		Component: component,
	}
}
