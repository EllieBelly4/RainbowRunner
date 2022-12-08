package objects

import (
	"RainbowRunner/internal/message"
	"RainbowRunner/internal/types"
	byter "RainbowRunner/pkg/byter"
	"RainbowRunner/pkg/datatypes"
	"errors"
	"fmt"
)

type UnitContainer struct {
	*Component

	Manipulator DRObject
	ActiveItem  DRObject
	Avatar      *Avatar
}

func (u *UnitContainer) ReadUpdate(body *byter.Byter) error {
	zone := u.RREntityProperties().Zone
	op := body.Byte()

	switch op {
	// Drop item on floor
	case 0x23:
		CEWriter := NewClientEntityWriterWithByter()

		item := u.ActiveItem

		if item == nil {
			return errors.New("cannot drop any item when no active item is selected")
		}

		u.SetActiveItem(nil)
		u.WriteClearActiveItem(CEWriter.Body)

		avatarUnitBehaviour := u.Avatar.GetUnitBehaviour()

		itemObject := NewItemObject("itemobject", item)
		itemObject.RREntityProperties().SetOwner(u.OwnerID())
		itemObject.WorldPosition = avatarUnitBehaviour.Position
		zone.AddEntity(types.UInt16(u.OwnerID()), itemObject)

		fmt.Printf("Avatar Pos: %d, %d", avatarUnitBehaviour.Position.X, avatarUnitBehaviour.Position.Y)

		CEWriter.Create(itemObject)
		CEWriter.Init(itemObject)

		Players.GetPlayer(uint16(u.OwnerID())).MessageQueue.Enqueue(
			message.QueueTypeClientEntity, CEWriter.Body, message.OpTypeInventoryItemDropResponse,
		)
	// Pickup item in inventory
	case 0x28:
		err := u.handlePickupItemFromInventory(body)

		if err != nil {
			return err
		}

	// Place item in inventory
	case 0x29:
		err := u.handlePlaceItemInInventory(body)

		if err != nil {
			return err
		}
	default:
		return errors.New(fmt.Sprintf("unhandled unitcontainer message: %d", op))
	}

	return nil
}

func (u *UnitContainer) handlePlaceItemInInventory(body *byter.Byter) error {
	inventoryID := body.Byte()
	inventory := u.GetInventoryByID(inventoryID)

	if inventory == nil {
		return errors.New(fmt.Sprintf("character does not have an inventory"))
	}

	CEWriter := NewClientEntityWriterWithByter()

	item := u.ActiveItem

	if item == nil {
		return errors.New(fmt.Sprintf("character does not have an active item"))
	}

	u.SetActiveItem(nil)
	u.WriteClearActiveItem(CEWriter.Body)

	x := body.Byte()
	y := body.Byte()

	inventory.AddChild(item)
	u.WriteAddItem(CEWriter.Body, item, inventory, x, y)

	Players.GetPlayer(uint16(u.OwnerID())).MessageQueue.Enqueue(
		message.QueueTypeClientEntity, CEWriter.Body, message.OpTypeInventoryItemClickResponse,
	)
	return nil
}

func (u *UnitContainer) handlePickupItemFromInventory(body *byter.Byter) error {
	index := body.UInt32()
	inventory := u.GetChildByGCType("avatar.base.inventory")
	inventoryCast, ok := inventory.(*Inventory)

	if !ok {
		return errors.New(fmt.Sprintf("character does not have an inventory"))
	}

	item := inventoryCast.RemoveItemByIndex(int(index))

	if item == nil {
		return errors.New(fmt.Sprintf("could not find item in inventory with index '%d'", index))
	}

	CEWriter := NewClientEntityWriterWithByter()

	u.WriteRemoveItem(CEWriter.Body, index)

	u.SetActiveItem(item)
	u.WriteSetActiveItem(CEWriter.Body)

	Players.GetPlayer(uint16(u.OwnerID())).MessageQueue.Enqueue(
		message.QueueTypeClientEntity, CEWriter.Body, message.OpTypeInventoryItemClickResponse,
	)
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

func (u *UnitContainer) WriteRemoveItem(body *byter.Byter, index uint32) {
	CEWriter := NewClientEntityWriter(body)
	CEWriter.BeginComponentUpdate(u)

	// 0x21 Remove Currency
	// 0x20 Add Currency
	// 0x1F Remove Item
	// 0x1E Add Item
	CEWriter.Body.WriteByte(0x1F)
	CEWriter.Body.WriteUInt32(index)

	CEWriter.EndComponentUpdate(u)
}

func (u *UnitContainer) WriteAddItem(body *byter.Byter, item DRObject, inventory *Inventory, x, y byte) {
	CEWriter := NewClientEntityWriter(body)
	CEWriter.BeginComponentUpdate(u)

	// 0x21 Remove Currency
	// 0x20 Add Currency
	// 0x1F Remove Item
	// 0x1E Add Item
	CEWriter.Body.WriteByte(0x1E)
	// Inventory ID (not the same as GCObject ID)
	CEWriter.Body.WriteByte(inventory.InventoryID)

	if drItem, ok := item.(DRItem); ok {
		drItem.SetInventoryPosition(datatypes.Vector2{
			X: int32(x),
			Y: int32(y),
		})
	}

	item.WriteInit(CEWriter.Body)
	CEWriter.EndComponentUpdate(u)
}

func (u *UnitContainer) GetInventoryByID(index byte) *Inventory {
	for _, child := range u.GCChildren {
		if inventory, ok := child.(*Inventory); ok {
			if inventory.InventoryID == index {
				return inventory
			}
		}
	}

	return nil
}

func NewUnitContainer(manipulator DRObject, name string, avatar *Avatar) *UnitContainer {
	container := NewComponent("unitcontainer", "UnitContainer")
	container.GCLabel = name

	//if manipulator.RREntityProperties().ID == 0 {
	//	panic("Register component before passing it to unit container")
	//}

	return &UnitContainer{
		Component:   container,
		Manipulator: manipulator,
		Avatar:      avatar,
	}
}
