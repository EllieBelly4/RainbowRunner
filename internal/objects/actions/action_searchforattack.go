package actions

import "RainbowRunner/pkg/byter"

type SearchForAttack struct {
}

func (d SearchForAttack) OpCode() BehaviourAction {
	return BehaviourActionSearchForAttack
}

func (d SearchForAttack) Init(body *byter.Byter) {
	panic("implement me")
}
