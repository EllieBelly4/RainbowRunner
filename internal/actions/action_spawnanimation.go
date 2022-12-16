package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../scripts/generatelua -type=ActionSpawnAnimation
type ActionSpawnAnimation struct {
}

func (a ActionSpawnAnimation) OpCode() BehaviourAction {
	return BehaviourActionSpawnAnimation
}

func (a ActionSpawnAnimation) Init(body *byter.Byter) {
	panic("implement me")
}

func NewActionSpawnAnimation() *ActionSpawnAnimation {
	return &ActionSpawnAnimation{}
}
