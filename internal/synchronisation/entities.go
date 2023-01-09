package synchronisation

import (
	"RainbowRunner/internal/message"
	"RainbowRunner/internal/objects"
	"RainbowRunner/pkg/events"
)

func Init() {
	events.RegisterHandler[objects.ExecuteActionEvent](func(event objects.ExecuteActionEvent) {
		behavior := event.UnitBehavior.GetUnitBehavior()
		zone := behavior.EntityProperties.Zone

		if zone == nil {
			return
		}

		players := zone.Players()

		CEWriter := objects.NewClientEntityWriterWithByter()
		CEWriter.BeginComponentUpdate(behavior)
		CEWriter.CreateActionComplete(event.Action)
		CEWriter.WriteSynch(behavior)

		for _, player := range players {
			//if player.Conn.GetID() == int(behavior.OwnerID()) {
			//	continue
			//}

			player.MessageQueue.Enqueue(message.QueueTypeClientEntity, CEWriter.Body, message.OpTypeBehaviourAction)
		}
	})

	events.RegisterHandler[objects.PlayerMoveEvent](onPlayerMove)
}
