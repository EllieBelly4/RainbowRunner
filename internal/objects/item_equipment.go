package objects

import (
	"RainbowRunner/internal/database"
	"RainbowRunner/internal/types"
	"RainbowRunner/pkg/byter"
	"fmt"
)

type ItemType string

const (
	EquipmentItemArmour       ItemType = "Armor"
	EquipmentItemMeleeWeapon  ItemType = "MeleeWeapon"
	EquipmentItemRangedWeapon ItemType = "RangedWeapon"
)

type Equipment struct {
	*GCObject
	Mod      string
	Slot     types.EquipmentSlot
	ModCount int
	ItemType ItemType
}

func (n *Equipment) WriteInit(b *byter.Byter) {
	b.WriteByte(0xFF) // GetType
	b.WriteCString(n.GCType)

	// Item::readData
	b.WriteUInt32(uint32(n.Slot))
	b.WriteByte(0xF0)
	b.WriteByte(0xF0)
	b.WriteByte(0x01)   // Item count
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

	// ItemModifier?
	itemModifierFlag := 0x01 | 0x02

	b.WriteByte(byte(itemModifierFlag))

	if itemModifierFlag&0x01 > 0 {
		b.WriteByte(0x15)
	}

	if itemModifierFlag&0x02 > 0 {
		b.WriteUInt32(0x11111111)
	}
}

func (n *Equipment) WriteManipulatorInit(b *byter.Byter) {
	// Manipulators::readInit
	// .text:004FD1AB
	if n.ItemType == EquipmentItemMeleeWeapon {
		// MeleeWeapon::readInit
		// Item::readInit nothing happens

		// MeleeWeapon::readInit
		b.WriteUInt16(0x01)
		b.WriteByte(0x02)

		unitIDMaybe := 0x00 // Guessing
		b.WriteUInt16(uint16(unitIDMaybe))

		// .text:00592438
		if unitIDMaybe > 0 {
			// do loads of stuff including checking if a type is a unit
		}
	}
}

func NewEquipment(itemGCType, itemModGCType string, itemType ItemType, slot types.EquipmentSlot) *Equipment {
	gcObject := NewGCObject(string(itemType))
	gcObject.GCType = itemGCType

	var drClass *database.DRClass

	if itemType == EquipmentItemArmour {
		drClass = database.Armour.Find(itemGCType)
	} else {
		drClass = database.Weapons.Find(itemGCType)
	}

	if drClass == nil {
		panic(fmt.Sprintf("equipment class not found in db %s", itemGCType))
	}

	return &Equipment{
		GCObject: gcObject,
		Mod:      itemModGCType,
		Slot:     slot,
		ModCount: drClass.ModCount(),
		ItemType: itemType,
	}
}
