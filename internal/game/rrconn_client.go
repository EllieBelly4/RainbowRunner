package game

import (
	"RainbowRunner/internal/game/messages"
	"RainbowRunner/internal/logging"
	"RainbowRunner/internal/objects"
	"RainbowRunner/pkg"
	"RainbowRunner/pkg/byter"
	"fmt"
	"time"
)

type RRConnClient struct {
	ID               int
	Conn             *RRConn
	Characters       []*objects.Player
	CurrentCharacter *objects.Player
	CurrentHP        uint32
	Zone             string
	IsSpawned        bool

	IsMoving             bool
	ServerUpdateNumber   uint8
	ClientUpdateNumber   byte
	Position             pkg.Vector3
	MoveUpdate           uint8
	Rotation             int32
	TicksSinceLastUpdate int
	MoveQueue            *MovementUpdateQueue
	LastPosition         pkg.Vector3
}

func NewRRConnClient(ID int, conn *RRConn) (p *RRConnClient) {
	p = &RRConnClient{
		ID:                 ID,
		Conn:               conn,
		ServerUpdateNumber: 0xFF,
		MoveQueue:          NewMovementUpdateQueue(1024),
	}

	return
}

func (p *RRConnClient) Warp(x int32, y int32, z int32) {
	p.Position.X = x
	p.Position.Y = x
	p.Position.Z = x

	id := p.getUnitBehaviourID()

	SendWarpTo(p.Conn, id, x, y, z)
	p.updated()
}

func (p *RRConnClient) getUnitBehaviourID() uint16 {
	avatar := p.CurrentCharacter.GetChildByGCNativeType("Avatar")
	unitContainer := avatar.GetChildByGCNativeType("UnitBehavior")
	id := unitContainer.RREntityProperties().ID
	return id
}

var i = byte(0)

func (p *RRConnClient) SendPosition(f byte) {
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

	body := byter.NewLEByter(make([]byte, 0))

	body.WriteByte(byte(messages.ClientEntityChannel))
	body.WriteByte(0x35)                     // ComponentUpdate
	body.WriteUInt16(p.getUnitBehaviourID()) // ComponentID
	body.WriteByte(0x65)                     // UnitMoverUpdate

	updateCount := 10

	// UnitBehavior::processUpdate
	body.WriteByte(0xFF)              // Unk
	body.WriteByte(byte(updateCount)) // Update count

	// UnitMoverUpdate::Read
	//body.WriteByte(0x08) // Not all values work
	//body.WriteByte(0x01) // Not all values work

	for i := 0; i < updateCount; i++ {
		body.WriteByte(0x08) // Not all values work
		body.WriteInt32(p.Rotation)
		body.WriteInt32(p.Position.X)
		body.WriteInt32(p.Position.Y)
	}

	//body.WriteInt32(0)
	//body.WriteInt32(0)
	//body.WriteInt32(0)

	body.WriteByte(0x02)
	body.WriteUInt32(uint32(time.Now().Unix())) // Random unk value

	//AddSynch(p.Conn, body)

	degrees := float32((float64(p.Rotation) / 0x17000) * 360)

	if logging.LoggingOpts.LogMoves {
		fmt.Printf(
			"Sending move 0x%x rotation 0x%x(%.2fdeg) (%d, %d) Hex (%x, %x)\n",
			f, p.Rotation, degrees, p.Position.X, p.Position.Y, p.Position.X, p.Position.Y,
		)
	}

	AddEntityUpdateStreamEnd(body)

	p.send(body)
	i++
}

// Move Move the entity
// I think this is only meant to be used for server controlled entities
func (p *RRConnClient) Move(x int32, y int32) {
	SendMoveTo(p.Conn, 0x2D, p.getUnitBehaviourID(), x, y)
	p.updated()
}

func (p *RRConnClient) SendFollowClient() {
	body := byter.NewLEByter(make([]byte, 0, 128))
	body.WriteByte(byte(messages.ClientEntityChannel))
	body.WriteByte(0x35)

	body.WriteUInt16(p.getUnitBehaviourID())

	body.WriteByte(0x64)
	body.WriteByte(0x01)

	AddSynch(p.Conn, body)

	//body := NewLEByterFromCommandString(`# UnitBehavior - FollowClient
	//07
	//35 # ComponentUpdate
	//05 00 # Component ID
	//# Command
	//64
	//01
	//
	//02 00 00 00 00 # Synch
	//06 # End`)

	AddEntityUpdateStreamEnd(body)
	p.send(body)
}

func (p *RRConnClient) updated() {
	p.TicksSinceLastUpdate = 0
}

func (p *RRConnClient) send(body *byter.Byter) {
	WriteCompressedA(p.Conn, 0x01, 0x0f, body)
	p.updated()
}

//CrashLog: ClientEntityManager::processComponentUpdate ERROR: Invalid ComponentID(5) from server. UpdateID(100)

func (p *RRConnClient) Tick() {
	//for {
	//	move := p.MoveQueue.Peek()
	//
	//	if move == nil {
	//		break
	//	}
	//
	//	p.MoveQueue.Dequeue()
	//
	//	p.Position.X = move.Position.X
	//	p.Position.Y = move.Position.Y
	//	p.Rotation = move.Rotation
	//
	//	p.SendPosition(0x00)
	//}

	if p.IsMoving {
		p.SendPosition(0x00)
	}

	p.TicksSinceLastUpdate++
}
