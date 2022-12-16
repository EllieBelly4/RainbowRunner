package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../scripts/generatelua -type=ActionKnockBack
type ActionKnockBack struct {
}

func (a ActionKnockBack) OpCode() BehaviourAction {
	return BehaviourActionKnockBack
}

func (a ActionKnockBack) Init(body *byter.Byter) {
	panic("implement me")
}

func NewActionKnockBack() *ActionKnockBack {
	return &ActionKnockBack{}
}
