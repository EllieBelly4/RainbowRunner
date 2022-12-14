package objects

import (
	"RainbowRunner/internal/types/drobjecttypes"
	"strings"
)

//go:generate go run ../../scripts/generatelua -type=Component -extends=GCObject
type Component struct {
	*GCObject
}

func (Component) Type() drobjecttypes.DRObjectType {
	return drobjecttypes.DRObjectComponent
}

func NewComponent(gcType string, nativeType string) *Component {
	gcObject := NewGCObject(nativeType)
	gcObject.GCType = strings.ToLower(gcType)

	return &Component{
		GCObject: gcObject,
	}
}
