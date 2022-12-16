package actions

import "RainbowRunner/pkg/byter"

type Flee struct {
}

func (d Flee) OpCode() BehaviourAction {
	return BehaviourActionFlee
}

func (d Flee) Init(body *byter.Byter) {
	panic("implement me")
}
