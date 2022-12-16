package actions

import "RainbowRunner/pkg/byter"

type KnockDown struct {
}

func (d KnockDown) OpCode() BehaviourAction {
	return BehaviourActionKnockDown
}

func (d KnockDown) Init(body *byter.Byter) {
	panic("implement me")
}
