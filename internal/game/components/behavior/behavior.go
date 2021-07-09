package behavior

import (
	byter "RainbowRunner/pkg/byter"
)

type Behavior struct {
}

func (b Behavior) Init(body *byter.Byter, action1, action2 Action) {
	body.WriteByte(0xFF)

	if action1 != nil {
		action1.Init(body)
	} else {
		body.WriteByte(0x00)
	}

	if action2 != nil {
		action2.Init(body)
	} else {
		body.WriteByte(0x00)
	}

	body.WriteByte(0x01)
}

func NewBehavior() *Behavior {
	return &Behavior{}
}
