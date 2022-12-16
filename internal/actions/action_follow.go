package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../scripts/generatelua -type=ActionFollow
type ActionFollow struct {
}

func (a ActionFollow) OpCode() BehaviourAction {
	return BehaviourActionFollow
}

func (a ActionFollow) Init(body *byter.Byter) {
	panic("implement me")
}

func NewActionFollow() *ActionFollow {
	return &ActionFollow{}
}
