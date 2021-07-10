package objects

import (
	"RainbowRunner/pkg/byter"
)

type Avatar struct {
	*GCObject
}

func NewAvatar(gcType string) *Avatar {
	a := &Avatar{
		GCObject: NewGCObject("Avatar"),
	}

	a.GCType = gcType
	a.GCName = "EllieAvatar"

	return a
}

func (a *Avatar) WriteFullGCObject(byter *byter.Byter) {
	//p.Properties = []GCObjectProperty{
	//	StringProp("Name", p.Name),
	//}

	a.GCObject.WriteFullGCObject(byter)
}

func (a Avatar) WriteInit(b *byter.Byter) {
	panic("implement me")
}

func (a Avatar) WriteUpdate(b *byter.Byter) {
	panic("implement me")
}
