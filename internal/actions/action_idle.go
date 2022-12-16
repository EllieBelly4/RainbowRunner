package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../../scripts/generatelua -type=ActionIdle
type ActionIdle struct {
}

func (a ActionIdle) OpCode() BehaviourAction {
	return BehaviourActionIdle
}

func (a ActionIdle) Init(body *byter.Byter) {
	panic("implement me")
}

func NewActionIdle() *ActionIdle {
	return &ActionIdle{}
}
