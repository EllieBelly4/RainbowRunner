package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../scripts/generatelua -type=ActionFaceTarget
type ActionFaceTarget struct {
	TargetID uint16
}

func (a ActionFaceTarget) OpCode() BehaviourAction {
	return BehaviourActionFaceTarget
}

func (a ActionFaceTarget) Init(body *byter.Byter) {
	body.WriteUInt16(a.TargetID)
}

func NewActionFaceTarget(targetID uint16) *ActionFaceTarget {
	return &ActionFaceTarget{
		TargetID: targetID,
	}
}
