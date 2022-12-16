package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../../scripts/generatelua -type=KnockBack
type KnockBack struct {
}

func (d KnockBack) OpCode() BehaviourAction {
	return BehaviourActionKnockBack
}

func (d KnockBack) Init(body *byter.Byter) {
	panic("implement me")
}
