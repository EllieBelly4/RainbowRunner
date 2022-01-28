package objects

import (
	"RainbowRunner/pkg"
	"RainbowRunner/pkg/byter"
)

type DRObjectType int

const (
	DRObjectEntity DRObjectType = iota
	DRObjectComponent
	DRObjectOther
	DRObjectManager
	DRObjectUnknown
)

type DRObject interface {
	RREntityProperties() *RREntityProperties

	WriteFullGCObject(b *byter.Byter)
	WriteInit(b *byter.Byter)
	WriteUpdate(b *byter.Byter)
	WriteSynch(b *byter.Byter)

	ReadUpdate(reader *byter.Byter) error

	AddChild(object DRObject)
	Children() []DRObject
	GetChildByGCType(s string) DRObject
	GetChildByGCNativeType(s string) DRObject

	Type() DRObjectType

	GetGCObject() *GCObject
	Tick()
	OwnerID() int
}

type DRItem interface {
	SetInventoryPosition(vector2 pkg.Vector2)
}
