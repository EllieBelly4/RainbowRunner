package actions

import "RainbowRunner/pkg/byter"

type FaceTarget struct {
}

func (d FaceTarget) OpCode() BehaviourAction {
	return BehaviourActionFaceTarget
}

func (d FaceTarget) Init(body *byter.Byter) {
	panic("implement me")
}
