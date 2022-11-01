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
			if _, ok := object.(*MonsterBehavior2); ok {
				return
			}

			switch object.Type() {
			case DRObjectComponent:
				CEWriter.CreateComponentAndInit(object, entity)
			}
		})

		CEWriter.Init(entity)
		p.MessageQueue.Enqueue(message.QueueTypeClientEntity, CEWriter.Body, message.OpTypeCreateNPC)
	}
}
