package actions

import "RainbowRunner/pkg/byter"

type Remove struct {
}

func (d Remove) OpCode() BehaviourAction {
	return BehaviourActionRemove
}

func (d Remove) Init(body *byter.Byter) {
	panic("implement me")
}
