package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../../scripts/generatelua -type=Stun
type Stun struct {
}

func (d Stun) OpCode() BehaviourAction {
	return BehaviourActionStun
}

func (d Stun) Init(body *byter.Byter) {
	panic("implement me")
}
