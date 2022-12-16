package actions

import "RainbowRunner/pkg/byter"

type Stun struct {
}

func (d Stun) OpCode() BehaviourAction {
	return BehaviourActionStun
}

func (d Stun) Init(body *byter.Byter) {
	panic("implement me")
}
