package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../../scripts/generatelua -type=UseItemTarget
type UseItemTarget struct {
}

func (d UseItemTarget) OpCode() BehaviourAction {
	return BehaviourActionUseItemTarget
}

func (d UseItemTarget) Init(body *byter.Byter) {
	panic("implement me")
}
