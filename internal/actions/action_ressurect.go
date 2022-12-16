package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../../scripts/generatelua -type=ActionRessurect
type ActionRessurect struct {
}

func (a ActionRessurect) OpCode() BehaviourAction {
	return BehaviourActionRessurect
}

func (a ActionRessurect) Init(body *byter.Byter) {
	panic("implement me")
}

func NewActionRessurect() *ActionRessurect {
	return &ActionRessurect{}
}
