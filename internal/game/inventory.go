package game

import (
	byter "RainbowRunner/pkg/byter"
)

type EquipmentSlot uint32

const (
	EquipmentSlotNone EquipmentSlot = iota
	EquipmentSlotAmulet
	EquipmentSlotHand
	EquipmentSlotLRing
	EquipmentSlotRRing
	EquipmentSlotHead
	EquipmentSlotTorso
	EquipmentSlotFoot
	EquipmentSlotShoulder
	EquipmentSlotNone2
	EquipmentSlotWeapon
	EquipmentSlotOffhand
)

// 0x00 None
// 0x01 Amulet
// 0x02 Hand
// 0x03 LRing
// 0x04 RRing
// 0x05 Head
// 0x06 Torso
// 0x07 Foot
// 0x08 Shoulder
// 0x09 None?
// 0x0a Weapon
// 0x0b Offhand/Shield

func AddInventoryItem(
	body *byter.Byter,
	item string,
	x, y byte,
	mod string,
) {
	body.WriteByte(0xFF)
	body.WriteCString(item)

	// Item::readData
	// Unk for inventory
	body.WriteUInt32(0x00)

	body.WriteByte(x) // Position X
	body.WriteByte(y) // Position Y

	body.WriteByte(0x00) // Item count
	body.WriteByte(0x00)

	// Flag?
	// 0x04 read 2 more bytes
	itemFlag := 0x01

	body.WriteByte(byte(itemFlag))

	if itemFlag&0x04 > 0 {
		body.WriteUInt16(0x1314)
	}

	if item == "LeatherArmor1PAL.LeatherArmor1-1" {
		// Required modifiers?
		// ItemModifier?
		itemModifierFlag1 := 0x01 | 0x02

		body.WriteByte(byte(itemModifierFlag1))

		if itemModifierFlag1&0x01 > 0 {
			body.WriteByte(0xFF)
		}

		if itemModifierFlag1&0x02 > 0 {
			body.WriteUInt32(0xFFFFFFFF)
		}

		//if mod != "" {
		// GCObject::readChildData<ItemModifier>
		body.WriteByte(0x01) // Count

		body.WriteByte(0xFF)
		body.WriteCString(mod)

		// ItemModifier?
		itemModifierFlag := 0x01 | 0x02

		body.WriteByte(byte(itemModifierFlag))

		if itemModifierFlag&0x01 > 0 {
			body.WriteByte(0x15)
		}

		if itemModifierFlag&0x02 > 0 {
			body.WriteUInt32(0x11111111)
		}
		//} else {
		//	body.WriteByte(0x00) // Count
		//}
	} else if item == "PlateMythicPAL.PlateMythicArmor1" || item == "PlateMythicPAL.PlateMythicBoots1" {
		// Required modifiers?
		// ItemModifier?
		itemModifierFlag1 := 0xFF

		// Each item has different numbers of required modifiers
		for i := 0; i < 5; i++ {
			body.WriteByte(byte(itemModifierFlag1))

			if itemModifierFlag1&0x01 > 0 {
				body.WriteByte(0xFF)
			}

			if itemModifierFlag1&0x02 > 0 {
				body.WriteUInt32(0xFFFFFFFF)
			}
		}

		//if mod != "" {
		// GCObject::readChildData<ItemModifier>
		body.WriteByte(0x01) // Count

		body.WriteByte(0xFF)
		body.WriteCString(mod)

		// ItemModifier?
		itemModifierFlag := 0x01 | 0x02

		body.WriteByte(byte(itemModifierFlag))

		if itemModifierFlag&0x01 > 0 {
			body.WriteByte(0x15)
		}

		if itemModifierFlag&0x02 > 0 {
			body.WriteUInt32(0x11111111)
		}
	} else {
		body.WriteByte(0x00) // Count
	}
}

func addEquippedItem(
	body *byter.Byter,
	item string,
	slot EquipmentSlot,
	armour bool,
	mod string,
) {
	body.WriteByte(0xFF) // GetType
	body.WriteCString(item)

	// Item::readData
	body.WriteUInt32(uint32(slot))
	body.WriteByte(0xF0)
	body.WriteByte(0xF0)
	body.WriteByte(0x01)   // Item count
	body.WriteByte(50 + 5) // Required level + 5

	// Flag?
	// 0x01 - Soulbound in 9 minutes, no idea where the time comes from
	// 0x02 - Not Sellable
	// 0x04 - +0x01 = Soulbound timer
	// 0x08 - Requires Membership
	itemFlag := 0x01 | 0x04 | 0x08

	body.WriteByte(byte(itemFlag))

	if itemFlag&0x04 > 0 {
		// Soulbind time
		// Minutes * 0x800 max 9
		body.WriteUInt16(0x800 * 7)
	}

	if item == "LeatherArmor1PAL.LeatherArmor1-1" || item == "ScaleArmor1PAL.ScaleArmor1-1" {
		// Required modifiers?
		// ItemModifier?
		itemModifierFlag1 := 0x01 | 0x02

		body.WriteByte(byte(itemModifierFlag1))

		if itemModifierFlag1&0x01 > 0 {
			body.WriteByte(0xFF)
		}

		if itemModifierFlag1&0x02 > 0 {
			body.WriteUInt32(0xFFFFFFFF)
		}

		//if mod != "" {
		// GCObject::readChildData<ItemModifier>
		body.WriteByte(0x01) // Count

		body.WriteByte(0xFF)
		body.WriteCString(mod)

		// ItemModifier?
		itemModifierFlag := 0x01 | 0x02

		body.WriteByte(byte(itemModifierFlag))

		if itemModifierFlag&0x01 > 0 {
			body.WriteByte(0x15)
		}

		if itemModifierFlag&0x02 > 0 {
			body.WriteUInt32(0x11111111)
		}
		//} else {
		//	body.WriteByte(0x00) // Count
		//}
	} else if item == "PlateMythicPAL.PlateMythicArmor1" || item == "PlateMythicPAL.PlateMythicBoots1" {
		// Required modifiers?
		// ItemModifier?
		itemModifierFlag1 := 0x00

		// Each item has different numbers of required modifiers
		for i := 0; i < 5; i++ {
			body.WriteByte(byte(itemModifierFlag1))

			if itemModifierFlag1&0x01 > 0 {
				body.WriteByte(0xFF)
			}

			if itemModifierFlag1&0x02 > 0 {
				body.WriteUInt32(0xFFFFFFFF)
			}
		}

		//if mod != "" {
		// GCObject::readChildData<ItemModifier>
		body.WriteByte(0x01) // Count

		body.WriteByte(0xFF)
		body.WriteCString(mod)

		// ItemModifier?
		itemModifierFlag := 0x01 | 0x02

		body.WriteByte(byte(itemModifierFlag))

		if itemModifierFlag&0x01 > 0 {
			body.WriteByte(0x15)
		}

		if itemModifierFlag&0x02 > 0 {
			body.WriteUInt32(0x11111111)
		}
	} else {
		panic("unhandled equipment")
	}
}
