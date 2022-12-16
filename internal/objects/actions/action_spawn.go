package actions

import "RainbowRunner/pkg/byter"

type Spawn struct {
}

func (d Spawn) OpCode() BehaviourAction {
	return BehaviourActionSpawn
}

func (d Spawn) Init(body *byter.Byter) {
	panic("implement me")
}
