package objects

import (
	byter "RainbowRunner/pkg/byter"
	"fmt"
)

type Inventory struct {
	*GCObject

	itemID int
}

func (i *Inventory) AddChild(child DRObject) {
	switch child.(type) {
	case *Equipment:
		child.(*Equipment).Index = i.itemID
	case *Item:
		child.(*Item).Index = i.itemID
	default:
		panic(fmt.Sprintf("cannot add non-item to Inventory: %s", child.GetGCObject().GCType))
	}

	i.GCObject.AddChild(child)
}

func (i *Inventory) WriteInit(body *byter.Byter) {
	body.WriteByte(0xFF)
	body.WriteCString(i.GCType)
	body.WriteByte(0x01)
	body.WriteByte(0x01)

	// GCObject::ReadChildData<Item>()
	body.WriteByte(byte(len(i.Children())))

	for _, item := range i.Children() {
		item.WriteInit(body)
	}

	//AddInventoryItem(body, "PlateMythicPAL.PlateMythicBoots1", 0, 0, "PlateMythicPAL.PlateMythicBoots1.Mod1")
}

func (i *Inventory) GetItemByIndex(index int) DRObject {
	for _, child := range i.Children() {
		foundIndex := 0

		switch child.(type) {
		case *Equipment:
			foundIndex = child.(*Equipment).Index
		case *Item:
			foundIndex = child.(*Item).Index
		default:
			panic(fmt.Sprintf("cannot add non-item to Inventory: %s", child.GetGCObject().GCType))
		}

		if foundIndex == index {
			return child
		}
	}

	return nil
}

func NewInventory(gcType string) *Inventory {
	gcObject := NewGCObject("Inventory")
	gcObject.GCType = gcType

	return &Inventory{
		GCObject: gcObject,
	}
}

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
