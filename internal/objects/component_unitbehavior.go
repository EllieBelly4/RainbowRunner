package objects

import (
	"RainbowRunner/internal/config"
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/game/components/behavior"
	"RainbowRunner/internal/message"
	"RainbowRunner/pkg/byter"
	"RainbowRunner/pkg/datatypes"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
)

type IUnitBehavior interface {
	GetUnitBehavior() *UnitBehavior
}

type UnitBehavior struct {
	*Component
	LastPosition   datatypes.Vector3Float32
	Position       datatypes.Vector3Float32
	Rotation       float32
	UnitMoverFlags byte
	Action1        behavior.Action
	Action2        behavior.Action
}

func (u *UnitBehavior) GetUnitBehavior() *UnitBehavior {
	return u
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
			Position: u.Position.ToVector2Float32(),
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
		b.WriteInt32(int32(position.Rotation * 256))
		b.WriteInt32(int32(position.Position.X * 256))
		b.WriteInt32(int32(position.Position.Y * 256))

		if config.Config.Logging.LogMoves {
			fmt.Printf(
				"Sending move rotation 0x%x(%.2fdeg) (%d, %d) Hex (%x, %x)\n",
				int32(position.Rotation*256), position.Rotation, position.Position.X, position.Position.Y, position.Position.X, position.Position.Y,
			)
		}
	}

	//b.WriteByte(0x02)
	//b.WriteUInt32(uint32(global.Tick)) // Random unk value

	oldLog := config.Config.Logging.LogGenericSent

	if !config.Config.Logging.LogMoves {
		config.Config.Logging.LogGenericSent = false
	}

	//if n.RREntityProperties().Zone != nil {
	//	n.RREntityProperties().Zone.SendToAll(b)
	//}

	config.Config.Logging.LogGenericSent = oldLog
}

func (u *UnitBehaviorHandler) WriteSynch(b *byter.Byter) {
	u.WriteSynch(b)
}

func (u *UnitBehaviorHandler) ReadUpdate(reader *byter.Byter) error {
	return u.ReadUpdate(reader)
}

func (u *UnitBehavior) WriteInit(b *byter.Byter) {
	behav := behavior.NewBehavior()
	behav.Init(b, u.Action1, u.Action2)

	// UnitMover::readInit()
	// Flags
	// 0x04
	// 0x01
	unitMover := u.UnitMoverFlags
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
	Position   datatypes.Vector2Float32
	Rotation   float32
	ResponseID byte
}

func (u *UnitBehavior) handleClientMove(conn connections.Connection, reader *byter.Byter) {
	// This increments each time the server sends a MoveTo message
	// The client will then increment by 1 for every individual movement performed (clicking)
	responseIDMaybe := reader.Byte()
	count := int(reader.Byte())
	pos := datatypes.Vector2Float32{}

	if config.Config.Logging.LogMoves {
		fmt.Printf("Received %d player moves unk val: %x\n", count, responseIDMaybe)
	}

	responseMoves := make([]UnitPathPosition, 0)

	avatar := Players.Players[conn.GetID()].CurrentCharacter.GetChildByGCNativeType("Avatar").(*Avatar)

	for i := 0; i < count; i++ {
		moveUpdateType := reader.Byte()     // Unk
		rotation := float32(reader.Int32()) // Seems to be rotation

		//degrees := float32((float64(rotation) / 0x17000) * 360)
		degrees := float32(rotation / 256)

		pos.X = float32(reader.Int32()) / 256
		pos.Y = float32(reader.Int32()) / 256

		avatar.ClientUpdateNumber = responseIDMaybe
		if config.Config.Logging.LogReceivedMoves {
			//xf := float32(pos.X>>8) + (float32(pos.X&0xFF) / 256)
			//yf := float32(pos.Y>>8) + (float32(pos.Y&0xFF) / 256)
			fmt.Printf(
				"Player move 0x%x rotation 0x%x(%.2fdeg) (%f, %f) Hex (%x, %x)\n",
				moveUpdateType, rotation, degrees, pos.X, pos.Y, pos.X, pos.Y,
			)
		}

		u.LastPosition = u.Position

		u.Position.X = pos.X
		u.Position.Y = pos.Y
		u.Position.Z = 0
		u.Rotation = degrees

		//conn.Player.SendPosition(moveUpdateType)

		//conn.Player.MoveQueue.Add(MovementUpdate{
		//	Position: pos,
		//	Rotation: rotation,
		//	Tick:     Tick,
		//})

		if moveUpdateType&0x02 > 0 {
			if config.Config.Logging.LogMoves {
				fmt.Println("player started moving")
			}
			avatar.IsMoving = true
			//conn.Player.SendPosition(0x02)
		} else if moveUpdateType&0x01 > 0 {
			if config.Config.Logging.LogMoves {
				fmt.Println("player finished moving")
			}
			avatar.IsMoving = false
			//avatar.SendPosition()
		}

		responseMoves = append(responseMoves, UnitPathPosition{
			ResponseID: responseIDMaybe,
			Rotation:   rotation,
			Position: datatypes.Vector2Float32{
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

	if config.Config.Logging.LogMoves {
		logrus.Infof("\n%s\n", hex.Dump(reader.Data()))
	}
}

func (u *UnitBehavior) ReadUpdate(reader *byter.Byter) error {
	subMessage := reader.Byte()
	switch int(subMessage) {
	case 0x01: // Execute Action?
		return u.handleExecuteAction(reader)
		//u.handleClientBlockMovement(reader)
		//u.handleClientAttack(reader)
	case 0x65:
		u.handleClientMove(u.EntityProperties.Conn, reader)
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

func (u *UnitBehavior) Warp(x, y, z float32) {
	u.Position.X = x
	u.Position.Y = y
	u.Position.Z = z

	if u.RREntityProperties().Conn != nil {
		u.sendWarpTo(x, y, z)
	}
}

func (u *UnitBehavior) sendWarpTo(posX, posY, posZ float32) {
	writer := NewClientEntityWriterWithByter()
	writer.BeginStream()
	writer.BeginComponentUpdate(u)

	writer.Body.WriteByte(0x04) // CreateAction1
	writer.Body.WriteByte(17)
	writer.Body.WriteByte(0x00)
	writer.Body.WriteInt32(int32(posX * 256))
	writer.Body.WriteInt32(int32(posY * 256))
	writer.Body.WriteInt32(int32(posZ * 256))

	writer.WriteSynch(u)
	writer.EndStream()

	if u.RREntityProperties().Zone != nil {
		u.RREntityProperties().Zone.SendToAll(writer.Body)
	}
}

//func (n *UnitBehavior) SendPositions(positions []UnitPathPosition) {
//
//}

func (u *UnitBehavior) handleClientAttack(reader *byter.Byter) {
	reader.DumpRemaining()

	writer := NewClientEntityWriterWithByter()

	writer.BeginStream()
	writer.BeginComponentUpdate(u)

	//00000000  07 34 b4 00 01 02 51 01  0a 76 b9 01 00 3d 4e ff  |.4....Q..v...=N.|
	//00000010  ff ec 31 00 00                                    |..1..|

	//00000000  02 51 01 0a 0b 90 01 00  88 4d ff ff ec 31 00 00  |.Q.......M...1..|

	writer.EndComponentUpdate(u)
	writer.EndStream()

	connections.WriteCompressedASimple(u.RREntityProperties().Conn, writer.Body)
}

func (u *UnitBehavior) WriteWarp(writer *ClientEntityWriter) {
	writer.BeginComponentUpdate(u)

	writer.Body.WriteByte(0x04) // CreateAction1
	writer.Body.WriteByte(17)
	writer.Body.WriteByte(0x00)
	writer.Body.WriteInt32(int32(u.Position.X * 256))
	writer.Body.WriteInt32(int32(u.Position.Y * 256))
	writer.Body.WriteInt32(int32(u.Position.Z * 256))

	writer.WriteSynch(u)
	//writer.EndComponentUpdate(u)
}

func (u *UnitBehavior) handleClientBlockMovement(reader *byter.Byter) {

}

func (u *UnitBehavior) handleExecuteAction(reader *byter.Byter) error {
	msgResponseIdMaybe := reader.Byte()
	action := behavior.BehaviourAction(reader.Byte())

	logrus.Infof("execute action %s, unk0 %d", action.String(), msgResponseIdMaybe)

	switch action {
	case behavior.BehaviourActionActivate:
		return u.handleExecuteActivate(reader, msgResponseIdMaybe)
	}

	return nil
}

func (u *UnitBehavior) handleExecuteActivate(reader *byter.Byter, responseID byte) error {
	msgIdMaybe := reader.Byte()

	logrus.Infof("execute Activate responseID %x", msgIdMaybe)

	targetID := reader.UInt16()
	targetEntity := Entities.FindByID(targetID)

	if targetEntity == nil {
		return errors.New(fmt.Sprintf("could not find target entity with ID %d", targetID))
	}

	CEWriter := NewClientEntityWriterWithByter()

	CEWriter.BeginComponentUpdate(u)

	CEWriter.CreateActionResponse(behavior.BehaviourActionActivate, responseID)

	activateAction := behavior.Activate{
		TargetEntityID: targetID,
	}

	activateAction.InitWithoutOpCode(CEWriter.Body)

	CEWriter.WriteSynch(u)

	Players.GetPlayer(u.OwnerID()).MessageQueue.Enqueue(
		message.QueueTypeClientEntity, CEWriter.Body, message.OpTypeBehaviourAction,
	)

	return nil
}

func NewUnitBehavior(gcType string) *UnitBehavior {
	component := NewComponent(gcType, "UnitBehavior")
	component.EntityHandler = &UnitBehaviorHandler{}

	return &UnitBehavior{
		Component: component,
	}
}
