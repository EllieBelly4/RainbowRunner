package objects

import (
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/game/messages"
	"RainbowRunner/internal/logging"
	"RainbowRunner/pkg"
	"RainbowRunner/pkg/byter"
	"fmt"
	"time"
)

type Player struct {
	*GCObject
	Name                 string
	IsMoving             bool
	Rotation             int32
	Position             pkg.Vector3
	ClientUpdateNumber   byte
	MoveUpdate           int
	TicksSinceLastUpdate int
	CurrentHP            uint32
	IsSpawned            bool
}

func (p *Player) WriteInit(b *byter.Byter) {
	// Init PLAYER /////////////////////////////////////////
	b.WriteCString("Ellie")
	b.WriteUInt32(0x01)
	b.WriteUInt32(0x01)
	b.WriteByte(0x01)

	b.WriteUInt32(0xFEEDBABA) // World ID
	b.WriteUInt32(1001)       // PvP wins
	b.WriteUInt32(1000)       // PvP rating?, 0 = ???

	// Here goes PvP Team
	// Null string
	b.WriteByte(0x00)

	// If player is in a PvP team then Avatar respawn will look for the team waypoints
	//b.WriteByte(0xFF)
	//b.WriteCString("pvp.DefaultTeamList.BlueTeam")

	b.WriteCString("Hello")
	b.WriteUInt32(0x01)

}

func (p *Player) WriteUpdate(b *byter.Byter) {
	// This maps to a specific event type for Player::processUpdate()
	// 0x01 - do nothing
	// 0x03 - Unk
	b.WriteByte(0x03)

	// 0x03 case
	b.WriteUInt16(0x02)
}

func (p *Player) WriteFullGCObject(byter *byter.Byter) {
	p.Properties = []GCObjectProperty{
		StringProp("Name", p.Name),
	}

	p.GCObject.WriteFullGCObject(byter)

	byter.WriteCString("Unk")  // Specific to player::readObject
	byter.WriteCString("Unk2") // Specific to player::readObject
	byter.WriteUInt32(0x01)    // Specific to player::readObject
	byter.WriteUInt32(0x01)    // Specific to player::readObject
}

func (p *Player) Tick() {
	if !p.IsSpawned {
		return
	}

	if p.TicksSinceLastUpdate >= 0x2D {
		p.SendPosition()
	}

	if p.IsMoving {
		p.SendPosition()
	}

	p.TicksSinceLastUpdate++
}

func (p *Player) updated() {
	p.TicksSinceLastUpdate = 0
}

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

	body.WriteByte(byte(messages.ClientEntityChannel))
	body.WriteByte(0x35)                     // ComponentUpdate
	body.WriteUInt16(p.GetUnitBehaviourID()) // ComponentID
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
			"Sending move rotation 0x%x(%.2fdeg) (%d, %d) Hex (%x, %x)\n",
			p.Rotation, degrees, p.Position.X, p.Position.Y, p.Position.X, p.Position.Y,
		)
	}

	// Stream end
	body.WriteByte(0x06)

	p.Send(body)
	p.updated()
	//p.RREntityProperties().Conn.Send(body)
}

func (p *Player) GetUnitBehaviourID() uint16 {
	avatar := p.GetChildByGCNativeType("Avatar")
	unitContainer := avatar.GetChildByGCNativeType("UnitBehavior")
	id := unitContainer.RREntityProperties().ID
	return id
}

func (p *Player) Send(body *byter.Byter) {
	connections.WriteCompressedA(p.RREntityProperties().Conn, 0x01, 0x0f, body)
}

func (p *Player) SendFollowClient() {
	body := byter.NewLEByter(make([]byte, 0, 128))
	body.WriteByte(byte(messages.ClientEntityChannel))
	body.WriteByte(0x35)

	body.WriteUInt16(p.GetUnitBehaviourID())

	body.WriteByte(0x64)
	body.WriteByte(0x01)

	// EntitySynchInfo::readFromStream
	body.WriteByte(0x02)
	body.WriteUInt32(p.CurrentHP)

	//AddSynch(p.Conn, body)

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

	// End Stream
	body.WriteByte(0x06)
	p.Send(body)
}

func (p *Player) Warp(x int32, y int32, z int32) {
	p.Position.X = x
	p.Position.Y = x
	p.Position.Z = x

	id := p.GetUnitBehaviourID()

	p.SendWarpTo(id, x, y, z)
	p.updated()
}
func (p *Player) SendWarpTo(compID uint16, posX, posY, posZ int32) {
	body := byter.NewLEByter(make([]byte, 0))

	body.WriteByte(byte(messages.ClientEntityChannel))
	body.WriteByte(0x35)
	body.WriteUInt16(compID) // UnitBehavior
	body.WriteByte(0x04)     // CreateAction1
	body.WriteByte(17)
	body.WriteByte(0x00)
	body.WriteInt32(posX)
	body.WriteInt32(posY)
	body.WriteInt32(posZ)

	// EntitySynchInfo::readFromStream
	body.WriteByte(0x02)
	body.WriteUInt32(p.CurrentHP)

	// EndStream
	body.WriteByte(0x06)

	connections.WriteCompressedA(p.RREntityProperties().Conn, 0x01, 0x0f, body)

	p.updated()
}

func (p *Player) SendMoveTo(unk uint8, compID uint16, posX, posY int32) {
	body := byter.NewLEByter(make([]byte, 0))

	body.WriteByte(byte(messages.ClientEntityChannel))
	body.WriteByte(0x35)
	body.WriteUInt16(compID) // UnitBehavior
	body.WriteByte(0x04)     // CreateAction1
	body.WriteByte(0x01)     // MoveTo
	body.WriteByte(unk)
	body.WriteInt32(posX)
	body.WriteInt32(posY)

	body.WriteByte(0x02)
	body.WriteUInt32(0x00)

	//AddSynch(conn, body)

	// EndStream
	body.WriteByte(0x06)

	connections.WriteCompressedA(p.RREntityProperties().Conn, 0x01, 0x0f, body)

	if logging.LoggingOpts.LogMoves {
		fmt.Printf("Send MoveTo %x (%d, %d) (%x, %x)\n", unk, posX, posY, posX, posY)
	}

	p.updated()
}

func NewPlayer(name string) (p *Player) {
	p = &Player{
		Name: name,
	}

	p.GCObject = NewGCObject("Player")
	p.GCName = "ElliePlayer"
	p.GCType = "player"

	return
}
