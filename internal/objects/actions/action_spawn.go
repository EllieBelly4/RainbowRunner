package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../../scripts/generatelua -type=Spawn
type Spawn struct {
}

func (d Spawn) OpCode() BehaviourAction {
	return BehaviourActionSpawn
}

func (d Spawn) Init(body *byter.Byter) {
	panic("implement me")
}
