package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../../scripts/generatelua -type=UseTarget
type UseTarget struct {
}

func (d UseTarget) OpCode() BehaviourAction {
	return BehaviourActionUseTarget
}

func (d UseTarget) Init(body *byter.Byter) {
	panic("implement me")
}
