package objects

import (
	byter "RainbowRunner/pkg/byter"
)

type UnitContainer struct {
	*Component

	Manipulator DRObject
	ActiveItem  *Equipment
}

func (a UnitContainer) WriteFullGCObject(byter *byter.Byter) {
	a.GCObject.WriteFullGCObject(byter)

	a.Manipulator.WriteFullGCObject(byter)
}

func (a *UnitContainer) SetActiveItem(item *Equipment) {
	a.ActiveItem = item
}

func (a UnitContainer) WriteSetActiveItem(body *byter.Byter) {
	CEWriter := NewClientEntityWriter(body)
	CEWriter.BeginComponentUpdate(a)
	// 0x29 clear item
	// 0x28 set active item
	CEWriter.Body.WriteByte(0x28)

	a.ActiveItem.WriteInit(CEWriter.Body)

	CEWriter.EndComponentUpdate(a)
}

func (a UnitContainer) WriteClearActiveItem(body *byter.Byter) {
	CEWriter := NewClientEntityWriter(body)
	CEWriter.BeginComponentUpdate(a)
	// 0x28 Add
	// 0x29 Remove
	CEWriter.Body.WriteByte(0x29)

	CEWriter.EndComponentUpdate(a)
}

func NewUnitContainer(manipulator DRObject, name string) *UnitContainer {
	container := NewComponent("unitcontainer", "UnitContainer")
	container.GCName = name

	if manipulator.RREntityProperties().ID == 0 {
		panic("Register component before passing it to unit container")
	}

	return &UnitContainer{
		Component:   container,
		Manipulator: manipulator,
	}
}
