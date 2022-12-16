package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../../scripts/generatelua -type=ActionDoEffect
type ActionDoEffect struct {
}

func (a ActionDoEffect) OpCode() BehaviourAction {
	return BehaviourActionDoEffect
}

func (a ActionDoEffect) Init(body *byter.Byter) {
	panic("implement me")
}

func NewActionDoEffect() *ActionDoEffect {
	return &ActionDoEffect{}
}
