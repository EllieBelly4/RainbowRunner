package objects

import (
	"strings"
)

//go:generate go run ../../scripts/generatelua -type=Component -extends=GCObject
type Component struct {
	*GCObject
}

func (Component) Type() DRObjectType {
	return DRObjectComponent
}

func NewComponent(gcType string, nativeType string) *Component {
	gcObject := NewGCObject(nativeType)
	gcObject.GCType = strings.ToLower(gcType)

	return &Component{
		GCObject: gcObject,
	}
}
