package actions

import "RainbowRunner/pkg/byter"

type UseItemPosition struct {
}

func (d UseItemPosition) OpCode() BehaviourAction {
	return BehaviourActionUseItemPosition
}

func (d UseItemPosition) Init(body *byter.Byter) {
	panic("implement me")
}
