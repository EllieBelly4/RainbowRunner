package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../../scripts/generatelua -type=SpawnAnimation
type SpawnAnimation struct {
}

func (d SpawnAnimation) OpCode() BehaviourAction {
	return BehaviourActionSpawnAnimation
}

func (d SpawnAnimation) Init(body *byter.Byter) {
	panic("implement me")
}
