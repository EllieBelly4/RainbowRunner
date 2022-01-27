package objects

import (
	"RainbowRunner/internal/database"
	"RainbowRunner/internal/types"
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

	var drClass *database.DRClass

	if itemType == ItemArmour {
		drClass = database.Armour.Find(itemGCType)
	} else {
		drClass = database.Weapons.Find(itemGCType)
	}

	item.Mod = itemModGCType
	item.Slot = slot
	item.ModCount = drClass.ModCount()
	item.ItemType = itemType

	if drClass == nil {
		panic(fmt.Sprintf("equipment class not found in db %s", itemGCType))
	}

	return &Equipment{
		Item: item,
	}
}
