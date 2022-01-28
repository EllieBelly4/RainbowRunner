package objects

import (
	"RainbowRunner/internal/config"
	"RainbowRunner/internal/game/messages"
	"RainbowRunner/internal/helpers"
	"RainbowRunner/internal/message"
	"RainbowRunner/pkg/byter"
	"fmt"
)

type Avatar struct {
	*GCObject
	IsMoving           bool
	Rotation           int32
	ClientUpdateNumber byte
	MoveUpdate         int
	IsSpawned          bool
}

func (p *Avatar) Type() DRObjectType {
	return DRObjectOther
}

func NewAvatar(gcType string) *Avatar {
	a := &Avatar{
		GCObject: NewGCObject("Avatar"),
	}

	a.GCType = gcType
	a.GCName = "EllieAvatar"

	return a
}

func (p *Avatar) WriteFullGCObject(byter *byter.Byter) {
	//p.Properties = []GCObjectProperty{
	//	StringProp("Name", p.Name),
	//}

	p.GCObject.WriteFullGCObject(byter)
}

func (p Avatar) WriteInit(b *byter.Byter) {
	panic("implement me")
}

func (p Avatar) WriteUpdate(b *byter.Byter) {
	panic("implement me")
}

func (p *Avatar) Tick() {
	if !p.IsSpawned {
		return
	}

	if config.Config.SendMovementMessages {
		player := Players.GetPlayer(uint16(p.OwnerID()))
		unitBehavior := p.GetChildByGCNativeType("UnitBehavior").(*UnitBehavior)

		CEWriter := NewClientEntityWriterWithByter()

		CEWriter.BeginComponentUpdate(unitBehavior)
		unitBehavior.WriteMoveUpdate(CEWriter.GetBody())
		CEWriter.EndComponentUpdate(unitBehavior)

		player.MessageQueue.Enqueue(message.QueueTypeClientEntity, CEWriter.Body, message.OpTypeAvatarMovement)
	}
}

//func (p *Avatar) SendPosition() {
//	unitBehavior := p.GetChildByGCNativeType("UnitBehavior").(*UnitBehavior)
//	unitBehavior.SendPositions([]UnitPathPosition{
//		{
//			Position: p.Position.ToVector2(),
//			Rotation: p.Rotation,
//		},
//	})
//	p.updated()
//	//p.RREntityProperties().Conn.Send(body)
//}

func (p *Avatar) GetUnitBehaviourID() uint16 {
	unitContainer := p.GetChildByGCNativeType("UnitBehavior")
	id := unitContainer.RREntityProperties().ID
	return id
}

func (p *Avatar) Send(body *byter.Byter) {
	helpers.WriteCompressedA(p.RREntityProperties().Conn, 0x01, 0x0f, body)
}

func (p *Avatar) SendFollowClient() {
	writer := NewClientEntityWriterWithByter()
	writer.BeginStream()
	writer.BeginComponentUpdate(p.GetChildByGCNativeType("UnitBehavior"))

	writer.Body.WriteByte(0x64)
	writer.Body.WriteByte(0x01)

	writer.WriteSynch(p)

	writer.EndStream()
	p.Send(writer.Body)
}

func (p *Avatar) Warp(x int32, y int32, z int32) {
	unitBehavior := p.GetChildByGCNativeType("UnitBehavior").(*UnitBehavior)
	unitBehavior.Warp(x, y, z)
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

	helpers.WriteCompressedA(p.RREntityProperties().Conn, 0x01, 0x0f, body)

	if config.Config.Logging.LogMoves {
		fmt.Printf("Send MoveTo %x (%d, %d) (%x, %x)\n", unk, posX, posY, posX, posY)
	}
}

func (p *Avatar) GetUnitContainer() *UnitContainer {
	return p.GetChildByGCNativeType("UnitContainer").(*UnitContainer)
}

func (p *Avatar) GetManipulators() *Manipulators {
	return p.GetChildByGCNativeType("Manipulators").(*Manipulators)
}

func (p *Avatar) GetUnitBehaviour() *UnitBehavior {
	unitBehaviour := p.GetChildByGCNativeType("UnitBehavior")
	return unitBehaviour.(*UnitBehavior)
}

//func (p *Avatar) SendPositions(positions []UnitPathPosition) {
//	unitBehavior := p.GetChildByGCNativeType("UnitBehavior").(*UnitBehavior)
//	unitBehavior.SendPositions(positions)
//	p.updated()
//	//p.RREntityProperties().Conn.Send(body)
//}
