package objects

import (
	"RainbowRunner/internal/types/drobjecttypes"
	byter "RainbowRunner/pkg/byter"
	"fmt"
)

//go:generate go run ../../scripts/generatelua -type=Inventory -extends=GCObject
type Inventory struct {
	*GCObject

	itemID      int
	InventoryID byte
	Items       []IItem
}

func (i *Inventory) AddItem(child drobjecttypes.DRObject) {
	switch child.(type) {
	case IItem:
		child.(IItem).GetItem().Index = i.itemID
	default:
		panic(fmt.Sprintf("cannot add non-item to Inventory: %s", child.(IGCObject).GetGCObject().GCType))
	}

	i.itemID++
	i.Items = append(i.Items, child.(IItem))
}

func (i *Inventory) WriteInit(body *byter.Byter) {
	body.WriteByte(0xFF)
	body.WriteCString(i.GCType)
	body.WriteByte(i.InventoryID)

	i.WriteInitData(body)

	//AddInventoryItem(body, "PlateMythicPAL.PlateMythicBoots1", 0, 0, "PlateMythicPAL.PlateMythicBoots1.Mod1")
}

func (i *Inventory) WriteInitData(body *byter.Byter) {
	body.WriteByte(0x01) // Cannot be 0

	// GCObject::ReadChildData<Item>()
	body.WriteByte(byte(len(i.Children())))

	for _, item := range i.Children() {
		item.WriteInit(body)
	}
}

func (i *Inventory) RemoveItemByIndex(index int) drobjecttypes.DRObject {
	toRemove := -1
	var toReturn drobjecttypes.DRObject = nil

	for li, child := range i.Children() {
		foundIndex := 0

		switch child.(type) {
		case IEquipment:
			foundIndex = child.(IEquipment).GetEquipment().Index
		case *Item:
			foundIndex = child.(*Item).Index
		default:
			panic(fmt.Sprintf("cannot remove non-item from Inventory: %s", child.(IGCObject).GetGCObject().GCType))
		}

		if foundIndex == index {
			toRemove = li
			toReturn = child
		}
	}

	if toRemove > -1 {
		i.GCChildren = append(i.GCChildren[:toRemove], i.GCChildren[toRemove+1:]...)
	}

	return toReturn
}

func NewInventory(gcType string, index byte) *Inventory {
	gcObject := NewGCObject("Inventory")
	gcObject.GCType = gcType

	return &Inventory{
		GCObject: gcObject,
		// TODO figure out how to set inventory ID properly, client is always using 1
		InventoryID: index,
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
