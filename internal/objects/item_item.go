package objects

import (
	"RainbowRunner/pkg/byter"
	"RainbowRunner/pkg/datatypes"
)

//go:generate go run ../../scripts/generatelua -type=Item -extends=Manipulator
type Item struct {
	*Manipulator
	ModCount          int
	Mod               string
	ItemType          ItemType
	InventoryPosition datatypes.Vector2
	Index             int
}

func (n *Item) SetInventoryPosition(vector2 datatypes.Vector2) {
	n.InventoryPosition = vector2
}

func (n *Item) WriteInit(b *byter.Byter) {
	// TODO remove this, I don't think this is in the right place as this should be handled in Manipulators or Inventory etc.
	b.WriteByte(0xFF) // GetType
	b.WriteCString(n.GCType)

	// Item::readData
	// This is the item index within the specific inventory
	// Equipment = Slots
	// Inventory = unique ID
	// TODO extract this into InventoryItem/EquippedItem
	b.WriteUInt32(uint32(n.Index))

	b.WriteByte(byte(n.InventoryPosition.X))
	b.WriteByte(byte(n.InventoryPosition.Y))

	b.WriteByte(0x01) // Item count

	b.WriteByte(50 + 5) // Required level + 5

	// Flag?
	// 0x01 - Soulbound in 9 minutes, no idea where the time comes from
	// 0x02 - Not Sellable
	// 0x04 - +0x01 = Soulbound timer
	// 0x08 - Requires Membership
	//itemFlag := 0x01 | 0x04 | 0x08
	itemFlag := 0x00

	b.WriteByte(byte(itemFlag))

	// This extra data cannot be written when this equipment is being sent as a Manipulator
	// Only for inventory/equipment, TODO work out a way to handle this well
	if itemFlag&0x04 > 0 {
		// Soulbind time
		// Minutes * 0x800 max 9
		b.WriteUInt16(0x800 * 7)
	}

	// Required modifiers?
	// ItemModifier?
	itemModifierFlag1 := 0x00

	//TODO Up to here itemObject is fine, but the required modifiers count is different 5

	// Each item has different numbers of required modifiers
	for i := 0; i < n.ModCount; i++ {
		b.WriteByte(byte(itemModifierFlag1))

		if itemModifierFlag1&0x01 > 0 {
			b.WriteByte(0xFF)
		}

		if itemModifierFlag1&0x02 > 0 {
			b.WriteUInt32(0xFFFFFFFF)
		}
	}

	// GCObject::readChildData<ItemModifier>
	b.WriteByte(0x01) // Count

	b.WriteByte(0xFF)
	b.WriteCString(n.Mod)

	// ItemModifier
	// ItemModifier::readData
	itemModifierFlag := 0x01 | 0x02

	b.WriteByte(byte(itemModifierFlag))

	if itemModifierFlag&0x01 > 0 {
		b.WriteByte(0x15)
	}

	if itemModifierFlag&0x02 > 0 {
		b.WriteUInt32(0x11111111)
	}
}

func NewItem(itemGCType string, itemType ItemType) *Item {
	manipulator := NewManipulator(itemGCType, string(itemType))

	return &Item{
		Manipulator: manipulator,
	}
}
