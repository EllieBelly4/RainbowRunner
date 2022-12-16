package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../scripts/generatelua -type=ActionWait
type ActionWait struct {
}

func (a ActionWait) OpCode() BehaviourAction {
	return BehaviourActionWait
}

func (a ActionWait) Init(body *byter.Byter) {
	panic("implement me")
}

func NewActionWait() *ActionWait {
	return &ActionWait{}
}
