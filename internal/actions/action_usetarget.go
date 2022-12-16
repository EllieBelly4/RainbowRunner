package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../scripts/generatelua -type=ActionUseTarget
type ActionUseTarget struct {
}

func (a ActionUseTarget) OpCode() BehaviourAction {
	return BehaviourActionUseTarget
}

func (a ActionUseTarget) Init(body *byter.Byter) {
	panic("implement me")
}

func NewActionUseTarget() *ActionUseTarget {
	return &ActionUseTarget{}
}
