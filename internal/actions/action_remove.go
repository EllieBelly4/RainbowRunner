package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../scripts/generatelua -type=ActionRemove
type ActionRemove struct {
}

func (a ActionRemove) OpCode() BehaviourAction {
	return BehaviourActionRemove
}

func (a ActionRemove) Init(body *byter.Byter) {
	panic("implement me")
}

func NewActionRemove() *ActionRemove {
	return &ActionRemove{}
}
