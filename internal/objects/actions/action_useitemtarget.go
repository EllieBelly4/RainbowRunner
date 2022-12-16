package actions

import "RainbowRunner/pkg/byter"

type UseItemTarget struct {
}

func (d UseItemTarget) OpCode() BehaviourAction {
	return BehaviourActionUseItemTarget
}

func (d UseItemTarget) Init(body *byter.Byter) {
	panic("implement me")
}
