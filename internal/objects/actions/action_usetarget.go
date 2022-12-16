package actions

import "RainbowRunner/pkg/byter"

type UseTarget struct {
}

func (d UseTarget) OpCode() BehaviourAction {
	return BehaviourActionUseTarget
}

func (d UseTarget) Init(body *byter.Byter) {
	panic("implement me")
}
