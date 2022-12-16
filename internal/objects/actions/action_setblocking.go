package actions

import "RainbowRunner/pkg/byter"

type SetBlocking struct {
}

func (d SetBlocking) OpCode() BehaviourAction {
	return BehaviourActionSetBlocking
}

func (d SetBlocking) Init(body *byter.Byter) {
	panic("implement me")
}
