package objects

import (
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/game/components/behavior"
	"RainbowRunner/internal/helpers"
	"RainbowRunner/internal/logging"
	"RainbowRunner/pkg"
	"RainbowRunner/pkg/byter"
	"encoding/hex"
	"errors"
	"fmt"
)

type UnitBehavior struct {
	*Component
	LastPosition   pkg.Vector3
	Position       pkg.Vector3
	Rotation       int32
	UnitMoverFlags byte
	Action1        behavior.Action
	Action2        behavior.Action
}

type UnitBehaviorHandler struct {
	*UnitBehavior
}

func (u *UnitBehaviorHandler) WriteInit(b *byter.Byter) {
	u.WriteInit(b)
}

func (u *UnitBehaviorHandler) WriteUpdate(b *byter.Byter) {
	u.WriteUpdate(b)
}

func (u *UnitBehavior) WriteMoveUpdate(b *byter.Byter) {
	positions := []UnitPathPosition{
		{
			Position: u.Position.ToVector2(),
			Rotation: u.Rotation,
		},
	}

	if len(positions) > 0xFF {
		panic("cannot send more than 255 position updates in a single message")
	}

	//writer := NewClientEntityWriterWithByter()
	//writer.BeginStream()
	//writer.BeginComponentUpdate(n)

	b.WriteByte(0x65) // UnitMoverUpdate

	// UnitBehavior::processUpdate
	// Potentially move type? only 0xFF works with players?
	b.WriteByte(0xFF) // Unk

	updateCount := byte(len(positions))
	b.WriteByte(updateCount) // Update count

	// UnitMoverUpdate::Read

	for _, position := range positions {
		// 0x08 - does not hit any flags
		// 0x01 - hits 1 flag
		// 0x02 - hits 1 flag
		// 0x04 - hits 1 flag BROKEN
		//
		// What does it do with this value?
		// var a = 0x08
		// var someVal = 0x78 (dynamic?)
		//
		// var result = a ^ someVal // 0x70
		// result = result & 0x07 // 0x00
		// result = result ^ someVal // 0x78
		//
		b.WriteByte(0x01 | 0x02) // Not all values work
		b.WriteInt32(position.Rotation)
		b.WriteInt32(position.Position.X)
		b.WriteInt32(position.Position.Y)

		degrees := float32((float64(position.Rotation) / 0x17000) * 360)

		if logging.LoggingOpts.LogMoves && logging.LoggingOpts.LogGenericSent {
			fmt.Printf(
				"Sending move rotation 0x%x(%.2fdeg) (%d, %d) Hex (%x, %x)\n",
				position.Rotation, degrees, position.Position.X, position.Position.Y, position.Position.X, position.Position.Y,
			)
		}
	}

	//b.WriteByte(0x02)
	//b.WriteUInt32(uint32(global.Tick)) // Random unk value

	oldLog := logging.LoggingOpts.LogGenericSent

	if !logging.LoggingOpts.LogMoves {
		logging.LoggingOpts.LogGenericSent = false
	}

	//if n.RREntityProperties().Zone != nil {
	//	n.RREntityProperties().Zone.SendToAll(b)
	//}

	logging.LoggingOpts.LogGenericSent = oldLog
}

func (u *UnitBehaviorHandler) WriteSynch(b *byter.Byter) {
	u.WriteSynch(b)
}

func (u *UnitBehaviorHandler) ReadUpdate(reader *byter.Byter) error {
	return u.ReadUpdate(reader)
}

func (n *UnitBehavior) WriteInit(b *byter.Byter) {
	behav := behavior.NewBehavior()
	behav.Init(b, n.Action1, n.Action2)

	// UnitMover::readInit()
	// Flags
	// 0x04
	// 0x01
	unitMover := n.UnitMoverFlags
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
	// TODO look into waypoints as movement targets, RTS movement would always be based on waypoints
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

type UnitPathPosition struct {
	Position pkg.Vector2
	Rotation int32
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

	responseMoves := make([]UnitPathPosition, 0)

	avatar := Players.Players[conn.GetID()].CurrentCharacter.GetChildByGCNativeType("Avatar").(*Avatar)

	for i := 0; i < count; i++ {
		moveUpdateType := reader.Byte() // Unk
		rotation := reader.Int32()      // Seems to be rotation

		//degrees := float32((float64(rotation) / 0x17000) * 360)
		degrees := float32(rotation / 256)

		pos.X = reader.Int32()
		pos.Y = reader.Int32()

		avatar.ClientUpdateNumber = updateNumber
		if logging.LoggingOpts.LogMoves {
			fmt.Printf(
				"Player move 0x%x rotation 0x%x(%.2fdeg) (%d, %d) Hex (%x, %x)\n",
				moveUpdateType, rotation, degrees, pos.X, pos.Y, pos.X, pos.Y,
			)
		}

		g.LastPosition = g.Position

		g.Position.X = pos.X
		g.Position.Y = pos.Y
		g.Position.Z = 0
		g.Rotation = rotation

		avatar.LastPosition = g.LastPosition
		avatar.Position = g.Position

		//conn.Player.SendPosition(moveUpdateType)

		//conn.Player.MoveQueue.Add(MovementUpdate{
		//	Position: pos,
		//	Rotation: rotation,
		//	Tick:     Tick,
		//})

		if moveUpdateType&0x02 > 0 {
			if logging.LoggingOpts.LogMoves {
				fmt.Println("player started moving")
			}
			avatar.IsMoving = true
			//conn.Player.SendPosition(0x02)
		} else if moveUpdateType&0x01 > 0 {
			if logging.LoggingOpts.LogMoves {
				fmt.Println("player finished moving")
			}
			avatar.IsMoving = false
			//avatar.SendPosition()
		}

		responseMoves = append(responseMoves, UnitPathPosition{
			Rotation: rotation,
			Position: pkg.Vector2{
				X: pos.X,
				Y: pos.Y,
			},
		})
	}

	//TODO replace this
	//avatar.SendPositions(responseMoves)

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
	switch int(subMessage) {
	case 0x01:
		g.handleClientAttack(reader)
	case 0x65:
		g.handleClientMove(g.EntityProperties.Conn, reader)
	// Potentially requesting current position because starting a new path
	case 0x03:
		fmt.Printf("player send move request\n")
		// This is required to be handled so the player can move after getting stuck due to attacking
		// TODO investigate this behaviour
		//Players.Players[g.RREntityProperties().Conn.GetID()].CurrentCharacter.GetChildByGCNativeType("Avatar").(*Avatar).SendPosition()
	default:
		fmt.Printf("unhandled client entity sub message %x\n", subMessage)
		return errors.New("unhandled unitbehavior update\n")
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
	writer.BeginComponentUpdate(n)

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

//func (n *UnitBehavior) SendPositions(positions []UnitPathPosition) {
//
//}

func (n *UnitBehavior) handleClientAttack(reader *byter.Byter) {
	reader.DumpRemaining()

	writer := NewClientEntityWriterWithByter()

	writer.BeginStream()
	writer.BeginComponentUpdate(n)

	//00000000  07 34 b4 00 01 02 51 01  0a 76 b9 01 00 3d 4e ff  |.4....Q..v...=N.|
	//00000010  ff ec 31 00 00                                    |..1..|

	//00000000  02 51 01 0a 0b 90 01 00  88 4d ff ff ec 31 00 00  |.Q.......M...1..|

	writer.EndComponentUpdate(n)
	writer.EndStream()

	helpers.WriteCompressedASimple(n.RREntityProperties().Conn, writer.Body)
}

func NewUnitBehavior(gcType string) *UnitBehavior {
	component := NewComponent(gcType, "UnitBehavior")
	component.EntityHandler = &UnitBehaviorHandler{}

	return &UnitBehavior{
		Component: component,
	}
}
