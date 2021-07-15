package objects

import (
	"RainbowRunner/pkg/byter"
)

type Item struct {
	*GCObject
}

func (n *Item) WriteInit(b *byter.Byter) {

}

func NewItem(itemGCType string, itemType ItemType) *Item {
	gcObject := NewGCObject(string(itemType))
	gcObject.GCType = itemGCType

	return &Item{
		GCObject: gcObject,
	}
}
