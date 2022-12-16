package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../../scripts/generatelua -type=PlayAnimation
type PlayAnimation struct {
}

func (d PlayAnimation) OpCode() BehaviourAction {
	return BehaviourActionPlayAnimation
}

func (d PlayAnimation) Init(body *byter.Byter) {
	panic("implement me")
}
