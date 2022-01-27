package objects

import (
	"RainbowRunner/internal/message"
	"RainbowRunner/internal/types"
	"RainbowRunner/pkg/byter"
	"errors"
	"fmt"
)

type InventoryEquipment struct {
	*Component
	Avatar *Avatar
}

func (n *InventoryEquipment) ReadUpdate(reader *byter.Byter) error {
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

		if slot != uint32(unitContainer.ActiveItem.Slot) {
			return errors.New(fmt.Sprintf("cannot equip item, wrong slot"))
		}

		err := n.addAddItemMessage(CEWriter, unitContainer.ActiveItem)

		if err != nil {
			return err
		}

		unitContainer.WriteClearActiveItem(CEWriter.Body)
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

func (n *InventoryEquipment) addRemoveItemMessage(CEWriter *ClientEntityWriter, item *Equipment) error {
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

func (n *InventoryEquipment) GetEquipmentBySlot(slot types.EquipmentSlot) *Equipment {
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

func (n *InventoryEquipment) GetEquipment() []*Equipment {
	items := make([]*Equipment, 0)

	for _, child := range n.Children() {
		switch child.(type) {
		case *Equipment:
			items = append(items, child.(*Equipment))
		}
	}

	return items
}

func (n *InventoryEquipment) addAddItemMessage(CEWriter *ClientEntityWriter, item *Equipment) error {
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

func NewInventoryEquipment(gcType string, avatar *Avatar) *InventoryEquipment {
	component := NewComponent(gcType, "Equipment")

	return &InventoryEquipment{
		Component: component,
		Avatar:    avatar,
	}
}
