package actions

import "RainbowRunner/pkg/byter"

type Wait struct {
}

func (d Wait) OpCode() BehaviourAction {
	return BehaviourActionWait
}

func (d Wait) Init(body *byter.Byter) {
	panic("implement me")
}
