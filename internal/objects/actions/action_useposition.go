package actions

import "RainbowRunner/pkg/byter"

type UsePosition struct {
}

func (d UsePosition) OpCode() BehaviourAction {
	return BehaviourActionUsePosition
}

func (d UsePosition) Init(body *byter.Byter) {
	panic("implement me")
}
