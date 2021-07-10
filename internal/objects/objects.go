package objects

import (
	"RainbowRunner/pkg/byter"
)

type DRObject interface {
	RREntityProperties() *RREntityProperties

	WriteFullGCObject(b *byter.Byter)
	WriteInit(b *byter.Byter)
	WriteUpdate(b *byter.Byter)

	AddChild(object DRObject)
	Children() []DRObject
	GetChildByGCType(s string) DRObject
	GetChildByGCNativeType(s string) DRObject

	GetGCObject() *GCObject
}
