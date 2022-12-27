package objects

import (
	"RainbowRunner/internal/message"
	"RainbowRunner/internal/types"
	"RainbowRunner/internal/types/drobjecttypes"
	"RainbowRunner/pkg/byter"
)

//go:generate go run ../../scripts/generatelua -type=Manipulators -extends=Component
type Manipulators struct {
	*Component
}

func (n *Manipulators) WriteInit(b *byter.Byter) {
	// Manipulators::readInit
	b.WriteByte(byte(len(n.Children()))) // Number of visually equipped items

	for _, item := range n.Children() {
		item.WriteInit(b)
	}
}

func (n *Manipulators) RemoveChildByID(id uint32) {
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

func (n *Manipulators) WriteAddItem(body *byter.Byter, item drobjecttypes.DRObject) {
	CEWriter := NewClientEntityWriter(body)
	CEWriter.BeginComponentUpdate(n)

	// 0x00 Add Item
	CEWriter.Body.WriteByte(0x00)

	item.WriteInit(CEWriter.Body)
	CEWriter.EndComponentUpdate(n)
}

func (m *Manipulators) AddChildAndUpdate(child drobjecttypes.DRObject) {
	m.AddChild(child)

	// TODO emit event
	m.sendAddItem(child)
}

func (m *Manipulators) sendAddItem(child drobjecttypes.DRObject) {
	CEWriter := NewClientEntityWriterWithByter()
	CEWriter.BeginComponentUpdate(m)

	// 0x00 Add Item
	CEWriter.Body.WriteByte(0x00)

	CEWriter.Body.WriteByte(0xFF) // GetType
	CEWriter.Body.WriteCString(child.(IGCObject).GetGCObject().GCType)

	child.WriteData(CEWriter.Body)
	CEWriter.WriteSynch(m)

	player := m.GetPlayerOwner()
	player.MessageQueue.EnqueueClientEntity(CEWriter.Body, message.OpTypeManipulators)
}

func (m *Manipulators) RemoveChildAndUpdate(skill drobjecttypes.DRObject) {
	m.RemoveChild(skill)

	// TODO emit event
	m.sendRemoveItem(skill)
}

func (m *Manipulators) sendRemoveItem(skill drobjecttypes.DRObject) {
	CEWriter := NewClientEntityWriterWithByter()
	CEWriter.BeginComponentUpdate(m)

	// 0x00 Add Item
	CEWriter.Body.WriteByte(0x01)
	CEWriter.Body.WriteUInt32(uint32(skill.(IRREntityPropertiesHaver).GetRREntityProperties().ID))

	player := m.GetPlayerOwner()
	player.MessageQueue.EnqueueClientEntity(CEWriter.Body, message.OpTypeManipulators)
}

func NewManipulators(gcType string) *Manipulators {
	component := NewComponent(gcType, "Manipulators")

	return &Manipulators{
		Component: component,
	}
}
