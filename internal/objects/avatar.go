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

type Avatar struct {
	*GCObject
	IsMoving             bool
	Rotation             int32
	Position             pkg.Vector3
	ClientUpdateNumber   byte
	MoveUpdate           int
	TicksSinceLastUpdate int
	IsSpawned            bool
	LastPosition         pkg.Vector3
}

func NewAvatar(gcType string) *Avatar {
	a := &Avatar{
		GCObject: NewGCObject("Avatar"),
	}

	a.GCType = gcType
	a.GCName = "EllieAvatar"

	return a
}

func (a *Avatar) WriteFullGCObject(byter *byter.Byter) {
	//p.Properties = []GCObjectProperty{
	//	StringProp("Name", p.Name),
	//}

	a.GCObject.WriteFullGCObject(byter)
}

func (a Avatar) WriteInit(b *byter.Byter) {
	panic("implement me")
}

func (a Avatar) WriteUpdate(b *byter.Byter) {
	panic("implement me")
}

func (p *Avatar) Tick() {
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

func (p *Avatar) updated() {
	p.TicksSinceLastUpdate = 0
}

func (p *Avatar) SendPosition() {
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
	writer.BeginComponentUpdate(p.GetUnitBehaviourID())

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
		writer.Body.WriteInt32(p.Rotation)
		writer.Body.WriteInt32(p.Position.X)
		writer.Body.WriteInt32(p.Position.Y)
	}

	//writer.Body.WriteInt32(0)
	//writer.Body.WriteInt32(0)
	//writer.Body.WriteInt32(0)

	writer.Body.WriteByte(0x02)
	writer.Body.WriteUInt32(uint32(time.Now().Unix())) // Random unk value

	//AddSynch(p.Conn, writer.Body)

	degrees := float32((float64(p.Rotation) / 0x17000) * 360)

	if logging.LoggingOpts.LogMoves {
		fmt.Printf(
			"Sending move rotation 0x%x(%.2fdeg) (%d, %d) Hex (%x, %x)\n",
			p.Rotation, degrees, p.Position.X, p.Position.Y, p.Position.X, p.Position.Y,
		)
	}

	writer.EndStream()

	p.Send(writer.Body)
	p.updated()
	//p.RREntityProperties().Conn.Send(body)
}

func (p *Avatar) GetUnitBehaviourID() uint16 {
	unitContainer := p.GetChildByGCNativeType("UnitBehavior")
	id := unitContainer.RREntityProperties().ID
	return id
}

func (p *Avatar) Send(body *byter.Byter) {
	connections.WriteCompressedA(p.RREntityProperties().Conn, 0x01, 0x0f, body)
}

func (p *Avatar) SendFollowClient() {
	writer := NewClientEntityWriterWithByter()
	writer.BeginStream()
	writer.BeginComponentUpdate(p.GetUnitBehaviourID())

	writer.Body.WriteByte(0x64)
	writer.Body.WriteByte(0x01)

	writer.WriteSynch(p)

	writer.EndStream()
	p.Send(writer.Body)
}

func (p *Avatar) Warp(x int32, y int32, z int32) {
	p.Position.X = x
	p.Position.Y = x
	p.Position.Z = x

	p.SendWarpTo(x, y, z)
	p.updated()
}

func (p *Avatar) SendWarpTo(posX, posY, posZ int32) {
	writer := NewClientEntityWriterWithByter()
	writer.BeginStream()
	writer.BeginComponentUpdate(p.GetUnitBehaviourID())

	writer.Body.WriteByte(0x04) // CreateAction1
	writer.Body.WriteByte(17)
	writer.Body.WriteByte(0x00)
	writer.Body.WriteInt32(posX)
	writer.Body.WriteInt32(posY)
	writer.Body.WriteInt32(posZ)

	writer.WriteSynch(p)
	writer.EndStream()

	connections.WriteCompressedA(p.RREntityProperties().Conn, 0x01, 0x0f, writer.Body)

	p.updated()
}

func (p *Avatar) SendMoveTo(unk uint8, compID uint16, posX, posY int32) {
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
