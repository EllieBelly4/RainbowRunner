package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../../scripts/generatelua -type=KnockDown
type KnockDown struct {
}

func (d KnockDown) OpCode() BehaviourAction {
	return BehaviourActionKnockDown
}

func (d KnockDown) Init(body *byter.Byter) {
	panic("implement me")
}
