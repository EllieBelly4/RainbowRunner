package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../scripts/generatelua -type=ActionUseItemTarget
type ActionUseItemTarget struct {
}

func (a ActionUseItemTarget) OpCode() BehaviourAction {
	return BehaviourActionUseItemTarget
}

func (a ActionUseItemTarget) Init(body *byter.Byter) {
	panic("implement me")
}

func NewActionUseItemTarget() *ActionUseItemTarget {
	return &ActionUseItemTarget{}
}
