package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../../scripts/generatelua -type=Follow
type Follow struct {
}

func (d Follow) OpCode() BehaviourAction {
	return BehaviourActionFollow
}

func (d Follow) Init(body *byter.Byter) {
	panic("implement me")
}
