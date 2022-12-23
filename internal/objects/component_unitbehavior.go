package objects

import (
	actions2 "RainbowRunner/internal/actions"
	"RainbowRunner/internal/config"
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/game/components/behavior"
	"RainbowRunner/internal/global"
	"RainbowRunner/internal/gosucks"
	"RainbowRunner/internal/message"
	"RainbowRunner/pkg/byter"
	"RainbowRunner/pkg/datatypes"
	"RainbowRunner/pkg/events"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"math"
)

//go:generate go run ../../scripts/generateLua/ -type=UnitBehavior -extends=Component
type UnitBehavior struct {
	*Component

	// Unknown if/where this is sent to the client
	Speed    int
	TurnRate int

	LastPosition   datatypes.Vector3Float32
	Position       datatypes.Vector3Float32
	Rotation       float32
	UnitMoverFlags byte
	Action1        actions2.Action
	Action2        actions2.Action

	// I don't know for certain if this is really called "SessionID" but I think it is
	// This increments every time a unit behavior action is executed and must be used for subsequent movement
	// updates and actions
	SessionID byte

	// unitMoverFlags & 0x04
	UnitMoverUnk0 byte

	// unitMoverFlags & 0x01
	UnitMoverUnk1 uint32
	UnitMoverUnk2 uint32

	UnitMoverUnk3 uint32
	UnitMoverUnk4 uint32

	// unitMoverFlags & 0x80
	UnitMoverUnk7 uint32

	UnitBehaviorUnk0 byte
	UnitBehaviorUnk1 byte
	UnitBehaviorUnk2 byte

	//Movement
	targetPosition datatypes.Vector2Float32
	IsMoving       bool
	//TODO add movement path
}

func (u *UnitBehavior) Tick() {
	if u.IsMoving {
		//TODO handle turning
		//turning is currently not 100% possible as the client is not recognising the unit behavior rotation
		//and is instead always defaulting to 0

		distanceToTarget := u.Position.ToVector2Float32().Distance(u.targetPosition)
		if distanceToTarget < 0.1 {
			//logrus.Info("reached target position")
			u.IsMoving = false
			return
		}

		dirToTarget := u.targetPosition.Sub(u.Position.ToVector2Float32()).Normalize()
		moveDistance := math.Min(distanceToTarget, float64(u.Speed)*global.GetDeltaTime())

		toMove := dirToTarget.Mul(float32(moveDistance)).ToVector3Float32()

		newPos := u.Position.Add(toMove)

		if u.RREntityProperties().Zone.PathMap != nil {
			newPos.Z = u.RREntityProperties().Zone.PathMap.HeightAt(newPos.ToVector2Float32())
		}

		u.Position = newPos
	}
}

func (u *UnitBehavior) WriteInit(b *byter.Byter) {
	behav := behavior.NewBehavior()
	behav.Init(b, u.Action1, u.Action2, u.SessionID)

	// UnitMover::readInit()
	// Flags
	// 0x04
	// 0x01
	unitMover := u.UnitMoverFlags
	b.WriteByte(unitMover)

	if unitMover&0x04 > 0 {
		b.WriteByte(u.UnitMoverUnk0)
	}

	if unitMover&0x01 > 0 {
		// 0x01 case
		b.WriteUInt32(u.UnitMoverUnk1)
		b.WriteUInt32(u.UnitMoverUnk2)
	}

	b.WriteUInt32(u.UnitMoverUnk3)
	b.WriteUInt32(u.UnitMoverUnk4)

	if unitMover&0x80 > 0 {
		b.WriteUInt32(u.UnitMoverUnk7)
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
	b.WriteByte(u.UnitBehaviorUnk0)
	b.WriteByte(u.UnitBehaviorUnk1)
	b.WriteByte(u.UnitBehaviorUnk2)
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
	b.WriteByte(u.SessionID) // Unk

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

type UnitPathPosition struct {
	Position   datatypes.Vector2Float32
	Rotation   float32
	ResponseID byte
}

func (u *UnitBehavior) handleClientMove(conn connections.Connection, reader *byter.Byter) {
	// This is the session ID that increments every time the client attempts to sync
	sessionID := reader.Byte()

	// TODO remove this, we should probably be correctly calculating the session ID
	u.SessionID = sessionID

	count := int(reader.Byte())
	pos := datatypes.Vector2Float32{}

	if config.Config.Logging.LogMoves {
		fmt.Printf("Received %d player moves unk val: %x\n", count, sessionID)
	}

	responseMoves := make([]UnitPathPosition, 0)

	currentZone := Players.Players[conn.GetID()].Zone()

	//TODO remove avatar object usage from here
	avatar := Players.Players[conn.GetID()].CurrentCharacter.GetChildByGCNativeType("Avatar").(*Avatar)

	for i := 0; i < count; i++ {
		moveUpdateType := reader.Byte()     // Unk
		rotation := float32(reader.Int32()) // Seems to be rotation

		//degrees := float32((float64(rotation) / 0x17000) * 360)
		degrees := float32(rotation / 256)

		pos.X = float32(reader.Int32()) / 256
		pos.Y = float32(reader.Int32()) / 256

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

		if currentZone.PathMap != nil {
			u.Position.Z = currentZone.PathMap.HeightAt(u.Position.ToVector2Float32())
		}

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
			ResponseID: sessionID,
			Rotation:   rotation,
			Position: datatypes.Vector2Float32{
				X: pos.X,
				Y: pos.Y,
			},
		})
	}

	if currentZone.PathMap != nil {
		//gridPos := currentZone.PathMap.WorldPosToGridCoords(u.Position)
		//logrus.Infof("Current coords: %d, %d height: %f", gridPos.X, gridPos.Y, u.Position.Z)
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
		// This seems to also be used when the player tries to interact with an object when it is too far away
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

func (u *UnitBehavior) WarpTo(pos datatypes.Vector3Float32) {
	u.Position = pos

	warpTo := &actions2.ActionWarpTo{
		Position: pos,
	}

	events.Emit(ExecuteActionEvent{Action: warpTo, UnitBehavior: u})
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
	responseId := reader.Byte()
	action := actions2.BehaviourAction(reader.Byte())
	someID := reader.Byte()

	logrus.Infof("execute action %s, unk0 %d sessionID %d\n", action.String(), responseId, someID)
	reader.DumpRemaining()

	var err error

	switch action {
	case actions2.BehaviourActionUsePosition:
		err = u.handleActionUsePosition(reader, responseId, someID)
	case actions2.BehaviourActionActivate:
		err = u.handleExecuteActivate(reader, responseId, someID)
	}

	u.SessionID++
	return err
}

func (u *UnitBehavior) handleExecuteActivate(reader *byter.Byter, responseID byte, sessionID byte) error {
	logrus.Infof("execute Activate responseID %x", sessionID)

	targetID := reader.UInt16()

	targetEntity := u.EntityProperties.Zone.FindEntityByID(targetID)
	//targetEntity := Entities.FindByID(targetID)

	if targetEntity == nil {
		return errors.New(fmt.Sprintf("could not find target entity with ID %d", targetID))
	}
	activateable, ok := targetEntity.(IActivatable)

	if !ok {
		logrus.Errorf("tried to activate non-activatable: %s", targetEntity.String())
		return nil
	}

	activateable.Activate(Players.GetPlayer(u.OwnerID()), u, responseID, sessionID)
	return nil
}

func (u *UnitBehavior) handleActionUsePosition(reader *byter.Byter, id byte, sessionID byte) error {
	//reader.Byte() // Some incrementing index
	actionID := reader.Byte()

	posX := float64(reader.Int32()) / 256
	posY := float64(reader.Int32()) / 256
	posZ := float64(reader.Int32()) / 256

	logrus.Infof("use position actionID %d\n%f,%f,%f", actionID, posX, posY, posZ)

	gosucks.VAR(posX, posY, posZ)

	CEWriter := NewClientEntityWriterWithByter()

	CEWriter.BeginComponentUpdate(u)
	CEWriter.CreateActionResponse(actions2.BehaviourActionUsePosition, id, sessionID)

	usePositionAction := actions2.ActionUsePosition{
		Position: datatypes.Vector3Float32{
			X: float32(posX),
			Y: float32(posY),
			Z: float32(posZ),
		},
		ActionID: actionID,
	}

	usePositionAction.Init(CEWriter.Body)

	CEWriter.WriteSynch(u)

	player := Players.GetPlayer(u.OwnerID())

	player.MessageQueue.Enqueue(
		message.QueueTypeClientEntity, CEWriter.Body, message.OpTypeBehaviourAction,
	)

	return nil
}

func (u *UnitBehavior) MoveTo(pos datatypes.Vector2Float32) {
	action := &actions2.ActionMoveTo{
		PosX: pos.X,
		PosY: pos.Y,
	}

	u.targetPosition = datatypes.Vector2Float32{X: pos.X, Y: pos.Y}
	u.IsMoving = true
	//u.Position = datatypes.Vector3Float32{X: pos.X, Y: pos.Y, Z: u.Position.Z}

	u.ExecuteAction(action)
}

func (u *UnitBehavior) MoveToEntity(g IWorldEntity) {
	targetPosition := datatypes.Vector2Float32{}
	nativeType := g.GetWorldEntity().GetChildByGCNativeType("UnitBehavior")

	set := false

	if nativeType != nil {
		if targetUnitBehav, ok := nativeType.(*UnitBehavior); ok {
			targetPosition.X = targetUnitBehav.Position.X
			targetPosition.Y = targetUnitBehav.Position.Y
			set = true
		}
	}

	if !set {
		we := g.GetWorldEntity()
		targetPosition.X = we.WorldPosition.X
		targetPosition.Y = we.WorldPosition.Y
	}

	if !set {
		logrus.Warningf("cannot move to target as it does not have a position")
		return
	}

	u.MoveTo(targetPosition)
}

// ExecuteAction should never be called directly as it is only used to emit actions
// you should call the direct action methods on this behaviour instead
// e.g. MoveTo, Attack, etc
func (u *UnitBehavior) ExecuteAction(action actions2.Action) {
	events.Emit(ExecuteActionEvent{
		Action:       action,
		UnitBehavior: u,
	})

	u.SessionID++
}

func (u *UnitBehavior) StopFollowClient() {
	CEWriter := NewClientEntityWriterWithByter()
	//writer.BeginStream()
	CEWriter.BeginComponentUpdate(u)

	CEWriter.Body.WriteByte(0x64) // Update type - something to do with client control
	CEWriter.Body.WriteByte(0x00) // Client control on or off

	CEWriter.WriteSynch(u)

	player := Players.GetPlayer(u.OwnerID())

	player.MessageQueue.Enqueue(message.QueueTypeClientEntity, CEWriter.Body, message.OpTypeOther)

}

func NewUnitBehavior(gcType string) *UnitBehavior {
	component := NewComponent(gcType, "UnitBehavior")
	component.EntityHandler = &UnitBehaviorHandler{}

	return &UnitBehavior{
		Component: component,
		SessionID: 0xFF,
	}
}
