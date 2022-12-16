package actions

import "RainbowRunner/pkg/byter"

type AttackTarget2 struct {
}

func (d AttackTarget2) OpCode() BehaviourAction {
	return BehaviourActionAttackTarget2
}

func (d AttackTarget2) Init(body *byter.Byter) {
	panic("implement me")
}
