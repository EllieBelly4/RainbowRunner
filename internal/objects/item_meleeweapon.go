package objects

import (
	"RainbowRunner/internal/types"
	"RainbowRunner/pkg/byter"
)

//go:generate go run ../../scripts/generatelua -type=MeleeWeapon -extends=Item
type MeleeWeapon struct {
	*Equipment
}

func (n *MeleeWeapon) WriteInit(b *byter.Byter) {
	n.Equipment.WriteInit(b)

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

func NewMeleeWeapon(gctype string, modGCType string) *MeleeWeapon {
	equipment := NewEquipment(gctype, modGCType, ItemMeleeWeapon, types.EquipmentSlotWeapon)
	return &MeleeWeapon{Equipment: equipment}
}
