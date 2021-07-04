package objects

import (
	"RainbowRunner/internal/byter"
)

type UnitContainer struct {
	gcobject *GCObject
}

func (a UnitContainer) Serialise(byter *byter.Byter) {
	a.gcobject.Serialise(byter)

	manipulator := NewGCObject("Manipulator")
	//manipulator.GCType = "base.MeleeUnit.Manipulators.PrimaryWeapon"
	manipulator.Serialise(byter)
}

func (a UnitContainer) AddChild(object IGCObject) {
	a.gcobject.AddChild(object)
}

func NewUnitContainer(id uint32, name string) *UnitContainer {
	container := NewGCObject("UnitContainer")
	container.ID = id
	container.Name = name

	return &UnitContainer{
		gcobject: container,
	}
}
