package objects

import (
	"RainbowRunner/internal/connections"
	"RainbowRunner/pkg/byter"
	"RainbowRunner/pkg/datatypes"
	"github.com/yuin/gopher-lua"
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

type DRItem interface {
	SetInventoryPosition(vector2 datatypes.Vector2)
}

type DRObject interface {
	WriteFullGCObject(b *byter.Byter)
	WriteInit(b *byter.Byter)
	WriteUpdate(b *byter.Byter)
	WriteSynch(b *byter.Byter)

	ReadUpdate(reader *byter.Byter) error

	AddChild(object DRObject)
	Children() []DRObject
	GetChildByGCType(s string) DRObject
	GetChildByGCNativeType(s string) DRObject
	GetParentEntity() IEntity

	Type() DRObjectType

	Tick()
	OwnerID() uint16
	SetVersion(version byte)
	ReadData(b *byter.Byter)
	WalkChildren(f func(object DRObject))
	RemoveChildAt(i int)
	SetOwner(conn *connections.RRConn)
	String() string
	ToLua(state *lua.LState) lua.LValue
	SetParent(g DRObject)
	Init()
}
