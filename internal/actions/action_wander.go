package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../../scripts/generatelua -type=ActionWander
type ActionWander struct {
}

func (a ActionWander) OpCode() BehaviourAction {
	return BehaviourActionWander
}

func (a ActionWander) Init(body *byter.Byter) {
	panic("implement me")
}

func NewActionWander() *ActionWander {
	return &ActionWander{}
}
