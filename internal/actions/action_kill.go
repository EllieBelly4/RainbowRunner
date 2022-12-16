package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../scripts/generatelua -type=ActionKill
type ActionKill struct {
}

func (a ActionKill) OpCode() BehaviourAction {
	return BehaviourActionKill
}

func (a ActionKill) Init(body *byter.Byter) {
	panic("implement me")
}

func NewActionKill() *ActionKill {
	return &ActionKill{}
}
