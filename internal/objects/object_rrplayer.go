package objects

import (
	"RainbowRunner/internal/config"
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/game/messages"
	"RainbowRunner/internal/message"
	"RainbowRunner/pkg/byter"
	"encoding/hex"
	"fmt"
	"github.com/sirupsen/logrus"
)

type RRPlayer struct {
	Conn               *connections.RRConn
	CurrentCharacter   *Player
	Characters         []*Player
	Zone               *Zone
	ClientEntityWriter *ClientEntityWriter
	MessageQueue       *message.Queue
	Spawned            bool
}

func (p *RRPlayer) OnZoneJoin() {
	p.Spawned = true
	entities := p.Zone.Entities()

	for _, entity := range entities {
		if int(entity.OwnerID()) == p.Conn.GetID() {
			continue
		}

		CEWriter := NewClientEntityWriterWithByter()
		CEWriter.Create(entity)

		entity.WalkChildren(func(object DRObject) {
			switch object.Type() {
			case DRObjectComponent:
				//if mb2, ok := object.(*MonsterBehavior2); ok {
				//	CEWriter.CreateComponentAndInit(object, entity)
				//}
				CEWriter.CreateComponentAndInit(object, entity)
			}
		})

		CEWriter.Init(entity)

		if unitBehavior, ok := entity.GetChildByGCNativeType("UnitBehavior").(IUnitBehavior); unitBehavior != nil && ok {
			unitBehavior.GetUnitBehavior().WriteWarp(CEWriter)
		}

		p.MessageQueue.Enqueue(message.QueueTypeClientEntity, CEWriter.Body, message.OpTypeCreateNPC)
	}
}

func (p *RRPlayer) LeaveCurrentZone() {
	p.Spawned = false
	p.Zone.RemovePlayer(int(p.CurrentCharacter.ID()))
	p.MessageQueue.Clear(message.QueueTypeClientEntity)

	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(messages.ZoneChannel))
	body.WriteByte(byte(messages.ZoneMessageDisconnected))
	body.WriteCString("zoneleaveuhh")
	connections.WriteCompressedA(p.Conn, 0x01, 0x0f, body)

	if config.Config.Logging.LogChangeZone {
		logrus.Info(fmt.Sprintf("Sent Leave Zone: \n%s", hex.Dump(body.Data())))
	}
}

func (p *RRPlayer) JoinZone(name string) {
	Zones.PlayerJoin(name, p)

	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(messages.ZoneChannel))
	body.WriteByte(byte(messages.ZoneMessageConnected))
	//body.WriteCString("TheHub")
	//body.WriteCString("Tutorial")
	body.WriteCString(name)
	body.WriteUInt32(0xBEEFBEEF)
	body.WriteByte(0x01)
	body.WriteByte(0xFF)
	body.WriteCString("world.town.quest.Q01_a1")
	body.WriteUInt32(0x01)
	connections.WriteCompressedA(p.Conn, 0x01, 0x0f, body)

	if config.Config.Logging.LogChangeZone {
		logrus.Info(fmt.Sprintf("Sent Change Zone: \n%s", hex.Dump(body.Data())))
	}
}
