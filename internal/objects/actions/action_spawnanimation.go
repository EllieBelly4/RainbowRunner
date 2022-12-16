package actions

import "RainbowRunner/pkg/byter"

type SpawnAnimation struct {
}

func (d SpawnAnimation) OpCode() BehaviourAction {
	return BehaviourActionSpawnAnimation
}

func (d SpawnAnimation) Init(body *byter.Byter) {
	panic("implement me")
}
