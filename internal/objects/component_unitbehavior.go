package objects

import (
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/game/components/behavior"
	"RainbowRunner/internal/logging"
	"RainbowRunner/pkg"
	"RainbowRunner/pkg/byter"
	"encoding/hex"
	"errors"
	"fmt"
)

type UnitBehavior struct {
	*GCObject
}

type UnitBehaviorHandler struct {
	Behaviour *UnitBehavior
}

func (u *UnitBehaviorHandler) WriteInit(b *byter.Byter) {
	u.WriteInit(b)
}

func (u *UnitBehaviorHandler) WriteUpdate(b *byter.Byter) {
	u.WriteUpdate(b)
}

func (u *UnitBehaviorHandler) WriteSynch(b *byter.Byter) {
	u.WriteSynch(b)
}

func (u *UnitBehaviorHandler) ReadUpdate(reader *byter.Byter) error {
	return u.ReadUpdate(reader)
}

func (n *UnitBehavior) WriteInit(b *byter.Byter) {
	behav := behavior.NewBehavior()
	behav.Init(b, nil, nil)

	// UnitMover::readInit()
	// Flags
	// 0x04
	// 0x01
	unitMover := byte(0x00)
	b.WriteByte(unitMover)

	if unitMover&0x04 > 0 {
		b.WriteByte(0xFF)
	}

	if unitMover&0x01 > 0 {
		// 0x01 case
		b.WriteUInt32(0x01)
		b.WriteUInt32(0x01)
	}

	b.WriteUInt32(0x00)
	b.WriteUInt32(0x00)

	if unitMover&0x80 > 0 {
		b.WriteUInt32(0x00)
	}

	// Set to 2 for waypoints
	unitMover2 := byte(0) // Could potentially be waypoints?

	b.WriteByte(unitMover2)

	if unitMover2 == 2 {
		waypointCount := uint16(0x0002)
		b.WriteUInt16(waypointCount)

		for i := 0; i < int(waypointCount); i++ {
			// Vector2
			b.WriteUInt32(uint32(1000 * i))   // X?
			b.WriteUInt32(uint32(100000 * i)) // Y?
		}
	}

	// UnitBehavior::readInit()
	b.WriteByte(0xFF)
	b.WriteByte(0xFF)
	b.WriteByte(0xFF)
}

func handleClientMove(conn connections.Connection, reader *byter.Byter) {
	// This increments each time the server sends a MoveTo message
	// The client will then increment by 1 for every individual movement performed (clicking)
	updateNumber := reader.Byte()
	count := int(reader.Byte())
	pos := pkg.Vector2{}

	if logging.LoggingOpts.LogMoves {
		fmt.Printf("Received %d player moves unk val: %x\n", count, updateNumber)
	}

	avatar := Players.Players[conn.GetID()].CurrentCharacter.GetChildByGCNativeType("Avatar").(*Avatar)

	for i := 0; i < count; i++ {
		unk := reader.Byte()       // Unk
		rotation := reader.Int32() // Seems to be rotation

		degrees := float32((float64(rotation) / 0x17000) * 360)

		pos.X = reader.Int32()
		pos.Y = reader.Int32()

		avatar.ClientUpdateNumber = updateNumber
		if logging.LoggingOpts.LogMoves {
			fmt.Printf(
				"Player move 0x%x rotation 0x%x(%.2fdeg) (%d, %d) Hex (%x, %x)\n",
				unk, rotation, degrees, pos.X, pos.Y, pos.X, pos.Y,
			)
		}

		avatar.LastPosition = avatar.Position

		avatar.Position.X = pos.X
		avatar.Position.Y = pos.Y
		avatar.Position.Z = 0
		avatar.Rotation = rotation

		//conn.Player.SendPosition(unk)

		//conn.Player.MoveQueue.Add(MovementUpdate{
		//	Position: pos,
		//	Rotation: rotation,
		//	Tick:     Tick,
		//})

		if unk&0x02 > 0 {
			if logging.LoggingOpts.LogMoves {
				fmt.Println("player started moving")
			}
			avatar.IsMoving = true
			//conn.Player.SendPosition(0x02)
		}

		if unk&0x01 > 0 {
			if logging.LoggingOpts.LogMoves {
				fmt.Println("player finished moving")
			}
			avatar.IsMoving = false
			avatar.SendPosition()
		}
	}

	if avatar.MoveUpdate >= 0x2D {
		//fmt.Printf(
		//	"sending move update %d, %d || %x, %x!\n",
		//	pos.X, pos.Y,
		//	pos.X, pos.Y,
		//)
		//conn.Player.Move(pos.X, pos.Y)
		//conn.Player.SendFollowClient()
		avatar.MoveUpdate = 0
	}

	if logging.LoggingOpts.LogMoves {
		fmt.Printf("%s\n", hex.Dump(reader.Data()))
	}
}

func (g *UnitBehavior) ReadUpdate(reader *byter.Byter) error {
	subMessage := reader.Byte()
	switch subMessage {
	case 0x65:
		handleClientMove(g.EntityProperties.Conn, reader)
	// Potentially requesting current position because starting a new path
	case 0x03:
		fmt.Printf("player sent pre-path")
		Players.Players[int(g.EntityProperties.ID)].CurrentCharacter.GetChildByGCNativeType("Avatar").(*Avatar).SendPosition()
	default:
		fmt.Printf("unhandled client entity sub message %x", subMessage)
		return errors.New("unhandled unitbehavior update")
	}

	return nil
}

func NewUnitBehavior(gcType string) *UnitBehavior {
	gcObject := NewGCObject("UnitBehavior")
	gcObject.GCType = gcType
	gcObject.EntityHandler = &UnitBehaviorHandler{}

	return &UnitBehavior{
		GCObject: gcObject,
	}
}
