package objects

import (
	byter "RainbowRunner/pkg/byter"
)

type UnitContainer struct {
	*GCObject

	Manipulator DRObject
}

func (a UnitContainer) WriteFullGCObject(byter *byter.Byter) {
	a.GCObject.WriteFullGCObject(byter)

	a.Manipulator.WriteFullGCObject(byter)
}

func NewUnitContainer(manipulator DRObject, name string) *UnitContainer {
	container := NewGCObject("UnitContainer")
	container.GCName = name

	if manipulator.RREntityProperties().ID == 0 {
		panic("Register component before passing it to unit container")
	}

	return &UnitContainer{
		GCObject:    container,
		Manipulator: manipulator,
	}
}
