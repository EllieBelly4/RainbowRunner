package objects

import (
	"RainbowRunner/internal/message"
	"RainbowRunner/internal/types"
	"RainbowRunner/pkg/byter"
	"errors"
	"fmt"
)

type EquipmentInventory struct {
	*Component
	Avatar *Avatar
	Slots  map[types.EquipmentSlot]*Equipment
}

func (n *EquipmentInventory) AddChild(child DRObject) {
	if _, ok := child.(*Equipment); !ok {
		panic(fmt.Sprintf("cannot add non-equipment item to EquipmentInventory: %s", child.GetGCObject().GCType))
	}

	equip := child.(*Equipment)

	if existing, ok := n.Slots[equip.Slot]; ok {
		panic(fmt.Sprintf("cannot add equipment '%s' to slot '%s' because '%s' is already equipped", equip.GCType, equip.Slot.String(), existing.GCType))
	}

	child.(*Equipment).Index = int(equip.Slot)

	n.GCObject.AddChild(child)
}

func (n *EquipmentInventory) ReadUpdate(reader *byter.Byter) error {
	reader.Dump()

	subType := reader.UInt8()
	switch subType {
	// Add equipped item
	case 0x28:
		slot := reader.UInt32()
		fmt.Printf("%d\n", slot)
		CEWriter := NewClientEntityWriterWithByter()

		unitContainer := n.Avatar.GetUnitContainer()

		if unitContainer == nil {
			return errors.New(fmt.Sprintf("could not find unit container for player"))
		}

		if unitContainer.ActiveItem == nil {
			return errors.New(fmt.Sprintf("cannot equip, no active item"))
		}

		equipment, ok := unitContainer.ActiveItem.(*Equipment)

		if !ok {
			return errors.New(fmt.Sprintf("cannot equip, active item '%s' is not Equipment", equipment.GCType))
		}

		if slot != uint32(equipment.Slot) {
			return errors.New(fmt.Sprintf("cannot equip item, wrong slot"))
		}

		err := n.addAddItemMessage(CEWriter, equipment)

		if err != nil {
			return err
		}

		unitContainer.WriteClearActiveItem(CEWriter.Body)
		unitContainer.SetActiveItem(nil)
		//n.addSetActiveItemMessage(CEWriter, unitContainer, slot)

		Players.GetPlayer(uint16(n.OwnerID())).MessageQueue.Enqueue(
			message.QueueTypeClientEntity, CEWriter.Body, message.OpTypeEquippedItemClickResponse,
		)
	// Remove equipped item
	case 0x29:
		slot := reader.UInt32()
		fmt.Printf("%d\n", slot)
		CEWriter := NewClientEntityWriterWithByter()

		unitContainer := n.Avatar.GetUnitContainer()

		if unitContainer == nil {
			return errors.New(fmt.Sprintf("could not find unit container for player"))
		}

		item := n.GetEquipmentBySlot(types.EquipmentSlot(slot))

		err := n.addRemoveItemMessage(CEWriter, item)

		if err != nil {
			return err
		}

		unitContainer.SetActiveItem(item)
		unitContainer.WriteSetActiveItem(CEWriter.Body)

		Players.GetPlayer(uint16(n.OwnerID())).MessageQueue.Enqueue(
			message.QueueTypeClientEntity, CEWriter.Body, message.OpTypeEquippedItemClickResponse,
		)
	default:
		return errors.New(fmt.Sprintf("Unknown inventory equipment message subtype %x", subType))
	}
	return nil
}

func (n *EquipmentInventory) addRemoveItemMessage(CEWriter *ClientEntityWriter, item *Equipment) error {
	CEWriter.BeginComponentUpdate(n)
	// 0x28 Add
	// 0x29 Remove
	CEWriter.Body.WriteByte(0x29)

	if item == nil {
		return errors.New(fmt.Sprintf("could not find item in slot %d", item.Slot))
	}

	CEWriter.Body.WriteUInt32(uint32(item.Slot))
	CEWriter.EndComponentUpdate(n)
	return nil
}

func (n *EquipmentInventory) GetEquipmentBySlot(slot types.EquipmentSlot) *Equipment {
	for _, child := range n.Children() {
		switch child.(type) {
		case *Equipment:
			equipment := child.(*Equipment)
			if equipment.Slot == slot {
				return equipment
			}
		}
	}

	return nil
}

func (n *EquipmentInventory) GetEquipment() []*Equipment {
	items := make([]*Equipment, 0)

	for _, child := range n.Children() {
		switch child.(type) {
		case *Equipment:
			items = append(items, child.(*Equipment))
		}
	}

	return items
}

func (n *EquipmentInventory) addAddItemMessage(CEWriter *ClientEntityWriter, item *Equipment) error {
	CEWriter.BeginComponentUpdate(n)
	// 0x28 Add
	// 0x29 Remove
	CEWriter.Body.WriteByte(0x28)

	if item == nil {
		return errors.New(fmt.Sprintf("could not find item in slot %d", item.Slot))
	}

	item.WriteInit(CEWriter.Body)
	CEWriter.EndComponentUpdate(n)
	return nil
}

func NewInventoryEquipment(gcType string, avatar *Avatar) *EquipmentInventory {
	component := NewComponent(gcType, "Equipment")

	return &EquipmentInventory{
		Component: component,
		Avatar:    avatar,
	}
}
