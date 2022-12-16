package actions

import "RainbowRunner/pkg/byter"

type DoEffect struct {
}

func (d DoEffect) OpCode() BehaviourAction {
	return BehaviourActionDoEffect
}

func (d DoEffect) Init(body *byter.Byter) {
	panic("implement me")
}
