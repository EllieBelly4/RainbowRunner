package objects

import (
	"RainbowRunner/internal/message"
	byter "RainbowRunner/pkg/byter"
	"errors"
	"fmt"
)

type UnitContainer struct {
	*Component

	Manipulator DRObject
	ActiveItem  DRObject
}

func (u *UnitContainer) ReadUpdate(body *byter.Byter) error {
	op := body.Byte()

	switch op {
	case 0x28:
		index := body.UInt32()
		inventory := u.GetChildByGCType("avatar.base.inventory")
		inventoryCast, ok := inventory.(*Inventory)

		if !ok {
			return errors.New(fmt.Sprintf("character does not have an inventory"))
		}

		item := inventoryCast.GetItemByIndex(int(index))

		if item == nil {
			return errors.New(fmt.Sprintf("could not find item in inventory with index '%d'", index))
		}

		CEWriter := NewClientEntityWriterWithByter()

		u.SetActiveItem(item)
		u.WriteSetActiveItem(CEWriter.Body)

		Players.GetPlayer(uint16(u.OwnerID())).MessageQueue.Enqueue(
			message.QueueTypeClientEntity, CEWriter.Body, message.OpTypeInventoryItemClickResponse,
		)

		// TODO remove item from inventory

	case 0x29:
		inventory := u.GetChildByGCType("avatar.base.inventory")
		_, ok := inventory.(*Inventory)

		if !ok {
			return errors.New(fmt.Sprintf("character does not have an inventory"))
		}

		CEWriter := NewClientEntityWriterWithByter()

		u.SetActiveItem(nil)
		u.WriteClearActiveItem(CEWriter.Body)

		Players.GetPlayer(uint16(u.OwnerID())).MessageQueue.Enqueue(
			message.QueueTypeClientEntity, CEWriter.Body, message.OpTypeInventoryItemClickResponse,
		)

		//TODO add item to inventory
	}

	body.Dump()

	return nil
}

func (u UnitContainer) WriteFullGCObject(byter *byter.Byter) {
	u.GCObject.WriteFullGCObject(byter)

	u.Manipulator.WriteFullGCObject(byter)
}

func (u *UnitContainer) SetActiveItem(item DRObject) {
	u.ActiveItem = item
}

func (u *UnitContainer) WriteSetActiveItem(body *byter.Byter) {
	CEWriter := NewClientEntityWriter(body)
	CEWriter.BeginComponentUpdate(u)
	// 0x29 clear item
	// 0x28 set active item
	CEWriter.Body.WriteByte(0x28)

	u.ActiveItem.WriteInit(CEWriter.Body)

	CEWriter.EndComponentUpdate(u)
}

func (u *UnitContainer) WriteClearActiveItem(body *byter.Byter) {
	CEWriter := NewClientEntityWriter(body)
	CEWriter.BeginComponentUpdate(u)
	// 0x28 Add
	// 0x29 Remove
	CEWriter.Body.WriteByte(0x29)

	CEWriter.EndComponentUpdate(u)
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
