package objects

import (
	"RainbowRunner/internal/game/components/behavior"
	"RainbowRunner/internal/message"
	"strings"
)

type Component struct {
	*GCObject
}

func (c Component) Activate(player *RRPlayer, responseID byte) {
	CEWriter := NewClientEntityWriterWithByter()

	CEWriter.BeginComponentUpdate(c)
	CEWriter.CreateActionResponse(behavior.BehaviourActionActivate, responseID)

	activateAction := behavior.Activate{
		TargetEntityID: uint16(c.EntityProperties.ID),
	}

	activateAction.InitWithoutOpCode(CEWriter.Body)
	CEWriter.WriteSynch(c)

	player.MessageQueue.Enqueue(
		message.QueueTypeClientEntity, CEWriter.Body, message.OpTypeBehaviourAction,
	)
}

func (Component) Type() DRObjectType {
	return DRObjectComponent
}

func NewComponent(gcType string, nativeType string) *Component {
	gcObject := NewGCObject(nativeType)
	gcObject.GCType = strings.ToLower(gcType)

	return &Component{
		GCObject: gcObject,
	}
}
