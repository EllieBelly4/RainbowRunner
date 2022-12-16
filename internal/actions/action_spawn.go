package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../scripts/generatelua -type=ActionSpawn
type ActionSpawn struct {
}

func (a ActionSpawn) OpCode() BehaviourAction {
	return BehaviourActionSpawn
}

func (a ActionSpawn) Init(body *byter.Byter) {
	panic("implement me")
}

func NewActionSpawn() *ActionSpawn {
	return &ActionSpawn{}
}
