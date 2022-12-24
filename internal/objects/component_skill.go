package objects

//go:generate go run ../../scripts/generatelua -type=Skill -extends=Component
type Skill struct {
	*Component
	Unk0         uint32
	Level        byte
	OriginalSlot int
	Slot         int

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

func NewSkill(gcType string) *Skill {
	component := NewComponent(gcType, "Skill")

	return &Skill{
		Component: component,
	}
}
