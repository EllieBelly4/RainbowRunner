package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../../scripts/generatelua -type=ActionFlee
type ActionFlee struct {
}

func (a ActionFlee) OpCode() BehaviourAction {
	return BehaviourActionFlee
}

func (a ActionFlee) Init(body *byter.Byter) {
	panic("implement me")
}

func NewActionFlee() *ActionFlee {
	return &ActionFlee{}
}
