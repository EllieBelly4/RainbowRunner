package behavior

import (
	"RainbowRunner/internal/actions"
	byter "RainbowRunner/pkg/byter"
)

// Deprecated: All components are now in objects/component_*
type Behavior struct {
}

func (b Behavior) Init(body *byter.Byter, action1, action2 actions.Action, sessionID byte) {
	body.WriteByte(0xFF)

	if action1 != nil {
		body.WriteByte(byte(action1.OpCode()))
		// If this doesn't work, it's probably because the init was refactored not taking this into account
		// There may a new method required for this
		action1.Init(body)
	} else {
		body.WriteByte(0x00)
	}

	if action2 != nil {
		body.WriteByte(byte(action2.OpCode()))
		action2.Init(body)
	} else {
		body.WriteByte(0x00)
	}

	body.WriteByte(0x01)
}

func NewBehavior() *Behavior {
	return &Behavior{}
}
