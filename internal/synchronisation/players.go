package synchronisation

import (
	"RainbowRunner/internal/message"
	"RainbowRunner/internal/objects"
)

func onPlayerMove(event objects.PlayerMoveEvent) {
	newPos := event.NewPosition.ToVector3DRFloat()

	CEWriter := objects.NewClientEntityWriterWithByter()

	// TODO remove duplication between here and UnitBehavior::WriteMoveUpdate
	// I am not doing this right now as things keep changing with movement.
	CEWriter.BeginComponentUpdate(event.UnitBehavior)

	CEWriter.Body.WriteByte(0x65) // UnitMoverUpdate

	CEWriter.Body.WriteByte(0x00) // Session ID

	CEWriter.Body.WriteByte(0x01)             // Update count
	CEWriter.Body.WriteByte(event.UpdateType) // Update type
	CEWriter.Body.WriteUInt32(event.NewHeading.ToWire())
	CEWriter.Body.WriteUInt32(newPos.X.ToWire())
	CEWriter.Body.WriteUInt32(newPos.Y.ToWire())

	CEWriter.EndComponentUpdate(event.UnitBehavior)

	//if math.Abs(float64(headingDiff)) > 0 {
	//	turnAction := actions.NewActionTurnAction()
	//
	//	// Avatar base turn rate, TODO read from config
	//	turnAction.Speed = drfloat.FromInt32(720)
	//
	//	CEWriter.BeginComponentUpdate(event.UnitBehavior)
	//	CEWriter.CreateActionComplete(turnAction)
	//	CEWriter.EndComponentUpdate(event.UnitBehavior)
	//}

	for _, player := range event.UnitBehavior.RREntityProperties().Zone.Players() {
		if uint16(player.Conn.GetID()) == event.UnitBehavior.OwnerID() {
			continue
		}

		player.MessageQueue.Enqueue(message.QueueTypeClientEntity, CEWriter.Body, message.OpTypeAvatarMovementOthers)
	}
}
