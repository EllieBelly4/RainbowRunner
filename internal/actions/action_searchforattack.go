package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../scripts/generatelua -type=ActionSearchForAttack
type ActionSearchForAttack struct {
}

func (a ActionSearchForAttack) OpCode() BehaviourAction {
	return BehaviourActionSearchForAttack
}

func (a ActionSearchForAttack) Init(body *byter.Byter) {
	panic("implement me")
}

func NewActionSearchForAttack() *ActionSearchForAttack {
	return &ActionSearchForAttack{}
}
