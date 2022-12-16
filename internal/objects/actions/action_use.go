package actions

import "RainbowRunner/pkg/byter"

type Use struct {
}

func (d Use) OpCode() BehaviourAction {
	return BehaviourActionUse
}

func (d Use) Init(body *byter.Byter) {
	panic("implement me")
}
