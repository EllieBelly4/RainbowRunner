package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../scripts/generatelua -type=ActionDodge
type ActionDodge struct {
}

func (a ActionDodge) OpCode() BehaviourAction {
	return BehaviourActionDodge
}

func (a ActionDodge) Init(body *byter.Byter) {
	panic("implement me")
}

func NewActionDodge() *ActionDodge {
	return &ActionDodge{}
}
