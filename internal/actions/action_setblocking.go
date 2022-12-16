package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../../scripts/generatelua -type=ActionSetBlocking
type ActionSetBlocking struct {
}

func (a ActionSetBlocking) OpCode() BehaviourAction {
	return BehaviourActionSetBlocking
}

func (a ActionSetBlocking) Init(body *byter.Byter) {
	panic("implement me")
}

func NewActionSetBlocking() *ActionSetBlocking {
	return &ActionSetBlocking{}
}
