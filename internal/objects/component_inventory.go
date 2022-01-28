package objects

import (
	byter "RainbowRunner/pkg/byter"
	"fmt"
)

type Inventory struct {
	*GCObject

	itemID      int
	InventoryID byte
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

	i.itemID++

	i.GCObject.AddChild(child)
}

func (i *Inventory) WriteInit(body *byter.Byter) {
	body.WriteByte(0xFF)
	body.WriteCString(i.GCType)
	body.WriteByte(i.InventoryID)
	body.WriteByte(0x01) // Cannot be 0

	// GCObject::ReadChildData<Item>()
	body.WriteByte(byte(len(i.Children())))

	for _, item := range i.Children() {
		item.WriteInit(body)
	}

	//AddInventoryItem(body, "PlateMythicPAL.PlateMythicBoots1", 0, 0, "PlateMythicPAL.PlateMythicBoots1.Mod1")
}

func (i *Inventory) RemoveItemByIndex(index int) DRObject {
	toRemove := -1
	var toReturn DRObject = nil

	for li, child := range i.Children() {
		foundIndex := 0

		switch child.(type) {
		case *Equipment:
			foundIndex = child.(*Equipment).Index
		case *Item:
			foundIndex = child.(*Item).Index
		default:
			panic(fmt.Sprintf("cannot remove non-item from Inventory: %s", child.GetGCObject().GCType))
		}

		if foundIndex == index {
			toRemove = li
			toReturn = child
		}
	}

	if toRemove > -1 {
		i.children = append(i.children[:toRemove], i.children[toRemove+1:]...)
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
