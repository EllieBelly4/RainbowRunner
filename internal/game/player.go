package game

import (
	"RainbowRunner/internal/byter"
)

type Vector3 struct {
	X, Y, Z int32
}

type Vector3Short struct {
	X, Y, Z int16
}

type Vector3Float32 struct {
	X, Y, Z float32
}

type Player struct {
	Conn      *RRConn
	CurrentHP uint32
	Zone      string
	IsSpawned bool

	IsMoving             bool
	LastMovementRequest  Vector3
	ServerUpdateNumber   uint8
	ClientUpdateNumber   byte
	Position             Vector3
	MoveUpdate           uint8
	Rotation             int32
	TicksSinceLastUpdate int
}

func NewPlayer(conn *RRConn) *Player {
	return &Player{
		Conn:               conn,
		ServerUpdateNumber: 0xFF,
	}
}

func (p *Player) Warp(x int32, y int32, z int32) {
	p.Position.X = x
	p.Position.Y = x
	p.Position.Z = x
	SendWarpTo(p.Conn, 0x05, x, y, z)
	p.updated()
}

var i = byte(0)

func (p *Player) SendPosition() {
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

	body.WriteByte(byte(ClientEntityChannel))
	body.WriteByte(0x35)   // ComponentUpdate
	body.WriteUInt16(0x05) // ComponentID
	body.WriteByte(0x65)   // UnitMoverUpdate

	// UnitBehavior::processUpdate
	body.WriteByte(0xFF) // Unk
	body.WriteByte(0x01) // Update count

	// UnitMoverUpdate::Read
	body.WriteByte(0x08) // Unk

	body.WriteInt32(p.Rotation)
	body.WriteInt32(p.Position.X)
	body.WriteInt32(p.Position.Y)

	body.WriteByte(0x02)
	body.WriteUInt32(uint32(p.ClientUpdateNumber))

	//AddSynch(p.Conn, body)

	AddEntityUpdateStreamEnd(body)

	p.send(body)
	i++
}

// Move Move the entity
// I think this is only meant to be used for server controlled entities
func (p *Player) Move(x int32, y int32) {
	SendMoveTo(p.Conn, 0x2D, 0x05, x, y)
	p.updated()
}

func (p *Player) SendFollowClient() {
	body := NewLEByterFromCommandString(`# UnitBehavior - FollowClient
07
35 # ComponentUpdate
05 00 # Component ID
# Command
64 
01

02 00 00 00 00 # Synch
06 # End`)

	p.send(body)
}

func (p *Player) updated() {
	p.TicksSinceLastUpdate = 0
}

func (p *Player) send(body *byter.Byter) {
	WriteCompressedA(p.Conn, 0x01, 0x0f, body)
	p.updated()
}

func (p *Player) Tick() {
	p.TicksSinceLastUpdate++
}
