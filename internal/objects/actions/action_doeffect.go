package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../../scripts/generatelua -type=DoEffect
type DoEffect struct {
}

func (d DoEffect) OpCode() BehaviourAction {
	return BehaviourActionDoEffect
}

func (d DoEffect) Init(body *byter.Byter) {
	panic("implement me")
}
