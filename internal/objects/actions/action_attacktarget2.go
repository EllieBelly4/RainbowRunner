package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../../scripts/generatelua -type=AttackTarget2
type AttackTarget2 struct {
}

func (d AttackTarget2) OpCode() BehaviourAction {
	return BehaviourActionAttackTarget2
}

func (d AttackTarget2) Init(body *byter.Byter) {
	panic("implement me")
}
