package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../scripts/generatelua -type=ActionKnockDown
type ActionKnockDown struct {
}

func (a ActionKnockDown) OpCode() BehaviourAction {
	return BehaviourActionKnockDown
}

func (a ActionKnockDown) Init(body *byter.Byter) {
	panic("implement me")
}

func NewActionKnockDown() *ActionKnockDown {
	return &ActionKnockDown{}
}
