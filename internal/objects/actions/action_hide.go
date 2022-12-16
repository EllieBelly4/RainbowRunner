package actions

import "RainbowRunner/pkg/byter"

type Hide struct {
}

func (d Hide) OpCode() BehaviourAction {
	return BehaviourActionHide
}

func (d Hide) Init(body *byter.Byter) {
	panic("implement me")
}
