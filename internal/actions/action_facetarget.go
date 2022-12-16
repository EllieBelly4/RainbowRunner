package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../../scripts/generatelua -type=ActionFaceTarget
type ActionFaceTarget struct {
}

func (a ActionFaceTarget) OpCode() BehaviourAction {
	return BehaviourActionFaceTarget
}

func (a ActionFaceTarget) Init(body *byter.Byter) {
	panic("implement me")
}

func NewActionFaceTarget() *ActionFaceTarget {
	return &ActionFaceTarget{}
}
