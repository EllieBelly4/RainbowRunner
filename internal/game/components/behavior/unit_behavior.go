package behavior

import (
	byter "RainbowRunner/pkg/byter"
)

type UnitBehavior struct {
	Action Action
	ID     uint16
}

func NewUnitBehavior(ID uint16) *UnitBehavior {
	return &UnitBehavior{
		ID: ID,
	}
}

func (u *UnitBehavior) AddUpdate(body *byter.Byter) {
	// UnitBehavior::processUpdate
	body.WriteUInt16(u.ID) // Component ID?

	// Command

	// 0x64 - something to do with client control
	// 0x65 - UnitMoverUpdate::read
	// 0x66 - Behavior::

	if u.Action == nil {
		// This appears to move the player onto the nearest valid position
		// 0x64 - something to do with client control
		body.WriteByte(0x64)
		body.WriteByte(0x01) // > 0 Causes the player to snap to a ground point sometimes, or go back to 0,0,0
	} else {
		body.WriteByte(0x66)
		u.Action.Init(body)
	}

	// UnitBehavior::processUpdate
	// 65 # Command: UnitMoverUpdate::read

	// 05 # Unk
	// 01 # Unk UnitBehavior::processUpdate, if 2 it fails

	// UnitMoverUpdate::Read
	// 06 # Unk
	// 02 02 02 02 # PosX?
	// 03 03 03 03 # PosY?
	// 04 04 04 04 # PosZ?

	// Synch
	// 02 00 00 00 00
	// 06 # End
}
