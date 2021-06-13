package objects

import (
	"RainbowRunner/internal/byter"
	"RainbowRunner/internal/logging"
	"fmt"
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
	Name  interface{}
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

func Uint32Prop(name interface{}, val uint32) GCObjectProperty {
	return GCObjectProperty{
		Name:  name,
		Value: val,
	}
}

func StringProp(name interface{}, val string) GCObjectProperty {
	return GCObjectProperty{
		Name:  name,
		Value: val,
	}
}

func NewGCObject(nativeType string) *GCObject {
	return &GCObject{
		// At version 2A or above you must use a hash I think
		//Version:    0x29, // No hash required
		Version:    0x2D,
		NativeType: nativeType,
		GCType:     strings.ToLower(nativeType),
	}
}

func (o GCObject) Serialise(byter *byter.Byter) {
	byter.WriteByte(o.Version)

	useHashes := o.Version >= 0x2a

	if useHashes {
		byter.WriteUInt32(GetTypeHash(o.NativeType))
	} else {
		byter.WriteCString(o.NativeType)
	}

	byter.WriteUInt32(o.ID)
	byter.WriteCString(o.Name)

	byter.WriteUInt32(uint32(len(o.Children)))

	for _, child := range o.Children {
		child.Serialise(byter)
	}

	if useHashes {
		byter.WriteUInt32(GetTypeHash(o.GCType))
	} else {
		byter.WriteCString(o.GCType)
	}

	for _, prop := range o.Properties {
		prop.Serialise(byter, useHashes)
	}

	byter.WriteUInt32(0)
}

func (o *GCObject) AddChild(child *GCObject) {
	if o.Children == nil {
		o.Children = make([]*GCObject, 0, 128)
	}

	o.Children = append(o.Children, child)
}

func (p GCObjectProperty) Serialise(b *byter.Byter, useHash bool) {
	switch name := p.Name.(type) {
	case string:
		if useHash {
			b.WriteUInt32(GetTypeHash(name))
		} else {
			b.WriteCString(name)
		}
	case int:
		b.WriteUInt32(uint32(name))
	case uint32:
		b.WriteUInt32(name)
	}

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

func GetTypeHash(name string) uint32 {
	result := uint32(5381) // eax

	a1 := len(name)

	if a1 > 0 {
		for _, v4 := range name {
			if v4 >= 0x41 && v4 <= 0x5A {
				v4 = v4 + 32
			}

			result += uint32(v4) + 32*result
		}

		if result == 0 {
			result = 1
		}
	}

	if logging.LoggingOpts.LogHashes {
		fmt.Printf("(%x) %s\n", result, name)
	}

	return result
}
