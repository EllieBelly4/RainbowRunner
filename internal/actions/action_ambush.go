package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../../scripts/generatelua -type=ActionAmbush
type ActionAmbush struct {
}

func (a ActionAmbush) OpCode() BehaviourAction {
	return BehaviourActionAmbush
}

func (a ActionAmbush) Init(body *byter.Byter) {
	panic("implement me")
}

func NewActionAmbush() *ActionAmbush {
	return &ActionAmbush{}
}
