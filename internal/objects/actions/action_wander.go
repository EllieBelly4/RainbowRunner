package actions

import "RainbowRunner/pkg/byter"

type Wander struct {
}

func (d Wander) OpCode() BehaviourAction {
	return BehaviourActionWander
}

func (d Wander) Init(body *byter.Byter) {
	panic("implement me")
}
