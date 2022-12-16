package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../scripts/generatelua -type=ActionAttackTarget2
type ActionAttackTarget2 struct {
}

func (a ActionAttackTarget2) OpCode() BehaviourAction {
	return BehaviourActionAttackTarget2
}

func (a ActionAttackTarget2) Init(body *byter.Byter) {
	panic("implement me")
}

func NewActionAttackTarget2() *ActionAttackTarget2 {
	return &ActionAttackTarget2{}
}
