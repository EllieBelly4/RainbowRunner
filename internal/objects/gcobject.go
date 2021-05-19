package objects

import (
	"RainbowRunner/internal/byter"
	"strings"
)

type GCObject struct {
	Version    uint8
	NativeType string
	ID         uint32
	Name       string
	Children   []*GCObject
	GCType     string
	Properties []GCObjectProperty
}

type GCObjectProperty struct {
	Name  string
	Value interface{}
}

//func Uint8Prop(name string, val uint8) GCObjectProperty {
//	return GCObjectProperty{
//		Name:  name,
//		Value: val,
//	}
//}
//
//func Uint16Prop(name string, val uint16) GCObjectProperty {
//	return GCObjectProperty{
//		Name:  name,
//		Value: val,
//	}
//}

func Uint32Prop(name string, val uint32) GCObjectProperty {
	return GCObjectProperty{
		Name:  name,
		Value: val,
	}
}

//func StringProp(name string, val string) GCObjectProperty {
//	return GCObjectProperty{
//		Name:  name,
//		Value: val,
//	}
//}

func NewGCObject(nativeType string) *GCObject {
	return &GCObject{
		// This is the latest version
		Version:    0x13,
		NativeType: nativeType,
		GCType:     strings.ToLower(nativeType),
	}
}

func (o GCObject) Serialise(byter *byter.Byter) {
	byter.WriteByte(o.Version)
	byter.WriteCString(o.NativeType)
	byter.WriteUInt32(o.ID)
	byter.WriteCString(o.Name)

	byter.WriteUInt32(uint32(len(o.Children)))

	for _, child := range o.Children {
		child.Serialise(byter)
	}

	byter.WriteCString(o.GCType)

	for _, prop := range o.Properties {
		prop.Serialise(byter)
	}

	byter.WriteNull()
}

func (o *GCObject) AddChild(modifiers *GCObject) {
	if o.Children == nil {
		o.Children = make([]*GCObject, 0, 128)
	}

	o.Children = append(o.Children, modifiers)
}

func (p GCObjectProperty) Serialise(b *byter.Byter) {
	b.WriteCString(p.Name)

	switch p.Value.(type) {
	case string:
		b.WriteCString(p.Value.(string))
	case uint32:
		b.WriteUInt32(p.Value.(uint32))
	case uint8:
		b.WriteByte(p.Value.(uint8))
	case uint16:
		b.WriteUInt16(p.Value.(uint16))
	}
}
