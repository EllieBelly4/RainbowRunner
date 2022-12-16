package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../scripts/generatelua -type=ActionHide
type ActionHide struct {
}

func (a ActionHide) OpCode() BehaviourAction {
	return BehaviourActionHide
}

func (a ActionHide) Init(body *byter.Byter) {
	panic("implement me")
}

func NewActionHide() *ActionHide {
	return &ActionHide{}
}
