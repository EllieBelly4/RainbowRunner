package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../scripts/generatelua -type=ActionStun
type ActionStun struct {
}

func (a ActionStun) OpCode() BehaviourAction {
	return BehaviourActionStun
}

func (a ActionStun) Init(body *byter.Byter) {
	panic("implement me")
}

func NewActionStun() *ActionStun {
	return &ActionStun{}
}
