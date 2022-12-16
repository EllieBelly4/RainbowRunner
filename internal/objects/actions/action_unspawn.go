package actions

import "RainbowRunner/pkg/byter"

type UnSpawn struct {
}

func (d UnSpawn) OpCode() BehaviourAction {
	return BehaviourActionUnSpawn
}

func (d UnSpawn) Init(body *byter.Byter) {
	panic("implement me")
}
