package actions

import "RainbowRunner/pkg/byter"

type Idle struct {
}

func (d Idle) OpCode() BehaviourAction {
	return BehaviourActionIdle
}

func (d Idle) Init(body *byter.Byter) {
	panic("implement me")
}
