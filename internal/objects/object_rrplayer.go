package objects

import (
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/message"
)

type RRPlayer struct {
	Conn               *connections.RRConn
	CurrentCharacter   *Player
	Characters         []*Player
	Zone               *Zone
	ClientEntityWriter *ClientEntityWriter
	MessageQueue       *message.Queue
}

func (p *RRPlayer) OnZoneJoin() {
	entities := p.Zone.Entities()

	for _, entity := range entities {
		if entity == p.CurrentCharacter.GetChildByGCNativeType("Avatar") {
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
	p.Zone.RemovePlayer(int(p.CurrentCharacter.ID()))
	p.MessageQueue.Clear(message.QueueTypeClientEntity)
}
