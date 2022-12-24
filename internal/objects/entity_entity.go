package objects

import "RainbowRunner/internal/types/drobjecttypes"

//go:generate go run ../../scripts/generatelua -type=Entity -extends=GCObject
type Entity struct {
	*GCObject
}

func (e *Entity) AddChild(child drobjecttypes.DRObject) {
	e.GCObject.AddChild(child)
	child.SetParent(e)
}

func NewEntity(gctype string) *Entity {
	return &Entity{
		GCObject: NewGCObject(gctype),
	}
}
