package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../../scripts/generatelua -type=ActionImmobilize
type ActionImmobilize struct {
}

func (a ActionImmobilize) OpCode() BehaviourAction {
	return BehaviourActionImmobilize
}

func (a ActionImmobilize) Init(body *byter.Byter) {
	panic("implement me")
}

func NewActionImmobilize() *ActionImmobilize {
	return &ActionImmobilize{}
}
