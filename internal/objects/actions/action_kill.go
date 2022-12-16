package actions

import "RainbowRunner/pkg/byter"

type Kill struct {
}

func (d Kill) OpCode() BehaviourAction {
	return BehaviourActionKill
}

func (d Kill) Init(body *byter.Byter) {
	panic("implement me")
}
