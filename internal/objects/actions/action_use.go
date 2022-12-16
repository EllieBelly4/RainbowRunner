package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../../scripts/generatelua -type=Use
type Use struct {
}

func (d Use) OpCode() BehaviourAction {
	return BehaviourActionUse
}

func (d Use) Init(body *byter.Byter) {
	panic("implement me")
}
