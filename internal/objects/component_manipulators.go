package objects

import (
	"RainbowRunner/internal/types"
	"RainbowRunner/pkg/byter"
	"fmt"
)

//go:generate go run ../../scripts/generatelua -type=Manipulators -extends=Component
type Manipulators struct {
	*Component
}

func (n *Manipulators) WriteInit(b *byter.Byter) {
	// Manipulators::readInit
	b.WriteByte(byte(len(n.Children()))) // Number of visually equipped items

	for _, equippedItem := range n.Children() {
		ok := false
		var equip *Equipment

		if equip, ok = equippedItem.(*Equipment); !ok {
			panic(fmt.Sprintf("cannot init manipulator as '%s' is not Equipment", equip.GCType))
		}

		equip.WriteInit(b)
		equip.WriteManipulatorInit(b)
	}
}

func (n *Manipulators) RemoveEquipmentByID(id uint32) {
	toRemove := -1

	for li, child := range n.Children() {
		if child.(IRREntityPropertiesHaver).GetRREntityProperties().ID == id {
			toRemove = li
		}
	}

	if toRemove > -1 {
		n.GCChildren = append(n.GCChildren[:toRemove], n.GCChildren[toRemove+1:]...)
	}
}

func (n *Manipulators) WriteRemoveItem(body *byter.Byter, id types.EquipmentSlot) {
	CEWriter := NewClientEntityWriter(body)
	CEWriter.BeginComponentUpdate(n)

	// 0x01 Remove item
	CEWriter.Body.WriteByte(0x01)
	CEWriter.Body.WriteUInt32(uint32(id))

	CEWriter.EndComponentUpdate(n)
}

func (n *Manipulators) WriteAddItem(body *byter.Byter, equipment *Equipment) {
	CEWriter := NewClientEntityWriter(body)
	CEWriter.BeginComponentUpdate(n)

	// 0x00 Add Item
	CEWriter.Body.WriteByte(0x00)

	equipment.WriteInit(CEWriter.Body)
	CEWriter.EndComponentUpdate(n)
}

func NewManipulators(gcType string) *Manipulators {
	component := NewComponent(gcType, "Manipulators")

	return &Manipulators{
		Component: component,
	}
}
