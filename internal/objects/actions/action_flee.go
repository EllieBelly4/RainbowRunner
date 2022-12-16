package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../../scripts/generatelua -type=Flee
type Flee struct {
}

func (d Flee) OpCode() BehaviourAction {
	return BehaviourActionFlee
}

func (d Flee) Init(body *byter.Byter) {
	panic("implement me")
}
