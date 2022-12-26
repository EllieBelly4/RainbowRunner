package objects

//go:generate go run ../../scripts/generatelua -type=SkillSlot -extends=Component
type SkillSlot struct {
	*Component
	SlotID   int
	SlotType uint32 // Unk

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
