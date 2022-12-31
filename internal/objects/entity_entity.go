package objects

import (
	"RainbowRunner/internal/types/drobjecttypes"
	"strings"
)

//go:generate go run ../../scripts/generatelua -type=Entity -extends=GCObject
type Entity struct {
	*GCObject
	Name string
}

func (e *Entity) GetName() string {
	return e.Name
}

func (e *Entity) AddChild(child drobjecttypes.DRObject) {
	e.GCObject.AddChild(child)
	child.SetParent(e)
}

func NewEntity(gctype string) *Entity {
	splitGCType := strings.Split(gctype, ".")

	return &Entity{
		GCObject: NewGCObject(gctype),
		Name:     splitGCType[len(splitGCType)-1],
	}
}
