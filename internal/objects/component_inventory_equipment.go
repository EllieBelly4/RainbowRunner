package objects

import (
	"RainbowRunner/internal/message"
	"RainbowRunner/internal/types"
	"RainbowRunner/internal/types/drobjecttypes"
	"RainbowRunner/pkg/byter"
	"errors"
	"fmt"
)

//go:generate go run ../../scripts/generatelua -type=EquipmentInventory -extends=Component
type EquipmentInventory struct {
	*Component
	Avatar *Avatar
	Slots  map[types.EquipmentSlot]IEquipment
}

func (n *EquipmentInventory) WriteInit(body *byter.Byter) {
	n.Component.WriteInit(body)

	equippedItems := n.GetEquipment()

	body.WriteByte(byte(len(equippedItems)))

	for _, equippedItem := range equippedItems {
		equippedItem.GetEquipment().WriteInit(body)
	}
}

func (n *EquipmentInventory) AddChild(child drobjecttypes.DRObject) {
	if _, ok := child.(IEquipment); !ok {
		panic(fmt.Sprintf("cannot add non-equipment item to EquipmentInventory: %s", child.(IGCObject).GetGCObject().GCType))
	}

	equip := child.(IEquipment).GetEquipment()

	if existing, ok := n.Slots[equip.Slot]; ok {
		panic(fmt.Sprintf("cannot add equipment '%s' to slot '%s' because '%s' is already equipped", equip.GCType, equip.Slot.String(), existing.GetEquipment().GCType))
	}

	equip.Index = int(equip.Slot)

	n.GCObject.AddChild(child)
}

func (n *EquipmentInventory) ReadUpdate(reader *byter.Byter) error {
	subType := reader.UInt8()
	switch subType {
	// Add equipped item
	case 0x28:
		err := n.handleAddEquippedItem(reader)

		if err != nil {
			return err
		}
	// Remove equipped item
	case 0x29:
		err := n.handleRemoveEquippedItem(reader)
		if err != nil {
			return err
		}
	default:
		return errors.New(fmt.Sprintf("Unknown inventory equipment message subtype %x", subType))
	}
	return nil
}

func (n *EquipmentInventory) handleRemoveEquippedItem(reader *byter.Byter) error {
	slot := reader.UInt32()
	fmt.Printf("%d\n", slot)
	CEWriter := NewClientEntityWriterWithByter()

	unitContainer := n.Avatar.GetUnitContainer()

	if unitContainer == nil {
		return errors.New(fmt.Sprintf("could not find unit container for player"))
	}

	manipulators := n.Avatar.GetManipulators()

	if manipulators == nil {
		return errors.New(fmt.Sprintf("could not find unit manipulators for player"))
	}

	item := n.RemoveEquipmentBySlot(types.EquipmentSlot(slot))
	err := n.addRemoveItemMessage(CEWriter, item.GetEquipment())

	if err != nil {
		return err
	}

	unitContainer.SetActiveItem(item.GetEquipment())
	unitContainer.WriteSetActiveItem(CEWriter.Body)

	manipulators.RemoveChildByID(uint32(item.GetEquipment().ID()))
	manipulators.WriteRemoveItem(CEWriter.Body, item.GetEquipment().Slot)

	Players.GetPlayer(n.OwnerID()).MessageQueue.Enqueue(
		message.QueueTypeClientEntity, CEWriter.Body, message.OpTypeEquippedItemClickResponse,
	)
	return nil
}

func (n *EquipmentInventory) handleAddEquippedItem(reader *byter.Byter) error {
	slot := reader.UInt32()
	CEWriter := NewClientEntityWriterWithByter()

	unitContainer := n.Avatar.GetUnitContainer()

	if unitContainer == nil {
		return errors.New(fmt.Sprintf("could not find unit container for player"))
	}

	if unitContainer.ActiveItem == nil {
		return errors.New(fmt.Sprintf("cannot equip, no active item"))
	}

	manipulators := n.Avatar.GetManipulators()

	if manipulators == nil {
		return errors.New(fmt.Sprintf("could not find unit manipulators for player"))
	}

	iEquipment, ok := unitContainer.ActiveItem.(IEquipment)

	if !ok {
		return errors.New(fmt.Sprintf("cannot equip, active item '%s' is not Equipment", unitContainer.ActiveItem.GetGCType()))
	}

	equipment := iEquipment.GetEquipment()

	if slot != uint32(equipment.Slot) {
		return errors.New(fmt.Sprintf("cannot equip item, wrong slot"))
	}

	n.AddChild(unitContainer.ActiveItem)

	err := n.addAddItemMessage(CEWriter, equipment)

	if err != nil {
		return err
	}

	unitContainer.WriteClearActiveItem(CEWriter.Body)
	unitContainer.SetActiveItem(nil)
	//n.addSetActiveItemMessage(CEWriter, unitContainer, slot)

	manipulators.AddChild(equipment)
	manipulators.WriteAddItem(CEWriter.Body, equipment)

	Players.GetPlayer(uint16(n.OwnerID())).MessageQueue.Enqueue(
		message.QueueTypeClientEntity, CEWriter.Body, message.OpTypeEquippedItemClickResponse,
	)
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

func (n *EquipmentInventory) RemoveEquipmentBySlot(slot types.EquipmentSlot) IEquipment {
	toRemove := -1
	var toReturn IEquipment = nil

	for li, child := range n.Children() {
		foundIndex := types.EquipmentSlot(0)

		switch child.(type) {
		case IEquipment:
			foundIndex = child.(IEquipment).GetEquipment().Slot
		default:
			panic(fmt.Sprintf("cannot add non-item to Inventory: %s", child.(IGCObject).GetGCObject().GCType))
		}

		if foundIndex == slot {
			toRemove = li
			toReturn = child.(IEquipment)
		}
	}

	if toRemove > -1 {
		n.GCChildren = append(n.GCChildren[:toRemove], n.GCChildren[toRemove+1:]...)
	}

	return toReturn
}

func (n *EquipmentInventory) GetEquipment() []IEquipment {
	items := make([]IEquipment, 0)

	for _, child := range n.Children() {
		switch child.(type) {
		case IEquipment:
			items = append(items, child.(IEquipment))
		}
	}

	return items
}

func (n *EquipmentInventory) addAddItemMessage(CEWriter *ClientEntityWriter, item IEquipment) error {
	CEWriter.BeginComponentUpdate(n)
	// 0x28 Add
	// 0x29 Remove
	CEWriter.Body.WriteByte(0x28)

	if item == nil {
		return errors.New(fmt.Sprintf("could not find item in slot %d", item.GetEquipment().Slot))
	}

	item.GetEquipment().WriteInit(CEWriter.Body)
	CEWriter.EndComponentUpdate(n)
	return nil
}

func NewEquipmentInventory(gcType string, avatar *Avatar) *EquipmentInventory {
	component := NewComponent(gcType, "Equipment")

	return &EquipmentInventory{
		Component: component,
		Avatar:    avatar,
	}
}
