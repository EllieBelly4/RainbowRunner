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
	"time"
)

type UnitBehavior struct {
	*GCObject
	LastPosition pkg.Vector3
	Position     pkg.Vector3
	Rotation     int32
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

func (g *UnitBehavior) handleClientMove(conn connections.Connection, reader *byter.Byter) {
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

		g.LastPosition = g.Position

		g.Position.X = pos.X
		g.Position.Y = pos.Y
		g.Position.Z = 0
		g.Rotation = rotation

		avatar.LastPosition = g.LastPosition
		avatar.Position = g.Position

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
		g.handleClientMove(g.EntityProperties.Conn, reader)
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

func (n *UnitBehavior) Warp(x int32, y int32, z int32) {
	n.Position.X = x
	n.Position.Y = y
	n.Position.Z = z

	if n.RREntityProperties().Conn != nil {
		n.sendWarpTo(x, y, z)
	}
}

func (n *UnitBehavior) sendWarpTo(posX, posY, posZ int32) {
	writer := NewClientEntityWriterWithByter()
	writer.BeginStream()
	writer.BeginComponentUpdate(n.RREntityProperties().ID)

	writer.Body.WriteByte(0x04) // CreateAction1
	writer.Body.WriteByte(17)
	writer.Body.WriteByte(0x00)
	writer.Body.WriteInt32(posX)
	writer.Body.WriteInt32(posY)
	writer.Body.WriteInt32(posZ)

	writer.WriteSynch(n)
	writer.EndStream()

	if n.RREntityProperties().Zone != nil {
		n.RREntityProperties().Zone.SendToAll(writer.Body)
	}
}

func (n *UnitBehavior) SendPosition() {
	//# UnitBehavior - UnitMoverUpdate::read
	//35 # ComponentUpdate
	//05 00 # Component ID
	//# Command
	//# 05 - Behavior::terminateAllActionsLocal
	//# 65 - UnitMoverUpdate::read
	//65 # Command
	//05 # Unk UnitBehavior::processUpdate
	//01 # Unk UnitBehavior::processUpdate, if 2 it fails
	//06 # Unk
	//10 10 00 00 # PosX?
	//00 10 10 00 # PosY?
	//00 10 00 00 # PosZ?
	//02 00 7e 04 00 # Synch
	//06 # End

	writer := NewClientEntityWriterWithByter()
	writer.BeginStream()
	writer.BeginComponentUpdate(n.RREntityProperties().ID)

	writer.Body.WriteByte(0x65) // UnitMoverUpdate

	updateCount := 10

	// UnitBehavior::processUpdate
	writer.Body.WriteByte(0xFF)              // Unk
	writer.Body.WriteByte(byte(updateCount)) // Update count

	// UnitMoverUpdate::Read
	//writer.Body.WriteByte(0x08) // Not all values work
	//writer.Body.WriteByte(0x01) // Not all values work

	for i := 0; i < updateCount; i++ {
		writer.Body.WriteByte(0x08) // Not all values work
		writer.Body.WriteInt32(n.Rotation)
		writer.Body.WriteInt32(n.Position.X)
		writer.Body.WriteInt32(n.Position.Y)
	}

	//writer.Body.WriteInt32(0)
	//writer.Body.WriteInt32(0)
	//writer.Body.WriteInt32(0)

	writer.Body.WriteByte(0x02)
	writer.Body.WriteUInt32(uint32(time.Now().Unix())) // Random unk value

	//AddSynch(p.Conn, writer.Body)

	degrees := float32((float64(n.Rotation) / 0x17000) * 360)

	if logging.LoggingOpts.LogMoves {
		fmt.Printf(
			"Sending move rotation 0x%x(%.2fdeg) (%d, %d) Hex (%x, %x)\n",
			n.Rotation, degrees, n.Position.X, n.Position.Y, n.Position.X, n.Position.Y,
		)
	}

	writer.EndStream()

	if n.RREntityProperties().Zone != nil {
		n.RREntityProperties().Zone.SendToAll(writer.Body)
	}
}

func NewUnitBehavior(gcType string) *UnitBehavior {
	gcObject := NewGCObject("UnitBehavior")
	gcObject.GCType = gcType
	gcObject.EntityHandler = &UnitBehaviorHandler{}

	return &UnitBehavior{
		GCObject: gcObject,
	}
}
