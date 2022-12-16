package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../scripts/generatelua -type=ActionUnSpawn
type ActionUnSpawn struct {
}

func (a ActionUnSpawn) OpCode() BehaviourAction {
	return BehaviourActionUnSpawn
}

func (a ActionUnSpawn) Init(body *byter.Byter) {
	panic("implement me")
}

func NewActionUnSpawn() *ActionUnSpawn {
	return &ActionUnSpawn{}
}
