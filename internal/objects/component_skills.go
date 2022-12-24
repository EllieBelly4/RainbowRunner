package objects

import (
	"RainbowRunner/internal/message"
	"RainbowRunner/internal/types/drobjecttypes"
	"RainbowRunner/pkg/byter"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"sort"
)

type SkillsUpdateType uint8

const (
	SkillsUpdateTypeUpdateSkill         SkillsUpdateType = 50
	SkillsUpdateTypeUpdateSkillPoints   SkillsUpdateType = 51
	SkillsUpdateTypeUpdateAddProfession SkillsUpdateType = 55
	SkillsUpdateTypeUpdateRemoveSkill   SkillsUpdateType = 56
	SkillsUpdateTypeUpdateSkillSlot     SkillsUpdateType = 57
)

type SkillsRequestType uint8

const (
	SkillsRequestTypeUnequipSkill  SkillsRequestType = 54
	SkillsRequestTypeEquipSkill    SkillsRequestType = 53
	SkillsRequestTypeBuySkillLevel SkillsRequestType = 50
)

//go:generate go run ../../scripts/generatelua -type=Skills -extends=Component
type Skills struct {
	*Component

	slots map[uint32]ISkill

	/*
		PropertySkillsSkillPoints
	*/
}

func (n *Skills) ReadUpdate(reader *byter.Byter) error {
	updateType := reader.Byte()

	switch SkillsRequestType(updateType) {
	case SkillsRequestTypeBuySkillLevel:
		n.handleBuySkillLevel(reader)
	case SkillsRequestTypeEquipSkill:
		n.handleEquipSkill(reader)
	case SkillsRequestTypeUnequipSkill:
		n.handleUnequipSkill(reader)
	default:
		log.Errorf("Unknown SkillsUpdateType: %d", updateType)
	}

	return nil
}

func (n *Skills) GetSkillInSlot(slot uint32) ISkill {
	return n.slots[slot]
}

func (n *Skills) WriteInit(b *byter.Byter) {
	// Skills::readInit()
	b.WriteUInt32(0xFFFFFFFF)

	skills := n.GetAllSkills()

	// GCObject::readChildData<Skill>
	b.WriteByte(byte(len(skills))) // Count

	for _, skillDRObject := range skills {
		skill := skillDRObject.(ISkill).GetSkill()

		b.WriteByte(0xFF)
		b.WriteCString(skill.GCType)
		b.WriteUInt32(skill.Unk0)
		b.WriteByte(skill.Level) // Level
	}

	//b.WriteByte(0xFF)
	//b.WriteCString("skills.generic.Butcher")
	//b.WriteUInt32(0x02)
	//b.WriteByte(0x03) // Level
	//
	//b.WriteByte(0xFF)
	//b.WriteCString("skills.generic.Stomp")
	//b.WriteUInt32(0x04)
	//b.WriteByte(0x05) // Level
	//
	//b.WriteByte(0xFF)
	//b.WriteCString("skills.generic.FighterClassPassive")
	//b.WriteUInt32(0x06)
	//b.WriteByte(0x07) // Level
	//
	//b.WriteByte(0xFF)
	//b.WriteCString("skills.generic.MeleeAttackSpeedModPassive")
	//b.WriteUInt32(0x08)
	//b.WriteByte(0x09) // Level

	// GCObject::readChildData<SkillProfession>
	b.WriteByte(0x01)
	b.WriteByte(0xFF)
	b.WriteCString("skills.professions.Warrior")
}

func (n *Skills) GetAllSkills() []ISkill {
	skills := make([]ISkill, 0)

	for _, skill := range n.slots {
		skills = append(skills, skill)
	}

	sort.SliceStable(skills, func(i, j int) bool {
		return skills[i].GetSkill().OriginalSlot <= skills[j].GetSkill().OriginalSlot
	})

	return skills
}

func (n *Skills) GetActivePassiveSkills() []drobjecttypes.DRObject {
	return n.GetChildrenByGCNativeType("ActivePassiveSkill")
}

func (n *Skills) GetPassiveSkills() []drobjecttypes.DRObject {
	return n.GetChildrenByGCNativeType("PassiveSkill")
}

func (n *Skills) GetActiveSkills() []drobjecttypes.DRObject {
	return n.GetChildrenByGCNativeType("ActiveSkill")
}

func (s *Skills) handleBuySkillLevel(reader *byter.Byter) {

}

func (s *Skills) handleEquipSkill(reader *byter.Byter) {
	slotNumber := reader.UInt32()

	skill := s.GetSkillByGCTypeRequest(reader)

	if skill == nil {
		log.Errorf("Skill not found")
		return
	}

	err := s.EquipSkill(skill.(ISkill), slotNumber)

	if err != nil {
		log.Error(err)
		return
	}

	log.Infof("Equip " + skill.(IGCObject).GetGCObject().GCType)
}

func (s *Skills) handleUnequipSkill(reader *byter.Byter) {
	slotNumber := reader.UInt32()

	skill, ok := s.slots[slotNumber]

	if !ok {
		log.Errorf("Skill not found in slot %d", slotNumber)
		return
	}

	err := s.UnequipSkill(skill)

	if err != nil {
		log.Error(err)
		return
	}

	log.Infof("Unequip " + skill.(IGCObject).GetGCObject().GCType)
}

func (n *Skills) UnequipSkill(skill ISkill) error {
	slot := skill.GetSkill().Slot

	if skillInSlot, ok := n.slots[uint32(slot)]; !ok || skillInSlot != skill {
		return errors.New(fmt.Sprintf("Slot %d is not occupied by skill %s", slot, skill.GetSkill().GCType))
	}

	originalSlot := skill.GetSkill().OriginalSlot

	delete(n.slots, uint32(slot))
	n.slots[uint32(originalSlot)] = skill
	skill.GetSkill().Slot = originalSlot

	n.sendUpdateSkillSlot(skill)

	return nil
}

func (n *Skills) EquipSkill(skill ISkill, slot uint32) error {
	if slot < 0x64 {
		return errors.New(fmt.Sprintf("Invalid equip slot %d", slot))
	}

	if skillInSlot, ok := n.slots[slot]; ok && skillInSlot != nil {
		return errors.New(fmt.Sprintf("Slot %d is already occupied by skill %s", slot, skillInSlot.GetSkill().GCType))
	}

	getSkill := skill.GetSkill()
	prevSlot := getSkill.Slot
	delete(n.slots, uint32(prevSlot))
	n.slots[slot] = skill
	prevSlot = int(slot)

	unit := n.GetParentEntity().(IUnit)
	manipulators := unit.GetUnit().GetChildByGCNativeType("Manipulators").(IManipulators).GetManipulators()

	manipulators.AddChildAndUpdate(getSkill)

	n.sendUpdateSkillSlot(skill)

	return nil
}

func (s *Skills) sendUpdateSkillSlot(skill ISkill) {
	CEWriter := NewClientEntityWriterWithByter()
	CEWriter.BeginComponentUpdate(s)
	body := CEWriter.Body

	body.WriteByte(byte(SkillsUpdateTypeUpdateSkillSlot))
	body.WriteByte(0xFF)
	body.WriteCString(skill.GetSkill().GCType)
	body.WriteByte(byte(skill.GetSkill().Slot))

	CEWriter.EndComponentUpdate(s)

	player := s.GetPlayerOwner()
	player.MessageQueue.EnqueueClientEntity(CEWriter.Body, message.OpTypeSkills)
}

func (s *Skills) AddSkill(skill *ActiveSkill, slot int) {
	s.slots[uint32(slot)] = skill
	skill.OriginalSlot = slot
	skill.Slot = slot
}

func (s *Skills) GetSkillByGCTypeRequest(reader *byter.Byter) ISkill {
	drObjects := make([]drobjecttypes.DRObject, 0)

	for _, skill := range s.GetAllSkills() {
		drObjects = append(drObjects, skill.(drobjecttypes.DRObject))
	}

	foundSkill := SelectFromGCTypeRequest(reader, drObjects)

	if foundSkill == nil {
		return nil
	}

	return foundSkill.(ISkill)
}

func NewSkills(gcType string) *Skills {
	component := NewComponent(gcType, "Skills")

	return &Skills{
		Component: component,
		slots:     map[uint32]ISkill{},
	}
}
