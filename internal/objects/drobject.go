package objects

import (
	"RainbowRunner/internal/connections"
	"RainbowRunner/pkg/byter"
	"RainbowRunner/pkg/datatypes"
)

//go:generate stringer -type=DRObjectType
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
	OwnerID() uint16
	SetVersion(version byte)
	ReadData(b *byter.Byter)
	WalkChildren(f func(object DRObject))
	RemoveChildAt(i int)
	SetOwner(conn *connections.RRConn)
	String() string
}

type DRItem interface {
	SetInventoryPosition(vector2 datatypes.Vector2)
}
