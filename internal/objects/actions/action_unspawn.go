package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../../scripts/generatelua -type=UnSpawn
type UnSpawn struct {
}

func (d UnSpawn) OpCode() BehaviourAction {
	return BehaviourActionUnSpawn
}

func (d UnSpawn) Init(body *byter.Byter) {
	panic("implement me")
}
