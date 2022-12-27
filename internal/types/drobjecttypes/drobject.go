package drobjecttypes

import (
	"RainbowRunner/internal/connections"
	"RainbowRunner/pkg/byter"
	"github.com/yuin/gopher-lua"
)

//go:generate stringer -type=DRObjectType
type DRObjectType int

type DRObject interface {
	GetGCType() string
	GetGCNativeType() string

	WriteFullGCObject(b *byter.Byter)
	WriteInit(b *byter.Byter)
	WriteData(b *byter.Byter)
	WriteUpdate(b *byter.Byter)
	WriteSynch(b *byter.Byter)

	ReadUpdate(reader *byter.Byter) error

	AddChild(object DRObject)
	Children() []DRObject
	GetChildrenFiltered(func(DRObject) bool) []DRObject
	GetChildByGCType(s string) DRObject
	GetChildByGCNativeType(s string) DRObject
	GetChildrenByGCNativeType(s string) []DRObject
	GetParentEntity() DRObject

	RemoveChild(object DRObject) bool

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
