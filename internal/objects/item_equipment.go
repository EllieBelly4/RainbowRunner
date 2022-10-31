package objects

import (
	"RainbowRunner/internal/database"
	"RainbowRunner/internal/types"
	"RainbowRunner/internal/types/configtypes"
	"RainbowRunner/pkg/byter"
	"fmt"
)

type ItemType string

const (
	ItemArmour       ItemType = "Armor"
	ItemMeleeWeapon  ItemType = "MeleeWeapon"
	ItemRangedWeapon ItemType = "RangedWeapon"
)

type Equipment struct {
	*Item
	Slot types.EquipmentSlot
}

func (n *Equipment) WriteInit(b *byter.Byter) {
	n.Item.WriteInit(b)
	// This does not always need to happen
	//n.WriteManipulatorInit(b)
}

func (n *Equipment) WriteManipulatorInit(b *byter.Byter) {
	// Manipulators::readInit
	// .text:004FD1AB
	if n.ItemType == ItemMeleeWeapon {
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
	item := NewItem(string(itemType), itemType)
	item.GCType = itemGCType

	var drClass *configtypes.DRClass

	if itemType == ItemArmour {
		drClass = database.FindItem(database.Armour, itemGCType)
	} else {
		drClass = database.FindItem(database.Weapons, itemGCType)
	}

	if drClass == nil {
		panic(fmt.Sprintf("could not find DRClass for equipment '%s'", itemGCType))
	}

	item.Mod = itemModGCType
	item.ModCount = drClass.ModCount()
	item.ItemType = itemType

	if drClass == nil {
		panic(fmt.Sprintf("equipment class not found in db %s", itemGCType))
	}

	return &Equipment{
		Item: item,
		Slot: slot,
	}
}
