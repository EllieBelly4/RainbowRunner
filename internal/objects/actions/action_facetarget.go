package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../../scripts/generatelua -type=FaceTarget
type FaceTarget struct {
}

func (d FaceTarget) OpCode() BehaviourAction {
	return BehaviourActionFaceTarget
}

func (d FaceTarget) Init(body *byter.Byter) {
	panic("implement me")
}
