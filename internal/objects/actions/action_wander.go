package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../../scripts/generatelua -type=Wander
type Wander struct {
}

func (d Wander) OpCode() BehaviourAction {
	return BehaviourActionWander
}

func (d Wander) Init(body *byter.Byter) {
	panic("implement me")
}
